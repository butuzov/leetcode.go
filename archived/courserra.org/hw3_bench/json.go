package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
)

// вам надо написать более быструю оптимальную этой функции
func JsonSearch(out io.Writer) {

	// ----- Part 1 - Read File with json parts..
	var err error
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalln(err)
	}

	// ----- Part 2a - Users - Users Count.
	var splits = make([]int, 0, 124)
	for i := 0; i < len(buf)-1; i++ {
		if buf[i] == 10 {
			splits = append(splits, i)
		}
	}

	// ----- Part 2b - Users - Get Users
	var ch = make(chan []byte)
	var Users = make([]User, 0, len(splits))
	var wg sync.WaitGroup

	// bytes to json
	go func(wg *sync.WaitGroup, users *[]User) {
		for buf := range ch {
			var user User
			json.Unmarshal(buf, &user)
			*users = append(*users, user)
			wg.Done()
		}
	}(&wg, &Users)

	var i, j int
	for _, j = range splits {
		wg.Add(1)
		ch <- buf[i:j]
		i = j
	}
	close(ch)
	wg.Wait()

	// ---- Part 3 & 4 - Browsers & Users
	var browsers = make(map[string]bool)

	fmt.Fprintln(out, "found users:")
	for i, user := range Users {

		var MsieUser bool
		var AndroidUser bool

		for _, browser := range user.Browsers {

			var msie = strings.Contains(browser, "MSIE")
			var android = strings.Contains(browser, "Android")

			if !(msie || android) {
				continue
			}

			if !MsieUser && msie {
				MsieUser = true
			}

			if !AndroidUser && android {
				AndroidUser = true
			}

			if _, ok := browsers[browser]; !ok {
				browsers[browser] = true
			}
		}

		if MsieUser && AndroidUser {
			fmt.Fprintf(out, "[%d] %s <%s>\n",
				i,
				user.Name,
				strings.Replace(user.Email, "@", " [at] ", 1),
			)
		}
	}

	fmt.Fprintln(out)

	// ---- Part 5 - Browsers Stats
	fmt.Fprintln(out, "Total unique browsers", len(browsers))
}

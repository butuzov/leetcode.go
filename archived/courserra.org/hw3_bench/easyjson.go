package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

// TODO - Optimize unmarshaler

func EasySearch(out io.Writer) {

	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalln(err)
	}

	var ch = make(chan []byte)
	var wg sync.WaitGroup

	var Browsers = [][]byte{
		[]byte("MSIE"),
		[]byte("Android"),
	}
	var (
		At      = []byte{'@'}
		NoSpam  = []byte(" [at] ")
		SqOpen  = []byte{'['}
		SqClose = []byte{']', ' '}
		COpen   = []byte{' ', '<'}
		CClose  = []byte{'>', '\n'}
	)

	var BrowsersUnique = make(map[string]bool)

	go func() {

		var m sync.Mutex
		var counter int

		for buf := range ch {

			if buf == nil {
				counter++
				continue
			}

			var user User

			err := user.UnmarshalJSON(buf)
			if err != nil {
				fmt.Println("easyjson", err)
				wg.Done()
				continue
			}

			// browsers
			m.Lock()
			var Android, MSIE bool
			for _, l := range user.Browsers {
				if strings.Contains(l, "Android") {
					Android = true
					BrowsersUnique[l] = true
					continue
				}

				if strings.Contains(l, "MSIE") {
					MSIE = true
					BrowsersUnique[l] = true
					continue
				}
			}

			m.Unlock()

			if MSIE && Android {

				out.Write(SqOpen)
				out.Write([]byte(strconv.Itoa(counter)))
				out.Write(SqClose)
				out.Write([]byte(user.Name))
				out.Write(COpen)
				out.Write(bytes.Replace([]byte(user.Email), At, NoSpam, 1))
				out.Write(CClose)

			}
			counter++
			wg.Done()
		}
	}()

	out.Write([]byte("found users:\n"))

	var counter int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		var b = scanner.Bytes()

		if !(bytes.Contains(b, Browsers[0]) || bytes.Contains(b, Browsers[1])) {
			counter++
			ch <- nil
			continue
		}

		var buf = make([]byte, len(b))
		copy(buf, b)

		wg.Add(1)
		ch <- buf
		counter++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	close(ch)
	wg.Wait()

	fmt.Fprintln(out, "\nTotal unique browsers", len(BrowsersUnique))
}

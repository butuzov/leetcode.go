package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/buger/jsonparser"
)

func ParserSearch(out io.Writer) {

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

			// browsers
			m.Lock()
			var Android, MSIE bool
			_, err2 := jsonparser.ArrayEach(buf,
				func(value []byte, dataType jsonparser.ValueType, offset int, errInner error) {
					if errInner != nil {
						return
					}

					if bytes.Contains(value, Browsers[0]) {
						Android = true
						BrowsersUnique[string(value)] = true

					} else if bytes.Contains(value, Browsers[1]) {
						MSIE = true
						BrowsersUnique[string(value)] = true
					}

				}, "browsers")
			m.Unlock()
			if err2 != nil {
				wg.Done()
				continue
			}

			if MSIE && Android {

				email, _, _, err := jsonparser.Get(buf, "email")
				if err != nil {
					wg.Done()
					continue
				}

				name, _, _, err := jsonparser.Get(buf, "name")
				if err != nil {
					wg.Done()
					continue
				}

				out.Write(SqOpen)
				out.Write([]byte(strconv.Itoa(counter)))
				out.Write(SqClose)
				out.Write(name)
				out.Write(COpen)
				out.Write(bytes.Replace(email, At, NoSpam, 1))
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

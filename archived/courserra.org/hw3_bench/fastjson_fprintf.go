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

	"github.com/valyala/fastjson"
)

// just to see overhead for single fprintf command

func FprintfFastJsonSearch(out io.Writer) {

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
		At     = []byte{'@'}
		NoSpam = []byte(" [at] ")
	)
	var BrowsersUnique = make(map[string]bool)

	go func() {
		var p fastjson.Parser
		var m sync.Mutex
		var counter int

		for buf := range ch {

			if buf == nil {
				counter++
				continue
			}

			v, err := p.ParseBytes(buf)
			if err != nil {
				fmt.Println("fastjson", err)
				wg.Done()
				continue
			}

			// browsers
			m.Lock()
			var Android, MSIE bool
			for i := 0; i < 1000; i++ {
				var l = v.GetStringBytes("browsers", strconv.Itoa(i))
				if len(l) > 0 {
					if bytes.Contains(l, Browsers[0]) {
						Android = true
						BrowsersUnique[string(l)] = true
						continue
					}

					if bytes.Contains(l, Browsers[1]) {
						MSIE = true
						BrowsersUnique[string(l)] = true
						continue
					}

					continue
				}
				break
			}
			m.Unlock()

			if MSIE && Android {
				fmt.Fprintf(out, "[%d] %s <%s>\n",
					counter,
					string(v.GetStringBytes("name")),
					string(bytes.Replace(v.GetStringBytes("email"), At, NoSpam, 1)),
				)
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

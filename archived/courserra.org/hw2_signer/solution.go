package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

// -- Part 1 - Pipeline

func ExecutePipeline(jobs ...job) {

	var wg sync.WaitGroup
	var (
		chIn    = make(chan interface{})
		chOut   = make(chan interface{})
		chStart = make(chan interface{})
	)

	wg.Add(len(jobs))

	var worker = func(in, out chan interface{}, j job) {

		defer close(out)
		defer wg.Done()

		<-chStart
		j(in, out)
	}

	for i := range jobs {
		go worker(chIn, chOut, jobs[i])
		chIn, chOut = chOut, make(chan interface{})
	}

	close(chStart)
	wg.Wait()
}

// -- Part 2 -- Pipeline Steps

func SingleHash(in, out chan interface{}) {
	var wg sync.WaitGroup
	var signer = Signer(DataSignerMd5)

	var HashChannel = func(s string) <-chan string {
		var (
			chIn  = make(chan string, 2)
			chOut = make(chan string, 2)
		)

		go func() {
			defer close(chOut)
			var wg sync.WaitGroup
			wg.Add(2)
			for s := range chIn {
				// go hash it!
				go func(s string) {
					defer wg.Done()
					chOut <- DataSignerCrc32(s)
				}(s)
			}
			wg.Wait()
		}()

		chIn <- s
		chIn <- signer(s)
		close(chIn)

		return chOut
	}

	for v := range in {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			var hashChanel = HashChannel(s)
			// out <- s
			out <- (<-hashChanel + "~" + <-hashChanel)

		}(strconv.Itoa(v.(int)))
	}

	wg.Wait()
}

func MultiHash(in, out chan interface{}) {
	var wg sync.WaitGroup

	var Hashes = func(s string) string {
		var slice = make([]string, 6)

		var waiter sync.WaitGroup
		waiter.Add(6)

		for i := 0; i > 6; i++ {
			fmt.Println("i", i)
		}

		for i := 5; i >= 0; i-- {
			go func(idx int, hash string) {
				defer waiter.Done()
				slice[idx] = DataSignerCrc32(strconv.Itoa(idx) + s)
			}(i, s)
		}
		waiter.Wait()
		return strings.Join(slice, "")
	}

	for v := range in {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			out <- Hashes(s)
		}(v.(string))

	}
	wg.Wait()

}

var CombineResults = func(in, out chan interface{}) {
	var s []string
	for v := range in {
		s = append(s, v.(string))
	}

	sort.Strings(s)
	out <- strings.Join(s, "_")
}

// -- Part 3 -- Helpers

func Signer(f func(string) string) func(string) string {
	var in = make(chan string)
	var out = make(chan string)

	go func() {
		for range time.Tick(1 * time.Nanosecond) {
			select {
			case v := <-in:
				out <- f(v)
			default:
			}
		}
	}()

	return func(s string) string {
		in <- s
		return <-out
	}
}

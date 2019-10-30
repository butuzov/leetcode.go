package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, e := ioutil.ReadFile("really_long_palindrom.txt")
	if e != nil {
		log.Fatal(e)
	}

	fmt.Println(string(b))
}

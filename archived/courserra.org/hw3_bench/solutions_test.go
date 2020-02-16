package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"testing"
)

type Search func(io.Writer)

func test(t *testing.T, fnc, etalonFnc Search) {
	slowOut := new(bytes.Buffer)
	etalonFnc(slowOut)
	slowResult := slowOut.String()

	fastOut := new(bytes.Buffer)
	fnc(fastOut)
	fastResult := fastOut.String()

	if slowResult != fastResult {
		t.Errorf("results not match\nGot:\n%v\nExpected:\n%v", fastResult, slowResult)
	}
}

// -- tests

func Test_FastJson_Vs_Slow(t *testing.T) {
	test(t, FastJsonSearch, SlowSearch)
}
func Test_FprintFastJson_Vs_Slow(t *testing.T) {
	test(t, FprintfFastJsonSearch, SlowSearch)
}
func Test_Easy_Vs_Slow(t *testing.T) {
	test(t, EasySearch, SlowSearch)
}
func Test_Json_Vs_Slow(t *testing.T) {
	test(t, JsonSearch, SlowSearch)
}
func Test_Parser_Vs_Slow(t *testing.T) {
	test(t, ParserSearch, SlowSearch)
}

// -- bechmarks

func Benchmark_SlowSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlowSearch(ioutil.Discard)
	}
}

func Benchmark_FastJsonSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastJsonSearch(ioutil.Discard)
	}
}

func Benchmark_FprintfFastJsonSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FprintfFastJsonSearch(ioutil.Discard)
	}
}

func Benchmark_EasySearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EasySearch(ioutil.Discard)
	}
}

func Benchmark_JsonSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JsonSearch(ioutil.Discard)
	}
}

func Benchmark_ParserSearch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ParserSearch(ioutil.Discard)
	}
}

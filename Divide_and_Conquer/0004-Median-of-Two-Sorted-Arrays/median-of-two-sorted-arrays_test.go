package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

/*******************************************************************************
  TestTable
*******************************************************************************/

var MessageError = "Fail: Input(%+v): Expected(%+v) vs Returend(%+v)"
var MessageOk = "OK: Input(%+v) as Expected(%+v)"

type TestCase struct {
	ArrayA  []int   `json:"a"`
	ArrayB  []int   `json:"b"`
	Medians []int   `json:"mi"`
	Median  float64 `json:"m"`
}

func getTestCases() []TestCase {
	// Reading json
	file, errJSONFile := ioutil.ReadFile("test-cases.json")
	if errJSONFile != nil {
		fmt.Printf("File error: %v\n", errJSONFile)
		os.Exit(1)
	}

	// Getting List of Problems from Json
	var cases = make([]TestCase, 10)
	json.Unmarshal(file, &cases)

	return cases
}

/*******************************************************************************
  Tests
*******************************************************************************/
func TestSortedArraysMedian(t *testing.T) {
	for _, test := range getTestCases() {
		median := findMedianSortedArrays(test.ArrayA, test.ArrayB)
		if median != test.Median {
			t.Errorf("Expected[%v] vs Returned[%v]", test.Median, median)
		}

	}
}

func TestSortedArraysMedianNaive(t *testing.T) {
	for _, test := range getTestCases() {
		median := findMedianSortedArraysNaive(test.ArrayA, test.ArrayB)
		if median != test.Median {
			t.Errorf("Expected[%v] vs Returned[%v]", test.Median, median)
		}

	}
}

func TestSortedArraysMedianTop(t *testing.T) {
	for _, test := range getTestCases() {
		median := findMedianSortedArraysTop(test.ArrayA, test.ArrayB)
		if median != test.Median {
			t.Errorf("Expected[%v] vs Returned[%v]", test.Median, median)
		}

	}
}

func TestSortedArraysMedians(t *testing.T) {
	for _, test := range getTestCases() {
		medians := medianaAssumptions(len(test.ArrayA), len(test.ArrayB))
		if !reflect.DeepEqual(test.Medians, medians) {
			t.Errorf("Expected medians [%v] vs [%v]", test.Medians, medians)
		}
	}
}

func BenchmarkProblem(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range getTestCases() {
			findMedianSortedArrays(test.ArrayA, test.ArrayB)
		}
	}
}

func BenchmarkProblemNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range getTestCases() {
			findMedianSortedArraysNaive(test.ArrayA, test.ArrayB)
		}
	}
}

func BenchmarkProblemTop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range getTestCases() {
			findMedianSortedArraysTop(test.ArrayA, test.ArrayB)
		}
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"
	"strings"
)

// Problem describe json structure.
type Problem struct {
	Name     string    `json:"name"`
	Slug     string    `json:"slug"`
	Problems []Problem `json:"unlocked"`
	Ready    bool
}

// List of Ready Exercises
var exercises []string

func main() {

	// Reading json
	file, err_json_file := ioutil.ReadFile("tracklist.json")
	if err_json_file != nil {
		fmt.Printf("File error: %v\n", err_json_file)
		os.Exit(1)
	}

	// Getting List of Exercises
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		if strings.Index(f.Name(), ".") < 0 {
			exercises = append(exercises, f.Name())
		}
	}

	// Getting List of Problems from Json
	var Problems []Problem
	json.Unmarshal(file, &Problems)

	// Maping ready
	Problems = isReady(Problems, exercises)

	// Getting list of tasks already finished.
	f, err_readme_file := os.Create("readme.md")
	if err_readme_file != nil {
		fmt.Printf("File write error: %v\n", err_readme_file)
		os.Exit(1)
	}

	// Saving Data to file.
	tpl := template.Must(template.ParseGlob("readme*"))
	tpl.ExecuteTemplate(f, "readme.md.tpl", Problems)
}

func isReady(p []Problem, ready []string) []Problem {
	for i, v := range p {

		for _, s := range ready {

			if len(v.Problems) > 0 {
				v.Problems = isReady(v.Problems, ready)
			}

			if s == v.Slug {
				v.Ready = true
				break
			}

		}

		p[i] = v
	}

	return p
}

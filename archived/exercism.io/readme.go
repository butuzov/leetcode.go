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
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Alias   string
	Level   string   `json:"level"`
	Topics  []string `json:"topics"`
	Topic   string
	Ready   bool
}

func main() {

	// Reading json
	file, err_json_file := ioutil.ReadFile("exercism.json")
	if err_json_file != nil {
		fmt.Printf("File error: %v\n", err_json_file)
		os.Exit(1)
	}

	// List of Ready Exercises
	var exercises []string

	// Getting List of Exercises
	files, _ := ioutil.ReadDir("./")
	for _, f := range files {
		if strings.Index(f.Name(), ".") < 0 {
			exercises = append(exercises, f.Name())
		}
	}

	// Getting List of Problems from Json
	var Problems = make(map[string]Problem)
	json.Unmarshal(file, &Problems)

	for k, v := range Problems {
		v.Alias = k

		for i := range v.Topics {
			v.Topics[i] = fmt.Sprintf("`%s`", v.Topics[i])
		}

		v.Topic = strings.Join(v.Topics, ", ")
		switch v.Level {
		case "easy":
			v.Level = "☆"
			break
		case "medium":

			v.Level = "☆☆"
			break
		case "hard":

			v.Level = "☆☆☆"
			break

		}
		Problems[k] = v
	}

	isReadyOrNot(Problems, exercises, true)
	isReadyOrNot(Problems, exercises, false)

	// Getting list of tasks already finished.
	f, errReadmeFile := os.Create("readme.md")
	if errReadmeFile != nil {
		fmt.Printf("File write error: %v\n", errReadmeFile)
		os.Exit(1)
	}

	// Saving Data to file.
	tpl := template.Must(template.ParseGlob("readme*"))
	tpl.ExecuteTemplate(f, "readme.md.tpl", Problems)
}

func isReadyOrNot(p map[string]Problem, exercises []string, ready bool) {

	for i, v := range p {

		var found bool

		for _, s := range exercises {
			if s == i {
				found = true
				break
			}
		}

		if found == ready {
			v.Ready = ready
			p[i] = v
		}
	}
}

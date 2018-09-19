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
	ID       int    `json:"Id"`
	Title    string `json:"Title"`
	Slug     string
	Level    int `json:"Level"`
	LevelStr string
	Topic    string

	Ready bool
}

func main() {

	// Reading json
	file, errJsonFile := ioutil.ReadFile("problems.json")
	if errJsonFile != nil {
		fmt.Printf("File error: %v\n", errJsonFile)
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

	for i, v := range Problems {

		// for i := range v.Topics {
		// 	v.Topics[i] = fmt.Sprintf("`%s`", v.Topics[i])
		// }

		switch v.Level {
		case 1:
			v.LevelStr = "☆"
			break
		case 2:
			v.LevelStr = "☆☆"
			break
		case 3:
			v.LevelStr = "☆☆☆"
			break
		}

		Problems[i] = v
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

		var fName = fmt.Sprintf("%.4d-%s", v.ID, strings.Replace(v.Title, " ", "-", -1))

		for i := range exercises {
			if exercises[i] == fName {
				found = true
				break
			}
		}

		if found == ready {
			v.Ready = ready

			if ready == true {
				v.Slug = fName
			}

			p[i] = v
		}
	}
}

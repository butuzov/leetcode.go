package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math"
	"os"
	"sort"
	"strings"
)

// Problem describe json structure.
type Problem struct {
	ID       int      `json:"Id"`
	Title    string   `json:"Title"`
	Level    int      `json:"Level"`
	Topics   []string `json:"topics"`
	Slug     string
	LevelStr string
	Topic    string

	Hard   bool
	Easy   bool
	Normal bool
	Ready  bool
}

type Problems []Problem

func main() {

	// Reading json
	file, errJsonFile := ioutil.ReadFile("problems.json")
	if errJsonFile != nil {
		fmt.Printf("File error: %v\n", errJsonFile)
		os.Exit(1)
	}

	// List of Ready Exercises
	var exercises = make(map[string]string)

	// Getting List of Exercises
	tags, _ := ioutil.ReadDir("./")
	for _, t := range tags {
		if strings.Index(t.Name(), ".") < 0 {

			files, _ := ioutil.ReadDir(fmt.Sprintf("./%s", t.Name()))
			for _, f := range files {
				if strings.Index(f.Name(), ".") < 0 {
					exercises[f.Name()] = t.Name() + "/" + f.Name()
				}
			}
		}
	}

	// Getting List of Problems from Json
	var ProblemsMap = make(map[string]Problem)

	json.Unmarshal(file, &ProblemsMap)

	for i, v := range ProblemsMap {

		for i := range v.Topics {
			v.Topics[i] = fmt.Sprintf("`%s`", v.Topics[i])
		}

		v.Topic = strings.Join(v.Topics, ", ")

		switch v.Level {
		case 1:
			v.LevelStr = "Easy"
			break
		case 2:
			v.LevelStr = "Normal"
			break
		case 3:
			v.LevelStr = "Hard"
			break
		}

		ProblemsMap[i] = v
	}

	isReadyOrNot(ProblemsMap, exercises)

	// Getting list of tasks already finished.
	f, errReadmeFile := os.Create("readme.md")
	if errReadmeFile != nil {
		fmt.Printf("File write error: %v\n", errReadmeFile)
		os.Exit(1)
	}

	// Saving Data to file.

	tmp, n := make(Problems, len(ProblemsMap)), 0
	for _, problem := range ProblemsMap {
		tmp[n] = problem
		n++
	}
	sort.Sort(tmp)

	solvedStat := map[string]int{"Easy": 0, "Normal": 0, "Hard": 0, "All": 0}
	totalStat := map[string]int{"Easy": 0, "Normal": 0, "Hard": 0, "All": 0}
	for i, p := range tmp {

		if p.Ready {
			solvedStat["All"]++
			solvedStat[p.LevelStr]++

			switch p.LevelStr {
			case "Easy":
				tmp[i].Easy = true
			case "Normal":
				tmp[i].Normal = true
			case "Hard":
				tmp[i].Hard = true
			}

		}
		totalStat[p.LevelStr]++
		totalStat["All"]++
	}

	var progressLen = 91
	solvedPercents := int(math.RoundToEven((float64(solvedStat["All"]) / float64(totalStat["All"])) * float64(progressLen)))

	tpl := template.Must(template.ParseGlob("readme*"))
	tpl.ExecuteTemplate(f, "readme.md.tpl", struct {
		List     Problems
		Progress string
		Stat     map[string]int
		Total    map[string]int
	}{
		tmp,
		strings.Repeat("▰", solvedPercents) + strings.Repeat("▱", progressLen-solvedPercents),
		solvedStat,
		totalStat,
	})
}

// Len is part of sort.Interface.
func (p Problems) Len() int {
	return len(p)
}

// Swap is part of sort.Interface.
func (p Problems) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (p Problems) Less(i, j int) bool {
	return p[i].ID > p[j].ID
}

func isReadyOrNot(p map[string]Problem, exercises map[string]string) {

	for i, v := range p {

		var fName = fmt.Sprintf("%.4d-%s", v.ID, strings.Replace(v.Title, " ", "-", -1))

		if location, ok := exercises[fName]; ok {
			v.Ready = true
			v.Slug = location
		} else {
			v.Ready = false
		}

		p[i] = v
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"
	"text/template"

	"github.com/butuzov/leetcode.go/internal/problems"
)

// Problem describe json structure.

func main() {
	// Reading json
	file, errJsonFile := ioutil.ReadFile("problems.json")
	if errJsonFile != nil {
		fmt.Printf("File error: %v\n", errJsonFile)
		os.Exit(1)
	}

	// List of Ready Exercises
	exercises := make(map[string]string)

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
	ProblemsMap := make(map[string]problems.Problem)

	json.Unmarshal(file, &ProblemsMap)

	problems.MapExisting(ProblemsMap, exercises)

	// Getting list of tasks already finished.
	f, errReadmeFile := os.Create("readme.md")
	if errReadmeFile != nil {
		fmt.Printf("File write error: %v\n", errReadmeFile)
		os.Exit(1)
	}

	// Saving Data to file.

	tmp := problems.ConvertMap(ProblemsMap)

	solvedStat := map[string]int{"Easy": 0, "Medium": 00, "Hard": 0, "All": 0}
	totalStat := map[string]int{"Easy": 0, "Medium": 0, "Hard": 0, "All": 0}
	percentsStat := map[string]int{"Easy": 0, "Medium": 0, "Hard": 0, "All": 0}

	for i, p := range tmp {

		if p.Ready {
			solvedStat["All"]++
			solvedStat[p.LevelStr]++

			switch p.LevelStr {
			case "Easy":
				tmp[i].Easy = true
			case "Medium":
				tmp[i].Medium = true
			case "Hard":
				tmp[i].Hard = true
			}

		}
		totalStat[p.LevelStr]++
		totalStat["All"]++
	}

	fmt.Println(totalStat)

	progressLen := 83
	solvedPercents := int(math.RoundToEven((float64(solvedStat["All"]) / float64(totalStat["All"])) * float64(progressLen)))

	for k := range percentsStat {
		fmt.Println(k, float64(solvedStat[k])/float64(totalStat[k]))
		percentsStat[k] = int(math.RoundToEven(float64(solvedStat[k]) / float64(totalStat[k]) * 100))
	}

	tpl, _ := template.New("readme.md.tpl").Parse(readmeTemplate)
	tpl.ExecuteTemplate(f, "readme.md.tpl", struct {
		List     problems.Problems
		Progress string
		Stat     map[string]int
		Total    map[string]int
		Percents map[string]int
	}{
		tmp,
		strings.Repeat("▰", solvedPercents) + strings.Repeat("▱", progressLen-solvedPercents),
		solvedStat,
		totalStat,
		percentsStat,
	})
}

package problems

import (
	"fmt"
	"strings"
)

type Problem struct {
	ID    int    `json:"Id"`
	Title string `json:"Title"`
	Slug  string

	Level    int    `json:"Level"`    // in json
	LevelStr string `json:"LevelStr"` // in readme

	Topics []string `json:"topics"` // in json
	Topic  string   // in readmemake readme

	Hard   bool
	Easy   bool
	Medium bool
	Ready  bool
}

func MapExisting(p map[string]Problem, exercises map[string]string) {
	for i, v := range p {

		fName := fmt.Sprintf("%.4d-%s", v.ID, strings.Replace(v.Title, " ", "-", -1))

		if location, ok := exercises[fName]; ok {
			v.Ready = true
			v.Slug = location
		} else {
			v.Ready = false
		}

		tags := []string{}
		for _, v := range v.Topics {
			tags = append(tags, fmt.Sprintf("`%s`", v))
		}
		v.Topic = strings.Join(tags, ", ")

		p[i] = v
	}
}

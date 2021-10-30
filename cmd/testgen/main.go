package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/butuzov/leetcode.go/internal/problems"
)

func main() {
	// read input
	var n int
	fmt.Println("Enter problem number")
	fmt.Scanf("%d", &n)

	// n := 37
	fmt.Println("out", n)

	// check problems
	data, err := getProblem(n)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	// generae case in the root of .sandbox

	d := fmt.Sprintf(".sandbox/%04d-%s", data.ID, strings.ReplaceAll(data.Title, " ", "-"))
	os.Mkdir(d, 0x755)

	time.Sleep(time.Second)
	// Getting list of tasks already finished.
	FilePutContents(d+"/readme.md", gen_readme(data))
	FilePutContents(d+"/main.go", gen_main(data))
	FilePutContents(d+"/main_test.go", gen_main_test(data))
}

func FilePutContents(loc string, b []byte) error {
	f, err := os.Create(loc)
	if err != nil {
		return err
	}

	f.Write(b)
	f.Close()
	return nil
}

func gen_readme(p *problems.Problem) []byte {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("# #%04d [%s](%s)\n", p.ID, p.Title, p.Slug))

	c := p.Content
	c = strings.ReplaceAll(c, "\n\n\n", "\n")
	c = strings.ReplaceAll(c, "\n\n", "\n")

	b.WriteString(c)
	return b.Bytes()
}

func gen_main_test(p *problems.Problem) []byte {
	var b bytes.Buffer

	reg := regexp.MustCompile("func ([a-zA-Z]{1,})\\((.*)\\)(.*){")

	res := reg.FindStringSubmatch(p.Code[0])
	// fmt.Printf("%#v\n", ))

	var argsTypes []string
	var testArgs []string
	args := strings.Split(res[2], ", ")
	for i := range args {

		// args
		a := strings.Split(args[i], " ")
		argsTypes = append(argsTypes, a[1])
		testArgs = append(testArgs, "test."+a[0])

		//
		args[i] = fmt.Sprintf("\t%s", strings.Trim(args[i], " "))

	}

	args = append(args, fmt.Sprintf("\texpected %s", strings.Trim(res[3], " ")))

	tpl, _ := template.New("main_test.go.tpl").Parse(testTemplate)
	tpl.ExecuteTemplate(&b, "main_test.go.tpl", struct {
		ReturnType      string
		Func            string
		FuncName        string
		FuncNameCapital string
		TestStruct      string
		TestArgs        string
	}{
		strings.Trim(res[3], " "),
		fmt.Sprintf("func(%s) %s", strings.Join(argsTypes, ", "), strings.Trim(res[3], " ")),
		res[1],
		strings.Title(res[1]),
		strings.Join(args, "\n"),
		strings.Join(testArgs, ", "),
	})

	return b.Bytes()
}

func gen_main(p *problems.Problem) []byte {
	var b bytes.Buffer

	b.WriteString("package main\n\n")
	b.WriteString("func main(){\n")
	b.WriteString("	// dude, write some code...\n")
	b.WriteString("}\n\n")

	b.WriteString(strings.ReplaceAll(p.Code[0], `\n`, "\n"))

	return b.Bytes()
}

func getProblem(n int) (*problems.Problem, error) {
	file, errProblemsFile := ioutil.ReadFile("problems.json")
	if errProblemsFile != nil {
		return nil, fmt.Errorf("read: %w", errProblemsFile)
	}
	// Getting List of Problems from Json
	ProblemsMap := make(map[string]problems.Problem)
	err := json.Unmarshal(file, &ProblemsMap)
	if err != nil {
		return nil, fmt.Errorf("json: %w", err)
	}

	for _, v := range ProblemsMap {
		if v.ID == n {
			return &v, nil
		}
	}

	return nil, errors.New("Problem Not found")
}

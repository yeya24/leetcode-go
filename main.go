package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

var valueMap = map[string]string{
	"<strong>":  "",
	"</strong>": "",
	"<em>":      "",
	"</em>":     "",
	"</p>":      "",
	"<p>":       "",
	"<b>":       "",
	"</b>":      "",
	"<pre>":     "",
	"</pre>":    "",
	"<ul>":      "",
	"</ul>":     "",
	"<li>":      "",
	"</li>":     "",
	"<code>":    "",
	"</code>":   "",
	"<i>":       "",
	"</i>":      "",
	"<sub>":     "",
	"</sub>":    "",
	"</sup>":    "",
	"<sup>":     "^",
	"&nbsp;":    " ",
	"&gt;":      ">",
	"&lt;":      "<",
	"&quot;":    "\"",
	"&minus;":   "-",
	"&#39;":     "'",
	"\n\n":      "\n",
	"\n":        "\n * ",
}

var (
	id    int
	title string
)

func main() {
	flag.IntVar(&id, "id", 1, "leetcode problem id")
	flag.StringVar(&title, "title", "", "leetcode problem title")
	flag.Parse()

	var (
		problem *problem
		err error
	)

	// use problem title
	if title != "" {
		problem, err = getProblemWithTitle(NewClient(15), title)
		if err != nil {
			fmt.Println("Cannot get the specific problem")
			os.Exit(1)
		}
	} else {
		problem, err = getProblemWithID(NewClient(15), id)
		if err != nil {
			fmt.Println("Cannot get the specific problem")
			os.Exit(1)
		}
	}

	filePath := fmt.Sprintf("tests/%04d_%s_test.go", problem.id, strings.Replace(problem.titleSlug, "-", "_", -1))
	if _, err := os.Stat(filePath); err == nil || os.IsExist(err) {
		fmt.Printf("File already exists", err)
		os.Exit(1)
	}

	sources, err := ioutil.ReadFile("template")
	if err != nil {
		fmt.Println("Error open problem template")
		os.Exit(1)
	}

	data := generateTemplate(string(sources), problem)
	// Generate code template.

	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		fmt.Println("Error generate problem template")
		os.Exit(1)
	}

	fmt.Println("generate problem template successfully")
}

func generateTemplate(data string, problem *problem) []byte {
	data = strings.Replace(data, "__PROBLEM_TITLE__", problem.title, -1)
	data = strings.Replace(data, "__PROBLEM_DESC__", buildDescription(problem.content), -1)
	data = strings.Replace(data, "__PROBLEM_ID__", strconv.Itoa(problem.id), -1)
	data = strings.Replace(data, "__PROBLEM_DEFAULT_CODE__", problem.Code, -1)
	data = strings.Replace(data, "__TEST_TITLE__", strings.ReplaceAll(problem.title, " ", ""), -1)
	return []byte(data)
}

func buildDescription(desc string) string {
	for k, v := range valueMap {
		desc = strings.Replace(desc, k, v, -1)
	}
	return desc
}

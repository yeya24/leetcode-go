package main

import (
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

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("You must specify Problem id")
		os.Exit(1)
	}

	id, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Problem id must be an integer!")
		os.Exit(1)
	}

	problem, err := getProblemWithID(NewClient(15), id)
	if err != nil {
		fmt.Println("Cannot get the specific problem")
		os.Exit(1)
	}

	fileName := fmt.Sprintf("%04d_%s", id, strings.Replace(problem.titleSlug, "-", "_", -1))
	filePath := "problems/" + fileName + ".go"

	if _, err := os.Stat(filePath); os.IsExist(err) {
		fmt.Printf("Error stat filepath, %v\n", err)
		os.Exit(1)
	}

	sources, err := ioutil.ReadFile("template")
	if err != nil {
		fmt.Println("Error open problem template")
		os.Exit(1)
	}

	data := generateTemplate(args[1], string(sources), problem)
	// Generate code template.

	if err := ioutil.WriteFile(filePath, data, 0644); err != nil {
		fmt.Println("Error generate problem template")
		os.Exit(1)
	}

	fmt.Println("generate problem template successfully")
}

func generateTemplate(id, data string, problem *problem) []byte {
	data = strings.Replace(data, "__PROBLEM_TITLE__", problem.title, -1)
	data = strings.Replace(data, "__PROBLEM_DESC__", buildDescription(problem.content), -1)
	data = strings.Replace(data, "__PROBLEM_ID__", id, -1)
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

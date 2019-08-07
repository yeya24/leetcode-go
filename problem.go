package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/machinebox/graphql"
)

const (
	GRAPHQL_URL  = "https://leetcode.com/graphql"
	PROBLEMS_URL = "https://leetcode.com/api/problems/all/"
)

var questionReq = graphql.NewRequest(`
query questionData($titleSlug: String!) {
    question(titleSlug: $titleSlug) {
        content
        stats
        codeDefinition
        sampleTestCase
        metaData
    }
}`)

type problem struct {
	id        int
	title     string
	titleSlug string
	Code      string
	content   string
}

func getProblemWithID(client *client, id int) (*problem, error) {
	problems, err := getProblems(client)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range problems.StatStatusPairs {
		if p.Stat.QuestionID == id {
			return getQuestionByID(client, p)
		}
	}

	return nil, errors.New(fmt.Sprintf("cannot find the specific problem id: %d\n", id))
}

func getProblemWithTitle(client *client, title string) (*problem, error) {
	problems, err := getProblems(client)
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range problems.StatStatusPairs {
		if p.Stat.QuestionTitle == title {
			return getQuestionByTitle(client, p)
		}
	}

	return nil, errors.New(fmt.Sprintf("cannot find the specific problem title: %s\n", title))
}

func getProblems(client *client) (*Problems, error) {
	data, err := client.Get(PROBLEMS_URL)
	if err != nil {

	}
	var problems Problems
	if err := json.Unmarshal(data, &problems); err != nil {
		return nil, err
	}
	return &problems, nil
}

func getQuestionByID(client *client, p ProblemStatStatus) (*problem, error) {
	// set any variables
	questionReq.Var("titleSlug", p.Stat.QuestionTitleSlug)
	rp := new(RawProblem)
	if err := client.gc.Run(context.Background(), questionReq, rp); err != nil {
		return nil, err
	}

	s, err := strconv.Unquote(strconv.Quote(rp.Question.CodeDefinition))
	if err != nil {
		return nil, err
	}

	problem := &problem{
		id:        p.Stat.QuestionID,
		title:     p.Stat.QuestionTitle,
		titleSlug: p.Stat.QuestionTitleSlug,
		content:   rp.Question.Content,
	}

	var cds Codes

	if err = json.Unmarshal([]byte(s), &cds); err != nil {
		return nil, err
	}

	problem.Code = cds.Code("go").DefaultCode

	return problem, nil
}

func getQuestionByTitle(client *client, p ProblemStatStatus) (*problem, error) {
	// set any variables
	questionReq.Var("titleSlug", p.Stat.QuestionTitleSlug)
	rp := new(RawProblem)
	if err := client.gc.Run(context.Background(), questionReq, rp); err != nil {
		return nil, err
	}

	s, err := strconv.Unquote(strconv.Quote(rp.Question.CodeDefinition))
	if err != nil {
		return nil, err
	}

	problem := &problem{
		id:        p.Stat.QuestionID,
		title:     p.Stat.QuestionTitle,
		titleSlug: p.Stat.QuestionTitleSlug,
		content:   rp.Question.Content,
	}

	var cds Codes

	if err = json.Unmarshal([]byte(s), &cds); err != nil {
		return nil, err
	}

	problem.Code = cds.Code("go").DefaultCode

	return problem, nil
}

type Problems struct {
	NumTotal        int                 `json:"num_total"`
	StatStatusPairs []ProblemStatStatus `json:"stat_status_pairs"`
}

// Code the struct of leetcode codes.
type Code struct {
	Text        string `json:"text"`
	Value       string `json:"value"`
	DefaultCode string `json:"defaultCode"`
}

// Codes the slice of Code
type Codes []*Code

// Code returns Code with lang.
func (cs Codes) Code(lang string) *Code {
	for _, c := range cs {
		if strings.ToLower(c.Text) == lang || c.Value == lang {
			return c
		}
	}
	return nil
}

type RawProblem struct {
	Question Question `json:"question"`
}

type Question struct {
	Content        string `json:"content"`
	Stats          string `json:"stats"`
	CodeDefinition string `json:"codeDefinition"`
	SampleTestCase string `json:"sampleTestCase"`
	MetaData       string `json:"metaData"`
}

type ProblemStat struct {
	QuestionTitle       string `json:"question__title"`
	QuestionArticleSlug string `json:"question__article__slug"`
	QuestionTitleSlug   string `json:"question__title_slug"`
	QuestionArticleLive bool   `json:"question__article__live"`
	QuestionID          int    `json:"question_id"`
}

type difficulty struct {
	Level int `json:"level"`
}

func (d difficulty) String() string {
	switch d.Level {
	case 1:
		return "Easy"
	case 2:
		return "Medium"
	case 3:
		return "Difficult"
	default:
		return "Unknown"
	}
}

type ProblemStatStatus struct {
	Stat       ProblemStat `json:"stat"`
	Difficulty difficulty  `json:"difficulty"`
}

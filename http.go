package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/machinebox/graphql"
)

type client struct {
	c *http.Client
	gc *graphql.Client
	timeout int
}

func NewClient(timeout int) *client {
	return &client{
		c:       http.DefaultClient,
		gc:      graphql.NewClient(GRAPHQL_URL),
		timeout: timeout,
	}
}

func (cli *client) Get(url string) ([]byte, error) {
	res, err := cli.c.Get(url)
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return nil, err
	}
	if len(data) <= 0 {
		return nil, errors.New("bad data length")
	}

	return data, nil
}

func (cli *client) Post(url string, data []byte) ([]byte, error) {
	var body []byte
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := cli.c.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode == http.StatusOK || res.StatusCode == http.StatusAccepted {
		body, err = ioutil.ReadAll(res.Body)
		if err != nil{
			return nil, err
		}
		return body, nil
	}
	return nil, errors.New(fmt.Sprintf("bad http status %d", res.StatusCode))
}

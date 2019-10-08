package github

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

type Client struct {
	BaseURL string
	Repo    string
	Owner   string
}

type GetIssuesQueryString struct {
	Filter    string `url:"filter"`
	State     string `url:"state"`
	Sort      string `url:"sort"`
	Direction string `url:"direction"`
}

func NewGithubClient(owner string, repo string) *Client {
	client := Client{}
	client.BaseURL = "https://api.github.com"
	client.Repo = repo
	client.Owner = owner

	return &client
}

func (client *Client) GetIssues(sort string, direction string) []RepoIssueResult {

	qsData := GetIssuesQueryString{
		Filter:    "assigned",
		State:     "open",
		Sort:      sort,
		Direction: direction,
	}
	qs, _ := query.Values(qsData)
	url := client.BaseURL + "/repos/" + client.Owner + "/" + client.Repo + "/issues" + "?" + qs.Encode()
	fmt.Printf("%s\r\n", url)

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var repoIssues []RepoIssueResult
	err = json.Unmarshal(body, &repoIssues)
	if err != nil {
		panic(err)
	}

	return repoIssues
}

func (client *Client) GetSingleIssue(id string) (SingleIssueResult, error) {

	url := client.BaseURL + "/repos/" + client.Owner + "/" + client.Repo + "/issues/" + id

	var issue SingleIssueResult
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	if res.StatusCode == 404 {
		return issue, errors.New("404 not found")
	}

	body, err := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &issue)
	if err != nil {
		panic(err)
	}
	return issue, nil
}

package repo

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fatih/structs"
	"github.com/ghub/gh"
	"github.com/ghub/util"
	"github.com/google/go-github/github"
	"log"
	//"github.com/olekukonko/tablewriter"
	//"os"
)

func getRepo(c *cli.Context) {
	o, r, err := util.SplitOrgRepoName(c.Args().Get(0))
	if err != nil {
		log.Fatal(err)
	}
	repo, _, err := gh.Client.Repositories.Get(o, r)
	if err != nil {
		log.Fatal(err)
	}

	//util.PrintJson(*repo)
	util.PrintTable(*repo)
	m := structs.Map(repo)
	for k, v := range m {
		fmt.Printf("Key type: %T, Value type: %T", k, v)
	}
}

func getRepoIssues(c *cli.Context) {
	owner, repo, err := util.SplitOrgRepoName(c.Args().Get(0))
	if err != nil {
		log.Fatal(err)
	}
	opts := &github.IssueListByRepoOptions{
		ListOptions: github.ListOptions{PerPage: 30},
	}
	issues, _, err := gh.Client.Issues.ListByRepo(owner, repo, opts)
	if err != nil {
		fmt.Println(err)
	}
	if len(issues) == 0 {
		fmt.Println("No issues found")
	}

	// PRs are considered issues. Remove all PRs from issue array.
	for index, issue := range issues {
		if issue.PullRequestLinks != nil {
			issues = append(issues[:index], issues[index+1:]...)
		}
	}

	util.PrintJson(issues)
}

func getRepoPullRequests(c *cli.Context) {
	owner, repo, err := util.SplitOrgRepoName(c.Args().Get(0))
	if err != nil {
		log.Fatal(err)
	}
	opts := &github.PullRequestListOptions{
		ListOptions: github.ListOptions{PerPage: 30},
	}
	prs, _, err := gh.Client.PullRequests.List(owner, repo, opts)
	if err != nil {
		fmt.Println(err)
	}
	if len(prs) == 0 {
		fmt.Println("No PRs found")
	}

	util.PrintJson(prs)
}

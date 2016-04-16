package repo

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ghub/gh"
	"github.com/ghub/util"
	"github.com/google/go-github/github"
	"log"
)

func getRepo(c *cli.Context) {
	if len(c.Args()) < 1 {
		log.Fatal("Usage: ghub org/repo")
	}
	o, r := util.SplitOrgRepoName(c.Args().Get(0))
	repo, _, err := gh.Client.Repositories.Get(o, r)
	if err != nil {
		log.Fatal(err)
	}

	util.PrintJson(*repo)
}

func getRepoIssues(c *cli.Context) {
	if len(c.Args()) < 1 {
		log.Fatal("Usage: ghub owner/repo")
	}
	owner, repo := util.SplitOrgRepoName(c.Args().Get(0))

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

	util.PrintJson(issues)
}

func getRepoPullRequests(c *cli.Context) {
	if len(c.Args()) < 1 {
		log.Fatal("Usage: ghub owner/repo")
	}
	owner, repo := util.SplitOrgRepoName(c.Args().Get(0))

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

package org

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ghub/gh"
	"github.com/ghub/util"
	"github.com/google/go-github/github"
)

// Return a single github.Organization struct
func getOrg(c *cli.Context) {
	org := c.Args().First()
	organization, _, err := gh.Client.Organizations.Get(org)
	if err != nil {
		fmt.Println(err)
	}

	util.PrintJson(organization)
}

func getOrgRepos(c *cli.Context) {
	org := c.Args().First()
	opts := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 30},
	}
	repos, _, err := gh.Client.Repositories.ListByOrg(org, opts)
	if err != nil {
		fmt.Println(err)
	}

	util.PrintJson(repos)
}

func getOrgTeams(c *cli.Context) {
	org := c.Args().First()
	opts := &github.ListOptions{PerPage: 30}
	teams, _, err := gh.Client.Organizations.ListTeams(org, opts)
	if err != nil {
		fmt.Println(err)
	}
	if len(teams) == 0 {
		fmt.Println("No teams found")
	}

	util.PrintJson(teams)
}

func getOrgMembers(c *cli.Context) {
	org := c.Args().First()
	opts := &github.ListMembersOptions{
		ListOptions: github.ListOptions{PerPage: 30},
	}
	members, _, err := gh.Client.Organizations.ListMembers(org, opts)
	if err != nil {
		fmt.Println(err)
	}

	util.PrintJson(members)
}

func getOrgHooks(c *cli.Context) {
	org := c.Args().First()
	opts := &github.ListOptions{PerPage: 30}

	hooks, _, err := gh.Client.Organizations.ListHooks(org, opts)
	if err != nil {
		fmt.Println(err)
	}
	if len(hooks) == 0 {
		fmt.Println("No hooks found")
	}

	util.PrintJson(hooks)
}

func getOrgIssues(c *cli.Context) {
	org := c.Args().First()
	opts := &github.IssueListOptions{
		ListOptions: github.ListOptions{PerPage: 30},
	}

	issues, _, err := gh.Client.Issues.ListByOrg(org, opts)
	if err != nil {
		fmt.Println(err)
	}
	if len(issues) == 0 {
		fmt.Println("No issues found")
	}

	util.PrintJson(issues)
}

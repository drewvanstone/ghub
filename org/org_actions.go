package org

import (
	"fmt"
    "os"
    "strconv"
    //"text/tabwriter"
    "github.com/olekukonko/tablewriter"
	"github.com/codegangsta/cli"
	"github.com/ghub/gh"
	"github.com/ghub/util"
	"github.com/google/go-github/github"
)

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

    table := tablewriter.NewWriter(os.Stdout)
    table.SetHeader([]string{"NAME", "STARS", "FORKS", "SIZE"})

    data := [][]string{}
    for _, r := range repos {
        d := []string{*r.Name, strconv.Itoa(*r.StargazersCount), strconv.Itoa(*r.ForksCount), strconv.Itoa(*r.Size)}
        data = append(data, d)
    }

    for _, v := range data {
            table.Append(v)
    }

    table.Render()

//    w := new(tabwriter.Writer)
//    w.Init(os.Stdout, 15, 8, 1, '\t', 0)
//    fmt.Fprintln(w, "NAME\tSTARS\tFORKS\tSIZE")
//    for _, r := range repos {
//        n := *r.Name
//        if len(n) > 10 {
//            fmt.Fprintf(w, "%s\t%d\t%d\t%d\f", n[:10], *r.StargazersCount, *r.ForksCount, *r.Size)
//        } else {
//            fmt.Fprintf(w, "%s\t%d\t%d\t%d\f", n, *r.StargazersCount, *r.ForksCount, *r.Size)
//        }
//    }
//    w.Flush()
//	//util.PrintJson(repos)
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

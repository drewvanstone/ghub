package org

import (
	"fmt"
	"sort"
	//"github.com/fatih/color"
	"github.com/codegangsta/cli"
	"github.com/ghub/gh"
	"github.com/ghub/util"
	"github.com/google/go-github/github"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
	"strings"
	//"text/tabwriter"
)

type repoRows [][]string

func (r repoRows) Len() int           { return len(r) }
func (r repoRows) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r repoRows) Less(i, j int) bool { return strings.Compare(r[i][0], r[j][0]) == -1 }

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

	var repositories []github.Repository
	for {
		repos, res, err := gh.Client.Repositories.ListByOrg(org, opts)
		if err != nil {
			fmt.Println(err)
		}

		repositories = append(repositories, repos...)
		if res.NextPage == 0 {
			break
		}
		opts.ListOptions.Page = res.NextPage
	}

	//grey := color.New(color.FgHiBlack).SprintFunc()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"NAME", "URL", "STARS", "FORKS", "SIZE", "PRIVATE", "CREATED"})

	repoRows := repoRows{}
	for _, r := range repositories {
		row := []string{
			*r.Name,
			*r.HTMLURL,
			strconv.Itoa(*r.StargazersCount),
			strconv.Itoa(*r.ForksCount),
			strconv.Itoa(*r.Size),
			strconv.FormatBool(*r.Private),
			r.CreatedAt.String(),
		}
		repoRows = append(repoRows, row)
	}

	sort.Sort(repoRows)

	for _, v := range repoRows {
		table.Append(v)
	}

	table.Render()

	// w := new(tabwriter.Writer)
	// w.Init(os.Stdout, 15, 8, 1, '\t', 0)
	// fmt.Fprintln(w, "NAME\tSTARS\tFORKS\tSIZE")
	// for _, r := range repos {
	// 	n := *r.Name
	// 	if len(n) > 10 {
	// 		fmt.Fprintf(w, "%s\t%d\t%d\t%d\f", n[:10], *r.StargazersCount, *r.ForksCount, *r.Size)
	// 	} else {
	// 		fmt.Fprintf(w, "%s\t%d\t%d\t%d\f", n, *r.StargazersCount, *r.ForksCount, *r.Size)
	// 	}
	// }
	// w.Flush()
	//util.PrintJson(repos)
}

//func getOrgTeams(c *cli.Context) {
//	org := c.Args().First()
//	opts := &github.ListOptions{PerPage: 30}
//	teams, _, err := gh.Client.Organizations.ListTeams(org, opts)
//	if err != nil {
//		fmt.Println(err)
//	}
//	if len(teams) == 0 {
//		fmt.Println("No teams found")
//	}
//
//	util.PrintJson(teams)
//}

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

	//util.PrintJson(teams)
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

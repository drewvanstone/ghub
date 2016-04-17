package issue

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ghub/gh"
	"github.com/ghub/util"
	"log"
	"strconv"
)

func getIssue(c *cli.Context) {
	owner, repo, err := util.SplitOrgRepoName(c.Args().Get(0))
	if err != nil {
		log.Fatal(err)
	}

	number, _ := strconv.Atoi(c.Args().Get(1))

	issue, _, err := gh.Client.Issues.Get(owner, repo, number)
	if err != nil {
		fmt.Println(err)
	}

	util.PrintJson(issue)
}

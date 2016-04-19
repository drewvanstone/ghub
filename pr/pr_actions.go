package pr

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ghub/gh"
	"github.com/ghub/util"
	"log"
	"strconv"
)

func getPr(c *cli.Context) {
	owner, repo, err := util.SplitOrgRepoName(c.Args().Get(0))
	if err != nil {
		log.Fatal(err)
	}

	number, _ := strconv.Atoi(c.Args().Get(1))

	pr, _, err := gh.Client.PullRequests.Get(owner, repo, number)
	if err != nil {
		fmt.Println(err)
	}

	util.PrintJson(pr)
}

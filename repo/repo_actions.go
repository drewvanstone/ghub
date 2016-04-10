package repo

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ghub/gh"
	"github.com/ghub/util"
	"log"
)

func getRepo(c *cli.Context) {
	if len(c.Args()) < 1 {
		log.Fatal("Usage: ghub org/repo")
	}
	o, r := util.SplitOrgRepoName(c.Args().Get(0))
	repo, resp, err := gh.Client.Repositories.Get(o, r)
	if err != nil {
		log.Fatal(err)
	}

	if resp != nil {
		util.PrintJson(*repo)
		fmt.Println(*resp)
	} else {
		fmt.Println("ERROR: Nil response")
	}
}

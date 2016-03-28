package org

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ghub/client"
	"github.com/ghub/util"
	"github.com/google/go-github/github"
	"log"
)

func getOrg(c *cli.Context) {
	if len(c.Args()) < 1 {
		log.Fatal("Some usage")
	}
	org := c.Args().First()

	gh := client.GitHub()
	o, _, _ := gh.Organizations.Get(org)

	util.PrintJson(*o)
}

func getReposInOrg(c *cli.Context) {
	if len(c.Args()) < 1 {
		log.Fatal("Usage: ghub get org {org}")
	}

	org := c.Args().First()
	opts := &github.RepositoryListByOrgOptions{
		ListOptions: github.ListOptions{PerPage: 30},
	}

	gh := client.GitHub()
	repos, _, _ := gh.Repositories.ListByOrg(org, opts)
	for _, repo := range repos {
		fmt.Println(*repo.Name)
	}
}

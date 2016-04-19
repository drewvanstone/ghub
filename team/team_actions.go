package team

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ghub/gh"
	"github.com/ghub/util"
	"strconv"
)

func getTeam(c *cli.Context) {
	teamNum, _ := strconv.Atoi(c.Args().First())
	team, _, err := gh.Client.Organizations.GetTeam(teamNum)
	if err != nil {
		fmt.Println(err)
	}

	util.PrintJson(team)
}

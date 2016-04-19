package user

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ghub/gh"
	"github.com/ghub/util"
)

func getUser(c *cli.Context) {
	u := c.Args().First()
	user, _, err := gh.Client.Users.Get(u)
	if err != nil {
		fmt.Println(err)
	}

	util.PrintJson(user)
}

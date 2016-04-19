package user

import (
	"github.com/codegangsta/cli"
	"github.com/ghub/util"
)

var Create cli.Command = cli.Command{
	Name:  "user",
	Usage: "Create a user (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("User created")
	},
}

var Delete cli.Command = cli.Command{
	Name:  "user",
	Usage: "Delete a user (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("User deleted")
	},
}

var Get cli.Command = cli.Command{
	Name:  "user",
	Usage: "Get a user",
	Action: func(c *cli.Context) {
		util.CheckCommandArgs(1, c)
		getUser(c)
	},
}

var Update cli.Command = cli.Command{
	Name:  "user",
	Usage: "Update a user (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("User updated")
	},
}

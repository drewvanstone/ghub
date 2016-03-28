package user

import (
	"github.com/codegangsta/cli"
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
	Usage: "Get a user (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("Got user")
	},
}

var Update cli.Command = cli.Command{
	Name:  "user",
	Usage: "Update a user (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("User updated")
	},
}

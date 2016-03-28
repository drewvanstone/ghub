package team

import (
	"github.com/codegangsta/cli"
)

var Create cli.Command = cli.Command{
	Name:  "team",
	Usage: "Create a team (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("team created")
	},
}

var Delete cli.Command = cli.Command{
	Name:  "team",
	Usage: "Delete a team (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("team deleted")
	},
}

var Get cli.Command = cli.Command{
	Name:  "team",
	Usage: "Get a team (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("Got team")
	},
}

var Update cli.Command = cli.Command{
	Name:  "team",
	Usage: "Update a team (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("team updated")
	},
}

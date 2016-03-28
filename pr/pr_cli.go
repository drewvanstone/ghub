package pr

import (
	"github.com/codegangsta/cli"
)

var Create cli.Command = cli.Command{
	Name:  "pr",
	Usage: "Create a pr (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("pr created")
	},
}

var Delete cli.Command = cli.Command{
	Name:  "pr",
	Usage: "Delete a pr (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("pr deleted")
	},
}

var Get cli.Command = cli.Command{
	Name:  "pr",
	Usage: "Get a pr (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("Got pr")
	},
}

var Update cli.Command = cli.Command{
	Name:  "pr",
	Usage: "Update a pr (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("pr updated")
	},
}

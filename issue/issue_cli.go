package issue

import (
	"github.com/codegangsta/cli"
	"github.com/ghub/util"
)

var Create cli.Command = cli.Command{
	Name:  "issue",
	Usage: "Create an issue (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("issue created")
	},
}

var Delete cli.Command = cli.Command{
	Name:  "issue",
	Usage: "Delete an issue (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("issue deleted")
	},
}

var Get cli.Command = cli.Command{
	Name:  "issue",
	Usage: "Get an issue (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		util.CheckCommandArgs(1, c)
		getIssue(c)
	},
}

var Update cli.Command = cli.Command{
	Name:  "issue",
	Usage: "Update an issue (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("issue updated")
	},
}

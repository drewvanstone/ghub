package repo

import (
	"github.com/codegangsta/cli"
)

var Create cli.Command = cli.Command{
	Name:  "repo",
	Usage: "Create a repo (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("repo created")
	},
}

var Delete cli.Command = cli.Command{
	Name:  "repo",
	Usage: "Delete a repo (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("repo deleted")
	},
}

var Get cli.Command = cli.Command{
	Name:  "repo",
	Usage: "ghub repo [org]/[repo]",
	Action: func(c *cli.Context) {
		getRepo(c)
	},
}

var Update cli.Command = cli.Command{
	Name:  "repo",
	Usage: "Update a repo (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("repo updated")
	},
}

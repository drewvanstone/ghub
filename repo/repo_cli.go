package repo

import (
	"github.com/codegangsta/cli"
	"github.com/ghub/util"
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
	Subcommands: []cli.Command{
		GetRepoIssuesCmd,
	},
	Action: func(c *cli.Context) {
		util.CheckCommandArgs(1, c)
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

var GetRepoIssuesCmd cli.Command = cli.Command{
	Name:  "issues",
	Usage: "Get all issues of a repo",
	Action: func(c *cli.Context) {
		util.CheckCommandArgs(1, c)
		getRepoIssues(c)
	},
}

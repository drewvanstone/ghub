package org

import (
	"github.com/codegangsta/cli"
)

var Create cli.Command = cli.Command{
	Name:  "org",
	Usage: "Create an org (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("org created")
	},
}

var Delete cli.Command = cli.Command{
	Name:  "org",
	Usage: "Delete an org (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("org deleted")
	},
}

var Get cli.Command = cli.Command{
	Name:  "org",
	Usage: "Get an org",
	Subcommands: []cli.Command{
		GetOrgRepos,
	},
	Action: func(c *cli.Context) {
		getOrg(c)
	},
}

var Update cli.Command = cli.Command{
	Name:  "org",
	Usage: "Update an org (NOT IMPLEMENTED)",
	Action: func(c *cli.Context) {
		println("org updated")
	},
}

// Repos subcommands
var GetOrgRepos cli.Command = cli.Command{
	Name:  "repos",
	Usage: "Get all repos belonging to org",
	Action: func(c *cli.Context) {
		getReposInOrg(c)
	},
}

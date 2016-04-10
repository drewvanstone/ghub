package org

import (
	"github.com/codegangsta/cli"
	"github.com/ghub/util"
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
		GetOrgReposCmd,
		GetOrgTeamsCmd,
		GetOrgMembersCmd,
		GetOrgHooksCmd,
	},
	Action: func(c *cli.Context) {
		util.CheckCommandArgs(1, c)
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
var GetOrgReposCmd cli.Command = cli.Command{
	Name:  "repos",
	Usage: "Get all repos belonging to an organization",
	Action: func(c *cli.Context) {
		util.CheckCommandArgs(1, c)
		getOrgRepos(c)
	},
}

var GetOrgTeamsCmd cli.Command = cli.Command{
	Name:  "teams",
	Usage: "Get all teams belonging to an organization",
	Action: func(c *cli.Context) {
		util.CheckCommandArgs(1, c)
		getOrgTeams(c)
	},
}

var GetOrgMembersCmd cli.Command = cli.Command{
	Name:  "members",
	Usage: "Get all members belonging to an organization",
	Action: func(c *cli.Context) {
		util.CheckCommandArgs(1, c)
		getOrgMembers(c)
	},
}

var GetOrgHooksCmd cli.Command = cli.Command{
	Name:  "hooks",
	Usage: "Get all hooks of an organization",
	Action: func(c *cli.Context) {
		util.CheckCommandArgs(1, c)
		getOrgHooks(c)
	},
}

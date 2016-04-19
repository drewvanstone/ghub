package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/ghub/gh"
	"github.com/ghub/issue"
	"github.com/ghub/org"
	"github.com/ghub/pr"
	"github.com/ghub/repo"
	"github.com/ghub/team"
	"github.com/ghub/user"
	//"github.com/ghub/util"
	"os"
)

func main() {

	gh.Init()

	// Autocomplete commands
	//var topCommands = []string{"create", "delete", "get", "update"}
	var getCommands = []string{"issue", "org", "pr", "repo", "team", "user"}

	ghub := cli.NewApp()
	ghub.Name = "Ghub"
	ghub.Usage = "Command line utility for GitHub"
	ghub.Email = "drewvanstone@gmail.com"
	ghub.Author = "Drew Flower"
	ghub.EnableBashCompletion = true
	ghub.Commands = []cli.Command{
		{
			Name:  "create",
			Usage: "Create [COMMAND] (user, repo, pr, etc)",
			Subcommands: []cli.Command{
				issue.Create,
				org.Create,
				pr.Create,
				repo.Create,
				team.Create,
				user.Create,
			},
		},
		{
			Name:  "delete",
			Usage: "Delete [COMMAND] (user, repo, pr, etc)",
			Subcommands: []cli.Command{
				issue.Delete,
				org.Delete,
				pr.Delete,
				repo.Delete,
				team.Delete,
				user.Delete,
			},
		},
		{
			Name:  "get",
			Usage: "Get    [COMMAND] (user, repo, pr, etc)",
			Subcommands: []cli.Command{
				issue.Get,
				org.Get,
				pr.Get,
				repo.Get,
				team.Get,
				user.Get,
			},
			BashComplete: func(c *cli.Context) {
				if c.NArg() > 0 {
					return
				}
				for _, c := range getCommands {
					fmt.Println(c)
				}
			},
		},
		{
			Name:  "update",
			Usage: "Update [COMMAND] (user, repo, pr, etc)",
			Subcommands: []cli.Command{
				issue.Update,
				org.Update,
				pr.Update,
				repo.Update,
				team.Update,
				user.Update,
			},
		},
	}

	ghub.Run(os.Args)
}

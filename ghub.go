package main

import (
	"github.com/codegangsta/cli"
	"github.com/ghub/issue"
	"github.com/ghub/org"
	"github.com/ghub/pr"
	"github.com/ghub/repo"
	"github.com/ghub/team"
	"github.com/ghub/user"
	"os"
)

func main() {

	ghub := cli.NewApp()
	ghub.Name = "Ghub"
	ghub.Usage = "Command line utility for GitHub"
	ghub.Email = "drewvanstone@gmail.com"
	ghub.Author = "Drew Flower"
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

package util

import (
	"github.com/codegangsta/cli"
	"log"
	"os"
	"strings"
)

func Check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Helper function to check that there are at least as many
// arguments as required, and if not, display subcommand help.
func CheckCommandArgs(minArgs int, c *cli.Context) {
	if len(c.Args()) < minArgs {
		cli.ShowSubcommandHelp(c)
		os.Exit(1)
	}
}

func SplitOrgRepoName(orgrepo string) (string, string) {
	parts := strings.Split(orgrepo, "/")
	if len(parts) != 2 {
		println("Repository must be in the form: org/reponame (ie. drewvanstone/ghub), got ", orgrepo)
	}
	return parts[0], parts[1]
}

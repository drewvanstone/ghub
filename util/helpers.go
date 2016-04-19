package util

import (
	"errors"
	"fmt"
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

func Autocomplete(commands []string) func(*cli.Context) {
	// This will complete if no args are passed
	return func(c *cli.Context) {
		if c.NArg() > 0 {
			return
		}
		for _, command := range commands {
			fmt.Println(command)
		}
	}
}

func CheckAutocomplete(commands []string, c *cli.Context) {
	if c.NArg() > 0 {
		return
	}
	for _, command := range commands {
		fmt.Println(command)
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

func SplitOrgRepoName(orgrepo string) (string, string, error) {
	parts := strings.Split(orgrepo, "/")
	if len(parts) != 2 {
		return "", "", errors.New("Failed to split org/repo")
	}
	return parts[0], parts[1], nil
}

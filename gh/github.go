package gh

import (
	"fmt"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"net/http"
	"os"
)

var Client *github.Client

func Init() {
	token := getToken()
	Client = gitHub(token)
}

func getToken() *http.Client {
	token := os.Getenv("GITHUB_TOKEN")
	if len(token) == 0 {
		fmt.Fprintln(os.Stderr, "[ghub] You must set the GITHUB_TOKEN environment variable.")
		os.Exit(1)
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return tc
}

func gitHub(tc *http.Client) *github.Client {
	client := github.NewClient(tc)
	return client
}

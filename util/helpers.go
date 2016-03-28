package util

import (
	"strings"
)

func SplitOrgRepoName(orgrepo string) (string, string) {
	parts := strings.Split(orgrepo, "/")
	if len(parts) != 2 {
		println("Repository must be in the form: org/reponame (ie. drewvanstone/ghub), got ", orgrepo)
	}
	return parts[0], parts[1]
}

package github

import (
	"errors"
	"fmt"

	"github.com/google/go-github/v47/github"
)

const (
	pullPermission     string = "pull"
	triagePermission   string = "triage"
	pushPermission     string = "push"
	maintainPermission string = "maintain"
	adminPermission    string = "admin"
	writePermission    string = "write"
	readPermission     string = "read"
)

func getRepoPermission(p map[string]bool) (string, error) {

	// Permissions are returned in this map format such that if you have a certain level
	// of permission, all levels below are also true. For example, if a team has push
	// permission, the map will be: {"pull": true, "push": true, "admin": false}
	if (p)[adminPermission] {
		return adminPermission, nil
	} else if (p)[maintainPermission] {
		return maintainPermission, nil
	} else if (p)[pushPermission] {
		return pushPermission, nil
	} else if (p)[triagePermission] {
		return triagePermission, nil
	} else {
		if (p)[pullPermission] {
			return pullPermission, nil
		}
		return "", errors.New("at least one permission expected from permissions map")
	}
}

func getInvitationPermission(i *github.RepositoryInvitation) (string, error) {
	// Permissions for some GitHub API routes are expressed as "read",
	// "write", and "admin"; in other places, they are expressed as "pull",
	// "push", and "admin".
	permissions := i.GetPermissions()
	if permissions == readPermission {
		return pullPermission, nil
	} else if permissions == writePermission {
		return pushPermission, nil
	} else if permissions == adminPermission {
		return adminPermission, nil
	} else if *i.Permissions == maintainPermission {
		return maintainPermission, nil
	} else if *i.Permissions == triagePermission {
		return triagePermission, nil
	}

	return "", fmt.Errorf("unexpected permission value: %v", permissions)
}

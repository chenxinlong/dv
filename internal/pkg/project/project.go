package project

import "fmt"

var (
	GITHASH   string
	GITBRANCH string
	BUILDTIME string
	GOVERSION string
	PROJDIR   string // project dir
)

func PrintlnBuildInfo() {
	fmt.Printf(`hash = %s, branch = %s, build time = %s, go version = %s, project dir = %s`, GITHASH, GITBRANCH, BUILDTIME, GOVERSION, PROJDIR)
}

package main

import (
	"os"

	"github.com/chenxinlong/dv/internal/app/dv"
	"github.com/chenxinlong/dv/internal/pkg/project"
)

func main() {
	if len(os.Args) == 2 && os.Args[1] == "version" {
		project.PrintlnBuildInfo()
		return
	}

	dv.Run()
}

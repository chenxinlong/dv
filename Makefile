.PHONY:build_darwin all cp

dir=$(shell pwd)
gomodule=`head -n 1 go.mod | cut -d ' ' -f 2`
buildbranch=`git symbolic-ref --short -q HEAD`
buildcommit=`git rev-parse HEAD | cut -c1-8`
buildtime=`date '+%Y%m%d.%H%M%S'`
buildgoversion=`go version|sed -e 's/go version //g'`
BuildFlags=-ldflags "-extldflags=-static -X '$(gomodule)/internal/pkg/project.PROJDIR=$(dir)' -X '$(gomodule)/internal/pkg/project.GITBRANCH=$(buildbranch)' -X $(gomodule)/internal/pkg/project.GITHASH=$(buildcommit) -X $(gomodule)/internal/pkg/project.BUILDTIME=$(buildtime) -X '$(gomodule)/internal/pkg/project.GOVERSION=$(buildgoversion)'"

all: build_darwin cp

build_darwin:
	@GOOS=darwin GOARCH=amd64 go build $(BuildFlags) -o ./build/dv ./cmd/dv/
cp:
	cp ./build/dv /usr/local/bin/dv

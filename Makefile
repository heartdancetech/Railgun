SHELL := /bin/bash
BASEDIR = $(shell pwd)

APP = "LastOrder"
BuildDIR = build
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0 | sed 's/v//g'; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)
versionDir = "github.com/MisakaSystem/LastOrder/cmd"
ldflags= "-X ${versionDir}.gitTag=${gitTag} \
-X ${versionDir}.buildDate=${buildDate} \
-X ${versionDir}.gitCommit=${gitCommit} \
-X ${versionDir}.gitTreeState=${gitTreeState}"

all: go-tool release
	@ls -al build/

clean:
	@rm -rvf build/
	@docker image prune
	@docker rmi --force ${APP}-${gitTag}:${gitTag}

gotool:
	gofmt -w .
	@for va in $(VETPACKAGES); do \
		go vet $$va; \
	done

release:
	# Build for linux
	go clean
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ${BuildDIR}/${APP}-${gitTag}-linux64-amd64
	# Build for win
	go clean
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -v -o ${BuildDIR}/${APP}-${gitTag}-windows-amd64.exe
	# Build for mac
	go clean
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o ${BuildDIR}/${APP}-${gitTag}-darwin-amd64

build-docker:
	docker build --build-arg BASE_BIN_NAME=${APP}-${gitTag}-linux64-amd64 -t ${APP}-${gitTag}:${gitTag} .


help:
	@echo "make - compile the source code to binary"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make release - build gateway binary"

.PHONY: clean go-tool help
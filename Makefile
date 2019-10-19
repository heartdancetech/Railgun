SHELL := /bin/bash
BASEDIR = $(shell pwd)

# build with verison infos
gatewayName = "LastOrder"
# versionDir = "dipole-gateway"
gitTag = $(shell if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ];then git describe --tags --abbrev=0 | sed 's/v//g'; else git log --pretty=format:'%h' -n 1; fi)
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
gitTreeState = $(shell if git status|grep -q 'clean';then echo clean; else echo dirty; fi)

ldflags="-w -X ${versionDir}.gitTag=${gitTag} -X ${versionDir}.buildDate=${buildDate} -X ${versionDir}.gitCommit=${gitCommit} -X ${versionDir}.gitTreeState=${gitTreeState}"

all: export GOOS=linux
all: export GOARCH=amd64
all: go-tool build-gateway

clean:
	rm -f ${gatewayName}
# 	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {}

go-tool:
	gofmt -w .
# 	go vet -v $(go list ./...| grep -v /vendor/)

build:
	@go build -v -ldflags ${ldflags}  -o ${gatewayName} main.go

build-gateway-docker:
	docker build -t ${gatewayName}:${gitTag} .
	rm -f ${gatewayName}


help:
	@echo "make - compile the source code to docker image"
	@echo "make clean - remove binary file and vim swp files"
	@echo "make gotool - run go tool 'fmt' and 'vet'"
	@echo "make build - build gateway binary"

.PHONY: clean go tool help
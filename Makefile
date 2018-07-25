default: help

help:   ## show this help
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sed 's/:.*##/#/' | column -t -c 2 -s '#'

clean-all:	## Go Clean
	go clean

test:	## Run Short tests
	go test -v ./... -short

build: ## Run dep ensure and build linux binary of all individual functions
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/auth functions/auth/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/privateFunc functions/privateFunc/main.go
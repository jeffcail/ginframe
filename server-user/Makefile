APP=server-user

.PHONY: help all build windows linux darwin

help:
	@echo "usage: make <option>"
	@echo "options and effects:"
	@echo "    help   : Show help"
	@echo "    all    : Build multiple binary of this project"
	@echo "    build  : Build the binary of this project for current platform"
	@echo "    windows: Build the windows binary of this project"
	@echo "    linux  : Build the linux binary of this project"
	@echo "    darwin : Build the darwin binary of this project"

all:build windows linux darwin

build:
	@go build -o ${APP}

windows:
	@GOOS=windows go build -o ${APP}-windows
linux:
	@GOOS=linux go build -o ${APP}-linux
darwin:
	@GOOS=darwin go build -o ${APP}-darwin
cnf ?= config.env
include $(cnf)
export $(shell sed 's/=.*//' $(cnf))

PROJECT=$(shell basename `pwd`)
GOBASE=$(shell pwd)
GOPATH="$(GOBASE)/main:$(GOBASE)"
GOBIN=$(GOBASE)/bin
GOFILES=$(wildcard *.go)
GOARCH=amd64
TAG=$(shell git rev-parse --short HEAD)
STDERR=/tmp/.$(PROJECT)-stderr.txt

## build: Compile the binary
build:
	@rm -f $(STDERR)
	@touch -f $(STDERR)
	@$(MAKE) -s go-build

## test: Run unit tests
.PHONY: test
test: go-test

## clean: Clean cache build files
.SILENT: clean
.PHONY: clean
clean: go-clean
	@rm -f $(GOBIN)/*

## docker: Build Docker container
.PHONY: docker
docker:
	@docker build -t $(PROJECT) .

## run: Execute application in container
.PHONY: run
run:
	@docker run -it --rm --env-file=./config.env -p=$(PORT):$(PORT) --name="$(PROJECT)" $(PROJECT)

go-compile: go-clean go-get go-build

go-build:
	@GOPATH=$(GOPATH) go build -o $(GOBIN)/$(PROJECT) $(GOBASE)/main/app.go
	@cp $(GOBASE)/data/*.txt $(GOBIN)/
	@chmod +x $(GOBIN)/$(PROJECT)

go-clean:
	@GOPATH=$(GOPATH) go clean

go-test:
	@GOPATH=$(GOPATH) go test -v ./...

.PHONY: help
all: help
help: Makefile
	@echo
	@echo " Command to run "$(PROJECT)":"
	@echo
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
	@echo
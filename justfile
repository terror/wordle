set dotenv-load

export EDITOR := 'nvim'

files := 'src/*'

alias f := fmt
alias r := run

default:
  just --list

all: lint forbid fmt-check

run:
	go run {{files}}

fmt:
	gofmt -w {{files}}
	just retab

fmt-check:
	gofmt -l .
	@echo formatting check done

forbid:
	./bin/forbid

lint:
  golangci-lint run {{files}}

retab:
	./bin/retab

dev-deps:
	brew install golangci-lint

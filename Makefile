GOVERSION=$(shell go version)
GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

# if tag doesnt exists, show revision
VERSION=$(shell git describe --tags --abbrev=7 --always)
REVISION=$(shell git rev-parse --short HEAD)

ifeq ($(VERSION),$(REVISION))
VERSION=v0.0.0
endif

# ANSI color
RED=\033[31m
GREEN=\033[32m
RESET=\033[0m

COLORIZE_PASS=sed ''/PASS/s//$$(printf "$(GREEN)PASS$(RESET)")/''
COLORIZE_FAIL=sed ''/FAIL/s//$$(printf "$(RED)FAIL$(RESET)")/''

NAME=exaauthd

.PHONY: \
	runner-test \
	install \
	build \
	runner \
	dep-clean

runner-test:
	GO111MODULE=on go test -v ./... | $(COLORIZE_PASS) | $(COLORIZE_FAIL)

runner:
	realize start

clean-dep:
	GO111MODULE=on go mod tidy

build:
	GO111MODULE=on go build -v -ldflags "-X github.com/smith-30/exaauth/cmd.version=$(VERSION) -X github.com/smith-30/exaauth/cmd.revision=$(REVISION) -X github.com/smith-30/exaauth/cmd.appName=$(NAME)" -o build/${GOOS}_${GOARCH}/${NAME} main.go

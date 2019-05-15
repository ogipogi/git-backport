# Go parameters
GOCMD=go
GOFMT=gofmt
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=git-backport
BINARY_UNIX=$(BINARY_NAME)_unix

all: test clean fmt build
build:
	$(GOBUILD) -o $(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
fmt:
	$(GOFMT) -s -w main.go
	$(GOFMT) -s -w backport_test.go

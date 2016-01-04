CGO_ENABLED=0
GOOS=linux
GOARCH=amd64
COMMIT=`git rev-parse --short HEAD`
APP=meld-agent
REPO?=factionlabs/$(APP)
TAG?=latest
export GO15VENDOREXPERIMENT=1

ifeq ($(GOOS), windows)
    BIN_NAME=meld-agent.exe
else
    BIN_NAME=meld-agent
endif

all: build

add-deps:
	@GO15VENDOREXPERIMENT=1 godep save -t ./...

build:
	@cd cmd/$(APP) && GO15VENDOREXPERIMENT=1 go build -ldflags "-w -X github.com/$(REPO)/version.GitCommit=$(COMMIT)" -o $(BIN_NAME) .

build-static:
	@cd cmd/$(APP) && GO15VENDOREXPERIMENT=1 go build -a -tags "netgo static_build" -installsuffix netgo -ldflags "-w -X github.com/$(REPO)/version.GitCommit=$(COMMIT)" -o $(BIN_NAME) .

test: build
	@GO15VENDOREXPERIMENT=1 go test -v ./...

clean:
	@rm -rf cmd/$(APP)/$(BIN_NAME)

.PHONY: add-deps build build-static test clean

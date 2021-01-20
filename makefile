##################
#   Variables    #
##################

GIT_VERSION = $(shell git rev-list -1 HEAD)

ifdef RELEASE
	ABSTRAKT_VERSION := $(RELEASE)
else
	ABSTRAKT_VERSION := edge
endif

ifdef ARCHIVE_OUTDIR
	OUTPUT_PATH := $(ARCHIVE_OUTDIR)
else
	OUTPUT_PATH := .
endif

LOCAL_OS := $(shell uname)
ifeq ($(LOCAL_OS),Linux)
   TARGET_OS_LOCAL = linux
else ifeq ($(LOCAL_OS),Darwin)
   TARGET_OS_LOCAL = darwin
else
   TARGET_OS_LOCAL ?= windows
endif
export GOOS ?= $(TARGET_OS_LOCAL)

ifeq ($(GOOS),windows)
	BINARY_EXT_LOCAL:=.exe
	GOLANGCI_LINT:=golangci-lint.exe
	export ARCHIVE_EXT = .zip
else
	BINARY_EXT_LOCAL:=
	GOLANGCI_LINT:=golangci-lint
	export ARCHIVE_EXT = .tar.gz
endif

export BINARY_EXT ?= $(BINARY_EXT_LOCAL)

##################
# Linting/Verify #
##################

#wire:
	#@echo "Running wire build"
	#wire ./internal/serviceLocator

test-watcher:
	@echo "Running test watcher"
	bash ./scripts/test_watcher.sh

lint-all: lint-prepare build lint vet

lint-prepare:
ifeq (,$(shell which golangci-lint))
	@echo "Installing golangci-lint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.23.8 > /dev/null 2>&1
	golangci-lint --version
else
	@echo "golangci-lint is installed"
endif

lint:
	@echo "Linting"
	golangci-lint run ./...

vet:
	@echo "Vetting"
	go vet ./...

##################
#    Testing     #
##################

test-prepare:
	go get github.com/stretchr/testify
	go get github.com/pmezard/go-difflib
	go get github.com/jstemmer/go-junit-report
	go get github.com/AlekSi/gocov-xml
	go get github.com/axw/gocov/gocov
	go get github.com/matm/gocov-html

test:
	go test -v ./... -cover -coverprofile=coverage.txt -race -covermode=atomic

test-export:
	go test -v ./... -cover -coverprofile=coverage.txt -race -covermode=atomic 2>&1 | $(GOPATH)/bin/go-junit-report > report.xml
	$(GOPATH)/bin/gocov convert coverage.txt > coverage.json
	$(GOPATH)/bin/gocov-xml < coverage.json > coverage.xml
	mkdir coverage | true
	$(GOPATH)/bin/gocov-html < coverage.json > coverage/index.html

test-all: test-prepare test

test-export-all: test-prepare test-export

##################
#     Build      #
##################

BASE_PACKAGE_NAME := github.com/microsoft/abstrakt
LDFLAGS:=-X $(BASE_PACKAGE_NAME)/cmd.commit=$(GIT_VERSION) -X $(BASE_PACKAGE_NAME)/cmd.version=$(ABSTRAKT_VERSION)

build::
	GOOS=$(GOOS) GOARCH=amd64 go build -ldflags="$(LDFLAGS)" -o abstrakt$(BINARY_EXT) main.go

##################
#    Release     #
##################

archive:
ifeq ("$(wildcard $(OUTPUT_PATH))", "")
	mkdir -p $(OUTPUT_PATH)
endif

ifeq ($(GOOS),windows)
	zip $(OUTPUT_PATH)/abstrakt_$(GOOS)_amd64$(ARCHIVE_EXT) abstrakt$(BINARY_EXT)
else
	tar -czvf "$(OUTPUT_PATH)/abstrakt_$(GOOS)_amd64$(ARCHIVE_EXT)" "abstrakt$(BINARY_EXT)"
endif

release: build archive generate-checksum

##################
#     Verify     #
##################

generate-checksum:
	cd $(OUTPUT_PATH)
	sha256sum abstrakt_$(GOOS)_amd64$(ARCHIVE_EXT) >> checksums.sha256

verify-checksum:
	sha256sum -c $(OUTPUT_PATH)/checksums.sha256

##################
#  Run Examples  #
##################

fmt:
	gofmt -s -w ./

visualise: build
	./abstrakt visualise -f ./examples/constellation/http_constellation.yaml | dot -Tpng > result.png

diff: build
	./abstrakt diff -o ./examples/constellation/sample_constellation.yaml -n ./examples/constellation/sample_constellation_changed.yaml | dot -Tpng > result.png

run-http-demo: http-demo http-demo-deploy

http-demo: build
	./abstrakt compose http-demo -f ./examples/constellation/http_constellation.yaml -m ./examples/constellation/http_constellation_maps.yaml -o ./output/http_sample

http-demo-deploy:
	helm install wormhole-http-demo ./output/http_sample/http-demo

http-demo-template:
	helm template wormhole-http-demo ./output/http_sample/http-demo

http-demo-delete:
	helm delete wormhole-http-demo

http-demo-template-all: http-demo http-demo-template

http-demo-deploy-all: http-demo-delete http-demo-deploy

update-golden-data:
	go test ./internal/chartservice -update -run TestUpdate

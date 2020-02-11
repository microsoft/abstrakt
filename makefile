##################
# Linting/Verify #
##################

#wire:
	#@echo "Running wire build"
	#wire ./internal/serviceLocator

test-watcher:
	@echo "Running test watcher"
	bash ./scripts/test_watcher.sh

lint-all: lint-prepare lint vet

lint-prepare:
ifeq (,$(shell which golangci-lint))
	@echo "Installing golangci-lint"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.23.1 > /dev/null 2>&1
else
	@echo "golangci-lint is installed"
endif

lint: build
	@echo "Linting"
	golangci-lint run ./...

vet:
	@echo "Vetting"
	go vet ./...

##################
#    Testing     #
##################

test-prepare:
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
#  Run Examples  #
##################

fmt:
	gofmt -s -w ./

build::
	go build -o abstrakt main.go

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

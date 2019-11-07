##################
# Linting/Verify #
##################

lint-all: lint-prepare lint vet

lint-prepare: 
	@echo "Installing golangci-lint"
	wget -O - -q https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.21.0

lint:
	@echo "Linting"
	./bin/golangci-lint run ./...

vet:
	@echo "Vetting"
	go vet ./...

##################
# Testing		 #
##################

test-prepare: 
	go get github.com/jstemmer/go-junit-report
	go get github.com/axw/gocov/gocov
	go get github.com/AlekSi/gocov-xml
	go get github.com/axw/gocov/gocov
	go get github.com/matm/gocov-html

test: 
	@go test -v ./... -cover -coverprofile=coverage.txt -race -covermode=atomic

test-export: 
	@go test -v ./... -cover -coverprofile=coverage.txt -race -covermode=atomic 2>&1 | $GOPATH/bin/go-junit-report > report.xml
	gocov convert coverage.txt > coverage.json
	gocov-xml < coverage.json > coverage.xml
	mkdir coverage | true
	gocov-html < coverage.json > coverage/index.html

test-all: test-prepare test

test-export-all: test-prepare test-export

##################
# Linting/Verify #
##################

#wire:
	#@echo "Running wire build"
	#wire ./internal/serviceLocator

test-watcher:
	@echo "Running test watcher"
	bash ./.scripts/test_watcher.sh

lint-all: lint-prepare lint vet

lint-prepare:
ifeq (,$(shell which golangci-lint))
	@echo "Installing golangci-lint"
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.21.0 > /dev/null 2>&1
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
# Testing		 #
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

update-golden-data:
	go test ./internal/chartservice -update -run TestUpdate

install-helm:
	curl https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3 > get_helm.sh
	chmod 700 get_helm.sh
	./get_helm.sh

create-kindcluster:
ifeq (,$(shell kind get clusters))
	@echo "no kind cluster"
else
	@echo "kind cluster is running, deleteing the current cluster"
	kind delete cluster 
endif
	@echo "creating kind cluster"
	kind create cluster

set-kindcluster: install-kind
ifeq (${shell kind get kubeconfig-path --name="kind"},${KUBECONFIG})
	@echo "kubeconfig-path points to kind path"
else
	@echo "please run below command in your shell and then re-run make set-kindcluster"
	@echo  "\e[31mexport KUBECONFIG=$(shell kind get kubeconfig-path --name="kind")\e[0m"
	@exit 111
endif
	make create-kindcluster
	kubectl apply -f /workspace/.scripts/rbac.yaml

install-kind:
ifeq (,$(shell which kind))
	@echo "installing kind"
	GO111MODULE="on" go get sigs.k8s.io/kind@v0.4.0
else
	@echo "kind has been installed"
endif

##################
#  Run Examples    		  #
##################

fmt:
	gofmt -s -w ./

build:	
	go build -o abstrakt main.go

visualise: build	
	./abstrakt visualise -f ./sample/constellation/http_constellation.yaml | dot -Tpng > result.png

diff: build
	./abstrakt diff -o ./sample/constellation/sample_constellation.yaml -n ./sample/constellation/sample_constellation_changed.yaml | dot -Tpng > result.png

run-http-demo: http-demo http-demo-deploy

http-demo: build	
	./abstrakt compose http-demo -f ./sample/constellation/http_constellation.yaml -m ./sample/constellation/http_constellation_maps.yaml -o ./output/http_sample

http-demo-deploy:
	helm install wormhole-http-demo ./output/http_sample/http-demo

http-demo-template:
	helm template wormhole-http-demo ./output/http_sample/http-demo

http-demo-delete:
	helm delete wormhole-http-demo

http-demo-template-all: http-demo http-demo-template

http-demo-deploy-all: http-demo-delete http-demo-deploy

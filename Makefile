
CODEGEN_PKG=${GOPATH}/src/k8s.io/code-generator

deps:
	go get -u golang.org/x/lint/golint
	go get -u github.com/kisielk/errcheck
	go get -u github.com/maxbrunsfeld/counterfeiter
	go get -u github.com/onsi/ginkgo/ginkgo
	go get -u golang.org/x/tools/cmd/goimports
	go get -u github.com/actgardner/gogen-avro/gogen-avro

precommit: ensure generate test check addlicense
	@echo "ready to commit"

ensure:
	GO111MODULE=on go mod tidy
	GO111MODULE=on go mod vendor

generate:
	go get github.com/maxbrunsfeld/counterfeiter
	go get github.com/actgardner/gogen-avro/gogen-avro
	rm -rf mocks avro
	go generate ./...

	go get -d k8s.io/code-generator/...
	${GOPATH}/src/k8s.io/code-generator/generate-groups.sh all \
	github.com/bborbe/kafka-k8s-topic-controller/client github.com/bborbe/kafka-k8s-topic-controller/apis \
	example.com:v1

test:
	go test -cover -race $(shell go list ./... | grep -v /vendor/)

check: format lint vet errcheck

format:
	@go get golang.org/x/tools/cmd/goimports
	@find . -type f -name '*.go' -not -path './vendor/*' -exec gofmt -w "{}" +
	@find . -type f -name '*.go' -not -path './vendor/*' -exec goimports -w "{}" +

lint:
	@go get golang.org/x/lint/golint
	@golint -min_confidence 1 $(shell go list ./... | grep -v /vendor/)

vet:
	@go vet $(shell go list ./... | grep -v /vendor/)

errcheck:
	@go get github.com/kisielk/errcheck
	@errcheck -ignore '(Close|Write|Fprint)' $(shell go list ./... | grep -v /vendor/)

addlicense:
	@go get github.com/google/addlicense
	@addlicense -c "Benjamin Borbe" -y 2019 -l bsd ./*.go


GO_VERSION := 1.19.5

setup:
	install-go
	init-go
	install-lint

install-go:
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

init-go:
	echo 'export PATH=$$PATH:/usr/local/go/bin' >> $$(HOME)/.bashrc
	echo 'export PATH=$$PATH:##(HOME)/go/bin' >> $$(HOME)/.bashrc

upgrade-go:
	sudo rm -rf /usr/bin/go
	wget "https://golang.org/dl/go$(GO_VERSION).linux-amd64.tar.gz"
	sudo tar -C /usr/local -xzf go$(GO_VERSION).linux-amd64.tar.gz
	rm go$(GO_VERSION).linux-amd64.tar.gz

install-lint:
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.50.1

build:
	go build -o api cmd/web/main.go

test :
	go test ./... -coverprofile=coverage.out

coverage:
	go tool cover -func coverage.out | grep "total:" | \
	awk '{print ((int($$3) > 80) != 1) }' 

report:
	go tool cover -html=coverage.out -o coverage.html

check-format:
	test -z $$(go fmt ./...)

static-check:
	golangci-lint run

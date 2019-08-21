PWD := $(shell pwd 2>/dev/null)
PKG_LIST := $(shell go list ./... | grep -v /vendor/ 2>/dev/null)

.PHONY: model
model:
	CGO_ENABLED=0 go build -o build/model-builder ./tools/model-builder/main.go
	./build/model-builder -path $(PWD) -p 12345678

.PHONY: lint
lint:
	golint -set_exit_status $(PKG_LIST)

.PHONY: gofmt
gofmt:
	go fmt $(PKG_LIST)

.PHONY: govet
govet:
	go vet $(PKG_LIST)

.PHONY: gotest
gotest:
	go test -race -short $(PKG_LIST)

.PHONY: coverage
coverage:
	sh ./coverage.sh

.PHONY: server
server:
	go build -race -ldflags "-extldflags '-static'" -o ./build/admin-server ./server/server.go




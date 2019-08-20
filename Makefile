PWD := $(shell pwd 2>/dev/null)

.PHONY: model
model:
	CGO_ENABLED=0 go build -o build/model-builder ./tools/model-builder/main.go
	./build/model-builder -path $(PWD) -p 12345678

.PHONY: server
server:
	go build -race -ldflags "-extldflags '-static'" -o ./build/admin-server ./server/server.go
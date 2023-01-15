CGO_ENABLED=1
GOOS=linux
GOARCH=amd64

.PHONY: build
build:
	GOOS=$(GOOS) CGO_ENABLED=$(CGO_ENABLED) GOARCH=$(GOARCH) \
	go build -o main cmd/main.go
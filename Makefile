.PHONY:

build:
	CGO_ENABLED=0 go build -o ./bin ./cmd/app/main.go

test:
	go test ./... -v
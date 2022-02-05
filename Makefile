.PHONY:

build:
	go mod tidy
	CGO_ENABLED=0 go build -o ./bin ./cmd/app/main.go

test:
	go mod tidy
	go test ./... -v
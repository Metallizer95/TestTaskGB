.PHONY:

build:
	go mod tidy
	CGO_ENABLED=0 go build -o ./bin ./cmd/app/main.go

run:
	go run cmd/app/main.go -path=ethConfig.yml

test:
	go mod tidy
	go test ./... -v
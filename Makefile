build:
	@go build -o bin/citizen ./main.go

run: build
	@./bin/citizen --config=./local.yaml

test:
	@go test -v ./...

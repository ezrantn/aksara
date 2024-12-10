build: 
	@go build -o bin/aksara examples/main.go

run: build
	@./bin/aksara

test:
	@go test -v ./...
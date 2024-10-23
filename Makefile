build:
	 @go build -o bin/app ./cmd/

run: build
	@./bin/app

# run tests with verbose output, with coverage
test:
	@go test -v ./... -cover

# run tests with output of coverage and show on browser 
test-preview:
	@go test ./filename/ -coverprofile=coverage.out 
	@go tool cover -html=coverage.out


# Phony targets to avoid conflicts with files or directories
.PHONY: run test

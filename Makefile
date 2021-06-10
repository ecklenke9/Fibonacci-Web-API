docker:
	@docker-compose up

local:
	@go run main.go

test:
	@go test -v ./...

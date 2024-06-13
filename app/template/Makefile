run:
	@go run main.go

service:
	@docker-compose up -d

service_stop:
	@docker-compose down

test:
	@go test ./... -v

.PHONY: run service service_stop test

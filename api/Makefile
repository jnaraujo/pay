# Build target
BINARY_NAME = app
BUILD_DIR = ./cmd/api/main.go

build:
	go build -o $(BINARY_NAME) $(BUILD_DIR)

run:
	air -c .air.toml

migration:
	migrate create -ext sql -dir ./migrations/ $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	go run ./cmd/migrate/main.go up

migrate-down:
	go run ./cmd/migrate/main.go down
install:
	go mod download
	go get github.com/google/wire/cmd/wire@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest

tidy:
	go mod tidy

swagger_gen:
	swag fmt
	swag init -q -g ./pkg/web/api/server/server.go -o ./pkg/web/api/api-docs

build: tidy swagger_gen
	wire cmd/main/wire.go

dev: build
	go run ./cmd/main .

lint:
	gosec ./...

run: dev
	echo "Running the project..."


install:
	go mod download

tidy:
	go mod tidy

swagger_gen:
	swag fmt
	swag init -q -g ./pkg/web/api/app.go -o ./pkg/web/api/api-docs

build: tidy swagger_gen
	wire cmd/main/wire.go

dev: build
	go run ./cmd/main .

run: dev
	echo "Running the project..."


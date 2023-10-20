install:
	go mod download
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.0

install_prod:
	go mod download

tidy:
	go mod tidy

swagger_gen:
	swag fmt
	swag init -q -g ./pkg/web/api/server/server.go -o ./pkg/web/api/api-docs

build_dev: tidy swagger_gen
	wire cmd/main/wire/wire.go

build_prod: install_prod
	CGO_ENABLED=0 GOOS=linux go build -o ./build/clean-architechture-api-go ./cmd/main/main.go

dev: build
	go run ./cmd/main .

prod: build_prod
	/build/clean-architechture-api-go

lint:
	gosec ./...

docker_build:
	docker build -t clean-architechture-api-go:latest .

docker_run:
	docker run -p 3002:3002 clean-architechture-api-go:latest

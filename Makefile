install:
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.0
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/securego/gosec/v2/cmd/gosec@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install golang.org/x/tools/cmd/goimports@latest
	go mod download


install_prod:
	go mod download

tidy:
	go mod tidy

generate_mocks:
	go run cmd/mocks/generate.go

swagger_gen:
	swag fmt
	swag init -q -g ./pkg/web/api/server/server.go -o ./pkg/web/api/api-docs

build_dev: tidy swagger_gen
	wire cmd/main/wire/wire.go
	mockgen -source=pkg/application/users/queries/get-user/get-user-by-id.go -destination=mocks/main.go

build_prod: install_prod
	CGO_ENABLED=0 GOOS=linux go build -o ./build/clean-architechture-api-go ./cmd/main/main.go

test: generate_mocks
	go test -coverpkg=./pkg/... ./pkg/... -v -covermode=count

dev: build_dev
	go run ./cmd/main .

prod: build_prod
	/build/clean-architechture-api-go

lint:
	gosec ./...

docker_build:
	docker build -t clean-architechture-api-go:latest .

docker_run:
	docker run -p 3002:3002 clean-architechture-api-go:latest

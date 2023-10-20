FROM golang:1.21 as base

RUN apt-get update && apt-get install -y make bash
WORKDIR /app

FROM base as copy-files

COPY cmd cmd
COPY pkg pkg
COPY go.mod ./
COPY go.sum ./
COPY build build
COPY .env ./
COPY mock-user-data.json ./

ENV GOPATH /app/go

FROM copy-files as run

EXPOSE 3002

CMD ["/app/build/clean-architechture-api-go"]

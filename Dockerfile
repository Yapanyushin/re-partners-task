FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./cmd/pack_calculator/bin/pack_calculator/pack_calculator -ldflags "-s -w" ./cmd/pack_calculator/


EXPOSE ${SERVER_PORT}

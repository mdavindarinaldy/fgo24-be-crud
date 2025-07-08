FROM golang:alpine AS build
WORKDIR /buildapp

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g main.go
RUN go build -o gocrud main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=build /buildapp/gocrud /app/

ENTRYPOINT [ "/app/gocrud" ]

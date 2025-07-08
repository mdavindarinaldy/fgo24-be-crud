FROM golang:alpine AS build
WORKDIR /buildapp

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

RUN go build -o gocrud main.go

FROM alpine:3.22

WORKDIR /app

COPY --from=build /buildapp/gocrud /app/
COPY .env /app/.env

ENTRYPOINT [ "/app/gocrud" ]

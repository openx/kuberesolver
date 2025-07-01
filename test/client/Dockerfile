FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o client ./test/client/main.go

RUN chmod +x ./client

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/client .

CMD ["./client"]

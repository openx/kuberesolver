FROM golang:1.22-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o server ./test/server/main.go

RUN chmod +x ./server

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/server ./

EXPOSE 50051

CMD ["./server"]
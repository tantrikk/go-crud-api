FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o go-crud-api ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/go-crud-api .

CMD ["./go-crud-api"]
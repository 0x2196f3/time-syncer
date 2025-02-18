FROM golang:alpine AS builder

WORKDIR /app

COPY go.mod main.go ./
RUN go mod download && go build -o main main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/main .
EXPOSE 30080

CMD ["./main"]

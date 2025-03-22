FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o /app/app cmd/main.go

FROM alpine:latest

RUN apk add --no-cache bash curl

WORKDIR /app

COPY --from=builder /app/app /app/app

EXPOSE 8888

RUN chmod +x /app/app

# Set the command to wait for PostgreSQL to be ready and then start the backend app
CMD ["./app"]

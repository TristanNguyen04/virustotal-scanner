# Build stage
FROM golang:1.24.2 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main

# Final minimal image
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/frontend ./frontend

# Make sure the uploads directory exists
RUN mkdir -p ./frontend/uploads

EXPOSE 8080

CMD ["./main"]
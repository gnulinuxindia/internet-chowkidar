FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM debian:bookworm-slim
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 9000
CMD ["./main"]

FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /calc_rest_api ./cmd/app/
FROM alpine:latest
WORKDIR /
COPY --from=builder /calc_rest_api .
COPY --from=builder /app/config.yaml .
EXPOSE 8080
CMD ["./calc_rest_api"]

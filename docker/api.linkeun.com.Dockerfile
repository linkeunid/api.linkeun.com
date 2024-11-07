FROM golang:1.23.2-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -ldflags="-s -w" -o ./api.linkeun.com ./cmd/api/

FROM alpine:latest AS api.linkeun.com
WORKDIR /app
COPY --from=builder /app/api.linkeun.com .
EXPOSE 4444
ENTRYPOINT ["./api.linkeun.com"]

FROM golang:1.23-alpine AS builder
WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./rillettes main.go

FROM alpine:latest AS final
COPY --from=builder /build/rillettes .

CMD ["./rillettes"]

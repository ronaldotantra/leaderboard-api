FROM golang:1.22.0 AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 go build -o main ./cmd/server

FROM alpine:3.21.0 

WORKDIR /

COPY --from=builder /app/main /main
COPY --from=builder /app/.env .

CMD ["/main"]

FROM golang:1.21-alpine as builder

RUN mkdir -p /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o brokerApp ./api

RUN chmod +x /app/brokerApp


FROM alpine:latest

RUN mkdir -p /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]

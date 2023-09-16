FROM golang:1.21-alpine as builder

RUN mkdir -p /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o loggerApp ./api

RUN chmod +x /app/loggerApp


FROM alpine:latest

RUN mkdir -p /app

COPY --from=builder /app/loggerApp /app

CMD ["/app/loggerApp"]

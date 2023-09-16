FROM golang:1.21-alpine as builder

RUN mkdir -p /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o authApp ./api

RUN chmod +x /app/authApp


FROM alpine:latest

RUN mkdir -p /app

COPY --from=builder /app/authApp /app

CMD ["/app/authApp"]

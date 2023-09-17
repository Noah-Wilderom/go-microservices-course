FROM golang:1.21-alpine as builder

RUN mkdir -p /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o listenerApp .

RUN chmod +x /app/listenerApp


FROM alpine:latest

RUN mkdir -p /app

COPY --from=builder /app/listenerApp /app

CMD ["/app/listenerApp"]

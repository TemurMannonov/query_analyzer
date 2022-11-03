FROM golang:1.19.1-alpine3.16 AS builder

WORKDIR /app   

COPY . .

RUN apk add curl
RUN go build -o main cmd/main.go
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz


FROM alpine:3.16
WORKDIR /app

COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY sample.env .
COPY start.sh .
COPY wait-for.sh .
COPY migrations ./migrations

EXPOSE 8000

# running built image
CMD ["/app/main"]
ENTRYPOINT ["/app/start.sh"]

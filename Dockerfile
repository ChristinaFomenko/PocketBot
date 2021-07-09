FROM golang:1.16.5-alpine3.14 AS builder

COPY . /github.com/KrisInferno/PocketBot/
WORKDIR /github.com/KrisInferno/PocketBot/

RUN go mod download
RUN go build -o ./bin/bot cmd/bot/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /github.com/KrisInferno/PocketBot/bin/bot .
COPY --from=0 /github.com/KrisInferno/PocketBot/configs configs/

EXPOSE 80

CMD ["./bot"]
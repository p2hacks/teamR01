FROM golang:1.13.5-alpine3.10 AS build-env

WORKDIR /Users/shotaro-yamada/go/src/github.com/p2hacks/teamR01/server/send-message
COPY ./ ./
RUN apk add --no-cache git
RUN go build -o server main.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /Users/shotaro-yamada/go/src/github.com/p2hacks/teamR01/server/send-message/server /usr/local/bin/server

EXPOSE 9003

CMD ["/usr/local/bin/server"]
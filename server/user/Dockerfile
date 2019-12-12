FROM golang:latest AS build-env

WORKDIR /Users/shotaro-yamada/go/src/github.com/p2hacks/teamR01/server/user
COPY ./ ./
RUN go build -o server main.go

FROM alpine:latest
RUN apk add --no-cache --update ca-certificates
COPY --from=build-env /Users/shotaro-yamada/go/src/github.com/p2hacks/teamR01/server/user /usr/local/bin/server

ENV DB_USER omamama
ENV DB_PASS 8Q*AvwuZKNcqdEPKJWy6ZHyNw_Th-fj_BRdtT_L-EUKQmGXVD3dF7ZsKVa4YfP!n
ENV DB_IP mysql
ENV DB_PORT 3306
ENV DB_NAME sign
ENV CHYPHER_KEY RaJ_GCdicC27ABxGQPD-7N@ri3RYqcgR

EXPOSE 9000

CMD ["/usr/local/bin/server"]
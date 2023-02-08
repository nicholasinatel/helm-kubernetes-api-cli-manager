FROM golang:1.18-alpine as susegoapp

WORKDIR /app

COPY . .

RUN apk add -u -t build-tools curl git && \
    go build -o susegoapp *.go && \
    apk del build-tools && \
    rm -rf /var/cache/apk/*

#
# Runtime container
#
FROM alpine:latest

WORKDIR /app

RUN apk --no-cache add ca-certificates bash

COPY --from=susegoapp /app/susegoapp /app/susegoapp

EXPOSE 8080

ENTRYPOINT ["/app/susegoapp"]

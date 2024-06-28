FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app

COPY content-fetcher /app/content-fetcher

WORKDIR /app

CMD ["/app/content-fetcher"]

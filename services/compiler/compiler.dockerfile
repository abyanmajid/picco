FROM alpine:latest

RUN apk update && apk add --no-cache \
    python3 \
    py3-pip \
    openjdk11-jdk \
    g++ \
    nodejs \
    npm \
    go

RUN mkdir /app

COPY compiler /app

CMD ["/app/compiler"]
FROM alpine:latest

RUN mkdir /app

COPY course /app

CMD ["/app/course"]
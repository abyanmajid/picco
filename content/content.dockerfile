FROM alpine:latest

RUN mkdir /app

COPY content /app

CMD ["/app/content"]
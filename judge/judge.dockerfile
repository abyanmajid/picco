FROM alpine:latest

RUN mkdir /app

COPY judge /app

CMD ["/app/judge"]
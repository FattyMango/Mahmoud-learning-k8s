FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY unicorn /app/unicorn


RUN chmod +x /app/unicorn


ENV PORT=8080

CMD ["./unicorn"]

FROM alpine:latest

RUN mkdir /app

COPY library /app

CMD [ "/app/library"]
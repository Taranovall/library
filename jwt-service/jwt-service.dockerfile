FROM alpine:latest

RUN mkdir /app

COPY jwtApp /app

CMD [ "/app/jwtApp"]
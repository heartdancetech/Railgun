FROM alpine:3.10.2

LABEL MAINTAINER="gsxhnd@gmail.com"

WORKDIR /app
ADD LastOrder-api-gateway /app
EXPOSE 8080
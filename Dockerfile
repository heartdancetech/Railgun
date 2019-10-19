FROM alpine:3.10.2

LABEL MAINTAINER="gsxhnd@gmail.com"

WORKDIR /app
ADD LastOrder /app
EXPOSE 8080
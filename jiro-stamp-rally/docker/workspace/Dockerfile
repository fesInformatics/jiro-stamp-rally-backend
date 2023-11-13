FROM golang:1.20.11-alpine

RUN apk --no-cache add curl && \
    curl -L https://github.com/golang-migrate/migrate/releases/download/v4.11.0/migrate.linux-amd64.tar.gz | tar xvz && \
    mv ./migrate.linux-amd64 /usr/bin/migrate

WORKDIR /workspace

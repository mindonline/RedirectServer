FROM alpine:3.12

COPY . /app
WORKDIR /app

ENV GOBIN=/app

# Setup Application & deps
RUN apk update && \
    apk upgrade && \
    apk add --no-cache --update --virtual .devs go && \
    go get ./ && go build -ldflags="-s -w" && \
    apk del .devs && \
    rm -rf /var/cache/apk/* && rm -rf /root/.cache/*

CMD ./app

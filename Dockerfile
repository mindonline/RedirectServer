FROM alpine:3.12

ARG PROJECTPATH=/go/src/tiny-server-go
ARG BIN=/app
ARG APPBIN=$BIN/server
ENV RUNTIME_APPBIN=$APPBIN

COPY . $PROJECTPATH
WORKDIR $PROJECTPATH

ENV GOBIN=$BIN
ENV GOPATH=/go

# Setup Application & deps
RUN apk update && \
    apk upgrade && \
    apk add --no-cache --update --virtual .devs go git && \
    go get -d ./ && go build -ldflags="-s -w" -o $APPBIN && \
    apk del .devs && \
    rm -rf /go && rm -rf /var/cache/apk/* && rm -rf /root/.cache/*

# Config files
COPY schema.json $BIN
COPY .env $BIN

#Prepare for binary run
WORKDIR $BIN
CMD $RUNTIME_APPBIN


FROM golang:1.22.4-alpine

WORKDIR $GOPATH/app

COPY go.mod ./
COPY go.sum ./


RUN go mod download

ENV SHORT_URL_SIZE=8
ENV APP_DOMAIN="http://localhost:9090/"

COPY ./ ./

CMD ["go","run","."]

FROM golang:latest

ENV GO111MODULE=on

COPY src/api /go/src/api

WORKDIR /go/src/api/

COPY src/api/go.mod .
COPY src/api/go.sum .

RUN go mod download

COPY . .

ENV PACKAGES mariadb-client
RUN apt-get update && apt-get install -y $PACKAGES


CMD ["sh", "./start_app.sh"]
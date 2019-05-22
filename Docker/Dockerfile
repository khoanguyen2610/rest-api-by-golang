FROM golang:1.11.0-alpine

RUN apk add git
WORKDIR /app

# ENV CGO_ENABLED=0

RUN go get github.com/githubnemo/CompileDaemon
RUN go get github.com/gorilla/mux

ENTRYPOINT CompileDaemon -build="go build" -command="./app" -directory="/app"

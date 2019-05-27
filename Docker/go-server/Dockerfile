FROM golang:1.12.5
COPY freshRunner.conf .
RUN mkdir -p /etc/go-app/fresh && \
    mv freshRunner.conf /etc/go-app/fresh/freshRunner.conf


WORKDIR /go/user-service
RUN go get github.com/githubnemo/CompileDaemon
RUN go mod download

ENTRYPOINT CompileDaemon -build="go build /go/user-service/cmd/http" -command="./http" -directory="/go/user-service" -color="true"

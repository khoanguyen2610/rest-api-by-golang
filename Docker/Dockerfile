FROM golang:1.12.5
COPY freshRunner.conf .
RUN mkdir -p /etc/go-app/fresh && \
    mv freshRunner.conf /etc/go-app/fresh/freshRunner.conf


WORKDIR /etc/go-app/app
RUN go get github.com/gorilla/mux
# RUN go get github.com/githubnemo/CompileDaemon
RUN go get github.com/gravityblast/fresh

# ENTRYPOINT CompileDaemon -build="go build" -command="./app" -directory="/app"
ENTRYPOINT fresh -c /etc/go-app/fresh/freshRunner.conf
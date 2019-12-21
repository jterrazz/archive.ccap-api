FROM golang:1.13-alpine

EXPOSE 8080

WORKDIR /app
COPY . /app

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build" --command=./ccap.live-api
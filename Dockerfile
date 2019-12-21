FROM golang:1.13-alpine

WORKDIR /app
COPY . /app

RUN go mod download
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o api" --command=./api
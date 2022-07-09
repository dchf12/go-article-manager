FROM golang:1.18-bullseye

ENV APP_HOME /go/src/go-article-manager
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"

COPY . $WORKDIR 
EXPOSE 8080
CMD go run ./
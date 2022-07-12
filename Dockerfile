# syntax=docker/dockerfile:1
FROM golang:1.18-bullseye

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /go-article-manager

EXPOSE 8080
CMD [ "/go-article-manager" ]
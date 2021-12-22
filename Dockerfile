# syntax=docker/dockerfile:1

# Alpine is chosen for its small footprint compared to Ubuntu
FROM golang:1.17-alpine

ENV GO111MODULE=on
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o /microservice

EXPOSE 9090

CMD [ "/microservice" ]
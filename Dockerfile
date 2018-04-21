FROM golang:1.10.1

MAINTAINER YuyaKODA

ADD webapp.go .

RUN go get -u github.com/gorilla/mux

CMD go run ./webapp.go

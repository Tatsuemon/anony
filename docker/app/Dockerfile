FROM golang:1.15.2-alpine

RUN apk --update --no-cache add git build-base curl

WORKDIR /go/src/github.com/Tatsuemon/anony

ENV GO111MODULE=on
ADD ./go.mod /go/src/github.com/Tatsuemon/anony/go.mod

RUN echo '[url "git@github.com:"]' > ~/.gitconfig \
  && echo "  insteadOf = https://github.com/" >> ~/.gitconfig \
  && go mod download

ADD . /go/src/github.com/Tatsuemon/anony

RUN go get -u github.com/cosmtrek/air && \
    go get bitbucket.org/liamstask/goose/cmd/goose

CMD air -c air.conf
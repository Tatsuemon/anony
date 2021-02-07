FROM golang:1.15.2-alpine

RUN apk --update --no-cache add git build-base curl

# Dockerize
ENV DOCKERIZE_VERSION v0.6.1
RUN apk add --no-cache openssl \
 && wget https://github.com/jwilder/dockerize/releases/download/$DOCKERIZE_VERSION/dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && tar -C /usr/local/bin -xzvf dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz \
 && rm dockerize-alpine-linux-amd64-$DOCKERIZE_VERSION.tar.gz

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
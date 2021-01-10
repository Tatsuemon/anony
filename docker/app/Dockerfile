FROM golang:1.15.2-alpine

RUN apk --update --no-cache add git curl

WORKDIR /go/src/github.com/Tatsuemon/shortURL

ENV GO111MODULE=on
ADD ./go.mod /go/src/github.com/Tatsuemon/shortURL/go.mod

RUN echo '[url "git@github.com:"]' > ~/.gitconfig \
  && echo "  insteadOf = https://github.com/" >> ~/.gitconfig \
  && go mod download

ADD . /go/src/github.com/Tatsuemon/shortURL

RUN go get -u github.com/cosmtrek/air

CMD air -c air.conf
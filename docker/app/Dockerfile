FROM golang:1.15.2

WORKDIR /go/src/github.com/Tatsuemon/shortURL

ENV GO111MODULE=on
ENV TZ=Asia/Tokyo

ADD . /go/src/github.com/Tatsuemon/shortURL

RUN go get -u github.com/cosmtrek/air

CMD air -c air.conf
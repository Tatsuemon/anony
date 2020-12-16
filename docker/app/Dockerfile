FROM golang:1.15.0
ENV GO111MODULE=on
WORKDIR /shortURL
COPY go.mod .
RUN go mod download
COPY . .

RUN go get -v github.com/rubenv/sql-migrate/...
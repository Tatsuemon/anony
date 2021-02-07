# DB Migration
install-goose:
	GO111MODULE=on go get bitbucket.org/liamstask/goose/cmd/goose;

# Protobuf
gen-proto-go:
	cd proto && \
		protoc -I. -I${GOPATH}/src *.proto --proto_path ../proto --go_out=plugins=grpc:../rpc --govalidators_out=paths=source_relative:../rpc;

install-protoc-gen-govalidators:
	GO111MODULE=off go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators;

# TEST
build-test:
	docker-compose -f docker-compose.test.yml build

up-test:
	docker-compose -f docker-compose.test.yml up -d

prepare-test-db:
	docker-compose -f docker-compose.test.yml exec -T app-test goose up

test:
	docker-compose -f docker-compose.test.yml exec -T app-test go test -v ./...
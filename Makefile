# DB Migration
install-goose:
	GO111MODULE=on go get bitbucket.org/liamstask/goose/cmd/goose;

# Develop
build-dev:
	docker-compose build

goose-up-dev:
	docker-compose run app goose up

goose-down-dev:
	docker-compose run app goose down

# Protobuf
gen-proto-go:
	cd proto && \
		protoc -I. -I${GOPATH}/src *.proto --proto_path ../proto --go_out=plugins=grpc:../rpc --govalidators_out=paths=source_relative:../rpc;

install-protoc-gen-govalidators:
	GO111MODULE=off go get github.com/mwitkow/go-proto-validators/protoc-gen-govalidators;

# TEST
build-test:
	docker-compose -f docker-compose.test.yml build

prepare-test-db:
	docker-compose -f docker-compose.test.yml run app-test goose up

test:
	docker-compose -f docker-compose.test.yml run app-test go test -v ./...

coverage:
	docker-compose -f docker-compose.test.yml run app-test go test -coverprofile=cover.out ./...
	go tool cover -html=cover.out
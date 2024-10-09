GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
DOCKER := $(shell type -P podman || type -P docker)



spicedb-up:
	./spicedb/start-spicedb.sh
.PHONY: spicedb-up

run:
	go run main.go
.PHONY: run
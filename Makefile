PROJECT_DIR := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))
SHELL = /bin/bash

build:
	cd ${PROJECT_DIR}
	mkdir -p dist/
	go build -ldflags '-extldflags "-static"' -o dist/writelog cmd/writelog.go

.PHONY: clean
clean:
	rm -rf ${PROJECT_DIR}/dist/ ${PROJECT_DIR}/.cache/
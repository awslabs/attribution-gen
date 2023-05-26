TOP := $(dir $(firstword $(MAKEFILE_LIST)))
SHELL := /bin/bash
GO111MODULE=on

OUT := ${TOP}/bin/gen-attributions


build:
	go build -o ${OUT} ${TOP}/cmd/*.go

generate: build
	${OUT} --depth 2 --output ${TOP}/ATTRIBUTIONS.md

cli-test:
	go test ${TOP}/test/...
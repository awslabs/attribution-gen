TOP := $(dir $(firstword $(MAKEFILE_LIST)))
SHELL := /bin/bash
GO111MODULE=on

OUT := ${TOP}/bin/attribution-gen


build:
	go build -o ${OUT} ${TOP}/cmd/*.go

generate: build
	${OUT} --depth 2 --output ${TOP}/ATTRIBUTIONS.md

cli-test:
	go test -v -count=1 ${TOP}/test/...
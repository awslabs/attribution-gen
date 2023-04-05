SHELL := /bin/bash
GO111MODULE=on

OUT := ./bin/gen-attributions

build:
	go build -o ${OUT} ./cmd/*.go

generate: build
	${OUT} --depth 2 --output ATTRIBUTIONS.md
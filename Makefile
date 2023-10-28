SHELL := /opt/homebrew/bin/bash
gen:
	buf generate
build:
	go build -o bin/app cmd/*

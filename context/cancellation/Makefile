.DEFAULT_GOAL := build

.PHONY: build run-ok run-cancel

build:
	cd ~/GolandProjects/go-learning/context/cancellation/
	go build

run-ok: build
	./cancellation false

run-cancel: build
	./cancellation true
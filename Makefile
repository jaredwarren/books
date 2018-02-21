#! /usr/bin/make
#

CURRENT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))

all: clean generate build

clean:
	@rm -rf app
	@rm -rf client
	@rm -rf tool
	@rm -rf swagger
	@rm -f books

generate:
	@rm -rf vendor/github.com/goadesign
	@goagen bootstrap -d github.com/jaredwarren/books/design

build:
	@CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -o books -ldflags '-w' .

ae-build:
	@if [ ! -d $(HOME)/books ]; then \
		mkdir $(HOME)/books; \
		ln -s $(CURRENT_DIR)/appengine.go $(HOME)/books/appengine.go; \
		ln -s $(CURRENT_DIR)/app.yaml     $(HOME)/books/app.yaml; \
	fi

ae-deploy: ae-build
	cd $(HOME)/books
	gcloud app deploy --project books
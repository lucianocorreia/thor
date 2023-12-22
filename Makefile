PACKAGES := $(shell go list ./...)
name := $(shell basename ${PWD})

## start: build and run local project
.PHONY: start
start:
	air

## test: run unit tests
.PHONY: test
test:
	go test -race -cover $(PACKAGES)

## css: build tailwindcss
.PHONY: css
css:
	./tailwindcss -i css/style.css -o static/css/style.css --minify

## css-watch: watch build tailwindcss
.PHONY: css-watch
css-watch:
	./tailwindcss -i css/style.css -o static/css/style.css --watch

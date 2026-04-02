# Makefile
setup:
	npm install
	npx husky init
	echo "npx --no -- commitlint --edit \$$1" > .husky/commit-msg
	chmod +x .husky/commit-msg

.PHONY: setup

run:
	cd getting-started/basic && go run main.go

.PHONY: run

build:
	cd getting-started/basic && go build main.go

.PHONY: build

run:
	cd getting-started/basic && go run main.go

.PHONY: run

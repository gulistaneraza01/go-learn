# Makefile
setup:
	npm install
	npx husky init
	echo "npx --no -- commitlint --edit \$$1" > .husky/commit-msg
	chmod +x .husky/commit-msg

.PHONY: setup

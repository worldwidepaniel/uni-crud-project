REPO_DIR=`pwd`
.PHONY: run
run:
	@docker-compose -f ./backend.yaml up
.PHONY: build
build:
	@docker build . -t custom-mongo-db

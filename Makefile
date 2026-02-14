APP_NAME=cli-todo
BIN_DIR=bin/

build:
	go build -o $(BIN_DIR)$(APP_NAME)

run:
	go run .




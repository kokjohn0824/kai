DIST_DIR=./dist
DEPLOY_DIR=~/bin
BINARY_NAME=gun
MAIN_PATH=./cmd/*.go

default: help

build: clean
	@go build -o ${DIST_DIR}/${BINARY_NAME} ${MAIN_PATH}

run: build 
	${DIST_DIR}/${BINARY_NAME}

clean:
	@go clean 
	@rm -rf ./${DIST_DIR}/**

deploy:
	go build -o ${DEPLOY_DIR}/${BINARY_NAME} ${MAIN_PATH}
	
help:
	@echo "Available targets:"
	@echo "  build:    Build the application"
	@echo "  run:      Build and run the application"
	@echo "  clean:    Clean up built binaries"
	@echo "  deploy:   Build and deploy the application"
	@echo "  help:     Show this help message"

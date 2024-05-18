.PHONY: default build run clean deploy help build-win

DIST_DIR=./dist
DEPLOY_DIR=~/bin
BINARY_NAME=kai
BINARY_EXE=kai.exe
MAIN_PATH=./cmd/*.go

default: help

run:  
	@go run ${MAIN_PATH}

clean:
	@go clean 
	@rm -rf ./${DIST_DIR}/**
	@echo cleaned up dist
	
build: clean
	@go build -ldflags="-s -w" -o ${DIST_DIR}/${BINARY_NAME} ${MAIN_PATH} && echo build complete

build-win: clean
	@GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o ${DIST_DIR}/$(BINARY_EXE) ./${MAIN_PATH}
	@upx ${DIST_DIR}/$(BINARY_EXE)
	@echo executable build complete

deploy: build
	@cp  ${DIST_DIR}/${BINARY_NAME} ${DEPLOY_DIR}/ && echo deploy complete

help:
	@echo "Available targets:"
	@echo "  build:    Build the application"
	@echo "  run:      Build and run the application"
	@echo "  clean:    Clean up built binaries"
	@echo "  deploy:   Build and deploy the application"
	@echo "  help:     Show this help message"

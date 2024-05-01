DIST_DIR=./dist
DEPLOY_DIR=~/bin
BINARY_NAME=gun

MAIN_PATH=./cmd/main.go


default: help

build:
	@go build -o ${DIST_DIR}/${BINARY_NAME} ${MAIN_PATH}

run: build 
	${DIST_DIR}/${BINARY_NAME}

clean:
	go clean 
	rm ${DIST_DIR}/*

deploy:
	go build -o ${DEPLOY_DIR}/${BINARY_NAME} ${MAIN_PATH}
	
help:
	@echo "projects for go "

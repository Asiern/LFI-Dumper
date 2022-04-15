PROJECT_NAME := lfidumper

all: build #test

build:
	go build -o ${PROJECT_NAME}.exe main.go

test:
	go test -v main.go

run:${PROJECT_NAME}
	./${PROJECT_NAME}

clean:
	go clean
	rm ${PROJECT_NAME}

.PHONY: build run clean test
PROJECT_NAME := lfidumper
FLAGS := -ldflags "-w -s"

ifeq ($(OS),Windows_NT)
	BINARY := ${PROJECT_NAME}.exe
else
	BINARY := ${PROJECT_NAME}
endif

all: build #test

build:
	go build ${FLAGS} -o ${BINARY} main.go 

test:
	go test -v main.go

run:${BINARY}
	./${BINARY}

clean:
	go clean
	rm ${PROJECT_NAME}

.PHONY: build run clean test

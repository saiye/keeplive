# Go参数
GO=go
GOCMD=$(GO)
GOBUILD=$(GOCMD) build
CGO_ENABLED=0
# 可执行文件名
# 读取系统环境
OS=$(shell uname -s)

BINARY_NAME=keeplive.exe
LINUX_BINARY_NAME=keeplive

all:
ifeq ($(OS),Linux)
	echo Running on Unix...
	make build-linux
else
	echo Running on Windows...
	make build
endif

build:
	$(GOBUILD) -o $(BINARY_NAME)

build-linux:
    GOOS=linux GOARCH=amd64 $(GOBUILD)  -o $(LINUX_BINARY_NAME)

clean:
	rm -f $(BINARY_NAME) $(LINUX_BINARY_NAME)
.PHONY: all build build-linux clean



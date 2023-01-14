all: build

build:
	docker build -t chaordic/collect-server-go .
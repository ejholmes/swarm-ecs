.PHONY: build

build:
	godep go build -o build/swarm-ecs ./cmd/swarm-ecs

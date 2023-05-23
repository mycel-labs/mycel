CONTAINER_NAME := mycel_i
DOCKERFILE := ./dockerfile-build
DOCKER_RUN_CMD := docker run --rm -it -v $(shell pwd):/mycel -w /mycel -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 --name mycel $(CONTAINER_NAME) 


test-all-keepers:
	go test ./x/.../keeper
test-all-types:
	go test ./x/.../types
test-all-keepers-types:
	go test ./x/.../keeper
	go test ./x/.../types
test-module-%:
	go test ./x/$*/keeper
	go test ./x/$*/types
docker-build:
	docker build -t $(CONTAINER_NAME) -f $(DOCKERFILE) .
docker-run:
	$(DOCKER_RUN_CMD)
serve:
	$(DOCKER_RUN_CMD) ignite chain serve -r
build:
	$(DOCKER_RUN_CMD) ignite chain build


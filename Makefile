CONTAINER_NAME := mycel_i
DOCKERFILE := ./dockerfile-build
DOCKER_RUN_CMD := docker run --rm -it -v $(shell pwd):/mycel -w /mycel -p 1317:1317 -p 3000:3000 -p 4500:4500 -p 5000:5000 -p 26657:26657 --name mycel $(CONTAINER_NAME) 

## Build
docker-build:
	docker build -t $(CONTAINER_NAME) -f $(DOCKERFILE) .
docker-run:
	$(DOCKER_RUN_CMD)
serve:
	$(DOCKER_RUN_CMD) ignite chain serve -r
build:
	$(DOCKER_RUN_CMD) ignite chain build

## Docker compose
docker-compose-build:
	docker compose build
docker-compose-up:
	docker compose up
docker-compose-build-arm:
	docker compose -f compose-arm.yml build
docker-compose-up-arm:
	docker compose -f compose-arm.yml up

## Test
test-all-module:
	go test ./x/.../
test-all-keepers:
	go test ./x/.../keeper
test-all-types:
	go test ./x/.../types
test-module-%:
	go test ./x/$*/.../


## Lint
golangci_lint_cmd=golangci-lint
golangci_version=v1.55.2

lint:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run --timeout=10m

lint-fix:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run --fix --out-format=tab --issues-exit-code=0

format:
	@go install mvdan.cc/gofumpt@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name "*.pb.go" -not -name "*.pb.gw.go" -not -name "*.pulsar.go" -not -path "./crypto/keys/secp256k1/*" | xargs gofumpt -w -l
	$(golangci_lint_cmd) run --fix
.PHONY: format

###############################################################################
###                                Protobuf                                 ###
###############################################################################

protoVer=0.14.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=docker run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

#? proto-all: Run make proto-format proto-lint proto-gen
proto-all: proto-format proto-lint proto-gen

#? proto-gen: Generate Protobuf files
proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh

#? proto-format: Format proto file
proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

#? proto-lint: Lint proto file
proto-lint:
	@$(protoImage) buf lint --error-format=json

.PHONY: proto-all proto-gen proto-format proto-lint

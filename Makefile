CONTAINER_NAME := mycel_i
DOCKER := $(shell which docker)
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
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

#? proto-all: Run make proto-format proto-lint proto-gen
proto-all: proto-format proto-lint proto-gen

#? proto-gen: Generate Protobuf files
proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh

#? proto-swagger-gen: Generate Protobuf Swagger
proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@$(protoImage) sh ./scripts/protoc-swagger-gen.sh

#? proto-format: Format proto file
proto-format:
	@$(protoImage) find ./ -name "*.proto" -exec clang-format -i {} \;

#? proto-lint: Lint proto file
proto-lint:
	@$(protoImage) buf lint --error-format=json

#? proto-check-breaking: Check proto file is breaking
proto-check-breaking:
	@$(protoImage) buf breaking --against $(HTTPS_GIT)#branch=main

CMT_URL              = https://raw.githubusercontent.com/cometbft/cometbft/v0.38.0/proto/tendermint

CMT_CRYPTO_TYPES     = proto/tendermint/crypto
CMT_ABCI_TYPES       = proto/tendermint/abci
CMT_TYPES            = proto/tendermint/types
CMT_VERSION          = proto/tendermint/version
CMT_LIBS             = proto/tendermint/libs/bits
CMT_P2P              = proto/tendermint/p2p

#? proto-update-deps: Update protobuf dependencies
proto-update-deps:
	@echo "Updating Protobuf dependencies"

	@mkdir -p $(CMT_ABCI_TYPES)
	@curl -sSL $(CMT_URL)/abci/types.proto > $(CMT_ABCI_TYPES)/types.proto

	@mkdir -p $(CMT_VERSION)
	@curl -sSL $(CMT_URL)/version/types.proto > $(CMT_VERSION)/types.proto

	@mkdir -p $(CMT_TYPES)
	@curl -sSL $(CMT_URL)/types/types.proto > $(CMT_TYPES)/types.proto
	@curl -sSL $(CMT_URL)/types/evidence.proto > $(CMT_TYPES)/evidence.proto
	@curl -sSL $(CMT_URL)/types/params.proto > $(CMT_TYPES)/params.proto
	@curl -sSL $(CMT_URL)/types/validator.proto > $(CMT_TYPES)/validator.proto
	@curl -sSL $(CMT_URL)/types/block.proto > $(CMT_TYPES)/block.proto

	@mkdir -p $(CMT_CRYPTO_TYPES)
	@curl -sSL $(CMT_URL)/crypto/proof.proto > $(CMT_CRYPTO_TYPES)/proof.proto
	@curl -sSL $(CMT_URL)/crypto/keys.proto > $(CMT_CRYPTO_TYPES)/keys.proto

	@mkdir -p $(CMT_LIBS)
	@curl -sSL $(CMT_URL)/libs/bits/types.proto > $(CMT_LIBS)/types.proto

	@mkdir -p $(CMT_P2P)
	@curl -sSL $(CMT_URL)/p2p/types.proto > $(CMT_P2P)/types.proto

	$(DOCKER) run --rm -v $(CURDIR)/proto:/workspace --workdir /workspace $(protoImageName) buf mod update

.PHONY: proto-all proto-gen proto-swagger-gen proto-format proto-lint proto-check-breaking proto-update-deps

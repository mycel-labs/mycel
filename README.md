# mycel
**mycel** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Dependencies
- Go `1.20.2`
- Ignite CLI `v0.26.1`
- Cosmos SDK `v0.46.7`

## Get started

for register .eth domain
```
export RPC_ENDPOINT_ETHEREUM_GOERLI="https://goerli.infura.io/v3/<YOUR_API_KEY>"
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Test
Test a specific module  
```
make test-module-{MODULE_NAME}
```

Test all keepers
```
make test-all-keepers
```

Test all types
```
make test-all-types
```

### Build with Docker
Build Docker Image
```
make docker-build
```

Build chain
```
make build
```

Serve Chains
```
make serve
```


### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

<!-- ### Web Frontend

Ignite CLI has scaffolded a React-based web app in the `react` directory. Run the following commands to install dependencies and start the app:

```
cd react
yarn
yarn dev
``` -->

For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Start Nodes with Docker
### Single node
Build
```
docker build . -t mycel -f dockerfile-node
```

Run
```
docker run -it --rm \
    -p26657:26657 \
    -p1317:1317 \
    -p4500:4500 \
    -v ~/.mycel:/root/.mycel \
    mycel
```
You can generate your `.mycel` config directory with `ignite chain init`

### Multiple nodes
#### Setup node1 using `docker compose`:
```
docker compose up
```

#### Setup node2:  
Initialize node2
```
docker compose exec node2 myceld init node2
```
Copy genesis.json
```
docker compose cp node1:/root/.mycel/config/genesis.json /tmp/genesis.json
docker compose cp /tmp/genesis.json node2:/root/.mycel/config/genesis.json
```
Update config.toml
```
docker compose exec node2 sed -i "s/persistent_peers = \"\"/persistent_peers = \"$(docker compose exec node1 myceld tendermint show-node-id)@node1:26656\"/g" /root/.mycel/config/config.toml
```
Setup key
```
docker compose exec node2 myceld keys add validator
NODE2_ADDR=$(docker compose exec node2 myceld keys show validator --output json | jq -r '.address') # enter password
```
Send stake token from node1
```
docker compose exec node1 myceld tx bank send alice $NODE2_ADDR 50000000stake
```
Stake
```
docker compose exec node2 myceld tx staking create-validator \
    --amount 50000000stake \
    --from validator --pubkey=$(docker compose exec node2 myceld tendermint show-validator) \
    --moniker="node2" \
    --commission-rate="0.1" \
    --commission-max-rate="0.2" \
    --commission-max-change-rate="0.01" \
    --min-self-delegation="50000000" \
    --node tcp://node1:26657
```

Check validators
```
docker compose exec node1 myceld q staking validators
```
Start node2
```
docker compose exec node2 myceld start
```



## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/mycel-domain/mycel@latest! | sudo bash
```

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)

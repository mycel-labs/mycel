#!/bin/bash

# domain registration
myceld tx registry register-domain foo cel 1 --from=alice -y
myceld tx registry register-domain hoge foo.cel 1 --from=alice -y
myceld tx registry register-domain bar cel 1 --from=alice -y
myceld tx registry register-domain piyo cel 1 --from=alice -y
myceld tx registry register-domain piyopiyo cel 1 --from=alice -y
myceld tx registry register-domain piyo piyo.cel 1 --from=alice -y

# update dns records
myceld tx registry update-dns-record foo cel A 192.168.0.1 --from=alice -y
myceld tx registry update-dns-record foo cel AAAA 2001:0db8:85a3:0000:0000:8a2e:0370:7334 --from=alice -y
myceld tx registry update-dns-record hoge foo.cel A 192.168.0.2 --from=alice -y
myceld tx registry update-dns-record bar cel A 10.0.0.1 --from=alice -y

# update wallet records
myceld tx registry update-wallet-record foo cel ETHEREUM_MAINNET 0x1234567890123456789012345678901234567890 --from=alice -y
myceld tx registry update-wallet-record foo cel POLYGON_MAINNET 0x1234567890123456789012345678901234567890 --from=alice -y
myceld tx registry update-wallet-record foo cel ETHEREUM_GOERLI 0x61b08DcE6751E5329984BD9098464Ee00c30984b --from=alice -y
myceld tx registry update-wallet-record hoge foo.cel ETHEREUM_MAINNET 0x1234567890123456789012345678901234567890 --from=alice -y
myceld tx registry update-wallet-record bar cel ETHEREUM_MAINNET 0x1234567890123456789012345678901234567890 --from=alice -y

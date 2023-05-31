# Incentives

## Abstract
Incentive module provides the following feature:
- Store incentives for each epochs
- Store validator and delegator incentives
- Distribute incentives to validator and delegator

![overview](https://user-images.githubusercontent.com/19934260/229605154-a9311e6c-1c55-43a5-88a8-2c2dd1020e16.png)

## Stores
[proto/mycel/incentives](https://github.com/mycel-domain/mycel/tree/main/proto/mycel/incentives)
### epoch_incentive.proto
```proto
message EpochIncentive {
  int64 epoch = 1; 
  repeated cosmos.base.v1beta1.Coin amount = 2 [
      (gogoproto.nullable) = false,
      (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins"
      ];
  bool isDistributed = 3; 
}
```

### validator_incentive.proto
```proto
message ValidatorIncentive {
  string address = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  repeated cosmos.base.v1beta1.Coin amount = 2 [
      (gogoproto.nullable) = false,
      (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins"
      ];
}
```

### delegator_incentive.proto
```proto
message DelegetorIncentive {
  string address = 1 [(cosmos_proto.scalar) = "cosmos.AddressString"];
  repeated cosmos.base.v1beta1.Coin amount = 2 [
      (gogoproto.nullable) = false,
      (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins"
      ];
}
```

## Events
This module emits the following events:

## Queries

### list-epoch-incentive
List all epoch incentive
```
myceld q incentives list-epoch-incentive
```

### show-epoch-incentive
```
myceld q incentives show-epoch-incentive [epoch]
```

### list-validator-incentive
List all validator incentive
```
myceld q incentives list-validator-incentive
```

### show-validator-incentive
```
myceld q incentives show-validator-incentive [address]
```

### list-delegator-incentive
List all delegator incentive
```
myceld q incentives list-delegator-incentive
```

### show-delegator-incentive
```
myceld q incentives show-delegator-incentive [address]
```

# Incentives

## Abstract
Incentive module provides the following feature:
- Store incentives for each epochs

## Stores
### epoch_incentives.proto
```proto
message Incentive {
  int64 epoch = 1; 
  repeated cosmos.base.v1beta1.Coin amount = 2 [
      (gogoproto.nullable) = false,
      (gogoproto.castrepeated) = "github.com/cosmos/cosmos-sdk/types.Coins"
      ];
  bool isDistributed = 3; 
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

### show-incentive
```
myceld q incentives show-epoch-incentive [epoch]
```

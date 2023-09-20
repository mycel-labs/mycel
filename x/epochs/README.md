# Epochs

## Abstract
[Osmosis's](https://github.com/osmosis-labs/osmosis/tree/x/epochs/v0.0.2/x/epochs) Epochs module provides the following feature:
- On-chain Timers that execute at fixed time intervals

## Stores
[proto/mycel/epochs](https://github.com/mycel-domain/mycel/tree/main/proto/mycel/epochs)
### epoch_info.proto
```proto
message EpochInfo {
  string identifier = 1; 
  google.protobuf.Timestamp startTime = 2 [
    (gogoproto.stdtime) = true,
    (gogoproto.nullable) = false,
    (gogoproto.moretags) = "yaml:\"start_time\""
  ];
  google.protobuf.Duration duration = 3[
    (gogoproto.nullable) = false,
    (gogoproto.stdduration) = true,
    (gogoproto.jsontag) = "duration,omitempty",
    (gogoproto.moretags) = "yaml:\"duration\""
  ]; 
  int64 currentEpoch = 4; 
  google.protobuf.Timestamp currentEpochStartTime = 5[
    (gogoproto.stdtime) = true,
    (gogoproto.nullable) = false,
    (gogoproto.moretags) = "yaml:\"current_epoch_start_time\""
  ]; 
  bool epochCountingStarted = 6; 
  int64 currentEpochStartHeight = 7; 
}

```

## Events
This module emits the following events:

### BeginBlocker
Event Type: `epoch-start`
Attributes:
- `start-time`: Epoch start time
- `epoch-number`: Epoch number

### EndBlocker
Event Type: `epoch-end`
Attributes:
- `epoch-number`: Epoch number


## Hooks
```go
  // the first block whose timestamp is after the duration is counted as the end of the epoch
  AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64)
  // new epoch is next block of epoch end block
  BeforeEpochStart(ctx sdk.Context, epochIdentifier string, epochNumber int64)
```

### How modules receive hooks

On hook receiver function of other modules, they need to filter epochIdentifier and only do executions for only specific epochIdentifier. Filtering epochIdentifier could be in Params of other modules so that they can be modified by governance.

This is the standard dev UX of this:
```go
func (k MyModuleKeeper) AfterEpochEnd(ctx sdk.Context, epochIdentifier string, epochNumber int64) {
    params := k.GetParams(ctx)
    if epochIdentifier == params.DistrEpochIdentifier {
    // my logic
  }
}
```

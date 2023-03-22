# Registry

## Abstract
Registry module provides the following feature:
- Register domains
- Update name records

## Stores
### domain.proto
```proto
enum DNSRecordType {
  A = 0;
  AAAA = 1;
  CNAME = 2;
  NS = 3;
  MX = 4;
  PTR = 5;
  SOA = 6;
  SRV = 7;
  TXT = 8;
}

enum WalletRecordType {
  ETHEREUM_MAINNET = 0;
  ETHEREUM_GOERLI = 1;
  POLYGON_MAINNET = 2;
  POLYGON_MUMBAI = 3;
}

enum WalletAddressFormat {
  ETHEREUM = 0;
}

message DNSRecord {
  DNSRecordType DNSRecordType = 1;
  string value = 2;
}

message WalletRecord {
  WalletRecordType walletRecordType = 1;
  WalletAddressFormat WalletAddressFormat = 2;
  string value = 3;
}

message Domain {
  string name = 1; 
  string parent = 2; 
  string owner = 3; 
  int64 expirationDate = 4; 
  map<string, DNSRecord> DNSRecords = 5;
  map<string, WalletRecord> walletRecords = 6;
  map<string, string> metadata = 7;
}

```

## Events
Registry module emits the following events:

### RegisterDomain
Event Type: `register-domain`  
Attributes:
- `name`: Domain name
- `parent`: Domain parent
- `registration-period-in-year`:  Registration period in year
- `expiration-date`: Expiration date in Unix time
- `domain-level`: Domain level

### UpdateWalletRecord
Event Type: `update-wallet-record`  
Attributes:
- `name`: Domain name
- `parent`: Domain parent
- `wallet-record-type`: Wallet record type
- `value`: Wallet address


## Transactions
### register-domain
Register domain to mycel  

```
myceld tx registry register-domain [name] [parent] [registration-period-in-year]
```

### update-wallet-record
Update wallet address record  

```
myceld tx registry update-wallet-record [name] [parent] [wallet-record-type] [value]
```

## Queries

### list-domain
List all domain
```
myceld q regisry list-domain
```
An example output:
```
domain:
- DNSRecords: {}
  expirationDate: "0"
  metadata: {}
  name: cel
  owner: ""
  parent: ""
  walletRecords: {}
- DNSRecords: {}
  expirationDate: "1711123442987026000"
  metadata: {}
  name: foo
  owner: cosmos1tk8gg20pcdp9alnnn6a84tdycf7pa2rjg8kwmc
  parent: cel
  walletRecords: {}
pagination:
  next_key: null
  total: "0"
```

### show-domain
```
myceld q regisry show-domain [name] [parent]
```

exmaple:  
Query `foo.cel`  
```
myceld q registry show-domain foo cel
```
Output: 
```
domain:
  DNSRecords: {}
  expirationDate: "1711123442987026000"
  metadata: {}
  name: foo
  owner: cosmos1tk8gg20pcdp9alnnn6a84tdycf7pa2rjg8kwmc
  parent: cel
  walletRecords:
    ETHEREUM_MAINNET:
      WalletAddressFormat: ETHEREUM
      value: 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
      walletRecordType: ETHEREUM_MAINNET
```

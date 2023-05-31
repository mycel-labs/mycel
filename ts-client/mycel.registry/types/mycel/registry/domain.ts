/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "mycel.registry";

export enum DnsRecordType {
  A = 0,
  AAAA = 1,
  CNAME = 2,
  NS = 3,
  MX = 4,
  PTR = 5,
  SOA = 6,
  SRV = 7,
  TXT = 8,
  UNRECOGNIZED = -1,
}

export function dnsRecordTypeFromJSON(object: any): DnsRecordType {
  switch (object) {
    case 0:
    case "A":
      return DnsRecordType.A;
    case 1:
    case "AAAA":
      return DnsRecordType.AAAA;
    case 2:
    case "CNAME":
      return DnsRecordType.CNAME;
    case 3:
    case "NS":
      return DnsRecordType.NS;
    case 4:
    case "MX":
      return DnsRecordType.MX;
    case 5:
    case "PTR":
      return DnsRecordType.PTR;
    case 6:
    case "SOA":
      return DnsRecordType.SOA;
    case 7:
    case "SRV":
      return DnsRecordType.SRV;
    case 8:
    case "TXT":
      return DnsRecordType.TXT;
    case -1:
    case "UNRECOGNIZED":
    default:
      return DnsRecordType.UNRECOGNIZED;
  }
}

export function dnsRecordTypeToJSON(object: DnsRecordType): string {
  switch (object) {
    case DnsRecordType.A:
      return "A";
    case DnsRecordType.AAAA:
      return "AAAA";
    case DnsRecordType.CNAME:
      return "CNAME";
    case DnsRecordType.NS:
      return "NS";
    case DnsRecordType.MX:
      return "MX";
    case DnsRecordType.PTR:
      return "PTR";
    case DnsRecordType.SOA:
      return "SOA";
    case DnsRecordType.SRV:
      return "SRV";
    case DnsRecordType.TXT:
      return "TXT";
    case DnsRecordType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum DnsRecordFormat {
  IPV4 = 0,
  IPV6 = 1,
  FQDN = 2,
  UNRECOGNIZED = -1,
}

export function dnsRecordFormatFromJSON(object: any): DnsRecordFormat {
  switch (object) {
    case 0:
    case "IPV4":
      return DnsRecordFormat.IPV4;
    case 1:
    case "IPV6":
      return DnsRecordFormat.IPV6;
    case 2:
    case "FQDN":
      return DnsRecordFormat.FQDN;
    case -1:
    case "UNRECOGNIZED":
    default:
      return DnsRecordFormat.UNRECOGNIZED;
  }
}

export function dnsRecordFormatToJSON(object: DnsRecordFormat): string {
  switch (object) {
    case DnsRecordFormat.IPV4:
      return "IPV4";
    case DnsRecordFormat.IPV6:
      return "IPV6";
    case DnsRecordFormat.FQDN:
      return "FQDN";
    case DnsRecordFormat.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum WalletRecordType {
  ETHEREUM_MAINNET = 0,
  ETHEREUM_GOERLI = 1,
  POLYGON_MAINNET = 2,
  POLYGON_MUMBAI = 3,
  GNOSIS_MAINNET = 4,
  GNOSIS_CHIADO = 5,
  UNRECOGNIZED = -1,
}

export function walletRecordTypeFromJSON(object: any): WalletRecordType {
  switch (object) {
    case 0:
    case "ETHEREUM_MAINNET":
      return WalletRecordType.ETHEREUM_MAINNET;
    case 1:
    case "ETHEREUM_GOERLI":
      return WalletRecordType.ETHEREUM_GOERLI;
    case 2:
    case "POLYGON_MAINNET":
      return WalletRecordType.POLYGON_MAINNET;
    case 3:
    case "POLYGON_MUMBAI":
      return WalletRecordType.POLYGON_MUMBAI;
    case 4:
    case "GNOSIS_MAINNET":
      return WalletRecordType.GNOSIS_MAINNET;
    case 5:
    case "GNOSIS_CHIADO":
      return WalletRecordType.GNOSIS_CHIADO;
    case -1:
    case "UNRECOGNIZED":
    default:
      return WalletRecordType.UNRECOGNIZED;
  }
}

export function walletRecordTypeToJSON(object: WalletRecordType): string {
  switch (object) {
    case WalletRecordType.ETHEREUM_MAINNET:
      return "ETHEREUM_MAINNET";
    case WalletRecordType.ETHEREUM_GOERLI:
      return "ETHEREUM_GOERLI";
    case WalletRecordType.POLYGON_MAINNET:
      return "POLYGON_MAINNET";
    case WalletRecordType.POLYGON_MUMBAI:
      return "POLYGON_MUMBAI";
    case WalletRecordType.GNOSIS_MAINNET:
      return "GNOSIS_MAINNET";
    case WalletRecordType.GNOSIS_CHIADO:
      return "GNOSIS_CHIADO";
    case WalletRecordType.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export enum WalletAddressFormat {
  ETHEREUM = 0,
  UNRECOGNIZED = -1,
}

export function walletAddressFormatFromJSON(object: any): WalletAddressFormat {
  switch (object) {
    case 0:
    case "ETHEREUM":
      return WalletAddressFormat.ETHEREUM;
    case -1:
    case "UNRECOGNIZED":
    default:
      return WalletAddressFormat.UNRECOGNIZED;
  }
}

export function walletAddressFormatToJSON(object: WalletAddressFormat): string {
  switch (object) {
    case WalletAddressFormat.ETHEREUM:
      return "ETHEREUM";
    case WalletAddressFormat.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface DnsRecord {
  dnsRecordType: DnsRecordType;
  dnsRecordFormat: DnsRecordFormat;
  value: string;
}

export interface WalletRecord {
  walletRecordType: WalletRecordType;
  walletAddressFormat: WalletAddressFormat;
  value: string;
}

export interface Domain {
  name: string;
  parent: string;
  owner: string;
  expirationDate: number;
  dnsRecords: { [key: string]: DnsRecord };
  walletRecords: { [key: string]: WalletRecord };
  metadata: { [key: string]: string };
}

export interface Domain_DnsRecordsEntry {
  key: string;
  value: DnsRecord | undefined;
}

export interface Domain_WalletRecordsEntry {
  key: string;
  value: WalletRecord | undefined;
}

export interface Domain_MetadataEntry {
  key: string;
  value: string;
}

function createBaseDnsRecord(): DnsRecord {
  return { dnsRecordType: 0, dnsRecordFormat: 0, value: "" };
}

export const DnsRecord = {
  encode(message: DnsRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.dnsRecordType !== 0) {
      writer.uint32(8).int32(message.dnsRecordType);
    }
    if (message.dnsRecordFormat !== 0) {
      writer.uint32(16).int32(message.dnsRecordFormat);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DnsRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDnsRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.dnsRecordType = reader.int32() as any;
          break;
        case 2:
          message.dnsRecordFormat = reader.int32() as any;
          break;
        case 3:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DnsRecord {
    return {
      dnsRecordType: isSet(object.dnsRecordType) ? dnsRecordTypeFromJSON(object.dnsRecordType) : 0,
      dnsRecordFormat: isSet(object.dnsRecordFormat) ? dnsRecordFormatFromJSON(object.dnsRecordFormat) : 0,
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: DnsRecord): unknown {
    const obj: any = {};
    message.dnsRecordType !== undefined && (obj.dnsRecordType = dnsRecordTypeToJSON(message.dnsRecordType));
    message.dnsRecordFormat !== undefined && (obj.dnsRecordFormat = dnsRecordFormatToJSON(message.dnsRecordFormat));
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DnsRecord>, I>>(object: I): DnsRecord {
    const message = createBaseDnsRecord();
    message.dnsRecordType = object.dnsRecordType ?? 0;
    message.dnsRecordFormat = object.dnsRecordFormat ?? 0;
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseWalletRecord(): WalletRecord {
  return { walletRecordType: 0, walletAddressFormat: 0, value: "" };
}

export const WalletRecord = {
  encode(message: WalletRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.walletRecordType !== 0) {
      writer.uint32(8).int32(message.walletRecordType);
    }
    if (message.walletAddressFormat !== 0) {
      writer.uint32(16).int32(message.walletAddressFormat);
    }
    if (message.value !== "") {
      writer.uint32(26).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): WalletRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseWalletRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.walletRecordType = reader.int32() as any;
          break;
        case 2:
          message.walletAddressFormat = reader.int32() as any;
          break;
        case 3:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): WalletRecord {
    return {
      walletRecordType: isSet(object.walletRecordType) ? walletRecordTypeFromJSON(object.walletRecordType) : 0,
      walletAddressFormat: isSet(object.walletAddressFormat)
        ? walletAddressFormatFromJSON(object.walletAddressFormat)
        : 0,
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: WalletRecord): unknown {
    const obj: any = {};
    message.walletRecordType !== undefined && (obj.walletRecordType = walletRecordTypeToJSON(message.walletRecordType));
    message.walletAddressFormat !== undefined
      && (obj.walletAddressFormat = walletAddressFormatToJSON(message.walletAddressFormat));
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<WalletRecord>, I>>(object: I): WalletRecord {
    const message = createBaseWalletRecord();
    message.walletRecordType = object.walletRecordType ?? 0;
    message.walletAddressFormat = object.walletAddressFormat ?? 0;
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseDomain(): Domain {
  return { name: "", parent: "", owner: "", expirationDate: 0, dnsRecords: {}, walletRecords: {}, metadata: {} };
}

export const Domain = {
  encode(message: Domain, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.parent !== "") {
      writer.uint32(18).string(message.parent);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    if (message.expirationDate !== 0) {
      writer.uint32(32).int64(message.expirationDate);
    }
    Object.entries(message.dnsRecords).forEach(([key, value]) => {
      Domain_DnsRecordsEntry.encode({ key: key as any, value }, writer.uint32(42).fork()).ldelim();
    });
    Object.entries(message.walletRecords).forEach(([key, value]) => {
      Domain_WalletRecordsEntry.encode({ key: key as any, value }, writer.uint32(50).fork()).ldelim();
    });
    Object.entries(message.metadata).forEach(([key, value]) => {
      Domain_MetadataEntry.encode({ key: key as any, value }, writer.uint32(58).fork()).ldelim();
    });
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Domain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDomain();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.parent = reader.string();
          break;
        case 3:
          message.owner = reader.string();
          break;
        case 4:
          message.expirationDate = longToNumber(reader.int64() as Long);
          break;
        case 5:
          const entry5 = Domain_DnsRecordsEntry.decode(reader, reader.uint32());
          if (entry5.value !== undefined) {
            message.dnsRecords[entry5.key] = entry5.value;
          }
          break;
        case 6:
          const entry6 = Domain_WalletRecordsEntry.decode(reader, reader.uint32());
          if (entry6.value !== undefined) {
            message.walletRecords[entry6.key] = entry6.value;
          }
          break;
        case 7:
          const entry7 = Domain_MetadataEntry.decode(reader, reader.uint32());
          if (entry7.value !== undefined) {
            message.metadata[entry7.key] = entry7.value;
          }
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Domain {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      parent: isSet(object.parent) ? String(object.parent) : "",
      owner: isSet(object.owner) ? String(object.owner) : "",
      expirationDate: isSet(object.expirationDate) ? Number(object.expirationDate) : 0,
      dnsRecords: isObject(object.dnsRecords)
        ? Object.entries(object.dnsRecords).reduce<{ [key: string]: DnsRecord }>((acc, [key, value]) => {
          acc[key] = DnsRecord.fromJSON(value);
          return acc;
        }, {})
        : {},
      walletRecords: isObject(object.walletRecords)
        ? Object.entries(object.walletRecords).reduce<{ [key: string]: WalletRecord }>((acc, [key, value]) => {
          acc[key] = WalletRecord.fromJSON(value);
          return acc;
        }, {})
        : {},
      metadata: isObject(object.metadata)
        ? Object.entries(object.metadata).reduce<{ [key: string]: string }>((acc, [key, value]) => {
          acc[key] = String(value);
          return acc;
        }, {})
        : {},
    };
  },

  toJSON(message: Domain): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.parent !== undefined && (obj.parent = message.parent);
    message.owner !== undefined && (obj.owner = message.owner);
    message.expirationDate !== undefined && (obj.expirationDate = Math.round(message.expirationDate));
    obj.dnsRecords = {};
    if (message.dnsRecords) {
      Object.entries(message.dnsRecords).forEach(([k, v]) => {
        obj.dnsRecords[k] = DnsRecord.toJSON(v);
      });
    }
    obj.walletRecords = {};
    if (message.walletRecords) {
      Object.entries(message.walletRecords).forEach(([k, v]) => {
        obj.walletRecords[k] = WalletRecord.toJSON(v);
      });
    }
    obj.metadata = {};
    if (message.metadata) {
      Object.entries(message.metadata).forEach(([k, v]) => {
        obj.metadata[k] = v;
      });
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Domain>, I>>(object: I): Domain {
    const message = createBaseDomain();
    message.name = object.name ?? "";
    message.parent = object.parent ?? "";
    message.owner = object.owner ?? "";
    message.expirationDate = object.expirationDate ?? 0;
    message.dnsRecords = Object.entries(object.dnsRecords ?? {}).reduce<{ [key: string]: DnsRecord }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = DnsRecord.fromPartial(value);
        }
        return acc;
      },
      {},
    );
    message.walletRecords = Object.entries(object.walletRecords ?? {}).reduce<{ [key: string]: WalletRecord }>(
      (acc, [key, value]) => {
        if (value !== undefined) {
          acc[key] = WalletRecord.fromPartial(value);
        }
        return acc;
      },
      {},
    );
    message.metadata = Object.entries(object.metadata ?? {}).reduce<{ [key: string]: string }>((acc, [key, value]) => {
      if (value !== undefined) {
        acc[key] = String(value);
      }
      return acc;
    }, {});
    return message;
  },
};

function createBaseDomain_DnsRecordsEntry(): Domain_DnsRecordsEntry {
  return { key: "", value: undefined };
}

export const Domain_DnsRecordsEntry = {
  encode(message: Domain_DnsRecordsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      DnsRecord.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Domain_DnsRecordsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDomain_DnsRecordsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = DnsRecord.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Domain_DnsRecordsEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? DnsRecord.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: Domain_DnsRecordsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? DnsRecord.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Domain_DnsRecordsEntry>, I>>(object: I): Domain_DnsRecordsEntry {
    const message = createBaseDomain_DnsRecordsEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? DnsRecord.fromPartial(object.value)
      : undefined;
    return message;
  },
};

function createBaseDomain_WalletRecordsEntry(): Domain_WalletRecordsEntry {
  return { key: "", value: undefined };
}

export const Domain_WalletRecordsEntry = {
  encode(message: Domain_WalletRecordsEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== undefined) {
      WalletRecord.encode(message.value, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Domain_WalletRecordsEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDomain_WalletRecordsEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = WalletRecord.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Domain_WalletRecordsEntry {
    return {
      key: isSet(object.key) ? String(object.key) : "",
      value: isSet(object.value) ? WalletRecord.fromJSON(object.value) : undefined,
    };
  },

  toJSON(message: Domain_WalletRecordsEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value ? WalletRecord.toJSON(message.value) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Domain_WalletRecordsEntry>, I>>(object: I): Domain_WalletRecordsEntry {
    const message = createBaseDomain_WalletRecordsEntry();
    message.key = object.key ?? "";
    message.value = (object.value !== undefined && object.value !== null)
      ? WalletRecord.fromPartial(object.value)
      : undefined;
    return message;
  },
};

function createBaseDomain_MetadataEntry(): Domain_MetadataEntry {
  return { key: "", value: "" };
}

export const Domain_MetadataEntry = {
  encode(message: Domain_MetadataEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.key !== "") {
      writer.uint32(10).string(message.key);
    }
    if (message.value !== "") {
      writer.uint32(18).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Domain_MetadataEntry {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDomain_MetadataEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.key = reader.string();
          break;
        case 2:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Domain_MetadataEntry {
    return { key: isSet(object.key) ? String(object.key) : "", value: isSet(object.value) ? String(object.value) : "" };
  },

  toJSON(message: Domain_MetadataEntry): unknown {
    const obj: any = {};
    message.key !== undefined && (obj.key = message.key);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Domain_MetadataEntry>, I>>(object: I): Domain_MetadataEntry {
    const message = createBaseDomain_MetadataEntry();
    message.key = object.key ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isObject(value: any): boolean {
  return typeof value === "object" && value !== null;
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

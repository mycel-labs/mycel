/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "mycel.registry";

export interface MsgUpdateWalletRecord {
  creator: string;
  name: string;
  parent: string;
  walletRecordType: string;
  value: string;
}

export interface MsgUpdateWalletRecordResponse {
}

export interface MsgRegisterDomain {
  creator: string;
  name: string;
  parent: string;
  registrationPeriodInYear: number;
}

export interface MsgRegisterDomainResponse {
}

function createBaseMsgUpdateWalletRecord(): MsgUpdateWalletRecord {
  return { creator: "", name: "", parent: "", walletRecordType: "", value: "" };
}

export const MsgUpdateWalletRecord = {
  encode(message: MsgUpdateWalletRecord, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.parent !== "") {
      writer.uint32(26).string(message.parent);
    }
    if (message.walletRecordType !== "") {
      writer.uint32(34).string(message.walletRecordType);
    }
    if (message.value !== "") {
      writer.uint32(42).string(message.value);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateWalletRecord {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateWalletRecord();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.parent = reader.string();
          break;
        case 4:
          message.walletRecordType = reader.string();
          break;
        case 5:
          message.value = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateWalletRecord {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      parent: isSet(object.parent) ? String(object.parent) : "",
      walletRecordType: isSet(object.walletRecordType) ? String(object.walletRecordType) : "",
      value: isSet(object.value) ? String(object.value) : "",
    };
  },

  toJSON(message: MsgUpdateWalletRecord): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.parent !== undefined && (obj.parent = message.parent);
    message.walletRecordType !== undefined && (obj.walletRecordType = message.walletRecordType);
    message.value !== undefined && (obj.value = message.value);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateWalletRecord>, I>>(object: I): MsgUpdateWalletRecord {
    const message = createBaseMsgUpdateWalletRecord();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.parent = object.parent ?? "";
    message.walletRecordType = object.walletRecordType ?? "";
    message.value = object.value ?? "";
    return message;
  },
};

function createBaseMsgUpdateWalletRecordResponse(): MsgUpdateWalletRecordResponse {
  return {};
}

export const MsgUpdateWalletRecordResponse = {
  encode(_: MsgUpdateWalletRecordResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateWalletRecordResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateWalletRecordResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateWalletRecordResponse {
    return {};
  },

  toJSON(_: MsgUpdateWalletRecordResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgUpdateWalletRecordResponse>, I>>(_: I): MsgUpdateWalletRecordResponse {
    const message = createBaseMsgUpdateWalletRecordResponse();
    return message;
  },
};

function createBaseMsgRegisterDomain(): MsgRegisterDomain {
  return { creator: "", name: "", parent: "", registrationPeriodInYear: 0 };
}

export const MsgRegisterDomain = {
  encode(message: MsgRegisterDomain, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.parent !== "") {
      writer.uint32(26).string(message.parent);
    }
    if (message.registrationPeriodInYear !== 0) {
      writer.uint32(32).uint64(message.registrationPeriodInYear);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterDomain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterDomain();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.parent = reader.string();
          break;
        case 4:
          message.registrationPeriodInYear = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRegisterDomain {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      name: isSet(object.name) ? String(object.name) : "",
      parent: isSet(object.parent) ? String(object.parent) : "",
      registrationPeriodInYear: isSet(object.registrationPeriodInYear) ? Number(object.registrationPeriodInYear) : 0,
    };
  },

  toJSON(message: MsgRegisterDomain): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.parent !== undefined && (obj.parent = message.parent);
    message.registrationPeriodInYear !== undefined
      && (obj.registrationPeriodInYear = Math.round(message.registrationPeriodInYear));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterDomain>, I>>(object: I): MsgRegisterDomain {
    const message = createBaseMsgRegisterDomain();
    message.creator = object.creator ?? "";
    message.name = object.name ?? "";
    message.parent = object.parent ?? "";
    message.registrationPeriodInYear = object.registrationPeriodInYear ?? 0;
    return message;
  },
};

function createBaseMsgRegisterDomainResponse(): MsgRegisterDomainResponse {
  return {};
}

export const MsgRegisterDomainResponse = {
  encode(_: MsgRegisterDomainResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRegisterDomainResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRegisterDomainResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRegisterDomainResponse {
    return {};
  },

  toJSON(_: MsgRegisterDomainResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRegisterDomainResponse>, I>>(_: I): MsgRegisterDomainResponse {
    const message = createBaseMsgRegisterDomainResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  UpdateWalletRecord(request: MsgUpdateWalletRecord): Promise<MsgUpdateWalletRecordResponse>;
  RegisterDomain(request: MsgRegisterDomain): Promise<MsgRegisterDomainResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.UpdateWalletRecord = this.UpdateWalletRecord.bind(this);
    this.RegisterDomain = this.RegisterDomain.bind(this);
  }
  UpdateWalletRecord(request: MsgUpdateWalletRecord): Promise<MsgUpdateWalletRecordResponse> {
    const data = MsgUpdateWalletRecord.encode(request).finish();
    const promise = this.rpc.request("mycel.registry.Msg", "UpdateWalletRecord", data);
    return promise.then((data) => MsgUpdateWalletRecordResponse.decode(new _m0.Reader(data)));
  }

  RegisterDomain(request: MsgRegisterDomain): Promise<MsgRegisterDomainResponse> {
    const data = MsgRegisterDomain.encode(request).finish();
    const promise = this.rpc.request("mycel.registry.Msg", "RegisterDomain", data);
    return promise.then((data) => MsgRegisterDomainResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "mycel.registry";

export interface OwnedDomain {
  name: string;
  parent: string;
}

export interface DomainOwnership {
  owner: string;
  domains: OwnedDomain[];
}

function createBaseOwnedDomain(): OwnedDomain {
  return { name: "", parent: "" };
}

export const OwnedDomain = {
  encode(message: OwnedDomain, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.parent !== "") {
      writer.uint32(18).string(message.parent);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OwnedDomain {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOwnedDomain();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.name = reader.string();
          break;
        case 2:
          message.parent = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OwnedDomain {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      parent: isSet(object.parent) ? String(object.parent) : "",
    };
  },

  toJSON(message: OwnedDomain): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.parent !== undefined && (obj.parent = message.parent);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<OwnedDomain>, I>>(object: I): OwnedDomain {
    const message = createBaseOwnedDomain();
    message.name = object.name ?? "";
    message.parent = object.parent ?? "";
    return message;
  },
};

function createBaseDomainOwnership(): DomainOwnership {
  return { owner: "", domains: [] };
}

export const DomainOwnership = {
  encode(message: DomainOwnership, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    for (const v of message.domains) {
      OwnedDomain.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DomainOwnership {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDomainOwnership();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        case 2:
          message.domains.push(OwnedDomain.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): DomainOwnership {
    return {
      owner: isSet(object.owner) ? String(object.owner) : "",
      domains: Array.isArray(object?.domains) ? object.domains.map((e: any) => OwnedDomain.fromJSON(e)) : [],
    };
  },

  toJSON(message: DomainOwnership): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    if (message.domains) {
      obj.domains = message.domains.map((e) => e ? OwnedDomain.toJSON(e) : undefined);
    } else {
      obj.domains = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DomainOwnership>, I>>(object: I): DomainOwnership {
    const message = createBaseDomainOwnership();
    message.owner = object.owner ?? "";
    message.domains = object.domains?.map((e) => OwnedDomain.fromPartial(e)) || [];
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

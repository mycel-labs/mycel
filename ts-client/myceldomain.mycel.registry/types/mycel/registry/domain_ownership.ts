/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "myceldomain.mycel.registry";

export interface DomainOwnership {
  owner: string;
  domains: string[];
}

function createBaseDomainOwnership(): DomainOwnership {
  return { owner: "", domains: [] };
}

export const DomainOwnership = {
  encode(message: DomainOwnership, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    for (const v of message.domains) {
      writer.uint32(18).string(v!);
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
          message.domains.push(reader.string());
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
      domains: Array.isArray(object?.domains) ? object.domains.map((e: any) => String(e)) : [],
    };
  },

  toJSON(message: DomainOwnership): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    if (message.domains) {
      obj.domains = message.domains.map((e) => e);
    } else {
      obj.domains = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<DomainOwnership>, I>>(object: I): DomainOwnership {
    const message = createBaseDomainOwnership();
    message.owner = object.owner ?? "";
    message.domains = object.domains?.map((e) => e) || [];
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

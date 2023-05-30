/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../cosmos/base/query/v1beta1/pagination";
import { Domain } from "./domain";
import { DomainOwnership } from "./domain_ownership";
import { Params } from "./params";

export const protobufPackage = "myceldomain.mycel.registry";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetDomainRequest {
  name: string;
  parent: string;
}

export interface QueryGetDomainResponse {
  domain: Domain | undefined;
}

export interface QueryAllDomainRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDomainResponse {
  domain: Domain[];
  pagination: PageResponse | undefined;
}

export interface QueryGetDomainOwnershipRequest {
  owner: string;
}

export interface QueryGetDomainOwnershipResponse {
  domainOwnership: DomainOwnership | undefined;
}

export interface QueryAllDomainOwnershipRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDomainOwnershipResponse {
  domainOwnership: DomainOwnership[];
  pagination: PageResponse | undefined;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryGetDomainRequest(): QueryGetDomainRequest {
  return { name: "", parent: "" };
}

export const QueryGetDomainRequest = {
  encode(message: QueryGetDomainRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.parent !== "") {
      writer.uint32(18).string(message.parent);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDomainRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDomainRequest();
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

  fromJSON(object: any): QueryGetDomainRequest {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      parent: isSet(object.parent) ? String(object.parent) : "",
    };
  },

  toJSON(message: QueryGetDomainRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.parent !== undefined && (obj.parent = message.parent);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDomainRequest>, I>>(object: I): QueryGetDomainRequest {
    const message = createBaseQueryGetDomainRequest();
    message.name = object.name ?? "";
    message.parent = object.parent ?? "";
    return message;
  },
};

function createBaseQueryGetDomainResponse(): QueryGetDomainResponse {
  return { domain: undefined };
}

export const QueryGetDomainResponse = {
  encode(message: QueryGetDomainResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.domain !== undefined) {
      Domain.encode(message.domain, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDomainResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDomainResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.domain = Domain.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDomainResponse {
    return { domain: isSet(object.domain) ? Domain.fromJSON(object.domain) : undefined };
  },

  toJSON(message: QueryGetDomainResponse): unknown {
    const obj: any = {};
    message.domain !== undefined && (obj.domain = message.domain ? Domain.toJSON(message.domain) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDomainResponse>, I>>(object: I): QueryGetDomainResponse {
    const message = createBaseQueryGetDomainResponse();
    message.domain = (object.domain !== undefined && object.domain !== null)
      ? Domain.fromPartial(object.domain)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDomainRequest(): QueryAllDomainRequest {
  return { pagination: undefined };
}

export const QueryAllDomainRequest = {
  encode(message: QueryAllDomainRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDomainRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDomainRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllDomainRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllDomainRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDomainRequest>, I>>(object: I): QueryAllDomainRequest {
    const message = createBaseQueryAllDomainRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDomainResponse(): QueryAllDomainResponse {
  return { domain: [], pagination: undefined };
}

export const QueryAllDomainResponse = {
  encode(message: QueryAllDomainResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.domain) {
      Domain.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDomainResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDomainResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.domain.push(Domain.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllDomainResponse {
    return {
      domain: Array.isArray(object?.domain) ? object.domain.map((e: any) => Domain.fromJSON(e)) : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDomainResponse): unknown {
    const obj: any = {};
    if (message.domain) {
      obj.domain = message.domain.map((e) => e ? Domain.toJSON(e) : undefined);
    } else {
      obj.domain = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDomainResponse>, I>>(object: I): QueryAllDomainResponse {
    const message = createBaseQueryAllDomainResponse();
    message.domain = object.domain?.map((e) => Domain.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryGetDomainOwnershipRequest(): QueryGetDomainOwnershipRequest {
  return { owner: "" };
}

export const QueryGetDomainOwnershipRequest = {
  encode(message: QueryGetDomainOwnershipRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.owner !== "") {
      writer.uint32(10).string(message.owner);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDomainOwnershipRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDomainOwnershipRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDomainOwnershipRequest {
    return { owner: isSet(object.owner) ? String(object.owner) : "" };
  },

  toJSON(message: QueryGetDomainOwnershipRequest): unknown {
    const obj: any = {};
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDomainOwnershipRequest>, I>>(
    object: I,
  ): QueryGetDomainOwnershipRequest {
    const message = createBaseQueryGetDomainOwnershipRequest();
    message.owner = object.owner ?? "";
    return message;
  },
};

function createBaseQueryGetDomainOwnershipResponse(): QueryGetDomainOwnershipResponse {
  return { domainOwnership: undefined };
}

export const QueryGetDomainOwnershipResponse = {
  encode(message: QueryGetDomainOwnershipResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.domainOwnership !== undefined) {
      DomainOwnership.encode(message.domainOwnership, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryGetDomainOwnershipResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryGetDomainOwnershipResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.domainOwnership = DomainOwnership.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetDomainOwnershipResponse {
    return {
      domainOwnership: isSet(object.domainOwnership) ? DomainOwnership.fromJSON(object.domainOwnership) : undefined,
    };
  },

  toJSON(message: QueryGetDomainOwnershipResponse): unknown {
    const obj: any = {};
    message.domainOwnership !== undefined
      && (obj.domainOwnership = message.domainOwnership ? DomainOwnership.toJSON(message.domainOwnership) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryGetDomainOwnershipResponse>, I>>(
    object: I,
  ): QueryGetDomainOwnershipResponse {
    const message = createBaseQueryGetDomainOwnershipResponse();
    message.domainOwnership = (object.domainOwnership !== undefined && object.domainOwnership !== null)
      ? DomainOwnership.fromPartial(object.domainOwnership)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDomainOwnershipRequest(): QueryAllDomainOwnershipRequest {
  return { pagination: undefined };
}

export const QueryAllDomainOwnershipRequest = {
  encode(message: QueryAllDomainOwnershipRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDomainOwnershipRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDomainOwnershipRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllDomainOwnershipRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllDomainOwnershipRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageRequest.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDomainOwnershipRequest>, I>>(
    object: I,
  ): QueryAllDomainOwnershipRequest {
    const message = createBaseQueryAllDomainOwnershipRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDomainOwnershipResponse(): QueryAllDomainOwnershipResponse {
  return { domainOwnership: [], pagination: undefined };
}

export const QueryAllDomainOwnershipResponse = {
  encode(message: QueryAllDomainOwnershipResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.domainOwnership) {
      DomainOwnership.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDomainOwnershipResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDomainOwnershipResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.domainOwnership.push(DomainOwnership.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllDomainOwnershipResponse {
    return {
      domainOwnership: Array.isArray(object?.domainOwnership)
        ? object.domainOwnership.map((e: any) => DomainOwnership.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDomainOwnershipResponse): unknown {
    const obj: any = {};
    if (message.domainOwnership) {
      obj.domainOwnership = message.domainOwnership.map((e) => e ? DomainOwnership.toJSON(e) : undefined);
    } else {
      obj.domainOwnership = [];
    }
    message.pagination !== undefined
      && (obj.pagination = message.pagination ? PageResponse.toJSON(message.pagination) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryAllDomainOwnershipResponse>, I>>(
    object: I,
  ): QueryAllDomainOwnershipResponse {
    const message = createBaseQueryAllDomainOwnershipResponse();
    message.domainOwnership = object.domainOwnership?.map((e) => DomainOwnership.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a list of Domain items. */
  Domain(request: QueryGetDomainRequest): Promise<QueryGetDomainResponse>;
  DomainAll(request: QueryAllDomainRequest): Promise<QueryAllDomainResponse>;
  /** Queries a list of DomainOwnership items. */
  DomainOwnership(request: QueryGetDomainOwnershipRequest): Promise<QueryGetDomainOwnershipResponse>;
  DomainOwnershipAll(request: QueryAllDomainOwnershipRequest): Promise<QueryAllDomainOwnershipResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Domain = this.Domain.bind(this);
    this.DomainAll = this.DomainAll.bind(this);
    this.DomainOwnership = this.DomainOwnership.bind(this);
    this.DomainOwnershipAll = this.DomainOwnershipAll.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("myceldomain.mycel.registry.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Domain(request: QueryGetDomainRequest): Promise<QueryGetDomainResponse> {
    const data = QueryGetDomainRequest.encode(request).finish();
    const promise = this.rpc.request("myceldomain.mycel.registry.Query", "Domain", data);
    return promise.then((data) => QueryGetDomainResponse.decode(new _m0.Reader(data)));
  }

  DomainAll(request: QueryAllDomainRequest): Promise<QueryAllDomainResponse> {
    const data = QueryAllDomainRequest.encode(request).finish();
    const promise = this.rpc.request("myceldomain.mycel.registry.Query", "DomainAll", data);
    return promise.then((data) => QueryAllDomainResponse.decode(new _m0.Reader(data)));
  }

  DomainOwnership(request: QueryGetDomainOwnershipRequest): Promise<QueryGetDomainOwnershipResponse> {
    const data = QueryGetDomainOwnershipRequest.encode(request).finish();
    const promise = this.rpc.request("myceldomain.mycel.registry.Query", "DomainOwnership", data);
    return promise.then((data) => QueryGetDomainOwnershipResponse.decode(new _m0.Reader(data)));
  }

  DomainOwnershipAll(request: QueryAllDomainOwnershipRequest): Promise<QueryAllDomainOwnershipResponse> {
    const data = QueryAllDomainOwnershipRequest.encode(request).finish();
    const promise = this.rpc.request("myceldomain.mycel.registry.Query", "DomainOwnershipAll", data);
    return promise.then((data) => QueryAllDomainOwnershipResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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

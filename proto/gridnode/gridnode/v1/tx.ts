/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { Params } from "./params";
import Long = require("long");

export const protobufPackage = "gridnode.gridnode.v1";

export interface Delegation {
  delegatorAddress: string;
  amount: number;
}

/** MsgGridnodeDelegate is the request type for the Msg/DelegateTokens RPC method. */
export interface MsgGridnodeDelegate {
  delegatorAddress: string;
  amount: number;
  /** Timestamp in Unix format or */
  timestamp: number;
  /** unique identifier UUID */
  uniqueId: string;
}

/** MsgGridnodeDelegateResponse is the response type for the Msg/DelegateTokens RPC method. */
export interface MsgGridnodeDelegateResponse {
  /** Transaction hash of the delegated tokens operation */
  txHash: string;
  /** Status of the operation, e.g., "success" or "failure" */
  status: string;
}

export interface MsgGridnodeUndelegate {
  delegatorAddress: string;
  amount: number;
  /** Timestamp in Unix format or */
  timestamp: number;
  /** unique identifier UUID */
  uniqueId: string;
}

export interface MsgGridnodeUndelegateResponse {
  /** Transaction hash of the delegated tokens operation */
  txHash: string;
  /** Status of the operation, e.g., "success" or "failure" */
  status: string;
}

/** MsgUpdateParams is the Msg/UpdateParams request type. */
export interface MsgUpdateParams {
  /** authority is the address that controls the module (defaults to x/gov unless overwritten). */
  authority: string;
  /**
   * params defines the module parameters to update.
   *
   * NOTE: All parameters must be supplied.
   */
  params: Params | undefined;
}

/**
 * MsgUpdateParamsResponse defines the response structure for executing a
 * MsgUpdateParams message.
 */
export interface MsgUpdateParamsResponse {
}

function createBaseDelegation(): Delegation {
  return { delegatorAddress: "", amount: 0 };
}

export const Delegation = {
  encode(message: Delegation, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.delegatorAddress !== "") {
      writer.uint32(10).string(message.delegatorAddress);
    }
    if (message.amount !== 0) {
      writer.uint32(16).int64(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Delegation {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDelegation();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.delegatorAddress = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.amount = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Delegation {
    return {
      delegatorAddress: isSet(object.delegatorAddress) ? globalThis.String(object.delegatorAddress) : "",
      amount: isSet(object.amount) ? globalThis.Number(object.amount) : 0,
    };
  },

  toJSON(message: Delegation): unknown {
    const obj: any = {};
    if (message.delegatorAddress !== "") {
      obj.delegatorAddress = message.delegatorAddress;
    }
    if (message.amount !== 0) {
      obj.amount = Math.round(message.amount);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<Delegation>, I>>(base?: I): Delegation {
    return Delegation.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<Delegation>, I>>(object: I): Delegation {
    const message = createBaseDelegation();
    message.delegatorAddress = object.delegatorAddress ?? "";
    message.amount = object.amount ?? 0;
    return message;
  },
};

function createBaseMsgGridnodeDelegate(): MsgGridnodeDelegate {
  return { delegatorAddress: "", amount: 0, timestamp: 0, uniqueId: "" };
}

export const MsgGridnodeDelegate = {
  encode(message: MsgGridnodeDelegate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.delegatorAddress !== "") {
      writer.uint32(10).string(message.delegatorAddress);
    }
    if (message.amount !== 0) {
      writer.uint32(16).int64(message.amount);
    }
    if (message.timestamp !== 0) {
      writer.uint32(24).int64(message.timestamp);
    }
    if (message.uniqueId !== "") {
      writer.uint32(34).string(message.uniqueId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgGridnodeDelegate {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgGridnodeDelegate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.delegatorAddress = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.amount = longToNumber(reader.int64() as Long);
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.timestamp = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.uniqueId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgGridnodeDelegate {
    return {
      delegatorAddress: isSet(object.delegatorAddress) ? globalThis.String(object.delegatorAddress) : "",
      amount: isSet(object.amount) ? globalThis.Number(object.amount) : 0,
      timestamp: isSet(object.timestamp) ? globalThis.Number(object.timestamp) : 0,
      uniqueId: isSet(object.uniqueId) ? globalThis.String(object.uniqueId) : "",
    };
  },

  toJSON(message: MsgGridnodeDelegate): unknown {
    const obj: any = {};
    if (message.delegatorAddress !== "") {
      obj.delegatorAddress = message.delegatorAddress;
    }
    if (message.amount !== 0) {
      obj.amount = Math.round(message.amount);
    }
    if (message.timestamp !== 0) {
      obj.timestamp = Math.round(message.timestamp);
    }
    if (message.uniqueId !== "") {
      obj.uniqueId = message.uniqueId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgGridnodeDelegate>, I>>(base?: I): MsgGridnodeDelegate {
    return MsgGridnodeDelegate.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgGridnodeDelegate>, I>>(object: I): MsgGridnodeDelegate {
    const message = createBaseMsgGridnodeDelegate();
    message.delegatorAddress = object.delegatorAddress ?? "";
    message.amount = object.amount ?? 0;
    message.timestamp = object.timestamp ?? 0;
    message.uniqueId = object.uniqueId ?? "";
    return message;
  },
};

function createBaseMsgGridnodeDelegateResponse(): MsgGridnodeDelegateResponse {
  return { txHash: "", status: "" };
}

export const MsgGridnodeDelegateResponse = {
  encode(message: MsgGridnodeDelegateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.txHash !== "") {
      writer.uint32(10).string(message.txHash);
    }
    if (message.status !== "") {
      writer.uint32(18).string(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgGridnodeDelegateResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgGridnodeDelegateResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.txHash = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.status = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgGridnodeDelegateResponse {
    return {
      txHash: isSet(object.txHash) ? globalThis.String(object.txHash) : "",
      status: isSet(object.status) ? globalThis.String(object.status) : "",
    };
  },

  toJSON(message: MsgGridnodeDelegateResponse): unknown {
    const obj: any = {};
    if (message.txHash !== "") {
      obj.txHash = message.txHash;
    }
    if (message.status !== "") {
      obj.status = message.status;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgGridnodeDelegateResponse>, I>>(base?: I): MsgGridnodeDelegateResponse {
    return MsgGridnodeDelegateResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgGridnodeDelegateResponse>, I>>(object: I): MsgGridnodeDelegateResponse {
    const message = createBaseMsgGridnodeDelegateResponse();
    message.txHash = object.txHash ?? "";
    message.status = object.status ?? "";
    return message;
  },
};

function createBaseMsgGridnodeUndelegate(): MsgGridnodeUndelegate {
  return { delegatorAddress: "", amount: 0, timestamp: 0, uniqueId: "" };
}

export const MsgGridnodeUndelegate = {
  encode(message: MsgGridnodeUndelegate, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.delegatorAddress !== "") {
      writer.uint32(10).string(message.delegatorAddress);
    }
    if (message.amount !== 0) {
      writer.uint32(16).int64(message.amount);
    }
    if (message.timestamp !== 0) {
      writer.uint32(24).int64(message.timestamp);
    }
    if (message.uniqueId !== "") {
      writer.uint32(34).string(message.uniqueId);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgGridnodeUndelegate {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgGridnodeUndelegate();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.delegatorAddress = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.amount = longToNumber(reader.int64() as Long);
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.timestamp = longToNumber(reader.int64() as Long);
          continue;
        case 4:
          if (tag !== 34) {
            break;
          }

          message.uniqueId = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgGridnodeUndelegate {
    return {
      delegatorAddress: isSet(object.delegatorAddress) ? globalThis.String(object.delegatorAddress) : "",
      amount: isSet(object.amount) ? globalThis.Number(object.amount) : 0,
      timestamp: isSet(object.timestamp) ? globalThis.Number(object.timestamp) : 0,
      uniqueId: isSet(object.uniqueId) ? globalThis.String(object.uniqueId) : "",
    };
  },

  toJSON(message: MsgGridnodeUndelegate): unknown {
    const obj: any = {};
    if (message.delegatorAddress !== "") {
      obj.delegatorAddress = message.delegatorAddress;
    }
    if (message.amount !== 0) {
      obj.amount = Math.round(message.amount);
    }
    if (message.timestamp !== 0) {
      obj.timestamp = Math.round(message.timestamp);
    }
    if (message.uniqueId !== "") {
      obj.uniqueId = message.uniqueId;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgGridnodeUndelegate>, I>>(base?: I): MsgGridnodeUndelegate {
    return MsgGridnodeUndelegate.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgGridnodeUndelegate>, I>>(object: I): MsgGridnodeUndelegate {
    const message = createBaseMsgGridnodeUndelegate();
    message.delegatorAddress = object.delegatorAddress ?? "";
    message.amount = object.amount ?? 0;
    message.timestamp = object.timestamp ?? 0;
    message.uniqueId = object.uniqueId ?? "";
    return message;
  },
};

function createBaseMsgGridnodeUndelegateResponse(): MsgGridnodeUndelegateResponse {
  return { txHash: "", status: "" };
}

export const MsgGridnodeUndelegateResponse = {
  encode(message: MsgGridnodeUndelegateResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.txHash !== "") {
      writer.uint32(10).string(message.txHash);
    }
    if (message.status !== "") {
      writer.uint32(18).string(message.status);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgGridnodeUndelegateResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgGridnodeUndelegateResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.txHash = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.status = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgGridnodeUndelegateResponse {
    return {
      txHash: isSet(object.txHash) ? globalThis.String(object.txHash) : "",
      status: isSet(object.status) ? globalThis.String(object.status) : "",
    };
  },

  toJSON(message: MsgGridnodeUndelegateResponse): unknown {
    const obj: any = {};
    if (message.txHash !== "") {
      obj.txHash = message.txHash;
    }
    if (message.status !== "") {
      obj.status = message.status;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgGridnodeUndelegateResponse>, I>>(base?: I): MsgGridnodeUndelegateResponse {
    return MsgGridnodeUndelegateResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgGridnodeUndelegateResponse>, I>>(
    object: I,
  ): MsgGridnodeUndelegateResponse {
    const message = createBaseMsgGridnodeUndelegateResponse();
    message.txHash = object.txHash ?? "";
    message.status = object.status ?? "";
    return message;
  },
};

function createBaseMsgUpdateParams(): MsgUpdateParams {
  return { authority: "", params: undefined };
}

export const MsgUpdateParams = {
  encode(message: MsgUpdateParams, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.authority !== "") {
      writer.uint32(10).string(message.authority);
    }
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateParams {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.authority = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.params = Params.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): MsgUpdateParams {
    return {
      authority: isSet(object.authority) ? globalThis.String(object.authority) : "",
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
    };
  },

  toJSON(message: MsgUpdateParams): unknown {
    const obj: any = {};
    if (message.authority !== "") {
      obj.authority = message.authority;
    }
    if (message.params !== undefined) {
      obj.params = Params.toJSON(message.params);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateParams>, I>>(base?: I): MsgUpdateParams {
    return MsgUpdateParams.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateParams>, I>>(object: I): MsgUpdateParams {
    const message = createBaseMsgUpdateParams();
    message.authority = object.authority ?? "";
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseMsgUpdateParamsResponse(): MsgUpdateParamsResponse {
  return {};
}

export const MsgUpdateParamsResponse = {
  encode(_: MsgUpdateParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgUpdateParamsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgUpdateParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(_: any): MsgUpdateParamsResponse {
    return {};
  },

  toJSON(_: MsgUpdateParamsResponse): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<MsgUpdateParamsResponse>, I>>(base?: I): MsgUpdateParamsResponse {
    return MsgUpdateParamsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<MsgUpdateParamsResponse>, I>>(_: I): MsgUpdateParamsResponse {
    const message = createBaseMsgUpdateParamsResponse();
    return message;
  },
};

/** Msg defines the Msg service for gridnode module. */
export interface Msg {
  DelegateTokens(request: MsgGridnodeDelegate): Promise<MsgGridnodeDelegateResponse>;
  UndelegateTokens(request: MsgGridnodeUndelegate): Promise<MsgGridnodeUndelegateResponse>;
  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse>;
}

export const MsgServiceName = "gridnode.gridnode.v1.Msg";
export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || MsgServiceName;
    this.rpc = rpc;
    this.DelegateTokens = this.DelegateTokens.bind(this);
    this.UndelegateTokens = this.UndelegateTokens.bind(this);
    this.UpdateParams = this.UpdateParams.bind(this);
  }
  DelegateTokens(request: MsgGridnodeDelegate): Promise<MsgGridnodeDelegateResponse> {
    const data = MsgGridnodeDelegate.encode(request).finish();
    const promise = this.rpc.request(this.service, "DelegateTokens", data);
    return promise.then((data) => MsgGridnodeDelegateResponse.decode(_m0.Reader.create(data)));
  }

  UndelegateTokens(request: MsgGridnodeUndelegate): Promise<MsgGridnodeUndelegateResponse> {
    const data = MsgGridnodeUndelegate.encode(request).finish();
    const promise = this.rpc.request(this.service, "UndelegateTokens", data);
    return promise.then((data) => MsgGridnodeUndelegateResponse.decode(_m0.Reader.create(data)));
  }

  UpdateParams(request: MsgUpdateParams): Promise<MsgUpdateParamsResponse> {
    const data = MsgUpdateParams.encode(request).finish();
    const promise = this.rpc.request(this.service, "UpdateParams", data);
    return promise.then((data) => MsgUpdateParamsResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends globalThis.Array<infer U> ? globalThis.Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(globalThis.Number.MAX_SAFE_INTEGER)) {
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

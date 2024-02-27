/* eslint-disable */
import * as _m0 from "protobufjs/minimal";
import { PageRequest, PageResponse } from "../../../cosmos/base/query/v1beta1/pagination";
import { Params } from "./params";
import Long = require("long");

export const protobufPackage = "gridnode.gridnode.v1";

export interface UnbondingEntry {
  /** This remains a string as it represents an address */
  account: string;
  /** Use the Coin type from Cosmos SDK */
  amount: number;
  /** Use the Timestamp type for time values */
  completionTime: number;
}

export interface SimpleUnbondingEntry {
  amount: number;
  completionTime: number;
}

export interface DelegationInfo {
  account: string;
  delegatedAmount: number;
  /** New field for unbonding entries */
  unbondingEntries: SimpleUnbondingEntry[];
}

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

/** QueryDelegatedAmountRequest is the request type for the Query/DelegatedAmount RPC method. */
export interface QueryDelegatedAmountRequest {
  delegatorAddress: string;
}

/** QueryDelegatedAmountResponse is the response type for the Query/DelegatedAmount RPC method. */
export interface QueryDelegatedAmountResponse {
  amount: number;
}

/** QueryUnbondingEntriesRequest is the request type for the Query/UnbondingEntries RPC method. */
export interface QueryUnbondingEntriesRequest {
  bondingAccountAddress: string;
  pagination: PageRequest | undefined;
}

/** QueryUnbondingEntriesResponse is the response type for the Query/UnbondingEntries RPC method. */
export interface QueryUnbondingEntriesResponse {
  unbondingEntries: UnbondingEntry[];
  pagination: PageResponse | undefined;
}

export interface QueryAllDelegationsRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllDelegationsResponse {
  delegations: DelegationInfo[];
  pagination: PageResponse | undefined;
}

function createBaseUnbondingEntry(): UnbondingEntry {
  return { account: "", amount: 0, completionTime: 0 };
}

export const UnbondingEntry = {
  encode(message: UnbondingEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== "") {
      writer.uint32(10).string(message.account);
    }
    if (message.amount !== 0) {
      writer.uint32(16).int64(message.amount);
    }
    if (message.completionTime !== 0) {
      writer.uint32(24).int64(message.completionTime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): UnbondingEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUnbondingEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.account = reader.string();
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

          message.completionTime = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): UnbondingEntry {
    return {
      account: isSet(object.account) ? globalThis.String(object.account) : "",
      amount: isSet(object.amount) ? globalThis.Number(object.amount) : 0,
      completionTime: isSet(object.completionTime) ? globalThis.Number(object.completionTime) : 0,
    };
  },

  toJSON(message: UnbondingEntry): unknown {
    const obj: any = {};
    if (message.account !== "") {
      obj.account = message.account;
    }
    if (message.amount !== 0) {
      obj.amount = Math.round(message.amount);
    }
    if (message.completionTime !== 0) {
      obj.completionTime = Math.round(message.completionTime);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<UnbondingEntry>, I>>(base?: I): UnbondingEntry {
    return UnbondingEntry.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<UnbondingEntry>, I>>(object: I): UnbondingEntry {
    const message = createBaseUnbondingEntry();
    message.account = object.account ?? "";
    message.amount = object.amount ?? 0;
    message.completionTime = object.completionTime ?? 0;
    return message;
  },
};

function createBaseSimpleUnbondingEntry(): SimpleUnbondingEntry {
  return { amount: 0, completionTime: 0 };
}

export const SimpleUnbondingEntry = {
  encode(message: SimpleUnbondingEntry, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.amount !== 0) {
      writer.uint32(8).int64(message.amount);
    }
    if (message.completionTime !== 0) {
      writer.uint32(16).int64(message.completionTime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): SimpleUnbondingEntry {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseSimpleUnbondingEntry();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.amount = longToNumber(reader.int64() as Long);
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.completionTime = longToNumber(reader.int64() as Long);
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): SimpleUnbondingEntry {
    return {
      amount: isSet(object.amount) ? globalThis.Number(object.amount) : 0,
      completionTime: isSet(object.completionTime) ? globalThis.Number(object.completionTime) : 0,
    };
  },

  toJSON(message: SimpleUnbondingEntry): unknown {
    const obj: any = {};
    if (message.amount !== 0) {
      obj.amount = Math.round(message.amount);
    }
    if (message.completionTime !== 0) {
      obj.completionTime = Math.round(message.completionTime);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<SimpleUnbondingEntry>, I>>(base?: I): SimpleUnbondingEntry {
    return SimpleUnbondingEntry.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<SimpleUnbondingEntry>, I>>(object: I): SimpleUnbondingEntry {
    const message = createBaseSimpleUnbondingEntry();
    message.amount = object.amount ?? 0;
    message.completionTime = object.completionTime ?? 0;
    return message;
  },
};

function createBaseDelegationInfo(): DelegationInfo {
  return { account: "", delegatedAmount: 0, unbondingEntries: [] };
}

export const DelegationInfo = {
  encode(message: DelegationInfo, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.account !== "") {
      writer.uint32(10).string(message.account);
    }
    if (message.delegatedAmount !== 0) {
      writer.uint32(16).int64(message.delegatedAmount);
    }
    for (const v of message.unbondingEntries) {
      SimpleUnbondingEntry.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): DelegationInfo {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseDelegationInfo();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.account = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.delegatedAmount = longToNumber(reader.int64() as Long);
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.unbondingEntries.push(SimpleUnbondingEntry.decode(reader, reader.uint32()));
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): DelegationInfo {
    return {
      account: isSet(object.account) ? globalThis.String(object.account) : "",
      delegatedAmount: isSet(object.delegatedAmount) ? globalThis.Number(object.delegatedAmount) : 0,
      unbondingEntries: globalThis.Array.isArray(object?.unbondingEntries)
        ? object.unbondingEntries.map((e: any) => SimpleUnbondingEntry.fromJSON(e))
        : [],
    };
  },

  toJSON(message: DelegationInfo): unknown {
    const obj: any = {};
    if (message.account !== "") {
      obj.account = message.account;
    }
    if (message.delegatedAmount !== 0) {
      obj.delegatedAmount = Math.round(message.delegatedAmount);
    }
    if (message.unbondingEntries?.length) {
      obj.unbondingEntries = message.unbondingEntries.map((e) => SimpleUnbondingEntry.toJSON(e));
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<DelegationInfo>, I>>(base?: I): DelegationInfo {
    return DelegationInfo.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<DelegationInfo>, I>>(object: I): DelegationInfo {
    const message = createBaseDelegationInfo();
    message.account = object.account ?? "";
    message.delegatedAmount = object.delegatedAmount ?? 0;
    message.unbondingEntries = object.unbondingEntries?.map((e) => SimpleUnbondingEntry.fromPartial(e)) || [];
    return message;
  },
};

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(base?: I): QueryParamsRequest {
    return QueryParamsRequest.fromPartial(base ?? ({} as any));
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
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
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

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    if (message.params !== undefined) {
      obj.params = Params.toJSON(message.params);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(base?: I): QueryParamsResponse {
    return QueryParamsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQueryDelegatedAmountRequest(): QueryDelegatedAmountRequest {
  return { delegatorAddress: "" };
}

export const QueryDelegatedAmountRequest = {
  encode(message: QueryDelegatedAmountRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.delegatorAddress !== "") {
      writer.uint32(10).string(message.delegatorAddress);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDelegatedAmountRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDelegatedAmountRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.delegatorAddress = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryDelegatedAmountRequest {
    return { delegatorAddress: isSet(object.delegatorAddress) ? globalThis.String(object.delegatorAddress) : "" };
  },

  toJSON(message: QueryDelegatedAmountRequest): unknown {
    const obj: any = {};
    if (message.delegatorAddress !== "") {
      obj.delegatorAddress = message.delegatorAddress;
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryDelegatedAmountRequest>, I>>(base?: I): QueryDelegatedAmountRequest {
    return QueryDelegatedAmountRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryDelegatedAmountRequest>, I>>(object: I): QueryDelegatedAmountRequest {
    const message = createBaseQueryDelegatedAmountRequest();
    message.delegatorAddress = object.delegatorAddress ?? "";
    return message;
  },
};

function createBaseQueryDelegatedAmountResponse(): QueryDelegatedAmountResponse {
  return { amount: 0 };
}

export const QueryDelegatedAmountResponse = {
  encode(message: QueryDelegatedAmountResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.amount !== 0) {
      writer.uint32(8).int64(message.amount);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryDelegatedAmountResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryDelegatedAmountResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
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

  fromJSON(object: any): QueryDelegatedAmountResponse {
    return { amount: isSet(object.amount) ? globalThis.Number(object.amount) : 0 };
  },

  toJSON(message: QueryDelegatedAmountResponse): unknown {
    const obj: any = {};
    if (message.amount !== 0) {
      obj.amount = Math.round(message.amount);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryDelegatedAmountResponse>, I>>(base?: I): QueryDelegatedAmountResponse {
    return QueryDelegatedAmountResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryDelegatedAmountResponse>, I>>(object: I): QueryDelegatedAmountResponse {
    const message = createBaseQueryDelegatedAmountResponse();
    message.amount = object.amount ?? 0;
    return message;
  },
};

function createBaseQueryUnbondingEntriesRequest(): QueryUnbondingEntriesRequest {
  return { bondingAccountAddress: "", pagination: undefined };
}

export const QueryUnbondingEntriesRequest = {
  encode(message: QueryUnbondingEntriesRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.bondingAccountAddress !== "") {
      writer.uint32(10).string(message.bondingAccountAddress);
    }
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUnbondingEntriesRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUnbondingEntriesRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.bondingAccountAddress = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.pagination = PageRequest.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryUnbondingEntriesRequest {
    return {
      bondingAccountAddress: isSet(object.bondingAccountAddress) ? globalThis.String(object.bondingAccountAddress) : "",
      pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryUnbondingEntriesRequest): unknown {
    const obj: any = {};
    if (message.bondingAccountAddress !== "") {
      obj.bondingAccountAddress = message.bondingAccountAddress;
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryUnbondingEntriesRequest>, I>>(base?: I): QueryUnbondingEntriesRequest {
    return QueryUnbondingEntriesRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryUnbondingEntriesRequest>, I>>(object: I): QueryUnbondingEntriesRequest {
    const message = createBaseQueryUnbondingEntriesRequest();
    message.bondingAccountAddress = object.bondingAccountAddress ?? "";
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryUnbondingEntriesResponse(): QueryUnbondingEntriesResponse {
  return { unbondingEntries: [], pagination: undefined };
}

export const QueryUnbondingEntriesResponse = {
  encode(message: QueryUnbondingEntriesResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.unbondingEntries) {
      UnbondingEntry.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryUnbondingEntriesResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryUnbondingEntriesResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.unbondingEntries.push(UnbondingEntry.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.pagination = PageResponse.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryUnbondingEntriesResponse {
    return {
      unbondingEntries: globalThis.Array.isArray(object?.unbondingEntries)
        ? object.unbondingEntries.map((e: any) => UnbondingEntry.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryUnbondingEntriesResponse): unknown {
    const obj: any = {};
    if (message.unbondingEntries?.length) {
      obj.unbondingEntries = message.unbondingEntries.map((e) => UnbondingEntry.toJSON(e));
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageResponse.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryUnbondingEntriesResponse>, I>>(base?: I): QueryUnbondingEntriesResponse {
    return QueryUnbondingEntriesResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryUnbondingEntriesResponse>, I>>(
    object: I,
  ): QueryUnbondingEntriesResponse {
    const message = createBaseQueryUnbondingEntriesResponse();
    message.unbondingEntries = object.unbondingEntries?.map((e) => UnbondingEntry.fromPartial(e)) || [];
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageResponse.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDelegationsRequest(): QueryAllDelegationsRequest {
  return { pagination: undefined };
}

export const QueryAllDelegationsRequest = {
  encode(message: QueryAllDelegationsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDelegationsRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDelegationsRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.pagination = PageRequest.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAllDelegationsRequest {
    return { pagination: isSet(object.pagination) ? PageRequest.fromJSON(object.pagination) : undefined };
  },

  toJSON(message: QueryAllDelegationsRequest): unknown {
    const obj: any = {};
    if (message.pagination !== undefined) {
      obj.pagination = PageRequest.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllDelegationsRequest>, I>>(base?: I): QueryAllDelegationsRequest {
    return QueryAllDelegationsRequest.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllDelegationsRequest>, I>>(object: I): QueryAllDelegationsRequest {
    const message = createBaseQueryAllDelegationsRequest();
    message.pagination = (object.pagination !== undefined && object.pagination !== null)
      ? PageRequest.fromPartial(object.pagination)
      : undefined;
    return message;
  },
};

function createBaseQueryAllDelegationsResponse(): QueryAllDelegationsResponse {
  return { delegations: [], pagination: undefined };
}

export const QueryAllDelegationsResponse = {
  encode(message: QueryAllDelegationsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.delegations) {
      DelegationInfo.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(message.pagination, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryAllDelegationsResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryAllDelegationsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.delegations.push(DelegationInfo.decode(reader, reader.uint32()));
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.pagination = PageResponse.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): QueryAllDelegationsResponse {
    return {
      delegations: globalThis.Array.isArray(object?.delegations)
        ? object.delegations.map((e: any) => DelegationInfo.fromJSON(e))
        : [],
      pagination: isSet(object.pagination) ? PageResponse.fromJSON(object.pagination) : undefined,
    };
  },

  toJSON(message: QueryAllDelegationsResponse): unknown {
    const obj: any = {};
    if (message.delegations?.length) {
      obj.delegations = message.delegations.map((e) => DelegationInfo.toJSON(e));
    }
    if (message.pagination !== undefined) {
      obj.pagination = PageResponse.toJSON(message.pagination);
    }
    return obj;
  },

  create<I extends Exact<DeepPartial<QueryAllDelegationsResponse>, I>>(base?: I): QueryAllDelegationsResponse {
    return QueryAllDelegationsResponse.fromPartial(base ?? ({} as any));
  },
  fromPartial<I extends Exact<DeepPartial<QueryAllDelegationsResponse>, I>>(object: I): QueryAllDelegationsResponse {
    const message = createBaseQueryAllDelegationsResponse();
    message.delegations = object.delegations?.map((e) => DelegationInfo.fromPartial(e)) || [];
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
  /** DelegatedAmount queries the amount delegated by a specific account. */
  DelegatedAmount(request: QueryDelegatedAmountRequest): Promise<QueryDelegatedAmountResponse>;
  UnbondingEntries(request: QueryUnbondingEntriesRequest): Promise<QueryUnbondingEntriesResponse>;
  AllDelegations(request: QueryAllDelegationsRequest): Promise<QueryAllDelegationsResponse>;
}

export const QueryServiceName = "gridnode.gridnode.v1.Query";
export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || QueryServiceName;
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.DelegatedAmount = this.DelegatedAmount.bind(this);
    this.UnbondingEntries = this.UnbondingEntries.bind(this);
    this.AllDelegations = this.AllDelegations.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(_m0.Reader.create(data)));
  }

  DelegatedAmount(request: QueryDelegatedAmountRequest): Promise<QueryDelegatedAmountResponse> {
    const data = QueryDelegatedAmountRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "DelegatedAmount", data);
    return promise.then((data) => QueryDelegatedAmountResponse.decode(_m0.Reader.create(data)));
  }

  UnbondingEntries(request: QueryUnbondingEntriesRequest): Promise<QueryUnbondingEntriesResponse> {
    const data = QueryUnbondingEntriesRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "UnbondingEntries", data);
    return promise.then((data) => QueryUnbondingEntriesResponse.decode(_m0.Reader.create(data)));
  }

  AllDelegations(request: QueryAllDelegationsRequest): Promise<QueryAllDelegationsResponse> {
    const data = QueryAllDelegationsRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "AllDelegations", data);
    return promise.then((data) => QueryAllDelegationsResponse.decode(_m0.Reader.create(data)));
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

/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Loan } from "./loan";
import { Params } from "./params";
import { User } from "./user";

export const protobufPackage = "credimint.credimint";

/** GenesisState defines the credimint module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  portId: string;
  loanList: Loan[];
  loanCount: number;
  /** this line is used by starport scaffolding # genesis/proto/state */
  userList: User[];
}

function createBaseGenesisState(): GenesisState {
  return { params: undefined, portId: "", loanList: [], loanCount: 0, userList: [] };
}

export const GenesisState = {
  encode(message: GenesisState, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    if (message.portId !== "") {
      writer.uint32(18).string(message.portId);
    }
    for (const v of message.loanList) {
      Loan.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    if (message.loanCount !== 0) {
      writer.uint32(32).uint64(message.loanCount);
    }
    for (const v of message.userList) {
      User.encode(v!, writer.uint32(42).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseGenesisState();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.portId = reader.string();
          break;
        case 3:
          message.loanList.push(Loan.decode(reader, reader.uint32()));
          break;
        case 4:
          message.loanCount = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.userList.push(User.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    return {
      params: isSet(object.params) ? Params.fromJSON(object.params) : undefined,
      portId: isSet(object.portId) ? String(object.portId) : "",
      loanList: Array.isArray(object?.loanList) ? object.loanList.map((e: any) => Loan.fromJSON(e)) : [],
      loanCount: isSet(object.loanCount) ? Number(object.loanCount) : 0,
      userList: Array.isArray(object?.userList) ? object.userList.map((e: any) => User.fromJSON(e)) : [],
    };
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    message.portId !== undefined && (obj.portId = message.portId);
    if (message.loanList) {
      obj.loanList = message.loanList.map((e) => e ? Loan.toJSON(e) : undefined);
    } else {
      obj.loanList = [];
    }
    message.loanCount !== undefined && (obj.loanCount = Math.round(message.loanCount));
    if (message.userList) {
      obj.userList = message.userList.map((e) => e ? User.toJSON(e) : undefined);
    } else {
      obj.userList = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<GenesisState>, I>>(object: I): GenesisState {
    const message = createBaseGenesisState();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    message.portId = object.portId ?? "";
    message.loanList = object.loanList?.map((e) => Loan.fromPartial(e)) || [];
    message.loanCount = object.loanCount ?? 0;
    message.userList = object.userList?.map((e) => User.fromPartial(e)) || [];
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

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}

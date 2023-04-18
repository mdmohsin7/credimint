/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "credimint.credimint";

export interface MsgRequestLoan {
  creator: string;
  amount: string;
  fee: string;
  collateral: string;
  deadline: string;
}

export interface MsgRequestLoanResponse {
}

export interface MsgApproveLoan {
  creator: string;
  id: number;
}

export interface MsgApproveLoanResponse {
}

export interface MsgRepayLoan {
  creator: string;
  id: number;
  repayTime: string;
}

export interface MsgRepayLoanResponse {
}

function createBaseMsgRequestLoan(): MsgRequestLoan {
  return { creator: "", amount: "", fee: "", collateral: "", deadline: "" };
}

export const MsgRequestLoan = {
  encode(message: MsgRequestLoan, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.fee !== "") {
      writer.uint32(26).string(message.fee);
    }
    if (message.collateral !== "") {
      writer.uint32(34).string(message.collateral);
    }
    if (message.deadline !== "") {
      writer.uint32(42).string(message.deadline);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRequestLoan {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRequestLoan();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        case 3:
          message.fee = reader.string();
          break;
        case 4:
          message.collateral = reader.string();
          break;
        case 5:
          message.deadline = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRequestLoan {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      fee: isSet(object.fee) ? String(object.fee) : "",
      collateral: isSet(object.collateral) ? String(object.collateral) : "",
      deadline: isSet(object.deadline) ? String(object.deadline) : "",
    };
  },

  toJSON(message: MsgRequestLoan): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    message.fee !== undefined && (obj.fee = message.fee);
    message.collateral !== undefined && (obj.collateral = message.collateral);
    message.deadline !== undefined && (obj.deadline = message.deadline);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRequestLoan>, I>>(object: I): MsgRequestLoan {
    const message = createBaseMsgRequestLoan();
    message.creator = object.creator ?? "";
    message.amount = object.amount ?? "";
    message.fee = object.fee ?? "";
    message.collateral = object.collateral ?? "";
    message.deadline = object.deadline ?? "";
    return message;
  },
};

function createBaseMsgRequestLoanResponse(): MsgRequestLoanResponse {
  return {};
}

export const MsgRequestLoanResponse = {
  encode(_: MsgRequestLoanResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRequestLoanResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRequestLoanResponse();
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

  fromJSON(_: any): MsgRequestLoanResponse {
    return {};
  },

  toJSON(_: MsgRequestLoanResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRequestLoanResponse>, I>>(_: I): MsgRequestLoanResponse {
    const message = createBaseMsgRequestLoanResponse();
    return message;
  },
};

function createBaseMsgApproveLoan(): MsgApproveLoan {
  return { creator: "", id: 0 };
}

export const MsgApproveLoan = {
  encode(message: MsgApproveLoan, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgApproveLoan {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgApproveLoan();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgApproveLoan {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
    };
  },

  toJSON(message: MsgApproveLoan): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgApproveLoan>, I>>(object: I): MsgApproveLoan {
    const message = createBaseMsgApproveLoan();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    return message;
  },
};

function createBaseMsgApproveLoanResponse(): MsgApproveLoanResponse {
  return {};
}

export const MsgApproveLoanResponse = {
  encode(_: MsgApproveLoanResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgApproveLoanResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgApproveLoanResponse();
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

  fromJSON(_: any): MsgApproveLoanResponse {
    return {};
  },

  toJSON(_: MsgApproveLoanResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgApproveLoanResponse>, I>>(_: I): MsgApproveLoanResponse {
    const message = createBaseMsgApproveLoanResponse();
    return message;
  },
};

function createBaseMsgRepayLoan(): MsgRepayLoan {
  return { creator: "", id: 0, repayTime: "" };
}

export const MsgRepayLoan = {
  encode(message: MsgRepayLoan, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    if (message.repayTime !== "") {
      writer.uint32(26).string(message.repayTime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRepayLoan {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRepayLoan();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.repayTime = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRepayLoan {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      id: isSet(object.id) ? Number(object.id) : 0,
      repayTime: isSet(object.repayTime) ? String(object.repayTime) : "",
    };
  },

  toJSON(message: MsgRepayLoan): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.repayTime !== undefined && (obj.repayTime = message.repayTime);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRepayLoan>, I>>(object: I): MsgRepayLoan {
    const message = createBaseMsgRepayLoan();
    message.creator = object.creator ?? "";
    message.id = object.id ?? 0;
    message.repayTime = object.repayTime ?? "";
    return message;
  },
};

function createBaseMsgRepayLoanResponse(): MsgRepayLoanResponse {
  return {};
}

export const MsgRepayLoanResponse = {
  encode(_: MsgRepayLoanResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgRepayLoanResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgRepayLoanResponse();
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

  fromJSON(_: any): MsgRepayLoanResponse {
    return {};
  },

  toJSON(_: MsgRepayLoanResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgRepayLoanResponse>, I>>(_: I): MsgRepayLoanResponse {
    const message = createBaseMsgRepayLoanResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  RequestLoan(request: MsgRequestLoan): Promise<MsgRequestLoanResponse>;
  ApproveLoan(request: MsgApproveLoan): Promise<MsgApproveLoanResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  RepayLoan(request: MsgRepayLoan): Promise<MsgRepayLoanResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.RequestLoan = this.RequestLoan.bind(this);
    this.ApproveLoan = this.ApproveLoan.bind(this);
    this.RepayLoan = this.RepayLoan.bind(this);
  }
  RequestLoan(request: MsgRequestLoan): Promise<MsgRequestLoanResponse> {
    const data = MsgRequestLoan.encode(request).finish();
    const promise = this.rpc.request("credimint.credimint.Msg", "RequestLoan", data);
    return promise.then((data) => MsgRequestLoanResponse.decode(new _m0.Reader(data)));
  }

  ApproveLoan(request: MsgApproveLoan): Promise<MsgApproveLoanResponse> {
    const data = MsgApproveLoan.encode(request).finish();
    const promise = this.rpc.request("credimint.credimint.Msg", "ApproveLoan", data);
    return promise.then((data) => MsgApproveLoanResponse.decode(new _m0.Reader(data)));
  }

  RepayLoan(request: MsgRepayLoan): Promise<MsgRepayLoanResponse> {
    const data = MsgRepayLoan.encode(request).finish();
    const promise = this.rpc.request("credimint.credimint.Msg", "RepayLoan", data);
    return promise.then((data) => MsgRepayLoanResponse.decode(new _m0.Reader(data)));
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

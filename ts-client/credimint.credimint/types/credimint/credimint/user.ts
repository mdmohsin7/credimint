/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "credimint.credimint";

export interface User {
  index: string;
  creditScore: number;
  timelyPayments: number;
  defaultRate: string;
  numberOfLoans: number;
  loanDuration: number;
  numberOfLoansFunded: number;
  loanFundedDuration: number;
  collateralPercent: string;
}

function createBaseUser(): User {
  return {
    index: "",
    creditScore: 0,
    timelyPayments: 0,
    defaultRate: "",
    numberOfLoans: 0,
    loanDuration: 0,
    numberOfLoansFunded: 0,
    loanFundedDuration: 0,
    collateralPercent: "",
  };
}

export const User = {
  encode(message: User, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.creditScore !== 0) {
      writer.uint32(16).uint64(message.creditScore);
    }
    if (message.timelyPayments !== 0) {
      writer.uint32(24).uint64(message.timelyPayments);
    }
    if (message.defaultRate !== "") {
      writer.uint32(34).string(message.defaultRate);
    }
    if (message.numberOfLoans !== 0) {
      writer.uint32(40).uint64(message.numberOfLoans);
    }
    if (message.loanDuration !== 0) {
      writer.uint32(48).uint64(message.loanDuration);
    }
    if (message.numberOfLoansFunded !== 0) {
      writer.uint32(56).uint64(message.numberOfLoansFunded);
    }
    if (message.loanFundedDuration !== 0) {
      writer.uint32(64).uint64(message.loanFundedDuration);
    }
    if (message.collateralPercent !== "") {
      writer.uint32(74).string(message.collateralPercent);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): User {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseUser();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.creditScore = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.timelyPayments = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.defaultRate = reader.string();
          break;
        case 5:
          message.numberOfLoans = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.loanDuration = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.numberOfLoansFunded = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.loanFundedDuration = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.collateralPercent = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): User {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      creditScore: isSet(object.creditScore) ? Number(object.creditScore) : 0,
      timelyPayments: isSet(object.timelyPayments) ? Number(object.timelyPayments) : 0,
      defaultRate: isSet(object.defaultRate) ? String(object.defaultRate) : "",
      numberOfLoans: isSet(object.numberOfLoans) ? Number(object.numberOfLoans) : 0,
      loanDuration: isSet(object.loanDuration) ? Number(object.loanDuration) : 0,
      numberOfLoansFunded: isSet(object.numberOfLoansFunded) ? Number(object.numberOfLoansFunded) : 0,
      loanFundedDuration: isSet(object.loanFundedDuration) ? Number(object.loanFundedDuration) : 0,
      collateralPercent: isSet(object.collateralPercent) ? String(object.collateralPercent) : "",
    };
  },

  toJSON(message: User): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.creditScore !== undefined && (obj.creditScore = Math.round(message.creditScore));
    message.timelyPayments !== undefined && (obj.timelyPayments = Math.round(message.timelyPayments));
    message.defaultRate !== undefined && (obj.defaultRate = message.defaultRate);
    message.numberOfLoans !== undefined && (obj.numberOfLoans = Math.round(message.numberOfLoans));
    message.loanDuration !== undefined && (obj.loanDuration = Math.round(message.loanDuration));
    message.numberOfLoansFunded !== undefined && (obj.numberOfLoansFunded = Math.round(message.numberOfLoansFunded));
    message.loanFundedDuration !== undefined && (obj.loanFundedDuration = Math.round(message.loanFundedDuration));
    message.collateralPercent !== undefined && (obj.collateralPercent = message.collateralPercent);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<User>, I>>(object: I): User {
    const message = createBaseUser();
    message.index = object.index ?? "";
    message.creditScore = object.creditScore ?? 0;
    message.timelyPayments = object.timelyPayments ?? 0;
    message.defaultRate = object.defaultRate ?? "";
    message.numberOfLoans = object.numberOfLoans ?? 0;
    message.loanDuration = object.loanDuration ?? 0;
    message.numberOfLoansFunded = object.numberOfLoansFunded ?? 0;
    message.loanFundedDuration = object.loanFundedDuration ?? 0;
    message.collateralPercent = object.collateralPercent ?? "";
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

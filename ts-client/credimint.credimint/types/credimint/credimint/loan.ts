/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "credimint.credimint";

export interface Loan {
  id: number;
  amoount: string;
  fee: string;
  collateral: string;
  deadline: string;
  state: string;
  borrower: string;
  lender: string;
  riskLevel: string;
  borrowerScore: string;
  lenderScore: string;
}

function createBaseLoan(): Loan {
  return {
    id: 0,
    amoount: "",
    fee: "",
    collateral: "",
    deadline: "",
    state: "",
    borrower: "",
    lender: "",
    riskLevel: "",
    borrowerScore: "",
    lenderScore: "",
  };
}

export const Loan = {
  encode(message: Loan, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== 0) {
      writer.uint32(8).uint64(message.id);
    }
    if (message.amoount !== "") {
      writer.uint32(18).string(message.amoount);
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
    if (message.state !== "") {
      writer.uint32(50).string(message.state);
    }
    if (message.borrower !== "") {
      writer.uint32(58).string(message.borrower);
    }
    if (message.lender !== "") {
      writer.uint32(66).string(message.lender);
    }
    if (message.riskLevel !== "") {
      writer.uint32(74).string(message.riskLevel);
    }
    if (message.borrowerScore !== "") {
      writer.uint32(82).string(message.borrowerScore);
    }
    if (message.lenderScore !== "") {
      writer.uint32(90).string(message.lenderScore);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Loan {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseLoan();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.amoount = reader.string();
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
        case 6:
          message.state = reader.string();
          break;
        case 7:
          message.borrower = reader.string();
          break;
        case 8:
          message.lender = reader.string();
          break;
        case 9:
          message.riskLevel = reader.string();
          break;
        case 10:
          message.borrowerScore = reader.string();
          break;
        case 11:
          message.lenderScore = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Loan {
    return {
      id: isSet(object.id) ? Number(object.id) : 0,
      amoount: isSet(object.amoount) ? String(object.amoount) : "",
      fee: isSet(object.fee) ? String(object.fee) : "",
      collateral: isSet(object.collateral) ? String(object.collateral) : "",
      deadline: isSet(object.deadline) ? String(object.deadline) : "",
      state: isSet(object.state) ? String(object.state) : "",
      borrower: isSet(object.borrower) ? String(object.borrower) : "",
      lender: isSet(object.lender) ? String(object.lender) : "",
      riskLevel: isSet(object.riskLevel) ? String(object.riskLevel) : "",
      borrowerScore: isSet(object.borrowerScore) ? String(object.borrowerScore) : "",
      lenderScore: isSet(object.lenderScore) ? String(object.lenderScore) : "",
    };
  },

  toJSON(message: Loan): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = Math.round(message.id));
    message.amoount !== undefined && (obj.amoount = message.amoount);
    message.fee !== undefined && (obj.fee = message.fee);
    message.collateral !== undefined && (obj.collateral = message.collateral);
    message.deadline !== undefined && (obj.deadline = message.deadline);
    message.state !== undefined && (obj.state = message.state);
    message.borrower !== undefined && (obj.borrower = message.borrower);
    message.lender !== undefined && (obj.lender = message.lender);
    message.riskLevel !== undefined && (obj.riskLevel = message.riskLevel);
    message.borrowerScore !== undefined && (obj.borrowerScore = message.borrowerScore);
    message.lenderScore !== undefined && (obj.lenderScore = message.lenderScore);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Loan>, I>>(object: I): Loan {
    const message = createBaseLoan();
    message.id = object.id ?? 0;
    message.amoount = object.amoount ?? "";
    message.fee = object.fee ?? "";
    message.collateral = object.collateral ?? "";
    message.deadline = object.deadline ?? "";
    message.state = object.state ?? "";
    message.borrower = object.borrower ?? "";
    message.lender = object.lender ?? "";
    message.riskLevel = object.riskLevel ?? "";
    message.borrowerScore = object.borrowerScore ?? "";
    message.lenderScore = object.lenderScore ?? "";
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

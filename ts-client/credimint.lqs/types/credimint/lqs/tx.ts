/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "credimint.lqs";

export interface MsgLiquid {
  creator: string;
  amount: string;
  validator: string;
}

export interface MsgLiquidResponse {
}

function createBaseMsgLiquid(): MsgLiquid {
  return { creator: "", amount: "", validator: "" };
}

export const MsgLiquid = {
  encode(message: MsgLiquid, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.validator !== "") {
      writer.uint32(26).string(message.validator);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgLiquid {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgLiquid();
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
          message.validator = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgLiquid {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      amount: isSet(object.amount) ? String(object.amount) : "",
      validator: isSet(object.validator) ? String(object.validator) : "",
    };
  },

  toJSON(message: MsgLiquid): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    message.validator !== undefined && (obj.validator = message.validator);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgLiquid>, I>>(object: I): MsgLiquid {
    const message = createBaseMsgLiquid();
    message.creator = object.creator ?? "";
    message.amount = object.amount ?? "";
    message.validator = object.validator ?? "";
    return message;
  },
};

function createBaseMsgLiquidResponse(): MsgLiquidResponse {
  return {};
}

export const MsgLiquidResponse = {
  encode(_: MsgLiquidResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgLiquidResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgLiquidResponse();
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

  fromJSON(_: any): MsgLiquidResponse {
    return {};
  },

  toJSON(_: MsgLiquidResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgLiquidResponse>, I>>(_: I): MsgLiquidResponse {
    const message = createBaseMsgLiquidResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Liquid(request: MsgLiquid): Promise<MsgLiquidResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Liquid = this.Liquid.bind(this);
  }
  Liquid(request: MsgLiquid): Promise<MsgLiquidResponse> {
    const data = MsgLiquid.encode(request).finish();
    const promise = this.rpc.request("credimint.lqs.Msg", "Liquid", data);
    return promise.then((data) => MsgLiquidResponse.decode(new _m0.Reader(data)));
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

/* eslint-disable */
import { Reader, Writer } from 'protobufjs/minimal'

export const protobufPackage = 'dekecreate.qbitgateway.xfer'

/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgAction {
  creator: string
  id: string
  action: string
}

export interface MsgActionResponse {}

const baseMsgAction: object = { creator: '', id: '', action: '' }

export const MsgAction = {
  encode(message: MsgAction, writer: Writer = Writer.create()): Writer {
    if (message.creator !== '') {
      writer.uint32(10).string(message.creator)
    }
    if (message.id !== '') {
      writer.uint32(18).string(message.id)
    }
    if (message.action !== '') {
      writer.uint32(26).string(message.action)
    }
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgAction {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgAction } as MsgAction
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string()
          break
        case 2:
          message.id = reader.string()
          break
        case 3:
          message.action = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgAction {
    const message = { ...baseMsgAction } as MsgAction
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator)
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = String(object.id)
    } else {
      message.id = ''
    }
    if (object.action !== undefined && object.action !== null) {
      message.action = String(object.action)
    } else {
      message.action = ''
    }
    return message
  },

  toJSON(message: MsgAction): unknown {
    const obj: any = {}
    message.creator !== undefined && (obj.creator = message.creator)
    message.id !== undefined && (obj.id = message.id)
    message.action !== undefined && (obj.action = message.action)
    return obj
  },

  fromPartial(object: DeepPartial<MsgAction>): MsgAction {
    const message = { ...baseMsgAction } as MsgAction
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator
    } else {
      message.creator = ''
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id
    } else {
      message.id = ''
    }
    if (object.action !== undefined && object.action !== null) {
      message.action = object.action
    } else {
      message.action = ''
    }
    return message
  }
}

const baseMsgActionResponse: object = {}

export const MsgActionResponse = {
  encode(_: MsgActionResponse, writer: Writer = Writer.create()): Writer {
    return writer
  },

  decode(input: Reader | Uint8Array, length?: number): MsgActionResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input
    let end = length === undefined ? reader.len : reader.pos + length
    const message = { ...baseMsgActionResponse } as MsgActionResponse
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgActionResponse {
    const message = { ...baseMsgActionResponse } as MsgActionResponse
    return message
  },

  toJSON(_: MsgActionResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial(_: DeepPartial<MsgActionResponse>): MsgActionResponse {
    const message = { ...baseMsgActionResponse } as MsgActionResponse
    return message
  }
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  Action(request: MsgAction): Promise<MsgActionResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
  }
  Action(request: MsgAction): Promise<MsgActionResponse> {
    const data = MsgAction.encode(request).finish()
    const promise = this.rpc.request('dekecreate.qbitgateway.xfer.Msg', 'Action', data)
    return promise.then((data) => MsgActionResponse.decode(new Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

type Builtin = Date | Function | Uint8Array | string | number | undefined
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

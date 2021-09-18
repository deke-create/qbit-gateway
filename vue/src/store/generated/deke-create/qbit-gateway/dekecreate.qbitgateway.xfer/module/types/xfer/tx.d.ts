import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "dekecreate.qbitgateway.xfer";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgAction {
    creator: string;
    action: Uint8Array;
}
export interface MsgActionResponse {
    id: string;
}
export declare const MsgAction: {
    encode(message: MsgAction, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAction;
    fromJSON(object: any): MsgAction;
    toJSON(message: MsgAction): unknown;
    fromPartial(object: DeepPartial<MsgAction>): MsgAction;
};
export declare const MsgActionResponse: {
    encode(message: MsgActionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgActionResponse;
    fromJSON(object: any): MsgActionResponse;
    toJSON(message: MsgActionResponse): unknown;
    fromPartial(object: DeepPartial<MsgActionResponse>): MsgActionResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    Action(request: MsgAction): Promise<MsgActionResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    Action(request: MsgAction): Promise<MsgActionResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

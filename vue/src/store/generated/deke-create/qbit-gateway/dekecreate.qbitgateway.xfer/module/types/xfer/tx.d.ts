import { Reader, Writer } from 'protobufjs/minimal';
export declare const protobufPackage = "dekecreate.qbitgateway.xfer";
/** this line is used by starport scaffolding # proto/tx/message */
export interface MsgAction {
    creator: string;
    id: string;
    action: string;
}
export interface MsgActionResponse {
}
export declare const MsgAction: {
    encode(message: MsgAction, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgAction;
    fromJSON(object: any): MsgAction;
    toJSON(message: MsgAction): unknown;
    fromPartial(object: DeepPartial<MsgAction>): MsgAction;
};
export declare const MsgActionResponse: {
    encode(_: MsgActionResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgActionResponse;
    fromJSON(_: any): MsgActionResponse;
    toJSON(_: MsgActionResponse): unknown;
    fromPartial(_: DeepPartial<MsgActionResponse>): MsgActionResponse;
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

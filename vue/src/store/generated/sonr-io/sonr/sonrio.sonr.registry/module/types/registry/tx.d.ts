import { Reader, Writer } from "protobufjs/minimal";
export declare const protobufPackage = "sonrio.sonr.registry";
export interface MsgRegisterName {
    creator: string;
    name: string;
    fingerprint: string;
}
export interface MsgRegisterNameResponse {
}
export declare const MsgRegisterName: {
    encode(message: MsgRegisterName, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterName;
    fromJSON(object: any): MsgRegisterName;
    toJSON(message: MsgRegisterName): unknown;
    fromPartial(object: DeepPartial<MsgRegisterName>): MsgRegisterName;
};
export declare const MsgRegisterNameResponse: {
    encode(_: MsgRegisterNameResponse, writer?: Writer): Writer;
    decode(input: Reader | Uint8Array, length?: number): MsgRegisterNameResponse;
    fromJSON(_: any): MsgRegisterNameResponse;
    toJSON(_: MsgRegisterNameResponse): unknown;
    fromPartial(_: DeepPartial<MsgRegisterNameResponse>): MsgRegisterNameResponse;
};
/** Msg defines the Msg service. */
export interface Msg {
    /** this line is used by starport scaffolding # proto/tx/rpc */
    RegisterName(request: MsgRegisterName): Promise<MsgRegisterNameResponse>;
}
export declare class MsgClientImpl implements Msg {
    private readonly rpc;
    constructor(rpc: Rpc);
    RegisterName(request: MsgRegisterName): Promise<MsgRegisterNameResponse>;
}
interface Rpc {
    request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}
declare type Builtin = Date | Function | Uint8Array | string | number | undefined;
export declare type DeepPartial<T> = T extends Builtin ? T : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>> : T extends {} ? {
    [K in keyof T]?: DeepPartial<T[K]>;
} : Partial<T>;
export {};

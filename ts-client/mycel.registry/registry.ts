import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRegisterDomain } from "./types/mycel/registry/tx";
import { MsgUpdateWalletRecord } from "./types/mycel/registry/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/mycel.registry.MsgRegisterDomain", MsgRegisterDomain],
    ["/mycel.registry.MsgUpdateWalletRecord", MsgUpdateWalletRecord],
    
];

export { msgTypes }
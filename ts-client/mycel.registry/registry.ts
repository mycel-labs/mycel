import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateDnsRecord } from "./types/mycel/registry/tx";
import { MsgUpdateWalletRecord } from "./types/mycel/registry/tx";
import { MsgRegisterDomain } from "./types/mycel/registry/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/mycel.registry.MsgUpdateDnsRecord", MsgUpdateDnsRecord],
    ["/mycel.registry.MsgUpdateWalletRecord", MsgUpdateWalletRecord],
    ["/mycel.registry.MsgRegisterDomain", MsgRegisterDomain],
    
];

export { msgTypes }
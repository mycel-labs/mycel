import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateWalletRecord } from "./types/mycel/registry/tx";
import { MsgRegisterDomain } from "./types/mycel/registry/tx";
import { MsgUpdateDnsRecord } from "./types/mycel/registry/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/mycel.registry.MsgUpdateWalletRecord", MsgUpdateWalletRecord],
    ["/mycel.registry.MsgRegisterDomain", MsgRegisterDomain],
    ["/mycel.registry.MsgUpdateDnsRecord", MsgUpdateDnsRecord],
    
];

export { msgTypes }
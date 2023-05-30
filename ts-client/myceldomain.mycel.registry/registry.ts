import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgUpdateWalletRecord } from "./types/mycel/registry/tx";
import { MsgUpdateDnsRecord } from "./types/mycel/registry/tx";
import { MsgRegisterDomain } from "./types/mycel/registry/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/myceldomain.mycel.registry.MsgUpdateWalletRecord", MsgUpdateWalletRecord],
    ["/myceldomain.mycel.registry.MsgUpdateDnsRecord", MsgUpdateDnsRecord],
    ["/myceldomain.mycel.registry.MsgRegisterDomain", MsgRegisterDomain],
    
];

export { msgTypes }
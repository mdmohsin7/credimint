import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgApproveLoan } from "./types/credimint/credimint/tx";
import { MsgRequestLoan } from "./types/credimint/credimint/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/credimint.credimint.MsgApproveLoan", MsgApproveLoan],
    ["/credimint.credimint.MsgRequestLoan", MsgRequestLoan],
    
];

export { msgTypes }
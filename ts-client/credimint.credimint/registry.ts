import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgLiquidateLoan } from "./types/credimint/credimint/tx";
import { MsgRequestLoan } from "./types/credimint/credimint/tx";
import { MsgRepayLoan } from "./types/credimint/credimint/tx";
import { MsgApproveLoan } from "./types/credimint/credimint/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/credimint.credimint.MsgLiquidateLoan", MsgLiquidateLoan],
    ["/credimint.credimint.MsgRequestLoan", MsgRequestLoan],
    ["/credimint.credimint.MsgRepayLoan", MsgRepayLoan],
    ["/credimint.credimint.MsgApproveLoan", MsgApproveLoan],
    
];

export { msgTypes }
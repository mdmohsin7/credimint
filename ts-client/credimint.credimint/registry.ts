import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgCancelLoan } from "./types/credimint/credimint/tx";
import { MsgLiquidateLoan } from "./types/credimint/credimint/tx";
import { MsgRepayLoan } from "./types/credimint/credimint/tx";
import { MsgApproveLoan } from "./types/credimint/credimint/tx";
import { MsgLiquidStake } from "./types/credimint/credimint/tx";
import { MsgRequestLoan } from "./types/credimint/credimint/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/credimint.credimint.MsgCancelLoan", MsgCancelLoan],
    ["/credimint.credimint.MsgLiquidateLoan", MsgLiquidateLoan],
    ["/credimint.credimint.MsgRepayLoan", MsgRepayLoan],
    ["/credimint.credimint.MsgApproveLoan", MsgApproveLoan],
    ["/credimint.credimint.MsgLiquidStake", MsgLiquidStake],
    ["/credimint.credimint.MsgRequestLoan", MsgRequestLoan],
    
];

export { msgTypes }
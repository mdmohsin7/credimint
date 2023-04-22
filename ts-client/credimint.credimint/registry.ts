import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgLiquidateLoan } from "./types/credimint/credimint/tx";
import { MsgApproveLoan } from "./types/credimint/credimint/tx";
import { MsgRepayLoan } from "./types/credimint/credimint/tx";
import { MsgLiquidStake } from "./types/credimint/credimint/tx";
import { MsgRequestLoan } from "./types/credimint/credimint/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/credimint.credimint.MsgLiquidateLoan", MsgLiquidateLoan],
    ["/credimint.credimint.MsgApproveLoan", MsgApproveLoan],
    ["/credimint.credimint.MsgRepayLoan", MsgRepayLoan],
    ["/credimint.credimint.MsgLiquidStake", MsgLiquidStake],
    ["/credimint.credimint.MsgRequestLoan", MsgRequestLoan],
    
];

export { msgTypes }
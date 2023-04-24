import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgLiquidStake } from "./types/credimint/credimint/tx";
import { MsgRepayLoan } from "./types/credimint/credimint/tx";
import { MsgRequestLoan } from "./types/credimint/credimint/tx";
import { MsgApproveLoan } from "./types/credimint/credimint/tx";
import { MsgLiquidateLoan } from "./types/credimint/credimint/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/credimint.credimint.MsgLiquidStake", MsgLiquidStake],
    ["/credimint.credimint.MsgRepayLoan", MsgRepayLoan],
    ["/credimint.credimint.MsgRequestLoan", MsgRequestLoan],
    ["/credimint.credimint.MsgApproveLoan", MsgApproveLoan],
    ["/credimint.credimint.MsgLiquidateLoan", MsgLiquidateLoan],
    
];

export { msgTypes }
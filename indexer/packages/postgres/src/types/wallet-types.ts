/* ------- WALLET TYPES ------- */

export interface WalletCreateObject {
  address: string,
  totalTradingRewards: string,
}

export enum WalletColumns {
  address = 'address',
  totalTradingRewards = 'totalTradingRewards',
}

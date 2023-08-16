package constants

import (
	"github.com/dydxprotocol/v4/dtypes"
	perptypes "github.com/dydxprotocol/v4/x/perpetuals/types"
)

func init() {
	_ = TestTxBuilder.SetMsgs(EmptyMsgAddPremiumVotes)
	EmptyMsgAddPremiumVotesTxBytes, _ = TestEncodingCfg.TxConfig.TxEncoder()(TestTxBuilder.GetTx())

	_ = TestTxBuilder.SetMsgs(ValidMsgAddPremiumVotes)
	ValidMsgAddPremiumVotesTxBytes, _ = TestEncodingCfg.TxConfig.TxEncoder()(TestTxBuilder.GetTx())

	_ = TestTxBuilder.SetMsgs(InvalidMsgAddPremiumVotes)
	InvalidMsgAddPremiumVotesTxBytes, _ = TestEncodingCfg.TxConfig.TxEncoder()(TestTxBuilder.GetTx())
}

// LiquidityTier objects.
var LiquidityTiers = []perptypes.LiquidityTier{
	{
		Id:                     0,
		Name:                   "0",
		InitialMarginPpm:       1_000_000,
		MaintenanceFractionPpm: 1_000_000,
		BasePositionNotional:   1_000_000,
		ImpactNotional:         500_000_000,
	},
	{
		Id:                     1,
		Name:                   "1",
		InitialMarginPpm:       1_000_000,
		MaintenanceFractionPpm: 750_000,
		BasePositionNotional:   1_000_000,
		ImpactNotional:         500_000_000,
	},
	{
		Id:                     2,
		Name:                   "2",
		InitialMarginPpm:       1_000_000,
		MaintenanceFractionPpm: 0,
		BasePositionNotional:   1_000_000,
		ImpactNotional:         500_000_000,
	},
	{
		Id:                     3,
		Name:                   "3",
		InitialMarginPpm:       200_000,
		MaintenanceFractionPpm: 500_000,
		BasePositionNotional:   100_000_000_000,
		ImpactNotional:         2_500_000_000,
	},
	{
		Id:                     4,
		Name:                   "4",
		InitialMarginPpm:       500_000,
		MaintenanceFractionPpm: 800_000,
		BasePositionNotional:   100_000_000_000,
		ImpactNotional:         1_000_000_000,
	},
	{
		Id:                     5,
		Name:                   "5",
		InitialMarginPpm:       500_000,
		MaintenanceFractionPpm: 600_000,
		BasePositionNotional:   1_000_000,
		ImpactNotional:         1_000_000_000,
	},
	{
		Id:                     6,
		Name:                   "6",
		InitialMarginPpm:       200_000,
		MaintenanceFractionPpm: 900_000,
		BasePositionNotional:   1_000_000,
		ImpactNotional:         2_500_000_000,
	},
	{
		Id:                     7,
		Name:                   "7",
		InitialMarginPpm:       0,
		MaintenanceFractionPpm: 0,
		BasePositionNotional:   100_000_000_000,
		ImpactNotional:         1_000_000_000,
	},
	{
		Id:                     8,
		Name:                   "8",
		InitialMarginPpm:       9_910, // 0.9910%
		MaintenanceFractionPpm: 1_000_000,
		BasePositionNotional:   100_000_000_000,
		ImpactNotional:         50_454_000_000,
	},
}

// Perpetual genesis parameters.
const TestFundingRateClampFactorPpm = 6_000_000
const TestPremiumVoteClampFactorPpm = 60_000_000
const TestMinNumVotesPerSample = 15

var PerpetualsGenesisParams = perptypes.Params{
	FundingRateClampFactorPpm: TestFundingRateClampFactorPpm,
	PremiumVoteClampFactorPpm: TestPremiumVoteClampFactorPpm,
	MinNumVotesPerSample:      TestMinNumVotesPerSample,
}

var Perpetuals_GenesisState_ParamsOnly = perptypes.GenesisState{
	Params: PerpetualsGenesisParams,
}

// Perpetual objects.
var (
	BtcUsd_InvalidMarketId = perptypes.Perpetual{
		Ticker:            "BTC-USD invalid market Id",
		MarketId:          uint32(9999),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-10),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(0),
	}
	BtcUsd_0DefaultFunding_0AtomicResolution = perptypes.Perpetual{
		Ticker:            "BTC-USD 0 percent default funding, 0 atomic resolution",
		MarketId:          uint32(0),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(0),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(2),
	}
	BtcUsd_NegativeDefaultFunding_10AtomicResolution = perptypes.Perpetual{
		Ticker:            "BTC-USD -0.001 percent percent default funding",
		MarketId:          uint32(0),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-10),
		DefaultFundingPpm: int32(-1_000),
		LiquidityTier:     uint32(1),
	}
	BtcUsd_0DefaultFunding_10AtomicResolution = perptypes.Perpetual{
		Ticker:            "BTC-USD 0 percent default funding",
		MarketId:          uint32(0),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-10),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(1),
	}
	BtcUsd_0DefaultFunding_10AtomicResolution_20IM_18MM = perptypes.Perpetual{
		Ticker:            "BTC-USD 0 percent default funding, 20% IM, 18% MM",
		MarketId:          uint32(0),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-10),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(6),
	}
	BtcUsd_0_001Percent_DefaultFunding_10AtomicResolution = perptypes.Perpetual{
		Ticker:            "BTC-USD 0.001 percent default funding",
		MarketId:          uint32(0),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-10),
		DefaultFundingPpm: int32(1000), // 0.001%
		LiquidityTier:     uint32(1),
	}
	BtcUsd_SmallMarginRequirement = perptypes.Perpetual{
		Ticker:            "BTC-USD small margin requirement",
		MarketId:          uint32(0),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-8),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(8),
	}
	BtcUsd_100PercentMarginRequirement = perptypes.Perpetual{
		Ticker:            "BTC-USD 100% margin requirement",
		MarketId:          uint32(0),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-8),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(0),
	}
	BtcUsd_50PercentInitial_40PercentMaintenance = perptypes.Perpetual{
		Ticker:            "BTC-USD 50/40 margin requirements",
		MarketId:          uint32(0),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-8),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(4),
	}
	BtcUsd_20PercentInitial_10PercentMaintenance = perptypes.Perpetual{
		Ticker:            "BTC-USD 20/10 margin requirements",
		MarketId:          uint32(0),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-8),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(3),
	}
	BtcUsd_NoMarginRequirement = perptypes.Perpetual{
		Ticker:            "BTC-USD no margin requirement",
		MarketId:          uint32(0),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-8),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(7),
	}
	EthUsd_0DefaultFunding_9AtomicResolution = perptypes.Perpetual{
		Ticker:            "ETH-USD default fundingm, -9 atomic resolution",
		MarketId:          uint32(1),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-9),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(5),
	}
	EthUsd_NoMarginRequirement = perptypes.Perpetual{
		Ticker:            "ETH-USD no margin requirement",
		MarketId:          uint32(1),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-9),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(7),
	}
	EthUsd_20PercentInitial_10PercentMaintenance = perptypes.Perpetual{
		Ticker:            "ETH-USD 20/10 margin requirements",
		MarketId:          uint32(1),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-9),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(3),
	}
	EthUsd_100PercentMarginRequirement = perptypes.Perpetual{
		Ticker:            "ETH-USD 100/100 margin requirements",
		MarketId:          uint32(1),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-9),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(0),
	}
	SolUsd_20PercentInitial_10PercentMaintenance = perptypes.Perpetual{
		Ticker:            "SOL-USD 20/10 margin requirements",
		MarketId:          uint32(2),
		FundingIndex:      dtypes.ZeroInt(),
		AtomicResolution:  int32(-9),
		DefaultFundingPpm: int32(0),
		LiquidityTier:     uint32(3),
	}
)

// AddPremiumVotes messages.
var (
	TestAddPremiumVotesMsg = &perptypes.MsgAddPremiumVotes{
		Votes: []perptypes.FundingPremium{
			{
				PerpetualId: 0,
				PremiumPpm:  1000,
			},
		},
	}

	EmptyMsgAddPremiumVotes        = &perptypes.MsgAddPremiumVotes{}
	EmptyMsgAddPremiumVotesTxBytes []byte

	ValidMsgAddPremiumVotes = &perptypes.MsgAddPremiumVotes{
		Votes: []perptypes.FundingPremium{
			{PerpetualId: 1, PremiumPpm: 1_000},
			{PerpetualId: 2, PremiumPpm: 2_000},
		},
	}
	ValidMsgAddPremiumVotesTxBytes []byte

	InvalidMsgAddPremiumVotes = &perptypes.MsgAddPremiumVotes{
		Votes: []perptypes.FundingPremium{
			{PerpetualId: 3, PremiumPpm: 3_000}, // descending order is incorrect.
			{PerpetualId: 2, PremiumPpm: 2_000},
		},
	}
	InvalidMsgAddPremiumVotesTxBytes []byte

	Perpetuals_DefaultGenesisState = perptypes.GenesisState{
		LiquidityTiers: []perptypes.LiquidityTier{
			{
				Id:                     uint32(0),
				Name:                   "Large-Cap",
				InitialMarginPpm:       200_000,
				MaintenanceFractionPpm: 500_000,
				BasePositionNotional:   1000_000_000,
				ImpactNotional:         2_500_000_000,
			},
			{
				Id:                     uint32(1),
				Name:                   "Mid-Cap",
				InitialMarginPpm:       300_000,
				MaintenanceFractionPpm: 600_000,
				BasePositionNotional:   500_000_000,
				ImpactNotional:         1_667_000_000,
			},
			{
				Id:                     uint32(2),
				Name:                   "Small-Cap",
				InitialMarginPpm:       400_000,
				MaintenanceFractionPpm: 700_000,
				BasePositionNotional:   250_000_000,
				ImpactNotional:         1_250_000_000,
			},
		},
		Params: PerpetualsGenesisParams,
		Perpetuals: []perptypes.Perpetual{
			{
				Id:            uint32(0),
				Ticker:        "genesis_test_ticker_0",
				FundingIndex:  dtypes.ZeroInt(),
				LiquidityTier: 0,
			},
			{
				Id:            uint32(1),
				Ticker:        "genesis_test_ticker_1",
				FundingIndex:  dtypes.ZeroInt(),
				LiquidityTier: 1,
			},
		},
	}
)

// Return a list of `count` constant funding premiums equal to `value`.
func GenerateConstantFundingPremiums(
	value int32,
	count uint32,
) (
	result []int32,
) {
	result = make([]int32, count)
	for i := uint32(0); i < count; i += 1 {
		result[i] = value
	}
	return result
}

// Returns a funding sample list of length `n = sum(counts)`, where
// each `values[i]` appears `counts[i]` times.
func GenerateFundingSamplesWithValues(
	values []int32,
	counts []uint32,
) (
	result []int32,
) {
	for i, count := range counts {
		result = append(result, GenerateConstantFundingPremiums(values[i], count)...)
	}

	return result
}

// Return a list of 60 funding premium samples, with 30 equal to
// 0.001% or 1000 in ppm, 15 equal to -0.1% and 15 equal to 0.1%.
func FundingSamples_Constant_0_001_Percent_Length_60_With_Noises() (
	result []int32,
) {
	result = make([]int32, 60)
	for i := 0; i < 30; i += 1 {
		result[i] = 1000
	}

	for i := 30; i < 45; i += 1 {
		result[i] = 100000
	}

	for i := 45; i < 60; i += 1 {
		result[i] = -100000
	}
	return result
}
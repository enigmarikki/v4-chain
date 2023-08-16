//go:build all || integration_test

package cli_test

import (
	"fmt"
	"math/big"
	"testing"

	networktestutil "github.com/cosmos/cosmos-sdk/testutil/network"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	"github.com/dydxprotocol/v4/app"
	daemonflags "github.com/dydxprotocol/v4/daemons/flags"
	"github.com/dydxprotocol/v4/lib"
	"github.com/dydxprotocol/v4/testutil/appoptions"
	testutil_bank "github.com/dydxprotocol/v4/testutil/bank"
	"github.com/dydxprotocol/v4/testutil/constants"
	testutil "github.com/dydxprotocol/v4/testutil/keeper"
	"github.com/dydxprotocol/v4/testutil/network"
	cli_testutil "github.com/dydxprotocol/v4/x/clob/client/testutil"
	"github.com/dydxprotocol/v4/x/clob/types"
	epochstypes "github.com/dydxprotocol/v4/x/epochs/types"
	feetierstypes "github.com/dydxprotocol/v4/x/feetiers/types"
	perptypes "github.com/dydxprotocol/v4/x/perpetuals/types"
	pricestypes "github.com/dydxprotocol/v4/x/prices/types"
	sa_testutil "github.com/dydxprotocol/v4/x/subaccounts/client/testutil"
	satypes "github.com/dydxprotocol/v4/x/subaccounts/types"
	"github.com/stretchr/testify/suite"
)

const (
	cancelsInitialQuoteBalance               = int64(1_000_000_000)  // $1,000.
	cancelsInitialSubaccountModuleAccBalance = int64(10_000_000_000) // $10,000.
	cancelsSubaccountNumberZero              = uint32(0)
	cancelsSubaccountNumberOne               = uint32(1)
)

type CancelOrderIntegrationTestSuite struct {
	suite.Suite

	validatorAddress sdk.AccAddress
	cfg              network.Config
	network          *network.Network
}

func TestCancelOrderIntegrationTestSuite(t *testing.T) {
	suite.Run(t, &CancelOrderIntegrationTestSuite{})
}

func (s *CancelOrderIntegrationTestSuite) SetupTest() {
	s.T().Log("setting up cancel order integration test")

	// Deterministic Mnemonic.
	validatorMnemonic := constants.AliceMnenomic

	// Generated from the above Mnemonic.
	s.validatorAddress = constants.AliceAccAddress

	appOptions := appoptions.NewFakeAppOptions()

	// Configure test network.
	s.cfg = network.DefaultConfig(&network.NetworkConfigOptions{
		AppOptions: appOptions,
		OnNewApp: func(val networktestutil.ValidatorI) {
			// Disable the Bridge and Price daemons in the integration tests.
			appOptions.Set(daemonflags.FlagPriceDaemonEnabled, false)
			appOptions.Set(daemonflags.FlagBridgeDaemonEnabled, false)
		},
	})

	s.cfg.Mnemonics = append(s.cfg.Mnemonics, validatorMnemonic)
	s.cfg.ChainID = app.AppName

	s.cfg.MinGasPrices = fmt.Sprintf("0%s", sdk.DefaultBondDenom)

	clobPair := constants.ClobPair_Btc
	state := types.GenesisState{}

	state.ClobPairs = append(state.ClobPairs, clobPair)
	state.LiquidationsConfig = types.LiquidationsConfig_Default

	perpstate := perptypes.GenesisState{}
	perpstate.LiquidityTiers = constants.LiquidityTiers
	perpstate.Params = constants.PerpetualsGenesisParams
	perpetual := constants.BtcUsd_50PercentInitial_40PercentMaintenance
	perpstate.Perpetuals = append(perpstate.Perpetuals, perpetual)

	pricesstate := constants.Prices_DefaultGenesisState

	buf, err := s.cfg.Codec.MarshalJSON(&state)
	s.NoError(err)
	s.cfg.GenesisState[types.ModuleName] = buf

	// Set the balances in the genesis state.
	s.cfg.GenesisState[banktypes.ModuleName] = cli_testutil.CreateBankGenesisState(
		s.T(),
		s.cfg,
		constants.QuoteBalance_OneDollar*10_000,
	)

	sastate := satypes.GenesisState{}
	sastate.Subaccounts = append(
		sastate.Subaccounts,
		satypes.Subaccount{
			Id: &satypes.SubaccountId{
				Owner:  s.validatorAddress.String(),
				Number: cancelsSubaccountNumberZero,
			},
			AssetPositions:     testutil.CreateUsdcAssetPosition(big.NewInt(cancelsInitialQuoteBalance)),
			PerpetualPositions: []*satypes.PerpetualPosition{},
		},
		satypes.Subaccount{
			Id: &satypes.SubaccountId{
				Owner:  s.validatorAddress.String(),
				Number: cancelsSubaccountNumberOne,
			},
			AssetPositions:     testutil.CreateUsdcAssetPosition(big.NewInt(cancelsInitialQuoteBalance)),
			PerpetualPositions: []*satypes.PerpetualPosition{},
		},
	)

	// Ensure that no funding payments will occur during this test.
	epstate := constants.GenerateEpochGenesisStateWithoutFunding()

	feeTiersState := feetierstypes.GenesisState{}
	feeTiersState.Params = constants.PerpetualFeeParams

	epbuf, err := s.cfg.Codec.MarshalJSON(&epstate)
	s.Require().NoError(err)
	s.cfg.GenesisState[epochstypes.ModuleName] = epbuf

	sabuf, err := s.cfg.Codec.MarshalJSON(&sastate)
	s.Require().NoError(err)
	s.cfg.GenesisState[satypes.ModuleName] = sabuf

	perpbuf, err := s.cfg.Codec.MarshalJSON(&perpstate)
	s.Require().NoError(err)
	s.cfg.GenesisState[perptypes.ModuleName] = perpbuf

	pricesbuf, err := s.cfg.Codec.MarshalJSON(&pricesstate)
	s.Require().NoError(err)
	s.cfg.GenesisState[pricestypes.ModuleName] = pricesbuf

	feeTiersBuf, err := s.cfg.Codec.MarshalJSON(&feeTiersState)
	s.Require().NoError(err)
	s.cfg.GenesisState[feetierstypes.ModuleName] = feeTiersBuf

	s.network = network.New(s.T(), s.cfg)

	_, err = s.network.WaitForHeight(1)
	s.Require().NoError(err)
}

// TestCLICancelPendingOrder places then cancels an order from a subaccount, and then places an order from
// a different subaccount (with the same owner and different numbers).
// The orders placed are expected to match, but should not due to the first order being canceled.
// Afterwards, an additional cancel of an unknown is order is made (expected to be a no-op).
// The subaccounts are then queried and assertions are performed on their QuoteBalance and PerpetualPositions.
// The account which places the orders is also the validator's AccAddress.
func (s *CancelOrderIntegrationTestSuite) TestCLICancelPendingOrder() {
	val := s.network.Validators[0]
	ctx := val.ClientCtx

	currentHeight, err := s.network.LatestHeight()
	s.Require().NoError(err)

	goodTilBlock := uint32(currentHeight) + types.ShortBlockWindow
	clientId := uint64(1)
	quantums := satypes.BaseQuantums(1_000)
	subticks := types.Subticks(50_000_000_000)

	// Place the first order.
	_, err = cli_testutil.MsgPlaceOrderExec(
		ctx,
		s.validatorAddress,
		cancelsSubaccountNumberZero,
		clientId,
		constants.ClobPair_Btc.Id,
		types.Order_SIDE_BUY,
		quantums,
		subticks.ToUint64(),
		goodTilBlock,
	)
	s.Require().NoError(err)

	// Cancel the first order.
	_, err = cli_testutil.MsgCancelOrderExec(
		ctx,
		s.validatorAddress,
		cancelsSubaccountNumberZero,
		clientId,
		goodTilBlock,
	)
	s.Require().NoError(err)

	// Place the second order.
	_, err = cli_testutil.MsgPlaceOrderExec(
		ctx,
		s.validatorAddress,
		cancelsSubaccountNumberOne,
		clientId,
		constants.ClobPair_Btc.Id,
		types.Order_SIDE_SELL,
		quantums,
		subticks.ToUint64(),
		goodTilBlock,
	)
	s.Require().NoError(err)

	// Cancel an unknown order.
	unknownClientId := uint64(10)
	_, err = cli_testutil.MsgCancelOrderExec(
		ctx,
		s.validatorAddress,
		cancelsSubaccountNumberZero,
		unknownClientId,
		goodTilBlock,
	)
	s.Require().NoError(err)

	currentHeight, err = s.network.LatestHeight()
	s.Require().NoError(err)

	// Wait for a few blocks.
	_, err = s.network.WaitForHeight(currentHeight + 3)
	s.Require().NoError(err)

	// Check that subaccounts balance have not changed, and no positions were opened.
	for _, subaccountNumber := range []uint32{cancelsSubaccountNumberZero, cancelsSubaccountNumberOne} {
		resp, err := sa_testutil.MsgQuerySubaccountExec(ctx, s.validatorAddress, subaccountNumber)
		s.Require().NoError(err)

		var subaccountResp satypes.QuerySubaccountResponse
		s.Require().NoError(s.network.Config.Codec.UnmarshalJSON(resp.Bytes(), &subaccountResp))
		subaccount := subaccountResp.Subaccount

		s.Require().Equal(
			new(big.Int).SetInt64(cancelsInitialQuoteBalance),
			subaccount.GetUsdcPosition(),
		)
		s.Require().Len(subaccount.PerpetualPositions, 0)

		s.Require().Equal(
			new(big.Int).SetInt64(cancelsInitialQuoteBalance),
			subaccount.GetUsdcPosition())
		s.Require().Len(subaccount.PerpetualPositions, 0)
	}

	// Check that the `subaccounts` module account balance has not changed.
	saModuleUSDCBalance, err := testutil_bank.GetModuleAccUsdcBalance(
		ctx,
		s.network.Config.Codec,
		satypes.ModuleName,
	)
	s.Require().NoError(err)
	s.Require().Equal(
		cancelsInitialSubaccountModuleAccBalance,
		saModuleUSDCBalance,
	)

	// Check that the `distribution` module account USDC balance has not changed.
	distrModuleUSDCBalance, err := testutil_bank.GetModuleAccUsdcBalance(
		ctx,
		s.network.Config.Codec,
		distrtypes.ModuleName,
	)

	s.Require().NoError(err)
	s.Require().Equal(int64(0), distrModuleUSDCBalance)
}

// TestCLICancelMatchingOrders places two matching orders from two different subaccounts (with the
// same owner and different numbers), then cancels the first matching order a few blocks later.
// The matching orders should not be canceled.
// The subaccounts are then queried and assertions are performed on their QuoteBalance and PerpetualPositions.
// The account which places the orders is also the validator's AccAddress.
func (s *CancelOrderIntegrationTestSuite) TestCLICancelMatchingOrders() {
	val := s.network.Validators[0]
	ctx := val.ClientCtx

	currentHeight, err := s.network.LatestHeight()
	s.Require().NoError(err)

	goodTilBlock := uint32(currentHeight) + types.ShortBlockWindow
	clientId := uint64(2)
	quantums := satypes.BaseQuantums(1_000)
	subticks := types.Subticks(50_000_000_000)

	// Place the first order.
	_, err = cli_testutil.MsgPlaceOrderExec(
		ctx,
		s.validatorAddress,
		cancelsSubaccountNumberZero,
		clientId,
		constants.ClobPair_Btc.Id,
		types.Order_SIDE_BUY,
		quantums,
		subticks.ToUint64(),
		goodTilBlock,
	)
	s.Require().NoError(err)

	// Place the second order.
	_, err = cli_testutil.MsgPlaceOrderExec(
		ctx,
		s.validatorAddress,
		cancelsSubaccountNumberOne,
		clientId,
		constants.ClobPair_Btc.Id,
		types.Order_SIDE_SELL,
		quantums,
		subticks.ToUint64(),
		goodTilBlock,
	)
	s.Require().NoError(err)

	currentHeight, err = s.network.LatestHeight()
	s.Require().NoError(err)

	// Wait for a few blocks.
	_, err = s.network.WaitForHeight(currentHeight + 3)
	s.Require().NoError(err)

	// Cancel the first order.
	_, err = cli_testutil.MsgCancelOrderExec(
		ctx,
		s.validatorAddress,
		cancelsSubaccountNumberZero,
		clientId,
		goodTilBlock,
	)
	s.Require().NoError(err)

	currentHeight, err = s.network.LatestHeight()
	s.Require().NoError(err)

	// Wait for a few blocks.
	_, err = s.network.WaitForHeight(currentHeight + 3)
	s.Require().NoError(err)

	// Query both subaccounts.
	resp, err := sa_testutil.MsgQuerySubaccountExec(ctx, s.validatorAddress, subaccountNumberZero)
	s.Require().NoError(err)

	var subaccountResp satypes.QuerySubaccountResponse
	s.Require().NoError(s.network.Config.Codec.UnmarshalJSON(resp.Bytes(), &subaccountResp))
	subaccountZero := subaccountResp.Subaccount

	resp, err = sa_testutil.MsgQuerySubaccountExec(ctx, s.validatorAddress, subaccountNumberOne)
	s.Require().NoError(err)

	s.Require().NoError(s.network.Config.Codec.UnmarshalJSON(resp.Bytes(), &subaccountResp))
	subaccountOne := subaccountResp.Subaccount

	// Compute the fill price so as to know how much QuoteBalance should be remaining.
	fillSizeQuoteQuantums := types.FillAmountToQuoteQuantums(
		subticks,
		quantums,
		constants.ClobPair_Btc.QuantumConversionExponent,
	).Int64()

	// Assert that both Subaccounts have the appropriate state.
	// Order could be maker or taker after Uncross, so assert that account could have been either.
	takerFee := fillSizeQuoteQuantums * int64(constants.TakerFeePpm) / int64(lib.OneMillion)
	makerFee := fillSizeQuoteQuantums * int64(constants.MakerFeePpm) / int64(lib.OneMillion)

	s.Require().Contains(
		[]*big.Int{
			new(big.Int).SetInt64(initialQuoteBalance - fillSizeQuoteQuantums - takerFee),
			new(big.Int).SetInt64(initialQuoteBalance - fillSizeQuoteQuantums - makerFee),
		},
		subaccountZero.GetUsdcPosition(),
	)
	s.Require().Len(subaccountZero.PerpetualPositions, 1)
	s.Require().Equal(quantums.ToBigInt(), subaccountZero.PerpetualPositions[0].GetBigQuantums())

	s.Require().Contains(
		[]*big.Int{
			new(big.Int).SetInt64(initialQuoteBalance + fillSizeQuoteQuantums - takerFee),
			new(big.Int).SetInt64(initialQuoteBalance + fillSizeQuoteQuantums - makerFee),
		},
		subaccountOne.GetUsdcPosition(),
	)
	s.Require().Len(subaccountOne.PerpetualPositions, 1)
	s.Require().Equal(new(big.Int).Neg(quantums.ToBigInt()), subaccountOne.PerpetualPositions[0].GetBigQuantums())

	// Check that the `subaccounts` module account has expected remaining USDC balance.
	saModuleUSDCBalance, err := testutil_bank.GetModuleAccUsdcBalance(
		ctx,
		s.network.Config.Codec,
		satypes.ModuleName,
	)
	s.Require().NoError(err)
	s.Require().Equal(
		initialSubaccountModuleAccBalance-makerFee-takerFee,
		saModuleUSDCBalance,
	)

	// Check that the `distribution` module account USDC balance has not changed.
	distrModuleUSDCBalance, err := testutil_bank.GetModuleAccUsdcBalance(
		ctx,
		s.network.Config.Codec,
		distrtypes.ModuleName,
	)

	s.Require().NoError(err)
	s.Require().Equal(makerFee+takerFee, distrModuleUSDCBalance)
}
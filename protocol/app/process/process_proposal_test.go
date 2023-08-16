package process_test

import (
	"testing"

	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/dydxprotocol/v4/app/process"
	"github.com/dydxprotocol/v4/mocks"
	"github.com/dydxprotocol/v4/testutil/constants"
	keepertest "github.com/dydxprotocol/v4/testutil/keeper"
	testmsgs "github.com/dydxprotocol/v4/testutil/msgs"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestProcessProposalHandler_Error(t *testing.T) {
	acceptResponse := abci.ResponseProcessProposal{
		Status: abci.ResponseProcessProposal_ACCEPT,
	}
	rejectResponse := abci.ResponseProcessProposal{
		Status: abci.ResponseProcessProposal_REJECT,
	}

	// Valid operations tx.
	validOperationsTx := constants.ValidEmptyMsgProposedOperationsTxBytes

	// Valid add funding tx.
	validAddFundingTx := constants.ValidMsgAddPremiumVotesTxBytes

	// Valid acknowledge bridges tx.
	validAcknowledgeBridgesTx := constants.MsgAcknowledgeBridges_Ids0_1_Height0_TxBytes
	validAcknowledgeBridgesTx_NoEvents := constants.MsgAcknowledgeBridges_NoEvents_TxBytes

	// Valid update price tx.
	validUpdatePriceTx := constants.ValidMsgUpdateMarketPricesTxBytes

	// Valid "other" single msg tx.
	validSingleMsgOtherTx := constants.Msg_Send_TxBytes

	// Valid "other" multi msgs tx.
	validMultiMsgOtherTx := constants.Msg_SendAndTransfer_TxBytes

	// Invalid update price tx.
	invalidUpdatePriceTx := constants.InvalidMsgUpdateMarketPricesStatelessTxBytes

	// Invalid acknowledge bridges txs.
	acknowledgeBridgesTx_IdsNotConsecutive := constants.MsgAcknowledgeBridges_Ids0_55_Height0_TxBytes
	acknowledgeBridgesTx_NotRecognized := constants.MsgAcknowledgeBridges_Id55_Height15_TxBytes
	acknowledgeBridgesTx_NotNextToAcknowledge := constants.MsgAcknowledgeBridges_Id1_Height0_TxBytes

	tests := map[string]struct {
		txsBytes [][]byte

		expectedResponse abci.ResponseProcessProposal
	}{
		"Reject: decode fails": {
			txsBytes:         [][]byte{{1}, {2}},
			expectedResponse: rejectResponse,
		},
		"Reject: invalid price tx": {
			txsBytes: [][]byte{
				validOperationsTx,
				validAcknowledgeBridgesTx,
				validAddFundingTx,
				invalidUpdatePriceTx, // invalid.
			},
			expectedResponse: rejectResponse,
		},
		"Reject: bridge event IDs not consecutive": {
			txsBytes: [][]byte{
				validOperationsTx,
				acknowledgeBridgesTx_IdsNotConsecutive,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: rejectResponse,
		},
		"Reject: bridge event ID not yet recognized": {
			txsBytes: [][]byte{
				validOperationsTx,
				acknowledgeBridgesTx_NotRecognized,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: rejectResponse,
		},
		"Reject: bridge event ID not next to acknowledge": {
			txsBytes: [][]byte{
				validOperationsTx,
				acknowledgeBridgesTx_NotNextToAcknowledge,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: rejectResponse,
		},
		"Error: place order type is not allowed": {
			txsBytes: [][]byte{
				validOperationsTx,
				constants.Msg_PlaceOrder_TxBtyes, // invalid other txs.
				validAcknowledgeBridgesTx,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: rejectResponse,
		},
		"Error: cancel order type is not allowed": {
			txsBytes: [][]byte{
				validOperationsTx,
				constants.Msg_CancelOrder_TxBtyes, // invalid other txs.
				validAcknowledgeBridgesTx,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: rejectResponse,
		},
		"Error: app-injected msg type is not allowed": {
			txsBytes: [][]byte{
				validOperationsTx,
				validUpdatePriceTx, // invalid other txs.
				validAcknowledgeBridgesTx,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: rejectResponse,
		},
		"Error: internal msg type is not allowed": {
			txsBytes: [][]byte{
				validOperationsTx,
				testmsgs.MsgSoftwareUpgradeTxBytes, // invalid other txs.
				validAcknowledgeBridgesTx,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: rejectResponse,
		},
		"Error: unsupported msg type is not allowed": {
			txsBytes: [][]byte{
				validOperationsTx,
				testmsgs.GovBetaMsgSubmitProposalTxBytes, // invalid other txs.
				validAcknowledgeBridgesTx,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: rejectResponse,
		},
		"Error: nested msg type with unsupported inner is not allowed": {
			txsBytes: [][]byte{
				validOperationsTx,
				testmsgs.MsgSubmitProposalWithUnsupportedInnerTxBytes, // invalid other txs.
				validAcknowledgeBridgesTx,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: rejectResponse,
		},
		"Error: nested msg type with app-injected inner is not allowed": {
			txsBytes: [][]byte{
				validOperationsTx,
				testmsgs.MsgSubmitProposalWithAppInjectedInnerTxBytes, // invalid other txs.
				validAcknowledgeBridgesTx,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: rejectResponse,
		},
		"Error: nested msg type with double-nested inner is not allowed": {
			txsBytes: [][]byte{
				validOperationsTx,
				testmsgs.MsgSubmitProposalWithDoubleNestedInnerTxBytes, // invalid other txs.
				validAcknowledgeBridgesTx,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: rejectResponse,
		},
		"Accept: bridge tx with no events": {
			txsBytes: [][]byte{
				validOperationsTx,
				validSingleMsgOtherTx,
				validAcknowledgeBridgesTx_NoEvents,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: acceptResponse,
		},
		"Accept: Valid txs": {
			txsBytes: [][]byte{
				validOperationsTx,
				validMultiMsgOtherTx,  // other txs.
				validSingleMsgOtherTx, // other txs.
				validAcknowledgeBridgesTx,
				validAddFundingTx,
				validUpdatePriceTx,
			},
			expectedResponse: acceptResponse,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			// Setup.
			ctx, pricesKeeper, _, indexPriceCache, marketToSmoothedPrices, mockTimeProvider := keepertest.PricesKeepers(t)
			keepertest.CreateTestMarkets(t, ctx, pricesKeeper)
			indexPriceCache.UpdatePrices(constants.AtTimeTSingleExchangePriceUpdate)
			mockTimeProvider.On("Now").Return(constants.TimeT)

			mockClobKeeper := &mocks.ProcessClobKeeper{}
			mockClobKeeper.On("RecordMevMetrics", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil)

			mockBridgeKeeper := &mocks.ProcessBridgeKeeper{}
			mockBridgeKeeper.On("GetAcknowledgedEventInfo", mock.Anything).Return(constants.AcknowledgedEventInfo_Id0_Height0)
			mockBridgeKeeper.On("GetRecognizedEventInfo", mock.Anything).Return(constants.RecognizedEventInfo_Id2_Height0)

			handler := process.ProcessProposalHandler(
				constants.TestEncodingCfg.TxConfig,
				mockBridgeKeeper,
				mockClobKeeper,
				&mocks.ProcessStakingKeeper{},
				&mocks.ProcessPerpetualKeeper{},
				pricesKeeper,
			)
			req := abci.RequestProcessProposal{Txs: tc.txsBytes}

			// Run.
			resp := handler(ctx, req)

			// Validate.
			require.Equal(t, tc.expectedResponse, resp)
			require.Equal(
				t,
				marketToSmoothedPrices.GetSmoothedPricesForTest(),
				constants.AtTimeTSingleExchangeSmoothedPrices,
			)
		})
	}
}
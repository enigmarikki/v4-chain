package types

import (
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/dydxprotocol/v4/lib/metrics"
)

const TypeMsgPlaceOrder = "place_order"

var _ sdk.Msg = &MsgPlaceOrder{}

func NewMsgPlaceOrder(order Order) *MsgPlaceOrder {
	return &MsgPlaceOrder{
		Order: order,
	}
}

func (msg *MsgPlaceOrder) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Order.OrderId.SubaccountId.Owner)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgPlaceOrder) ValidateBasic() (err error) {
	defer func() {
		if err != nil {
			telemetry.IncrCounterWithLabels(
				[]string{ModuleName, metrics.PlaceOrder, metrics.ValidateBasic, metrics.Error, metrics.Count},
				1,
				msg.Order.GetOrderLabels(),
			)
		}
	}()

	err = msg.Order.OrderId.SubaccountId.Validate()
	if err != nil {
		return err
	}

	if _, exists := Order_Side_name[int32(msg.Order.Side)]; !exists {
		return sdkerrors.Wrapf(ErrInvalidOrderSide, "invalid order side (%s)", msg.Order.Side)
	}

	if msg.Order.Side == Order_SIDE_UNSPECIFIED {
		return sdkerrors.Wrapf(ErrInvalidOrderSide, "UNSPECIFIED is not a valid order side")
	}

	if msg.Order.Quantums == uint64(0) {
		return sdkerrors.Wrapf(ErrInvalidOrderQuantums, "order size quantums cannot be 0")
	}

	orderId := msg.Order.GetOrderId()
	if orderId.IsShortTermOrder() {
		// This also implicitly verifies that GoodTilBlockTime is not set / is zero for short-term orders.
		if msg.Order.GetGoodTilBlock() == uint32(0) {
			return sdkerrors.Wrapf(ErrInvalidOrderGoodTilBlock, "order goodTilBlock cannot be 0")
		}
	} else if orderId.IsStatefulOrder() {
		if msg.Order.GetGoodTilBlockTime() == uint32(0) {
			return sdkerrors.Wrapf(
				ErrInvalidStatefulOrderGoodTilBlockTime,
				"stateful order goodTilBlockTime cannot be 0",
			)
		}
	} else {
		return sdkerrors.Wrapf(ErrInvalidOrderFlag, "invalid order flag %v", orderId.OrderFlags)
	}

	if orderId.IsLongTermOrder() && msg.Order.RequiresImmediateExecution() {
		return ErrStatefulOrdersCannotRequireImmediateExecution
	}

	if msg.Order.ReduceOnly {
		return sdkerrors.Wrapf(ErrReduceOnlyDisabled, "reduce-only is temporarily disabled")
	}

	if msg.Order.Subticks == uint64(0) {
		return sdkerrors.Wrapf(ErrInvalidOrderSubticks, "order subticks cannot be 0")
	}

	if orderId.IsConditionalOrder() {
		if msg.Order.ConditionType == Order_CONDITION_TYPE_UNSPECIFIED {
			return sdkerrors.Wrapf(ErrInvalidConditionType, "condition type cannot be unspecified")
		}

		if msg.Order.ConditionalOrderTriggerSubticks == uint64(0) {
			return sdkerrors.Wrapf(ErrInvalidConditionalOrderTriggerSubticks, "conditional order trigger subticks cannot be 0")
		}
	} else {
		if msg.Order.ConditionType != Order_CONDITION_TYPE_UNSPECIFIED {
			return sdkerrors.Wrapf(ErrInvalidConditionType, "condition type specified for non-conditional order")
		}

		if msg.Order.ConditionalOrderTriggerSubticks != uint64(0) {
			return sdkerrors.Wrapf(
				ErrInvalidConditionalOrderTriggerSubticks,
				"conditional order trigger subticks greater than 0 for non-conditional order",
			)
		}
	}

	return nil
}
import { Order, OrderSDKType, OrderId, OrderIdSDKType } from "./order";
import { ClobMatch, ClobMatchSDKType } from "./matches";
import { OrderRemoval, OrderRemovalSDKType } from "./order_removals";
import * as _m0 from "protobufjs/minimal";
import { DeepPartial } from "../../helpers";
/**
 * MsgProposedOperations is a message injected by block proposers to
 * specify the operations that occurred in a block.
 */

export interface MsgProposedOperations {
  /** The list of operations proposed by the block proposer. */
  operationsQueue: OperationRaw[];
}
/**
 * MsgProposedOperations is a message injected by block proposers to
 * specify the operations that occurred in a block.
 */

export interface MsgProposedOperationsSDKType {
  /** The list of operations proposed by the block proposer. */
  operations_queue: OperationRawSDKType[];
}
/**
 * MsgProposedOperationsResponse is the response type of the message injected
 * by block proposers to specify the operations that occurred in a block.
 */

export interface MsgProposedOperationsResponse {}
/**
 * MsgProposedOperationsResponse is the response type of the message injected
 * by block proposers to specify the operations that occurred in a block.
 */

export interface MsgProposedOperationsResponseSDKType {}
/** MsgPlaceOrder is a request type used for placing orders. */

export interface MsgPlaceOrder {
  /** MsgPlaceOrder is a request type used for placing orders. */
  order?: Order;
}
/** MsgPlaceOrder is a request type used for placing orders. */

export interface MsgPlaceOrderSDKType {
  /** MsgPlaceOrder is a request type used for placing orders. */
  order?: OrderSDKType;
}
/** MsgPlaceOrderResponse is a response type used for placing orders. */

export interface MsgPlaceOrderResponse {}
/** MsgPlaceOrderResponse is a response type used for placing orders. */

export interface MsgPlaceOrderResponseSDKType {}
/** MsgCancelOrder is a request type used for canceling orders. */

export interface MsgCancelOrder {
  orderId?: OrderId;
  /**
   * The last block this order cancellation can be executed at.
   * Used only for Short-Term orders and must be zero for stateful orders.
   */

  goodTilBlock?: number;
  /**
   * good_til_block_time represents the unix timestamp (in seconds) at which a
   * stateful order cancellation will be considered expired. The
   * good_til_block_time is always evaluated against the previous block's
   * `BlockTime` instead of the block in which the order is committed.
   * This value must be zero for Short-Term orders.
   */

  goodTilBlockTime?: number;
}
/** MsgCancelOrder is a request type used for canceling orders. */

export interface MsgCancelOrderSDKType {
  order_id?: OrderIdSDKType;
  /**
   * The last block this order cancellation can be executed at.
   * Used only for Short-Term orders and must be zero for stateful orders.
   */

  good_til_block?: number;
  /**
   * good_til_block_time represents the unix timestamp (in seconds) at which a
   * stateful order cancellation will be considered expired. The
   * good_til_block_time is always evaluated against the previous block's
   * `BlockTime` instead of the block in which the order is committed.
   * This value must be zero for Short-Term orders.
   */

  good_til_block_time?: number;
}
/** MsgCancelOrderResponse is a response type used for canceling orders. */

export interface MsgCancelOrderResponse {}
/** MsgCancelOrderResponse is a response type used for canceling orders. */

export interface MsgCancelOrderResponseSDKType {}
/**
 * OperationRaw represents an operation in the proposed operations.
 * Note that the `order_placement` operation is a signed message.
 */

export interface OperationRaw {
  match?: ClobMatch;
  shortTermOrderPlacement?: Uint8Array;
  orderRemoval?: OrderRemoval;
}
/**
 * OperationRaw represents an operation in the proposed operations.
 * Note that the `order_placement` operation is a signed message.
 */

export interface OperationRawSDKType {
  match?: ClobMatchSDKType;
  short_term_order_placement?: Uint8Array;
  order_removal?: OrderRemovalSDKType;
}

function createBaseMsgProposedOperations(): MsgProposedOperations {
  return {
    operationsQueue: []
  };
}

export const MsgProposedOperations = {
  encode(message: MsgProposedOperations, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.operationsQueue) {
      OperationRaw.encode(v!, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgProposedOperations {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgProposedOperations();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.operationsQueue.push(OperationRaw.decode(reader, reader.uint32()));
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgProposedOperations>): MsgProposedOperations {
    const message = createBaseMsgProposedOperations();
    message.operationsQueue = object.operationsQueue?.map(e => OperationRaw.fromPartial(e)) || [];
    return message;
  }

};

function createBaseMsgProposedOperationsResponse(): MsgProposedOperationsResponse {
  return {};
}

export const MsgProposedOperationsResponse = {
  encode(_: MsgProposedOperationsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgProposedOperationsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgProposedOperationsResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgProposedOperationsResponse>): MsgProposedOperationsResponse {
    const message = createBaseMsgProposedOperationsResponse();
    return message;
  }

};

function createBaseMsgPlaceOrder(): MsgPlaceOrder {
  return {
    order: undefined
  };
}

export const MsgPlaceOrder = {
  encode(message: MsgPlaceOrder, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.order !== undefined) {
      Order.encode(message.order, writer.uint32(10).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPlaceOrder {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPlaceOrder();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.order = Order.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgPlaceOrder>): MsgPlaceOrder {
    const message = createBaseMsgPlaceOrder();
    message.order = object.order !== undefined && object.order !== null ? Order.fromPartial(object.order) : undefined;
    return message;
  }

};

function createBaseMsgPlaceOrderResponse(): MsgPlaceOrderResponse {
  return {};
}

export const MsgPlaceOrderResponse = {
  encode(_: MsgPlaceOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgPlaceOrderResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgPlaceOrderResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgPlaceOrderResponse>): MsgPlaceOrderResponse {
    const message = createBaseMsgPlaceOrderResponse();
    return message;
  }

};

function createBaseMsgCancelOrder(): MsgCancelOrder {
  return {
    orderId: undefined,
    goodTilBlock: undefined,
    goodTilBlockTime: undefined
  };
}

export const MsgCancelOrder = {
  encode(message: MsgCancelOrder, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.orderId !== undefined) {
      OrderId.encode(message.orderId, writer.uint32(10).fork()).ldelim();
    }

    if (message.goodTilBlock !== undefined) {
      writer.uint32(16).uint32(message.goodTilBlock);
    }

    if (message.goodTilBlockTime !== undefined) {
      writer.uint32(29).fixed32(message.goodTilBlockTime);
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancelOrder {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancelOrder();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.orderId = OrderId.decode(reader, reader.uint32());
          break;

        case 2:
          message.goodTilBlock = reader.uint32();
          break;

        case 3:
          message.goodTilBlockTime = reader.fixed32();
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<MsgCancelOrder>): MsgCancelOrder {
    const message = createBaseMsgCancelOrder();
    message.orderId = object.orderId !== undefined && object.orderId !== null ? OrderId.fromPartial(object.orderId) : undefined;
    message.goodTilBlock = object.goodTilBlock ?? undefined;
    message.goodTilBlockTime = object.goodTilBlockTime ?? undefined;
    return message;
  }

};

function createBaseMsgCancelOrderResponse(): MsgCancelOrderResponse {
  return {};
}

export const MsgCancelOrderResponse = {
  encode(_: MsgCancelOrderResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCancelOrderResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCancelOrderResponse();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(_: DeepPartial<MsgCancelOrderResponse>): MsgCancelOrderResponse {
    const message = createBaseMsgCancelOrderResponse();
    return message;
  }

};

function createBaseOperationRaw(): OperationRaw {
  return {
    match: undefined,
    shortTermOrderPlacement: undefined,
    orderRemoval: undefined
  };
}

export const OperationRaw = {
  encode(message: OperationRaw, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.match !== undefined) {
      ClobMatch.encode(message.match, writer.uint32(10).fork()).ldelim();
    }

    if (message.shortTermOrderPlacement !== undefined) {
      writer.uint32(18).bytes(message.shortTermOrderPlacement);
    }

    if (message.orderRemoval !== undefined) {
      OrderRemoval.encode(message.orderRemoval, writer.uint32(26).fork()).ldelim();
    }

    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): OperationRaw {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseOperationRaw();

    while (reader.pos < end) {
      const tag = reader.uint32();

      switch (tag >>> 3) {
        case 1:
          message.match = ClobMatch.decode(reader, reader.uint32());
          break;

        case 2:
          message.shortTermOrderPlacement = reader.bytes();
          break;

        case 3:
          message.orderRemoval = OrderRemoval.decode(reader, reader.uint32());
          break;

        default:
          reader.skipType(tag & 7);
          break;
      }
    }

    return message;
  },

  fromPartial(object: DeepPartial<OperationRaw>): OperationRaw {
    const message = createBaseOperationRaw();
    message.match = object.match !== undefined && object.match !== null ? ClobMatch.fromPartial(object.match) : undefined;
    message.shortTermOrderPlacement = object.shortTermOrderPlacement ?? undefined;
    message.orderRemoval = object.orderRemoval !== undefined && object.orderRemoval !== null ? OrderRemoval.fromPartial(object.orderRemoval) : undefined;
    return message;
  }

};
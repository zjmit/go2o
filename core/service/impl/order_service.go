/**
 * Copyright 2014 @ to2.net.
 * name :
 * author : jarryliu
 * date : 2013-12-05 17:53
 * description :
 * history :
 */

package impl

import (
	"bytes"
	"context"
	"go2o/core/domain/interface/cart"
	"go2o/core/domain/interface/item"
	"go2o/core/domain/interface/member"
	"go2o/core/domain/interface/merchant"
	"go2o/core/domain/interface/merchant/shop"
	"go2o/core/domain/interface/order"
	"go2o/core/domain/interface/product"
	orderImpl "go2o/core/domain/order"
	"go2o/core/dto"
	"go2o/core/query"
	"go2o/core/service/parser"
	"go2o/core/service/proto"
	"strconv"
)

var _ proto.OrderServiceServer = new(orderServiceImpl)

type orderServiceImpl struct {
	repo       order.IOrderRepo
	prodRepo   product.IProductRepo
	itemRepo   item.IGoodsItemRepo
	cartRepo   cart.ICartRepo
	mchRepo    merchant.IMerchantRepo
	shopRepo   shop.IShopRepo
	manager    order.IOrderManager
	memberRepo member.IMemberRepo
	orderQuery *query.OrderQuery
	serviceUtil
}


func NewShoppingService(r order.IOrderRepo,
	cartRepo cart.ICartRepo, memberRepo member.IMemberRepo,
	prodRepo product.IProductRepo, goodsRepo item.IGoodsItemRepo,
	mchRepo merchant.IMerchantRepo, shopRepo shop.IShopRepo,
	orderQuery *query.OrderQuery) *orderServiceImpl {
	return &orderServiceImpl{
		repo:       r,
		prodRepo:   prodRepo,
		cartRepo:   cartRepo,
		memberRepo: memberRepo,
		itemRepo:   goodsRepo,
		mchRepo:    mchRepo,
		shopRepo:   shopRepo,
		manager:    r.Manager(),
		orderQuery: orderQuery,
	}
}

//  获取购物车
func (s *orderServiceImpl) getShoppingCart(buyerId int64, code string) cart.ICart {
	var c cart.ICart
	var cc cart.ICart
	if len(code) > 0 {
		cc = s.cartRepo.GetShoppingCartByKey(code)
	}
	// 如果传入会员编号，则合并购物车
	if buyerId > 0 {
		c = s.cartRepo.GetMyCart(buyerId, cart.KNormal)
		if cc != nil {
			rc := c.(cart.INormalCart)
			rc.Combine(cc)
			c.Save()
		}
		return c
	}
	// 如果只传入code,且购物车存在，直接返回。
	if cc != nil {
		return cc
	}
	// 不存在，则新建购物车
	c = s.cartRepo.NewNormalCart(code)
	//_, err := c.Save()
	//domain.HandleError(err, "service")
	return c
}

// 提交订单
func (s *orderServiceImpl) SubmitOrderV1(_ context.Context, r *proto.SubmitOrderRequest) (*proto.StringMap, error) {
	c := s.cartRepo.GetMyCart(r.BuyerId, cart.KWholesale)
	iData := orderImpl.NewPostedData(r.Data)
	rd, err := s.repo.Manager().SubmitWholesaleOrder(c, iData)
	if err != nil {
		return &proto.StringMap{Value: map[string]string{
			"error": err.Error(),
		}}, nil
	}
	return &proto.StringMap{Value: rd}, nil
}

func (s *orderServiceImpl) PrepareOrder(buyerId int64, addressId int64,
	cartCode string) (*order.ComplexOrder, error) {
	ic := s.getShoppingCart(buyerId, cartCode)
	o, err := s.manager.PrepareNormalOrder(ic)
	if err == nil {
		no := o.(order.INormalOrder)
		if addressId > 0 {
			err = no.SetAddress(addressId)
		} else {
			arr := s.memberRepo.GetDeliverAddress(buyerId)
			if len(arr) > 0 {
				err = no.SetAddress(arr[0].ID)
			}
		}
	}
	if err == nil {
		//log.Println("-------",o == nil,err)
		return o.Complex(), err
	}
	return nil, err
}

// 预生成订单，使用优惠券
func (s *orderServiceImpl) PrepareOrderWithCoupon(buyerId int64, cartCode string,
	addressId int64, subject string, couponCode string) (map[string]interface{}, error) {
	cart := s.getShoppingCart(buyerId, cartCode)
	o, err := s.manager.PrepareNormalOrder(cart)
	if err != nil {
		return nil, err
	}
	no := o.(order.INormalOrder)
	no.SetAddress(addressId)
	//todo: 应用优惠码
	v := o.Complex()
	buf := bytes.NewBufferString("")

	if o.Type() != order.TRetail {
		panic("not support order type")
	}
	io := o.(order.INormalOrder)
	for _, v := range io.GetCoupons() {
		buf.WriteString(v.GetDescribe())
		buf.WriteString("\n")
	}

	discountFee := v.ItemAmount - v.FinalAmount + v.DiscountAmount
	data := make(map[string]interface{})

	//　取消优惠券
	data["totalFee"] = v.ItemAmount
	data["fee"] = v.ItemAmount
	data["payFee"] = v.FinalAmount
	data["discountFee"] = discountFee
	data["expressFee"] = v.ExpressFee

	// 设置优惠券的信息
	if couponCode != "" {
		// 优惠券没有减金额
		if v.DiscountAmount == 0 {
			data["result"] = v.DiscountAmount != 0
			data["message"] = "优惠券无效"
		} else {
			// 成功应用优惠券
			data["couponFee"] = v.DiscountAmount
			data["couponDescribe"] = buf.String()
		}
	}

	return data, err
}

func (s *orderServiceImpl) SubmitOrder_V1(buyerId int64, cartCode string,
	addressId int64, subject string, couponCode string, balanceDiscount bool) (*order.SubmitReturnData, error) {
	c := s.getShoppingCart(buyerId, cartCode)
	_, rd, err := s.manager.SubmitOrder(c, addressId, couponCode, balanceDiscount)
	return rd, err
}


// 根据编号获取订单
func (s *orderServiceImpl) GetParentOrder(c context.Context, id *proto.OrderNoV2) (*proto.SParentOrder, error) {
	//c := s.manager.Unified(id.Value,false).Complex()
	//if c != nil {
	//	return parser.OrderDto(c), nil
	//}
	return nil, nil
}



// 获取订单和商品项信息
func (s *orderServiceImpl) GetOrder(_ context.Context, id *proto.OrderNoV2) (*proto.SSingleOrder, error) {
	c := s.manager.Unified(id.Value, true).Complex()
	if c != nil {
		return parser.OrderDto(c), nil
	}
	return nil, nil
}

// 获取子订单
func (s *orderServiceImpl) GetSubOrder(_ context.Context, id *proto.Int64) (*proto.SSingleOrder, error) {
	o := s.repo.GetSubOrder(id.Value)
	if o != nil {
		return parser.SubOrderDto(o), nil
	}
	return nil, nil
}


// 根据编号获取订单
func (s *orderServiceImpl) GetOrderById(id int64) *order.ComplexOrder {
	o := s.manager.GetOrderById(id)
	if o != nil {
		return o.Complex()
	}
	return nil
}

func (s *orderServiceImpl) GetOrderByNo(orderNo string) *order.ComplexOrder {
	o := s.manager.GetOrderByNo(orderNo)
	if o != nil {
		return o.Complex()
	}
	return nil
}

// 根据订单号获取子订单
func (s *orderServiceImpl) GetSubOrderByNo(_ context.Context, orderNo *proto.String) (*proto.SSingleOrder, error) {
	orderId := s.repo.GetOrderId(orderNo.Value, true)
	o := s.repo.GetSubOrder(orderId)
	if o != nil {
		return parser.SubOrderDto(o), nil
	}
	return nil, nil
}

// 获取订单商品项
func (s *orderServiceImpl) GetSubOrderItems(_ context.Context, subOrderId *proto.Int64) (*proto.ComplexItemsResponse, error) {
	list := s.repo.GetSubOrderItems(subOrderId.Value)
	arr := make([]*proto.SOrderItem, len(list))
	for i, v := range list {
		arr[i] = parser.SubOrderItemDto(v)
	}
	return &proto.ComplexItemsResponse{Value: arr}, nil
}

// 获取子订单及商品项
func (s *orderServiceImpl) GetSubOrderAndItems(id int64) (*order.NormalSubOrder, []*dto.OrderItem) {
	o := s.repo.GetSubOrder(id)
	if o == nil {
		return o, []*dto.OrderItem{}
	}
	return o, s.orderQuery.QueryOrderItems(id)
}

// 获取子订单及商品项
func (s *orderServiceImpl) GetSubOrderAndItemsByNo(orderNo string) (*order.NormalSubOrder, []*dto.OrderItem) {
	orderId := s.repo.GetOrderId(orderNo, true)
	o := s.repo.GetSubOrder(orderId)
	if o == nil {
		return o, []*dto.OrderItem{}
	}
	return o, s.orderQuery.QueryOrderItems(orderId)
}

// 提交订单
func (s *orderServiceImpl) SubmitTradeOrder(_ context.Context, r *proto.TradeOrderSubmitRequest) (*proto.Result, error) {
	if r.Order.ShopId <= 0 {
		mch := s.mchRepo.GetMerchant(int(r.Order.SellerId))
		if mch != nil {
			sp := mch.ShopManager().GetOnlineShop()
			if sp != nil {
				r.Order.ShopId = int64(sp.GetDomainId())
			} else {
				r.Order.ShopId = 1
			}
		}
	}
	io, err := s.manager.SubmitTradeOrder(parser.Order(r.Order), r.Rate)
	rs := s.result(err)
	rs.Data = map[string]string{
		"OrderId": strconv.Itoa(int(io.GetAggregateRootId())),
	}
	if err == nil {
		// 返回支付单号
		ro := io.(order.ITradeOrder)
		rs.Data["OrderNo"] = io.OrderNo()
		rs.Data["PaymentOrderNo"] = ro.GetPaymentOrder().TradeNo()
	}
	return rs, nil
}

// 交易单现金支付
func (s *orderServiceImpl) TradeOrderCashPay(_ context.Context, orderId *proto.Int64) (ro *proto.Result, err error) {
	o := s.manager.GetOrderById(orderId.Value)
	if o == nil || o.Type() != order.TTrade {
		err = order.ErrNoSuchOrder
	} else {
		io := o.(order.ITradeOrder)
		err = io.CashPay()
	}
	return s.result(err), nil
}

// 上传交易单发票
func (s *orderServiceImpl) TradeOrderUpdateTicket(_ context.Context, r *proto.TradeOrderTicketRequest) (rs *proto.Result, err error) {
	o := s.manager.GetOrderById(r.OrderId)
	if o == nil || o.Type() != order.TTrade {
		err = order.ErrNoSuchOrder
	} else {
		io := o.(order.ITradeOrder)
		err = io.UpdateTicket(r.Img)
	}
	return s.result(err), nil
}

// 取消订单
func (s *orderServiceImpl) CancelOrder(_ context.Context, r *proto.CancelOrderRequest) (*proto.Result, error) {
	c := s.manager.Unified(r.OrderNo, r.Sub)
	err := c.Cancel(r.Reason)
	return s.error(err), nil
}

// 确定订单
func (s *orderServiceImpl) ConfirmOrder(_ context.Context, r *proto.OrderNo) (*proto.Result, error) {
	c := s.manager.Unified(r.OrderNo, r.Sub)
	err := c.Confirm()
	return s.error(err), nil
}

// 备货完成
func (s *orderServiceImpl) PickUp(_ context.Context, r *proto.OrderNo) (*proto.Result, error) {
	c := s.manager.Unified(r.OrderNo, r.Sub)
	err := c.PickUp()
	return s.error(err), nil
}

// 订单发货,并记录配送服务商编号及单号
func (s *orderServiceImpl) Ship(_ context.Context, r *proto.OrderShipmentRequest) (*proto.Result, error) {
	c := s.manager.Unified(r.OrderNo, r.Sub)
	err := c.Ship(int32(r.ProviderId), r.ShipOrderNo)
	return s.error(err), nil
}

// 买家收货
func (s *orderServiceImpl) BuyerReceived(_ context.Context, r *proto.OrderNo) (*proto.Result, error) {
	c := s.manager.Unified(r.OrderNo, r.Sub)
	err := c.BuyerReceived()
	return s.error(err), nil
}

// 获取订单日志
func (s *orderServiceImpl) LogBytes(_ context.Context, r *proto.OrderNo) (*proto.String, error) {
	c := s.manager.Unified(r.OrderNo, r.Sub)
	return &proto.String{
		Value: string(c.LogBytes()),
	}, nil
}

//
//// 根据商品快照获取订单项
//func (s *orderServiceImpl) GetOrderItemBySnapshotId(orderId int64, snapshotId int32) *order.SubOrderItem {
//	return s.repo.GetOrderItemBySnapshotId(orderId, snapshotId)
//}

//// 根据商品快照获取订单项数据传输对象
//func (s *orderServiceImpl) GetOrderItemDtoBySnapshotId(orderId int64, snapshotId int32) *dto.OrderItem {
//	return s.repo.GetOrderItemDtoBySnapshotId(orderId, snapshotId)
//}

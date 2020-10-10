package order

import (
	"fmt"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-order/business/crm"
	b_order "github.com/gingerxman/ginger-order/business/order"
)

type PayedOrder struct {
	eel.RestResource
}

func (this *PayedOrder) Resource() string {
	return "order.payed_order"
}

func (this *PayedOrder) GetParameters() map[string][]string {
	return map[string][]string{
		"PUT": []string{
			"bid",
			"?channel",
		},
	}
}

func (this *PayedOrder) Put(ctx *eel.Context) {
	req := ctx.Request
	bid := req.GetString("bid")
	
	bCtx := ctx.GetBusinessContext()
	order := b_order.NewOrderRepository(bCtx).GetOrderByBid(bid)
	
	if order == nil {
		ctx.Response.Error("payed_order:invalid_order", fmt.Sprintf("invalid bid(%s)", bid))
	} else {
		channel := req.GetString("channel", "weixin")
		order.Pay(channel)
		
		crm.NewCrmService(bCtx).RecordOrder(order.Bid, order.Money.FinalMoney)
		
		ctx.Response.JSON(eel.Map{})
	}
}

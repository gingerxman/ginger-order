package routers

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/eel/handler/rest/console"
	"github.com/gingerxman/eel/handler/rest/op"
	"github.com/gingerxman/ginger-order/rest/consumption"
	"github.com/gingerxman/ginger-order/rest/dev"
	"github.com/gingerxman/ginger-order/rest/mall"
	"github.com/gingerxman/ginger-order/rest/mall/ship_info"
	"github.com/gingerxman/ginger-order/rest/order"
)

func init() {
	eel.RegisterResource(&console.Console{})
	eel.RegisterResource(&op.Health{})
	
	/*
	 order
	 */
	eel.RegisterResource(&order.Order{})
	eel.RegisterResource(&order.PayedOrder{})
	eel.RegisterResource(&order.CanceledOrder{})
	eel.RegisterResource(&order.ConfirmedInvoice{})
	eel.RegisterResource(&order.CanceledInvoice{})
	eel.RegisterResource(&order.ShippedInvoice{})
	eel.RegisterResource(&order.FinishedInvoice{})
	eel.RegisterResource(&order.Orders{})
	eel.RegisterResource(&order.CorpInvoices{})
	eel.RegisterResource(&order.UserOrders{})
	eel.RegisterResource(&order.OrderStatus{})
	eel.RegisterResource(&order.OrderRemark{})
	
	/*
	 consumption
	 */
	eel.RegisterResource(&consumption.UserConsumptionRecords{})
	
	/*
	 mall
	 */
	eel.RegisterResource(&mall.PurchaseData{})
	//ship_info
	eel.RegisterResource(&ship_info.ShipInfo{})
	eel.RegisterResource(&ship_info.ShipInfos{})
	eel.RegisterResource(&ship_info.DefaultShipInfo{})

	/*
	 dev
	 */
	eel.RegisterResource(&dev.BDDReset{})
}
package routers

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/eel/handler/rest/console"
	"github.com/gingerxman/eel/handler/rest/op"
	"github.com/gingerxman/ginger-order/rest/area"
	"github.com/gingerxman/ginger-order/rest/consumption"
	"github.com/gingerxman/ginger-order/rest/dev"
	"github.com/gingerxman/ginger-order/rest/mall"
	"github.com/gingerxman/ginger-order/rest/mall/ship_info"
	"github.com/gingerxman/ginger-order/rest/material"
	"github.com/gingerxman/ginger-order/rest/order"
)

func init() {
	eel.RegisterResource(&console.Console{})
	eel.RegisterResource(&op.Health{})
	
	/*
	 product
	 */
	//category
	//eel.RegisterResource(&product.Category{})
	//eel.RegisterResource(&product.DisabledCategory{})
	//eel.RegisterResource(&product.SubCategories{})
	////label
	//eel.RegisterResource(&product.ProductLabel{})
	//eel.RegisterResource(&product.ProductLabels{})
	//eel.RegisterResource(&product.CorpProductLabels{})
	//eel.RegisterResource(&product.DisabledCategory{})
	////property
	//eel.RegisterResource(&product.ProductProperty{})
	//eel.RegisterResource(&product.ProductPropertyValue{})
	//eel.RegisterResource(&product.ProductProperties{})
	//eel.RegisterResource(&product.CorpProductProperties{})
	////product
	//eel.RegisterResource(&product.Product{})
	//eel.RegisterResource(&product.OffshelfProducts{})
	//eel.RegisterResource(&product.OnshelfProducts{})
	//eel.RegisterResource(&product.CorpProducts{})
	//eel.RegisterResource(&product.CreateOptions{})
	
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
	 material
	*/
	eel.RegisterResource(&material.Image{})
	
	/*
	 area
	 */
	eel.RegisterResource(&area.Area{})
	eel.RegisterResource(&area.AreaCode{})
	eel.RegisterResource(&area.YouzanAreaList{})

	/*
	 dev
	 */
	eel.RegisterResource(&dev.BDDReset{})
}
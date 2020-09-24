package order

import (
	"context"
	
	"github.com/gingerxman/eel"
	m_order "github.com/gingerxman/ginger-order/models/order"
)

type OrderProductRepository struct {
	eel.ServiceBase
}

func NewOrderProductRepository(ctx context.Context) *OrderProductRepository {
	service := new(OrderProductRepository)
	service.Ctx = ctx
	return service
}

func (this *OrderProductRepository) GetOrderProducts(invoiceIds []int) []*OrderProduct {
	o := eel.GetOrmFromContext(this.Ctx)
	
	var models []*m_order.OrderHasProduct
	db := o.Model(&m_order.OrderHasProduct{}).Where("order_id__in", invoiceIds).Find(&models)
	err := db.Error
	
	if err != nil {
		eel.Logger.Error(err)
		return make([]*OrderProduct, 0)
	}
	
	orderProducts := make([]*OrderProduct, 0)
	productIds := make([]int, 0)
	for _, model := range models {
		product := &OrderProduct{}
		product.ProductId = model.ProductId
		product.Name = model.ProductName
		product.Sku = model.ProductSkuName
		product.SkuDisplayName = model.ProductSkuDisplayName
		product.Price = model.Price
		product.PurchaseCount = model.Count
		product.Thumbnail = model.Thumbnail
		product.Weight = model.Weight
		product.OrderId = model.OrderId
		product.SupplierId = model.SupplierId
		
		orderProducts = append(orderProducts, product)
		productIds = append(productIds, model.ProductId)
	}
	
	return orderProducts
}

func init() {
}

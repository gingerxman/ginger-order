package resource

import (
	"context"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-order/business/product"
	
	"github.com/gingerxman/ginger-order/business"
)

type ProductResourceAllocator struct {
	eel.ServiceBase
}

func NewProductResourceAllocator(ctx context.Context) business.IResourceAllocator {
	service := new(ProductResourceAllocator)
	service.Ctx = ctx
	return service
}

//Allocate 申请商品资源，减少库存
func (this *ProductResourceAllocator) Allocate(resource business.IResource, newOrder business.IOrder) error {
	productResource := resource.(*ProductResource)
	sku := productResource.GetProduct().GetSku(productResource.Sku)
	
	err := product.NewProductRepository(this.Ctx).UseSkuStocks(sku.Id, productResource.Count)
	if err != nil {
		eel.Logger.Error(err)
		return err
	}
	
	return nil
}

//Release 释放商品资源，恢复库存
func (this *ProductResourceAllocator) Release(resource business.IResource) {
	productResource := resource.(*ProductResource)
	sku := productResource.GetProduct().GetSku(productResource.Sku)
	
	err := product.NewProductRepository(this.Ctx).AddSkuStocks(sku.Id, productResource.Count)
	if err != nil {
		eel.Logger.Error(err)
	}
}


func init() {
}

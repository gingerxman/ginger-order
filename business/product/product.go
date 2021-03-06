package product

import (
	"github.com/gingerxman/eel"
)

type sku struct {
	Id int
	Name string
	DisplayName string
	Code string
	Price int
	Stocks int
}

func (this *sku) CanAffordStock(stock int) bool {
	if this.Stocks >= stock {
		return true
	}
	
	return false
}

type sLogisticsInfo struct {
	LimitZoneType string `json:"limit_zone_type"`
	PostageType string `json:"postage_type"`
	UnifiedPostageMoney int `json:"unified_postage_money"`
}

type sPointProductInfo struct {
	IsPointProduct bool
	PointPrice int
	MoneyPrice int
}

type Product struct {
	eel.EntityBase
	Id int
	SourceProductId int
	RawProductId int
	CorpId int
	SupplierId int
	Status string
	
	Name string
	Thumbnail string
	IsDeleted bool
	
	//logistics info
	LogisticsInfo *sLogisticsInfo

	//foreign key
	ProductUsableImoneyId int //refer to product_usable_imoney
	//ProductUsableImoney *ProductUsableImoney
	Skus []*sku
	
	// 积分商品信息
	PointProductInfo *sPointProductInfo
}

func (this *Product) GetSku(skuName string) *sku {
	for _, sku := range this.Skus {
		if sku.Name == skuName {
			return sku
		}
	}
	
	return nil
}

func (this *Product) UseUnifiedPostage() bool {
	return this.LogisticsInfo.PostageType == "unified"
}

func (this *Product) GetUnifiedPostageMoney() int {
	return this.LogisticsInfo.UnifiedPostageMoney
}

func (this *Product) CanPurchase() bool {
	return this.Status == "on_shelf"
}

// IsPointProduct: 是否是积分商品
func (this *Product) IsPointProduct() bool {
	return this.PointProductInfo != nil
}

func (this *Product) GetPromotionType() string {
	if this.IsPointProduct() {
		return "point_product"
	}
	
	return ""
}

func init() {
}

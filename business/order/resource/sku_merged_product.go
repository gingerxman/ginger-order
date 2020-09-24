package resource

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-order/business/product"
)

type SkuMergedProduct struct {
	eel.EntityBase
	
	TotalCount int
	TotalWeight float64
	TotalPrice float64
	PoolProduct *product.Product
}

func init() {
}

package order

import (
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/config"
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-order/business/account"
	"github.com/gingerxman/ginger-order/business/order"
	sentinel "github.com/alibaba/sentinel-golang/api"
)

type Orders struct {
	eel.RestResource
}

func (this *Orders) Resource() string {
	return "order.orders"
}

func (this *Orders) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"?filters:json"},
	}
}


func (this *Orders) Get(ctx *eel.Context) {
	req := ctx.Request
	bCtx := ctx.GetBusinessContext()

	filters := req.GetOrmFilters()
	pageInfo := req.GetPageInfo()

	user := account.GetUserFromContext(bCtx)
	if v, ok := filters["target"]; !ok || v.(string) != "all_corps"{
		filters["user_id"] = user.Id
	}

	orders, nextPageInfo := order.NewOrderRepository(bCtx).GetPagedOrders(filters, pageInfo, "-created_at")

	//fillOp
	//fillOptions := req.GetJSON("fill_options")
	//if fillOptions["with_delivery_items"] == nil && fillOptions["with_invoice"] == nil{
	//	fillOptions["with_invoice"] = map[string]interface{}{
	//		"with_products": true,
	//	}
	//}
	fillOptions := eel.Map{}
	fillOptions["with_invoice"] = map[string]interface{}{
		"with_products": true,
	}
	order.NewFillOrderService(bCtx).Fill(orders, fillOptions)

	rows := order.NewEncodeOrderService(bCtx).EncodeMany(orders)
	ctx.Response.JSON(eel.Map{
		"orders": rows,
		"pageinfo": nextPageInfo.ToMap(),
	})
}

func init() {
	err := sentinel.InitWithConfig(&config.Entity{
		Version: "v1",
		Sentinel: config.SentinelConfig{
			App: struct {
				Name string
				Type int32
			}{
				Name: "ginger-order",
				Type: 1,
			},
			Log: config.LogConfig{
				UsePid: false,
				Metric: config.MetricLogConfig{
					SingleFileMaxSize: config.DefaultMetricLogSingleFileMaxSize,
					MaxFileCount:      config.DefaultMetricLogMaxFileAmount,
					FlushIntervalSec:  config.DefaultMetricLogFlushIntervalSec,
				},
			},
			Stat: config.StatConfig{
				GlobalStatisticSampleCountTotal: base.DefaultSampleCountTotal,
				GlobalStatisticIntervalMsTotal:  base.DefaultIntervalMsTotal,
				MetricStatisticSampleCount:      base.DefaultSampleCount,
				MetricStatisticIntervalMs:       base.DefaultIntervalMs,
				System: config.SystemStatConfig{
					CollectIntervalMs: config.DefaultSystemStatCollectIntervalMs,
				},
			},
			UseCacheTime: true,
		},
	})
	
	if err != nil {
		eel.Logger.Error(err)
	}
}

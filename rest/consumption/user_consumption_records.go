package consumption

import (
	"github.com/gingerxman/eel"
	"github.com/gingerxman/ginger-order/business/consumption"
)

type UserConsumptionRecords struct {
	eel.RestResource
}

func (this *UserConsumptionRecords) Resource() string {
	return "consumption.user_consumption_records"
}

func (this *UserConsumptionRecords) GetParameters() map[string][]string {
	return map[string][]string{
		"GET": []string{"user_ids:json-array"},
	}
}


func (this *UserConsumptionRecords) Get(ctx *eel.Context) {
	req := ctx.Request
	bCtx := ctx.GetBusinessContext()

	userIds := req.GetIntArray("user_ids")
	records := consumption.NewConsumptionRecordRepository(bCtx).GetRecordsForUsers(userIds)

	consumption.NewFillConsumptionRecordService(bCtx).Fill(records, eel.Map{})
	datas := consumption.NewEncodeConsumptionRecordService(bCtx).EncodeMany(records)

	ctx.Response.JSON(eel.Map{
		"records": datas,
	})
}


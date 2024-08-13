package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraType"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type MultiSpectraTypeSearch struct {
	MultiSpectraType.MultiSpectraType
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/NestAirlinePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type NestAirlineSearch struct {
	NestAirlinePkg.NestAirline
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	NestId         string     `json:"nestId" form:"nestId"`
	request.PageInfo
}

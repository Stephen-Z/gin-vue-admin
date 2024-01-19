package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/ProblemRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type ProblemRecordSearch struct {
	ProblemRecordPkg.ProblemRecord
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

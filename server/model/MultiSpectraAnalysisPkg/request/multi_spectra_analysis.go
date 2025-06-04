package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraAnalysisPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type MultiSpectraAnalysisSearch struct {
	MultiSpectraAnalysisPkg.MultiSpectraAnalysis
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	request.PageInfo
}

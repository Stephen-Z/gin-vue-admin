package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/AerialPhotographyResult"
	"github.com/flipped-aurora/gin-vue-admin/server/router/AerialPhotographyResultPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/FlyResultPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/MultiSpectraAnalysisPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/MultiSpectraPlate"
	"github.com/flipped-aurora/gin-vue-admin/server/router/MultiSpectraType"
	"github.com/flipped-aurora/gin-vue-admin/server/router/NestAirlinePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/NestExecRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/NestInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Nestrolepkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/ProblemRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/Theme"
	"github.com/flipped-aurora/gin-vue-admin/server/router/UserTeemlinkPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System                     system.RouterGroup
	Example                    example.RouterGroup
	NestInfo                   NestInfo.RouterGroup
	Nestrolepkg                Nestrolepkg.RouterGroup
	NestAirlinePkg             NestAirlinePkg.RouterGroup
	NestExecRecordPkg          NestExecRecordPkg.RouterGroup
	FlyResultPkg               FlyResultPkg.RouterGroup
	UserTeemlinkPkg            UserTeemlinkPkg.RouterGroup
	AerialPhotographyResult    AerialPhotographyResult.RouterGroup
	AerialPhotographyResultPkg AerialPhotographyResultPkg.RouterGroup
	Theme                      Theme.RouterGroup
	ProblemRecordPkg           ProblemRecordPkg.RouterGroup
	MultiSpectraAnalysisPkg    MultiSpectraAnalysisPkg.RouterGroup
	MultiSpectraPlate          MultiSpectraPlate.RouterGroup
	MultiSpectraType           MultiSpectraType.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

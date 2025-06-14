package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/AerialPhotographyResultPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/FlyResultPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/MultiSpectraAnalysisPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/MultiSpectraPlate"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/MultiSpectraType"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/NestAirlinePkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/NestExecRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/NestInfo"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Nestrolepkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/ProblemRecordPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/Theme"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/UserTeemlinkPkg"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup                     system.ApiGroup
	ExampleApiGroup                    example.ApiGroup
	NestInfoApiGroup                   NestInfo.ApiGroup
	NestrolepkgApiGroup                Nestrolepkg.ApiGroup
	NestAirlinePkgApiGroup             NestAirlinePkg.ApiGroup
	NestExecRecordPkgApiGroup          NestExecRecordPkg.ApiGroup
	FlyResultPkgApiGroup               FlyResultPkg.ApiGroup
	UserTeemlinkPkgApiGroup            UserTeemlinkPkg.ApiGroup
	AerialPhotographyResultPkgApiGroup AerialPhotographyResultPkg.ApiGroup
	ThemeApiGroup                      Theme.ApiGroup
	ProblemRecordPkgApiGroup           ProblemRecordPkg.ApiGroup
	MultiSpectraAnalysisPkgApiGroup    MultiSpectraAnalysisPkg.ApiGroup
	MultiSpectraPlateApiGroup          MultiSpectraPlate.ApiGroup
	MultiSpectraTypeApiGroup           MultiSpectraType.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

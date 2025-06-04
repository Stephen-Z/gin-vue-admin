package MultiSpectraAnalysisPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MultiSpectraAnalysisRouter struct {
}

// InitMultiSpectraAnalysisRouter 初始化 MultiSpectraAnalysis 路由信息
func (s *MultiSpectraAnalysisRouter) InitMultiSpectraAnalysisRouter(Router *gin.RouterGroup) {
	MtSpectraAlyRouter := Router.Group("MtSpectraAly").Use(middleware.OperationRecord())
	MtSpectraAlyRouterWithoutRecord := Router.Group("MtSpectraAly")
	var MtSpectraAlyApi = v1.ApiGroupApp.MultiSpectraAnalysisPkgApiGroup.MultiSpectraAnalysisApi
	{
		MtSpectraAlyRouter.POST("createMultiSpectraAnalysis", MtSpectraAlyApi.CreateMultiSpectraAnalysis)             // 新建MultiSpectraAnalysis
		MtSpectraAlyRouter.DELETE("deleteMultiSpectraAnalysis", MtSpectraAlyApi.DeleteMultiSpectraAnalysis)           // 删除MultiSpectraAnalysis
		MtSpectraAlyRouter.DELETE("deleteMultiSpectraAnalysisByIds", MtSpectraAlyApi.DeleteMultiSpectraAnalysisByIds) // 批量删除MultiSpectraAnalysis
		MtSpectraAlyRouter.PUT("updateMultiSpectraAnalysis", MtSpectraAlyApi.UpdateMultiSpectraAnalysis)              // 更新MultiSpectraAnalysis
	}
	{
		MtSpectraAlyRouterWithoutRecord.GET("findMultiSpectraAnalysis", MtSpectraAlyApi.FindMultiSpectraAnalysis)       // 根据ID获取MultiSpectraAnalysis
		MtSpectraAlyRouterWithoutRecord.GET("getMultiSpectraAnalysisList", MtSpectraAlyApi.GetMultiSpectraAnalysisList) // 获取MultiSpectraAnalysis列表
	}
}

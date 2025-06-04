package MultiSpectraType

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MultiSpectraTypeRouter struct {
}

// InitMultiSpectraTypeRouter 初始化 MultiSpectraType 路由信息
func (s *MultiSpectraTypeRouter) InitMultiSpectraTypeRouter(Router *gin.RouterGroup) {
	MtSpectraTypeRouter := Router.Group("MtSpectraType").Use(middleware.OperationRecord())
	MtSpectraTypeRouterWithoutRecord := Router.Group("MtSpectraType")
	var MtSpectraTypeApi = v1.ApiGroupApp.MultiSpectraTypeApiGroup.MultiSpectraTypeApi
	{
		MtSpectraTypeRouter.POST("createMultiSpectraType", MtSpectraTypeApi.CreateMultiSpectraType)             // 新建MultiSpectraType
		MtSpectraTypeRouter.DELETE("deleteMultiSpectraType", MtSpectraTypeApi.DeleteMultiSpectraType)           // 删除MultiSpectraType
		MtSpectraTypeRouter.DELETE("deleteMultiSpectraTypeByIds", MtSpectraTypeApi.DeleteMultiSpectraTypeByIds) // 批量删除MultiSpectraType
		MtSpectraTypeRouter.PUT("updateMultiSpectraType", MtSpectraTypeApi.UpdateMultiSpectraType)              // 更新MultiSpectraType
	}
	{
		MtSpectraTypeRouterWithoutRecord.GET("findMultiSpectraType", MtSpectraTypeApi.FindMultiSpectraType)              // 根据ID获取MultiSpectraType
		MtSpectraTypeRouterWithoutRecord.GET("getMultiSpectraTypeList", MtSpectraTypeApi.GetMultiSpectraTypeList)        // 获取MultiSpectraType列表
		MtSpectraTypeRouterWithoutRecord.POST("createMultiSpectraTypeZhnc", MtSpectraTypeApi.CreateMultiSpectraTypeZhnc) // 获取MultiSpectraType列表
	}
}

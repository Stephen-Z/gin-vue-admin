package MultiSpectraPlate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MultiSpectraPlateRouter struct {
}

// InitMultiSpectraPlateRouter 初始化 MultiSpectraPlate 路由信息
func (s *MultiSpectraPlateRouter) InitMultiSpectraPlateRouter(Router *gin.RouterGroup) {
	MtpectraPlateRouter := Router.Group("MtpectraPlate").Use(middleware.OperationRecord())
	MtpectraPlateRouterWithoutRecord := Router.Group("MtpectraPlate")
	var MtpectraPlateApi = v1.ApiGroupApp.MultiSpectraPlateApiGroup.MultiSpectraPlateApi
	{
		MtpectraPlateRouter.POST("createMultiSpectraPlate", MtpectraPlateApi.CreateMultiSpectraPlate)             // 新建MultiSpectraPlate
		MtpectraPlateRouter.DELETE("deleteMultiSpectraPlate", MtpectraPlateApi.DeleteMultiSpectraPlate)           // 删除MultiSpectraPlate
		MtpectraPlateRouter.DELETE("deleteMultiSpectraPlateByIds", MtpectraPlateApi.DeleteMultiSpectraPlateByIds) // 批量删除MultiSpectraPlate
		MtpectraPlateRouter.PUT("updateMultiSpectraPlate", MtpectraPlateApi.UpdateMultiSpectraPlate)              // 更新MultiSpectraPlate
	}
	{
		MtpectraPlateRouterWithoutRecord.GET("findMultiSpectraPlate", MtpectraPlateApi.FindMultiSpectraPlate)       // 根据ID获取MultiSpectraPlate
		MtpectraPlateRouterWithoutRecord.GET("getMultiSpectraPlateList", MtpectraPlateApi.GetMultiSpectraPlateList) // 获取MultiSpectraPlate列表
	}
}

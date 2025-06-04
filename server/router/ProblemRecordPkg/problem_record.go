package ProblemRecordPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProblemRecordRouter struct {
}

// InitProblemRecordRouter 初始化 ProblemRecord 路由信息
func (s *ProblemRecordRouter) InitProblemRecordRouter(Router *gin.RouterGroup) {
	pbRecordRouter := Router.Group("pbRecord").Use(middleware.OperationRecord())
	pbRecordRouterWithoutRecord := Router.Group("pbRecord")
	var pbRecordApi = v1.ApiGroupApp.ProblemRecordPkgApiGroup.ProblemRecordApi
	{
		pbRecordRouter.POST("createProblemRecord", pbRecordApi.CreateProblemRecord)             // 新建ProblemRecord
		pbRecordRouter.DELETE("deleteProblemRecord", pbRecordApi.DeleteProblemRecord)           // 删除ProblemRecord
		pbRecordRouter.DELETE("deleteProblemRecordByIds", pbRecordApi.DeleteProblemRecordByIds) // 批量删除ProblemRecord
		pbRecordRouter.PUT("updateProblemRecord", pbRecordApi.UpdateProblemRecord)              // 更新ProblemRecord
	}
	{
		pbRecordRouterWithoutRecord.GET("findProblemRecord", pbRecordApi.FindProblemRecord)                   // 根据ID获取ProblemRecord
		pbRecordRouterWithoutRecord.GET("getProblemRecordList", pbRecordApi.GetProblemRecordList)             // 获取ProblemRecord列表
		pbRecordRouterWithoutRecord.GET("getProblemRecordListByUser", pbRecordApi.GetProblemRecordListByUser) // 根据用户权限分页获取ProblemRecord列表
	}
}

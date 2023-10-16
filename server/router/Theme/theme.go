package Theme

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ThemeRouter struct {
}

// InitThemeRouter 初始化 Theme 路由信息
func (s *ThemeRouter) InitThemeRouter(Router *gin.RouterGroup) {
	themeRouter := Router.Group("theme").Use(middleware.OperationRecord())
	themeRouterWithoutRecord := Router.Group("theme")
	var themeApi = v1.ApiGroupApp.ThemeApiGroup.ThemeApi
	{
		themeRouter.POST("createTheme", themeApi.CreateTheme)             // 新建Theme
		themeRouter.DELETE("deleteTheme", themeApi.DeleteTheme)           // 删除Theme
		themeRouter.DELETE("deleteThemeByIds", themeApi.DeleteThemeByIds) // 批量删除Theme
		themeRouter.PUT("updateTheme", themeApi.UpdateTheme)              // 更新Theme
	}
	{
		themeRouterWithoutRecord.GET("findTheme", themeApi.FindTheme)                 // 根据ID获取Theme
		themeRouterWithoutRecord.GET("getThemeList", themeApi.GetThemeList)           // 获取Theme列表
		themeRouterWithoutRecord.GET("findThemeByRoleId", themeApi.FindThemeByRoleId) // 获取角色id获取Theme
	}
}

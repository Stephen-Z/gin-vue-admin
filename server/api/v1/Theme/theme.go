package Theme

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/Theme"
	ThemeReq "github.com/flipped-aurora/gin-vue-admin/server/model/Theme/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ThemeApi struct {
}

var themeService = service.ServiceGroupApp.ThemeServiceGroup.ThemeService

// CreateTheme 创建Theme
// @Tags Theme
// @Summary 创建Theme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Theme.Theme true "创建Theme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /theme/createTheme [post]
func (themeApi *ThemeApi) CreateTheme(c *gin.Context) {
	var theme Theme.Theme
	err := c.ShouldBindJSON(&theme)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := themeService.CreateTheme(&theme); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteTheme 删除Theme
// @Tags Theme
// @Summary 删除Theme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Theme.Theme true "删除Theme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /theme/deleteTheme [delete]
func (themeApi *ThemeApi) DeleteTheme(c *gin.Context) {
	var theme Theme.Theme
	err := c.ShouldBindJSON(&theme)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := themeService.DeleteTheme(theme); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteThemeByIds 批量删除Theme
// @Tags Theme
// @Summary 批量删除Theme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Theme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /theme/deleteThemeByIds [delete]
func (themeApi *ThemeApi) DeleteThemeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := themeService.DeleteThemeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateTheme 更新Theme
// @Tags Theme
// @Summary 更新Theme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body Theme.Theme true "更新Theme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /theme/updateTheme [put]
func (themeApi *ThemeApi) UpdateTheme(c *gin.Context) {
	var theme Theme.Theme
	err := c.ShouldBindJSON(&theme)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := themeService.UpdateTheme(theme); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindTheme 用id查询Theme
// @Tags Theme
// @Summary 用id查询Theme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Theme.Theme true "用id查询Theme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /theme/findTheme [get]
func (themeApi *ThemeApi) FindTheme(c *gin.Context) {
	var theme Theme.Theme
	err := c.ShouldBindQuery(&theme)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retheme, err := themeService.GetTheme(theme.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retheme": retheme}, c)
	}
}

// GetThemeList 分页获取Theme列表
// @Tags Theme
// @Summary 分页获取Theme列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ThemeReq.ThemeSearch true "分页获取Theme列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /theme/getThemeList [get]
func (themeApi *ThemeApi) GetThemeList(c *gin.Context) {
	var pageInfo ThemeReq.ThemeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := themeService.GetThemeInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// FindThemes 用角色id查询Theme
// @Tags Themes
// @Summary 用id查询Themes
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query Theme.Theme true "用角色id查询Theme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /theme/findTheme [get]
func (themeApi *ThemeApi) FindThemeByRoleId(c *gin.Context) {
	var theme Theme.Theme
	err := c.ShouldBindQuery(&theme)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if retheme, err := themeService.GetThemeByRoleId(theme.UserRoles); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"retheme": retheme}, c)
	}
}

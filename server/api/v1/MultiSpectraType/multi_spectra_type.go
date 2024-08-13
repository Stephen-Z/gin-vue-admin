package MultiSpectraType

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraType"
	MultiSpectraTypeReq "github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraType/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MultiSpectraTypeApi struct {
}

var MtSpectraTypeService = service.ServiceGroupApp.MultiSpectraTypeServiceGroup.MultiSpectraTypeService

// CreateMultiSpectraType 创建MultiSpectraType
// @Tags MultiSpectraType
// @Summary 创建MultiSpectraType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MultiSpectraType.MultiSpectraType true "创建MultiSpectraType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MtSpectraType/createMultiSpectraType [post]
func (MtSpectraTypeApi *MultiSpectraTypeApi) CreateMultiSpectraType(c *gin.Context) {
	var MtSpectraType MultiSpectraType.MultiSpectraType
	err := c.ShouldBindJSON(&MtSpectraType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MtSpectraTypeService.CreateMultiSpectraType(&MtSpectraType); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMultiSpectraType 删除MultiSpectraType
// @Tags MultiSpectraType
// @Summary 删除MultiSpectraType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MultiSpectraType.MultiSpectraType true "删除MultiSpectraType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MtSpectraType/deleteMultiSpectraType [delete]
func (MtSpectraTypeApi *MultiSpectraTypeApi) DeleteMultiSpectraType(c *gin.Context) {
	var MtSpectraType MultiSpectraType.MultiSpectraType
	err := c.ShouldBindJSON(&MtSpectraType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MtSpectraTypeService.DeleteMultiSpectraType(MtSpectraType); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMultiSpectraTypeByIds 批量删除MultiSpectraType
// @Tags MultiSpectraType
// @Summary 批量删除MultiSpectraType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MultiSpectraType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /MtSpectraType/deleteMultiSpectraTypeByIds [delete]
func (MtSpectraTypeApi *MultiSpectraTypeApi) DeleteMultiSpectraTypeByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MtSpectraTypeService.DeleteMultiSpectraTypeByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMultiSpectraType 更新MultiSpectraType
// @Tags MultiSpectraType
// @Summary 更新MultiSpectraType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MultiSpectraType.MultiSpectraType true "更新MultiSpectraType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MtSpectraType/updateMultiSpectraType [put]
func (MtSpectraTypeApi *MultiSpectraTypeApi) UpdateMultiSpectraType(c *gin.Context) {
	var MtSpectraType MultiSpectraType.MultiSpectraType
	err := c.ShouldBindJSON(&MtSpectraType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MtSpectraTypeService.UpdateMultiSpectraType(MtSpectraType); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMultiSpectraType 用id查询MultiSpectraType
// @Tags MultiSpectraType
// @Summary 用id查询MultiSpectraType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MultiSpectraType.MultiSpectraType true "用id查询MultiSpectraType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MtSpectraType/findMultiSpectraType [get]
func (MtSpectraTypeApi *MultiSpectraTypeApi) FindMultiSpectraType(c *gin.Context) {
	var MtSpectraType MultiSpectraType.MultiSpectraType
	err := c.ShouldBindQuery(&MtSpectraType)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reMtSpectraType, err := MtSpectraTypeService.GetMultiSpectraType(MtSpectraType.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reMtSpectraType": reMtSpectraType}, c)
	}
}

// GetMultiSpectraTypeList 分页获取MultiSpectraType列表
// @Tags MultiSpectraType
// @Summary 分页获取MultiSpectraType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MultiSpectraTypeReq.MultiSpectraTypeSearch true "分页获取MultiSpectraType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MtSpectraType/getMultiSpectraTypeList [get]
func (MtSpectraTypeApi *MultiSpectraTypeApi) GetMultiSpectraTypeList(c *gin.Context) {
	var pageInfo MultiSpectraTypeReq.MultiSpectraTypeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := MtSpectraTypeService.GetMultiSpectraTypeInfoList(pageInfo); err != nil {
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

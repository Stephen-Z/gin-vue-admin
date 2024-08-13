package MultiSpectraPlate

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraPlate"
	MultiSpectraPlateReq "github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraPlate/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MultiSpectraPlateApi struct {
}

var MtpectraPlateService = service.ServiceGroupApp.MultiSpectraPlateServiceGroup.MultiSpectraPlateService

// CreateMultiSpectraPlate 创建MultiSpectraPlate
// @Tags MultiSpectraPlate
// @Summary 创建MultiSpectraPlate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MultiSpectraPlate.MultiSpectraPlate true "创建MultiSpectraPlate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MtpectraPlate/createMultiSpectraPlate [post]
func (MtpectraPlateApi *MultiSpectraPlateApi) CreateMultiSpectraPlate(c *gin.Context) {
	var MtpectraPlate MultiSpectraPlate.MultiSpectraPlate
	err := c.ShouldBindJSON(&MtpectraPlate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MtpectraPlateService.CreateMultiSpectraPlate(&MtpectraPlate); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMultiSpectraPlate 删除MultiSpectraPlate
// @Tags MultiSpectraPlate
// @Summary 删除MultiSpectraPlate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MultiSpectraPlate.MultiSpectraPlate true "删除MultiSpectraPlate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MtpectraPlate/deleteMultiSpectraPlate [delete]
func (MtpectraPlateApi *MultiSpectraPlateApi) DeleteMultiSpectraPlate(c *gin.Context) {
	var MtpectraPlate MultiSpectraPlate.MultiSpectraPlate
	err := c.ShouldBindJSON(&MtpectraPlate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MtpectraPlateService.DeleteMultiSpectraPlate(MtpectraPlate); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMultiSpectraPlateByIds 批量删除MultiSpectraPlate
// @Tags MultiSpectraPlate
// @Summary 批量删除MultiSpectraPlate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MultiSpectraPlate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /MtpectraPlate/deleteMultiSpectraPlateByIds [delete]
func (MtpectraPlateApi *MultiSpectraPlateApi) DeleteMultiSpectraPlateByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MtpectraPlateService.DeleteMultiSpectraPlateByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMultiSpectraPlate 更新MultiSpectraPlate
// @Tags MultiSpectraPlate
// @Summary 更新MultiSpectraPlate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MultiSpectraPlate.MultiSpectraPlate true "更新MultiSpectraPlate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MtpectraPlate/updateMultiSpectraPlate [put]
func (MtpectraPlateApi *MultiSpectraPlateApi) UpdateMultiSpectraPlate(c *gin.Context) {
	var MtpectraPlate MultiSpectraPlate.MultiSpectraPlate
	err := c.ShouldBindJSON(&MtpectraPlate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MtpectraPlateService.UpdateMultiSpectraPlate(MtpectraPlate); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMultiSpectraPlate 用id查询MultiSpectraPlate
// @Tags MultiSpectraPlate
// @Summary 用id查询MultiSpectraPlate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MultiSpectraPlate.MultiSpectraPlate true "用id查询MultiSpectraPlate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MtpectraPlate/findMultiSpectraPlate [get]
func (MtpectraPlateApi *MultiSpectraPlateApi) FindMultiSpectraPlate(c *gin.Context) {
	var MtpectraPlate MultiSpectraPlate.MultiSpectraPlate
	err := c.ShouldBindQuery(&MtpectraPlate)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reMtpectraPlate, err := MtpectraPlateService.GetMultiSpectraPlate(MtpectraPlate.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reMtpectraPlate": reMtpectraPlate}, c)
	}
}

// GetMultiSpectraPlateList 分页获取MultiSpectraPlate列表
// @Tags MultiSpectraPlate
// @Summary 分页获取MultiSpectraPlate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MultiSpectraPlateReq.MultiSpectraPlateSearch true "分页获取MultiSpectraPlate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MtpectraPlate/getMultiSpectraPlateList [get]
func (MtpectraPlateApi *MultiSpectraPlateApi) GetMultiSpectraPlateList(c *gin.Context) {
	var pageInfo MultiSpectraPlateReq.MultiSpectraPlateSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := MtpectraPlateService.GetMultiSpectraPlateInfoList(pageInfo); err != nil {
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

package MultiSpectraAnalysisPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraAnalysisPkg"
	MultiSpectraAnalysisPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/MultiSpectraAnalysisPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MultiSpectraAnalysisApi struct {
}

var MtSpectraAlyService = service.ServiceGroupApp.MultiSpectraAnalysisPkgServiceGroup.MultiSpectraAnalysisService

// CreateMultiSpectraAnalysis 创建MultiSpectraAnalysis
// @Tags MultiSpectraAnalysis
// @Summary 创建MultiSpectraAnalysis
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MultiSpectraAnalysisPkg.MultiSpectraAnalysis true "创建MultiSpectraAnalysis"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MtSpectraAly/createMultiSpectraAnalysis [post]
func (MtSpectraAlyApi *MultiSpectraAnalysisApi) CreateMultiSpectraAnalysis(c *gin.Context) {
	var MtSpectraAly MultiSpectraAnalysisPkg.MultiSpectraAnalysis
	err := c.ShouldBindJSON(&MtSpectraAly)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MtSpectraAlyService.CreateMultiSpectraAnalysis(&MtSpectraAly); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteMultiSpectraAnalysis 删除MultiSpectraAnalysis
// @Tags MultiSpectraAnalysis
// @Summary 删除MultiSpectraAnalysis
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MultiSpectraAnalysisPkg.MultiSpectraAnalysis true "删除MultiSpectraAnalysis"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MtSpectraAly/deleteMultiSpectraAnalysis [delete]
func (MtSpectraAlyApi *MultiSpectraAnalysisApi) DeleteMultiSpectraAnalysis(c *gin.Context) {
	var MtSpectraAly MultiSpectraAnalysisPkg.MultiSpectraAnalysis
	err := c.ShouldBindJSON(&MtSpectraAly)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MtSpectraAlyService.DeleteMultiSpectraAnalysis(MtSpectraAly); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteMultiSpectraAnalysisByIds 批量删除MultiSpectraAnalysis
// @Tags MultiSpectraAnalysis
// @Summary 批量删除MultiSpectraAnalysis
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MultiSpectraAnalysis"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /MtSpectraAly/deleteMultiSpectraAnalysisByIds [delete]
func (MtSpectraAlyApi *MultiSpectraAnalysisApi) DeleteMultiSpectraAnalysisByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MtSpectraAlyService.DeleteMultiSpectraAnalysisByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateMultiSpectraAnalysis 更新MultiSpectraAnalysis
// @Tags MultiSpectraAnalysis
// @Summary 更新MultiSpectraAnalysis
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body MultiSpectraAnalysisPkg.MultiSpectraAnalysis true "更新MultiSpectraAnalysis"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MtSpectraAly/updateMultiSpectraAnalysis [put]
func (MtSpectraAlyApi *MultiSpectraAnalysisApi) UpdateMultiSpectraAnalysis(c *gin.Context) {
	var MtSpectraAly MultiSpectraAnalysisPkg.MultiSpectraAnalysis
	err := c.ShouldBindJSON(&MtSpectraAly)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := MtSpectraAlyService.UpdateMultiSpectraAnalysis(MtSpectraAly); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindMultiSpectraAnalysis 用id查询MultiSpectraAnalysis
// @Tags MultiSpectraAnalysis
// @Summary 用id查询MultiSpectraAnalysis
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MultiSpectraAnalysisPkg.MultiSpectraAnalysis true "用id查询MultiSpectraAnalysis"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MtSpectraAly/findMultiSpectraAnalysis [get]
func (MtSpectraAlyApi *MultiSpectraAnalysisApi) FindMultiSpectraAnalysis(c *gin.Context) {
	var MtSpectraAly MultiSpectraAnalysisPkg.MultiSpectraAnalysis
	err := c.ShouldBindQuery(&MtSpectraAly)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reMtSpectraAly, err := MtSpectraAlyService.GetMultiSpectraAnalysis(MtSpectraAly.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reMtSpectraAly": reMtSpectraAly}, c)
	}
}

// GetMultiSpectraAnalysisList 分页获取MultiSpectraAnalysis列表
// @Tags MultiSpectraAnalysis
// @Summary 分页获取MultiSpectraAnalysis列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query MultiSpectraAnalysisPkgReq.MultiSpectraAnalysisSearch true "分页获取MultiSpectraAnalysis列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MtSpectraAly/getMultiSpectraAnalysisList [get]
func (MtSpectraAlyApi *MultiSpectraAnalysisApi) GetMultiSpectraAnalysisList(c *gin.Context) {
	var pageInfo MultiSpectraAnalysisPkgReq.MultiSpectraAnalysisSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := MtSpectraAlyService.GetMultiSpectraAnalysisInfoList(pageInfo, c); err != nil {
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

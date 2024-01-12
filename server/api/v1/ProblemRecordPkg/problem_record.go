package ProblemRecordPkg

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ProblemRecordPkg"
	ProblemRecordPkgReq "github.com/flipped-aurora/gin-vue-admin/server/model/ProblemRecordPkg/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProblemRecordApi struct {
}

var pbRecordService = service.ServiceGroupApp.ProblemRecordPkgServiceGroup.ProblemRecordService

// CreateProblemRecord 创建ProblemRecord
// @Tags ProblemRecord
// @Summary 创建ProblemRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ProblemRecordPkg.ProblemRecord true "创建ProblemRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pbRecord/createProblemRecord [post]
func (pbRecordApi *ProblemRecordApi) CreateProblemRecord(c *gin.Context) {
	var pbRecord ProblemRecordPkg.ProblemRecord
	err := c.ShouldBindJSON(&pbRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pbRecordService.CreateProblemRecord(&pbRecord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProblemRecord 删除ProblemRecord
// @Tags ProblemRecord
// @Summary 删除ProblemRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ProblemRecordPkg.ProblemRecord true "删除ProblemRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pbRecord/deleteProblemRecord [delete]
func (pbRecordApi *ProblemRecordApi) DeleteProblemRecord(c *gin.Context) {
	var pbRecord ProblemRecordPkg.ProblemRecord
	err := c.ShouldBindJSON(&pbRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pbRecordService.DeleteProblemRecord(pbRecord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteProblemRecordByIds 批量删除ProblemRecord
// @Tags ProblemRecord
// @Summary 批量删除ProblemRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProblemRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /pbRecord/deleteProblemRecordByIds [delete]
func (pbRecordApi *ProblemRecordApi) DeleteProblemRecordByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pbRecordService.DeleteProblemRecordByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProblemRecord 更新ProblemRecord
// @Tags ProblemRecord
// @Summary 更新ProblemRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ProblemRecordPkg.ProblemRecord true "更新ProblemRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pbRecord/updateProblemRecord [put]
func (pbRecordApi *ProblemRecordApi) UpdateProblemRecord(c *gin.Context) {
	var pbRecord ProblemRecordPkg.ProblemRecord
	err := c.ShouldBindJSON(&pbRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := pbRecordService.UpdateProblemRecord(pbRecord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProblemRecord 用id查询ProblemRecord
// @Tags ProblemRecord
// @Summary 用id查询ProblemRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ProblemRecordPkg.ProblemRecord true "用id查询ProblemRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pbRecord/findProblemRecord [get]
func (pbRecordApi *ProblemRecordApi) FindProblemRecord(c *gin.Context) {
	var pbRecord ProblemRecordPkg.ProblemRecord
	err := c.ShouldBindQuery(&pbRecord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repbRecord, err := pbRecordService.GetProblemRecord(pbRecord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repbRecord": repbRecord}, c)
	}
}

// GetProblemRecordList 分页获取ProblemRecord列表
// @Tags ProblemRecord
// @Summary 分页获取ProblemRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ProblemRecordPkgReq.ProblemRecordSearch true "分页获取ProblemRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pbRecord/getProblemRecordList [get]
func (pbRecordApi *ProblemRecordApi) GetProblemRecordList(c *gin.Context) {
	var pageInfo ProblemRecordPkgReq.ProblemRecordSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := pbRecordService.GetProblemRecordInfoList(pageInfo); err != nil {
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

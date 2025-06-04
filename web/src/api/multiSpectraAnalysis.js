import service from '@/utils/request'

// @Tags MultiSpectraAnalysis
// @Summary 创建MultiSpectraAnalysis
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiSpectraAnalysis true "创建MultiSpectraAnalysis"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MtSpectraAly/createMultiSpectraAnalysis [post]
export const createMultiSpectraAnalysis = (data) => {
  return service({
    url: '/MtSpectraAly/createMultiSpectraAnalysis',
    method: 'post',
    data
  })
}

// @Tags MultiSpectraAnalysis
// @Summary 删除MultiSpectraAnalysis
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiSpectraAnalysis true "删除MultiSpectraAnalysis"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MtSpectraAly/deleteMultiSpectraAnalysis [delete]
export const deleteMultiSpectraAnalysis = (data) => {
  return service({
    url: '/MtSpectraAly/deleteMultiSpectraAnalysis',
    method: 'delete',
    data
  })
}

// @Tags MultiSpectraAnalysis
// @Summary 删除MultiSpectraAnalysis
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MultiSpectraAnalysis"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MtSpectraAly/deleteMultiSpectraAnalysis [delete]
export const deleteMultiSpectraAnalysisByIds = (data) => {
  return service({
    url: '/MtSpectraAly/deleteMultiSpectraAnalysisByIds',
    method: 'delete',
    data
  })
}

// @Tags MultiSpectraAnalysis
// @Summary 更新MultiSpectraAnalysis
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiSpectraAnalysis true "更新MultiSpectraAnalysis"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MtSpectraAly/updateMultiSpectraAnalysis [put]
export const updateMultiSpectraAnalysis = (data) => {
  return service({
    url: '/MtSpectraAly/updateMultiSpectraAnalysis',
    method: 'put',
    data
  })
}

// @Tags MultiSpectraAnalysis
// @Summary 用id查询MultiSpectraAnalysis
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MultiSpectraAnalysis true "用id查询MultiSpectraAnalysis"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MtSpectraAly/findMultiSpectraAnalysis [get]
export const findMultiSpectraAnalysis = (params) => {
  return service({
    url: '/MtSpectraAly/findMultiSpectraAnalysis',
    method: 'get',
    params
  })
}

// @Tags MultiSpectraAnalysis
// @Summary 分页获取MultiSpectraAnalysis列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MultiSpectraAnalysis列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MtSpectraAly/getMultiSpectraAnalysisList [get]
export const getMultiSpectraAnalysisList = (params) => {
  return service({
    url: '/MtSpectraAly/getMultiSpectraAnalysisList',
    method: 'get',
    params
  })
}

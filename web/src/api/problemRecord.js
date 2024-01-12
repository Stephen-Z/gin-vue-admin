import service from '@/utils/request'

// @Tags ProblemRecord
// @Summary 创建ProblemRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProblemRecord true "创建ProblemRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pbRecord/createProblemRecord [post]
export const createProblemRecord = (data) => {
  return service({
    url: '/pbRecord/createProblemRecord',
    method: 'post',
    data
  })
}

// @Tags ProblemRecord
// @Summary 删除ProblemRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProblemRecord true "删除ProblemRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pbRecord/deleteProblemRecord [delete]
export const deleteProblemRecord = (data) => {
  return service({
    url: '/pbRecord/deleteProblemRecord',
    method: 'delete',
    data
  })
}

// @Tags ProblemRecord
// @Summary 删除ProblemRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ProblemRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /pbRecord/deleteProblemRecord [delete]
export const deleteProblemRecordByIds = (data) => {
  return service({
    url: '/pbRecord/deleteProblemRecordByIds',
    method: 'delete',
    data
  })
}

// @Tags ProblemRecord
// @Summary 更新ProblemRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ProblemRecord true "更新ProblemRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /pbRecord/updateProblemRecord [put]
export const updateProblemRecord = (data) => {
  return service({
    url: '/pbRecord/updateProblemRecord',
    method: 'put',
    data
  })
}

// @Tags ProblemRecord
// @Summary 用id查询ProblemRecord
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ProblemRecord true "用id查询ProblemRecord"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /pbRecord/findProblemRecord [get]
export const findProblemRecord = (params) => {
  return service({
    url: '/pbRecord/findProblemRecord',
    method: 'get',
    params
  })
}

// @Tags ProblemRecord
// @Summary 分页获取ProblemRecord列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ProblemRecord列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /pbRecord/getProblemRecordList [get]
export const getProblemRecordList = (params) => {
  return service({
    url: '/pbRecord/getProblemRecordList',
    method: 'get',
    params
  })
}

import service from '@/utils/request'

// @Tags MultiSpectraType
// @Summary 创建MultiSpectraType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiSpectraType true "创建MultiSpectraType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MtSpectraType/createMultiSpectraType [post]
export const createMultiSpectraType = (data) => {
  return service({
    url: '/MtSpectraType/createMultiSpectraType',
    method: 'post',
    data
  })
}

// @Tags MultiSpectraType
// @Summary 删除MultiSpectraType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiSpectraType true "删除MultiSpectraType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MtSpectraType/deleteMultiSpectraType [delete]
export const deleteMultiSpectraType = (data) => {
  return service({
    url: '/MtSpectraType/deleteMultiSpectraType',
    method: 'delete',
    data
  })
}

// @Tags MultiSpectraType
// @Summary 删除MultiSpectraType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MultiSpectraType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MtSpectraType/deleteMultiSpectraType [delete]
export const deleteMultiSpectraTypeByIds = (data) => {
  return service({
    url: '/MtSpectraType/deleteMultiSpectraTypeByIds',
    method: 'delete',
    data
  })
}

// @Tags MultiSpectraType
// @Summary 更新MultiSpectraType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiSpectraType true "更新MultiSpectraType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MtSpectraType/updateMultiSpectraType [put]
export const updateMultiSpectraType = (data) => {
  return service({
    url: '/MtSpectraType/updateMultiSpectraType',
    method: 'put',
    data
  })
}

// @Tags MultiSpectraType
// @Summary 用id查询MultiSpectraType
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MultiSpectraType true "用id查询MultiSpectraType"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MtSpectraType/findMultiSpectraType [get]
export const findMultiSpectraType = (params) => {
  return service({
    url: '/MtSpectraType/findMultiSpectraType',
    method: 'get',
    params
  })
}

// @Tags MultiSpectraType
// @Summary 分页获取MultiSpectraType列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MultiSpectraType列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MtSpectraType/getMultiSpectraTypeList [get]
export const getMultiSpectraTypeList = (params) => {
  return service({
    url: '/MtSpectraType/getMultiSpectraTypeList',
    method: 'get',
    params
  })
}

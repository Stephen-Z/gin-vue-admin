import service from '@/utils/request'

// @Tags MultiSpectraPlate
// @Summary 创建MultiSpectraPlate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiSpectraPlate true "创建MultiSpectraPlate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MtpectraPlate/createMultiSpectraPlate [post]
export const createMultiSpectraPlate = (data) => {
  return service({
    url: '/MtpectraPlate/createMultiSpectraPlate',
    method: 'post',
    data
  })
}

// @Tags MultiSpectraPlate
// @Summary 删除MultiSpectraPlate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiSpectraPlate true "删除MultiSpectraPlate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MtpectraPlate/deleteMultiSpectraPlate [delete]
export const deleteMultiSpectraPlate = (data) => {
  return service({
    url: '/MtpectraPlate/deleteMultiSpectraPlate',
    method: 'delete',
    data
  })
}

// @Tags MultiSpectraPlate
// @Summary 删除MultiSpectraPlate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除MultiSpectraPlate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /MtpectraPlate/deleteMultiSpectraPlate [delete]
export const deleteMultiSpectraPlateByIds = (data) => {
  return service({
    url: '/MtpectraPlate/deleteMultiSpectraPlateByIds',
    method: 'delete',
    data
  })
}

// @Tags MultiSpectraPlate
// @Summary 更新MultiSpectraPlate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.MultiSpectraPlate true "更新MultiSpectraPlate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /MtpectraPlate/updateMultiSpectraPlate [put]
export const updateMultiSpectraPlate = (data) => {
  return service({
    url: '/MtpectraPlate/updateMultiSpectraPlate',
    method: 'put',
    data
  })
}

// @Tags MultiSpectraPlate
// @Summary 用id查询MultiSpectraPlate
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.MultiSpectraPlate true "用id查询MultiSpectraPlate"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /MtpectraPlate/findMultiSpectraPlate [get]
export const findMultiSpectraPlate = (params) => {
  return service({
    url: '/MtpectraPlate/findMultiSpectraPlate',
    method: 'get',
    params
  })
}

// @Tags MultiSpectraPlate
// @Summary 分页获取MultiSpectraPlate列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取MultiSpectraPlate列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /MtpectraPlate/getMultiSpectraPlateList [get]
export const getMultiSpectraPlateList = (params) => {
  return service({
    url: '/MtpectraPlate/getMultiSpectraPlateList',
    method: 'get',
    params
  })
}

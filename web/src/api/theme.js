import service from '@/utils/request'

// @Tags Theme
// @Summary 创建Theme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Theme true "创建Theme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /theme/createTheme [post]
export const createTheme = (data) => {
  return service({
    url: '/theme/createTheme',
    method: 'post',
    data
  })
}

// @Tags Theme
// @Summary 删除Theme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Theme true "删除Theme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /theme/deleteTheme [delete]
export const deleteTheme = (data) => {
  return service({
    url: '/theme/deleteTheme',
    method: 'delete',
    data
  })
}

// @Tags Theme
// @Summary 删除Theme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Theme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /theme/deleteTheme [delete]
export const deleteThemeByIds = (data) => {
  return service({
    url: '/theme/deleteThemeByIds',
    method: 'delete',
    data
  })
}

// @Tags Theme
// @Summary 更新Theme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Theme true "更新Theme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /theme/updateTheme [put]
export const updateTheme = (data) => {
  return service({
    url: '/theme/updateTheme',
    method: 'put',
    data
  })
}

// @Tags Theme
// @Summary 用id查询Theme
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Theme true "用id查询Theme"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /theme/findTheme [get]
export const findTheme = (params) => {
  return service({
    url: '/theme/findTheme',
    method: 'get',
    params
  })
}

// @Tags Theme
// @Summary 分页获取Theme列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Theme列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /theme/getThemeList [get]
export const getThemeList = (params) => {
  return service({
    url: '/theme/getThemeList',
    method: 'get',
    params
  })
}

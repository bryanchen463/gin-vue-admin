import service from '@/utils/request'

// @Tags Speed
// @Summary 创建耗时
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Speed true "创建耗时"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /metric/createSpeed [post]
export const createSpeed = (data) => {
  return service({
    url: '/metric/createSpeed',
    method: 'post',
    data
  })
}

// @Tags Speed
// @Summary 删除耗时
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Speed true "删除耗时"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /metric/deleteSpeed [delete]
export const deleteSpeed = (params) => {
  return service({
    url: '/metric/deleteSpeed',
    method: 'delete',
    params
  })
}

// @Tags Speed
// @Summary 批量删除耗时
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除耗时"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /metric/deleteSpeed [delete]
export const deleteSpeedByIds = (params) => {
  return service({
    url: '/metric/deleteSpeedByIds',
    method: 'delete',
    params
  })
}

// @Tags Speed
// @Summary 更新耗时
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Speed true "更新耗时"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /metric/updateSpeed [put]
export const updateSpeed = (data) => {
  return service({
    url: '/metric/updateSpeed',
    method: 'put',
    data
  })
}

// @Tags Speed
// @Summary 用id查询耗时
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Speed true "用id查询耗时"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /metric/findSpeed [get]
export const findSpeed = (params) => {
  return service({
    url: '/metric/findSpeed',
    method: 'get',
    params
  })
}

// @Tags Speed
// @Summary 分页获取耗时列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取耗时列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /metric/getSpeedList [get]
export const getSpeedList = (params) => {
  return service({
    url: '/metric/getSpeedList',
    method: 'get',
    params
  })
}

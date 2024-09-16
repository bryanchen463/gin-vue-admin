package metrics

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/metrics"
    metricsReq "github.com/flipped-aurora/gin-vue-admin/server/model/metrics/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type SpeedApi struct {
}

var metricService = service.ServiceGroupApp.MetricsServiceGroup.SpeedService


// CreateSpeed 创建耗时
// @Tags Speed
// @Summary 创建耗时
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body metrics.Speed true "创建耗时"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /metric/createSpeed [post]
func (metricApi *SpeedApi) CreateSpeed(c *gin.Context) {
	var metric metrics.Speed
	err := c.ShouldBindJSON(&metric)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := metricService.CreateSpeed(&metric); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSpeed 删除耗时
// @Tags Speed
// @Summary 删除耗时
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body metrics.Speed true "删除耗时"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /metric/deleteSpeed [delete]
func (metricApi *SpeedApi) DeleteSpeed(c *gin.Context) {
	ID := c.Query("ID")
	if err := metricService.DeleteSpeed(ID); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSpeedByIds 批量删除耗时
// @Tags Speed
// @Summary 批量删除耗时
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /metric/deleteSpeedByIds [delete]
func (metricApi *SpeedApi) DeleteSpeedByIds(c *gin.Context) {
	IDs := c.QueryArray("IDs[]")
	if err := metricService.DeleteSpeedByIds(IDs); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSpeed 更新耗时
// @Tags Speed
// @Summary 更新耗时
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body metrics.Speed true "更新耗时"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /metric/updateSpeed [put]
func (metricApi *SpeedApi) UpdateSpeed(c *gin.Context) {
	var metric metrics.Speed
	err := c.ShouldBindJSON(&metric)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := metricService.UpdateSpeed(metric); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSpeed 用id查询耗时
// @Tags Speed
// @Summary 用id查询耗时
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query metrics.Speed true "用id查询耗时"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /metric/findSpeed [get]
func (metricApi *SpeedApi) FindSpeed(c *gin.Context) {
	ID := c.Query("ID")
	if remetric, err := metricService.GetSpeed(ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"remetric": remetric}, c)
	}
}

// GetSpeedList 分页获取耗时列表
// @Tags Speed
// @Summary 分页获取耗时列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query metricsReq.SpeedSearch true "分页获取耗时列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /metric/getSpeedList [get]
func (metricApi *SpeedApi) GetSpeedList(c *gin.Context) {
	var pageInfo metricsReq.SpeedSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := metricService.GetSpeedInfoList(pageInfo); err != nil {
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

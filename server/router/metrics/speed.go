package metrics

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type SpeedRouter struct {
}

// InitSpeedRouter 初始化 耗时 路由信息
func (s *SpeedRouter) InitSpeedRouter(Router *gin.RouterGroup) {
	metricRouter := Router.Group("metric")
	metricRouterWithoutRecord := Router.Group("metric")
	var metricApi = v1.ApiGroupApp.MetricsApiGroup.SpeedApi
	{
		metricRouter.POST("createSpeed", metricApi.CreateSpeed)             // 新建耗时
		metricRouter.DELETE("deleteSpeed", metricApi.DeleteSpeed)           // 删除耗时
		metricRouter.DELETE("deleteSpeedByIds", metricApi.DeleteSpeedByIds) // 批量删除耗时
		metricRouter.PUT("updateSpeed", metricApi.UpdateSpeed)              // 更新耗时
	}
	{
		metricRouterWithoutRecord.GET("findSpeed", metricApi.FindSpeed)       // 根据ID获取耗时
		metricRouterWithoutRecord.GET("getSpeedList", metricApi.GetSpeedList) // 获取耗时列表
	}
}

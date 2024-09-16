package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/factor_web"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/metrics"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup     system.ApiGroup
	ExampleApiGroup    example.ApiGroup
	Factor_webApiGroup factor_web.ApiGroup
	MetricsApiGroup    metrics.ApiGroup
}

var ApiGroupApp = new(ApiGroup)

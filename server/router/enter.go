package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/factor_web"
	"github.com/flipped-aurora/gin-vue-admin/server/router/metrics"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System     system.RouterGroup
	Example    example.RouterGroup
	Factor_web factor_web.RouterGroup
	Metrics    metrics.RouterGroup
}

var RouterGroupApp = new(RouterGroup)

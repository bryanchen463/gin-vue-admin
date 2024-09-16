package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
	
)

type SpeedSearch struct{
    
        StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
        EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    
                      Metric_name  string `json:"metric_name" form:"metric_name" `
                      Machine_ip  string `json:"machine_ip" form:"machine_ip" `
                      Sample_time  *time.Time `json:"sample_time" form:"sample_time" `
                      Exchange  string `json:"exchange" form:"exchange" `
                      Symbol  string `json:"symbol" form:"symbol" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}

// 自动生成模板Speed
package metrics

import (
	"time"
)

// 耗时 结构体  Speed
type Speed struct {
	Id          string     `gorm:"primarykey" json:"ID" bson:"_id"`
	Metric_name string     `json:"metric_name" form:"metric_name" bson:"metric_name"` //指标名
	Machine_ip  string     `json:"machine_ip" form:"machine_ip" bson:"machine_ip"`    //机器ip
	Sample_time *time.Time `json:"sample_time" form:"sample_time" bson:"sample_time"` //创建时间
	Count       *int       `json:"count" form:"count" bson:"count"`                   //次数
	Exchange    string     `json:"exchange" form:"exchange" bson:"exchange"`          //交易所
	Mean        *float64   `json:"mean" form:"mean" bson:"mean"`                      //耗时(平均)
	Mid         *float64   `json:"mid" form:"mid" bson:"mid"`                         //耗时(平均)
	Std         *float64   `json:"std" form:"std" bson:"std"`                         //标准差
	Symbol      string     `json:"symbol" form:"symbol" bson:"symbol"`                //Symbol
}

// TableName 耗时 Speed自定义表名 speed
func (Speed) TableName() string {
	return "speed"
}

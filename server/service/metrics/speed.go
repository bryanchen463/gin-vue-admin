package metrics

import (
	"context"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/metrics"
	metricsReq "github.com/flipped-aurora/gin-vue-admin/server/model/metrics/request"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type SpeedService struct {
}

// CreateSpeed 创建耗时记录
// Author [piexlmax](https://github.com/piexlmax)
func (metricService *SpeedService) CreateSpeed(metric *metrics.Speed) (err error) {
	result, err := global.GVA_MONGO.InsertOne(context.Background(), metric)
	global.GVA_LOG.Log(zapcore.DebugLevel, "CreateSpeed", zap.Any("record", result))
	return err
}

// DeleteSpeed 删除耗时记录
// Author [piexlmax](https://github.com/piexlmax)
func (metricService *SpeedService) DeleteSpeed(ID string) (err error) {
	err = global.GVA_MONGO.RemoveId(context.Background(), ID)
	return err
}

// DeleteSpeedByIds 批量删除耗时记录
// Author [piexlmax](https://github.com/piexlmax)
func (metricService *SpeedService) DeleteSpeedByIds(IDs []string) (err error) {
	var ids []primitive.ObjectID
	for _, v := range IDs {
		id, err := primitive.ObjectIDFromHex(v)
		if err != nil {
			return err
		}
		ids = append(ids, id)
	}
	err = global.GVA_MONGO.Remove(context.Background(), bson.M{"_id": bson.M{"$in": ids}})
	return err
}

// UpdateSpeed 更新耗时记录
// Author [piexlmax](https://github.com/piexlmax)
func (metricService *SpeedService) UpdateSpeed(metric metrics.Speed) (err error) {

	return global.GVA_MONGO.UpdateId(context.Background(), metric.Id, bson.M{"$set": metric})
}

// GetSpeed 根据ID获取耗时记录
// Author [piexlmax](https://github.com/piexlmax)
func (metricService *SpeedService) GetSpeed(ID string) (metric metrics.Speed, err error) {
	id, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		global.GVA_LOG.Log(zap.ErrorLevel, "getSpeed", zap.Error(err))
		return
	}
	err = global.GVA_MONGO.Find(context.Background(), bson.M{"_id": id}).One(&metric)
	return
}

// GetSpeedInfoList 分页获取耗时记录
// Author [piexlmax](https://github.com/piexlmax)
func (metricService *SpeedService) GetSpeedInfoList(info metricsReq.SpeedSearch) (list []metrics.Speed, total int64, err error) {
	ctx := context.Background()
	// limit := info.PageSize
	// offset := info.PageSize * (info.Page - 1)
	// 创建db
	querySon := bson.D{}

	// 时间范围查询
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		timeQuery := bson.D{
			{Key: "$gte", Value: info.StartCreatedAt},
			{Key: "$lt", Value: info.EndCreatedAt},
		}
		querySon = append(querySon, bson.E{Key: "sample_time", Value: timeQuery})
	}

	// 其他条件查询
	if info.Metric_name != "" {
		querySon = append(querySon, bson.E{Key: "metric_name", Value: info.Metric_name})
	}
	if info.Machine_ip != "" {
		querySon = append(querySon, bson.E{Key: "machine_ip", Value: info.Machine_ip})
	}
	if info.Exchange != "" {
		querySon = append(querySon, bson.E{Key: "exchange", Value: info.Exchange})
	}
	if info.Symbol != "" {
		querySon = append(querySon, bson.E{Key: "symbol", Value: info.Symbol})
	}

	aggrateI := global.GVA_MONGO.Aggregate(ctx, qmgo.Pipeline{bson.D{{Key: "$match", Value: querySon}}})
	err = aggrateI.All(&list)
	if err != nil {
		return
	}
	// }
	// orderMap := make(map[string]bool)
	// orderMap["sample_time"] = true
	// orderMap["mean"] = true
	// orderMap["mid"] = true
	// orderMap["std"] = true
	// orderMap["symbol"] = true
	// if orderMap[info.Sort] {
	// 	query = query.Sort("sample_time")
	// }

	// if limit != 0 {
	// 	query = query.Limit(int64(limit)).Skip(int64(offset))
	// }

	// err = query.All(&metrics)
	total = int64(len(list))
	return list, total, err
}

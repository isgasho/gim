package cache

import (
	"gim/logic/model"
	"gim/public/logger"
	"strconv"
	"time"

	"github.com/go-redis/redis"
)

const (
	AppKey    = "app:"
	AppExpire = 24 * time.Hour
)

type appCache struct{}

var AppCache = new(appCache)

// Get 获取设备缓存
func (c *appCache) Get(appId int64) (*model.App, error) {
	var app model.App
	err := get(AppKey+strconv.FormatInt(appId, 10), &app)
	if err != nil && err != redis.Nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	if err == redis.Nil {
		return nil, nil
	}
	return &app, nil
}

// Set 设置app缓存
func (c *appCache) Set(app *model.App) error {
	err := set(AppKey+strconv.FormatInt(app.Id, 10), app, AppExpire)
	if err != nil {
		logger.Sugar.Error(err)
		return err
	}
	return nil
}

package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/linzijie1998/bytedance_camp_douyin/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	DOUYIN_VIPER  *viper.Viper
	DOUYIN_CONFIG config.Config
	DOUYIN_DB     *gorm.DB
	DOUYIN_REDIS  *redis.Client
	DOUYIN_LOGGER *zap.Logger
)

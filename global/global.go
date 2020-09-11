package global

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"go-demo/config"
)

var (
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	GVA_DB     *gorm.DB
	GVA_REDIS  *redis.Client
)

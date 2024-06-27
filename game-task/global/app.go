package global

import (
	"my-gin/config"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConifgViper   *viper.Viper
	Config        config.Configuration
	Log           *zap.Logger
	DB            *gorm.DB
	Redis         *redis.Client
	ElasticSearch *elasticsearch.Client
}

var App = new(Application)

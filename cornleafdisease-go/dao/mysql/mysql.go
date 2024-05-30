package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project/cornleafdisease/setting"
)

var DB *gorm.DB
var err error

func init() {
	setting.Init()
	cfg := setting.Conf.MySQLConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DB,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		QueryFields: true,
		//SkipDefaultTransaction: true,
	})
	if err != nil {
		zap.L().Error("connect DB failed, err:%v\n", zap.Error(err))
		return
	}
}

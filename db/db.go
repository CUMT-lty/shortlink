package db

import (
	"github.com/lty/shortlink/config"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func init() {
	initDB()
}

func initDB() {
	var err error
	dsn := config.Conf.Common.CommonMysql.DSN
	var newLogger = logger.New( //
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      logger.Info,
			Colorful:      true,
		})
	// TODO: gorm 内部使用的就是数据库连接池，对于一个 gorm 实例，内部其实自己维护了一个链接池
	// TODO: 因此用户不需要关注对数据库连接的管理
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		logrus.Error("connect db fail:%s", err.Error())
	}
}

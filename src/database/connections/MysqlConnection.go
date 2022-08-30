package connections

import (
	"fmt"
	"log"

	"github.com/yeisonLucio/shopping-cart/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func CreateConnection() {
	var err error
	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.App.MysqlUser,
		config.App.MysqlPassword,
		config.App.MysqlHost,
		config.App.MysqlPort,
		config.App.MysqlDbName,
	)

	gormConfig := &gorm.Config{}

	if !config.App.EnableGormLog {
		gormConfig.Logger = logger.Default.LogMode(logger.Silent)
	}

	DB, err = gorm.Open(mysql.Open(dsn), gormConfig)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Database connected")
	}

}

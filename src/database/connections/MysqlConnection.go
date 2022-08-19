package connections

import (
	"fmt"
	"log"

	"github.com/yeisonLucio/shopping-cart/src/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Print("Database connected")
	}

}

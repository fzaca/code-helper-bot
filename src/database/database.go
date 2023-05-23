package database

import (
	"fmt"

	"github.com/Xukay101/code-helper-bot/src/config"
	"github.com/Xukay101/code-helper-bot/src/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True",
		config.GetConfig().Database.User,
		config.GetConfig().Database.Password,
		config.GetConfig().Database.Host,
		config.GetConfig().Database.Port,
		config.GetConfig().Database.Name,
	)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	utils.FatalOnError("Error getting database", err)

	return conn
}

package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitConfig() {
	// load the config file
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read config failed, err:", err)
	}
	fmt.Println("config app: ", viper.Get("app"))
	fmt.Println("config mysql: ", viper.Get("mysql"))
}

func InitMySQL() {
	// create custom sql log, print sql query
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)
	// connect to mysql
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")),
		&gorm.Config{Logger: newLogger})
	// print sql is initiating
	fmt.Println("Initiating MySQL...")

}

package config

import (
	"log"
	"strings"
	"unisun/api/promotion-price-mapping-processor-schedule/src/constants"
	"unisun/api/promotion-price-mapping-processor-schedule/src/entitys"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	str := []string{
		"host=" + viper.GetString(constants.DB_HOST),
		"user=" + viper.GetString(constants.DB_USER),
		"password=" + viper.GetString(constants.DB_PASS),
		"dbname=" + viper.GetString(constants.DB_NAME),
		"port=" + viper.GetString(constants.DB_PORT),
		"TimeZone=" + viper.GetString(constants.DB_TIMEZONE),
	}
	dsn := strings.Join(str, " ")
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln("Failed to connect to database!")
	}

	var ClassRoomPriceEntity entitys.ClassRoomPrice
	database.AutoMigrate(&ClassRoomPriceEntity)
	DB = database
}

package main

import (
	"log"
	"unisun/api/class-room-price-mapping-processor-schedule/src"
	config "unisun/api/class-room-price-mapping-processor-schedule/src/configs"
	"unisun/api/class-room-price-mapping-processor-schedule/src/constants"

	"github.com/spf13/viper"
)

func main() {
	envService := config.New(constants.ENV_FILE, constants.ENV_PATH)
	if err := envService.ConfigENV(); err != nil {
		log.Panic(err)
	}

	r := src.App()
	port := viper.GetString(constants.ENV_PORT)
	if port == "" {
		r.Run(":8080")
	} else {
		r.Run(":" + port)
	}
}

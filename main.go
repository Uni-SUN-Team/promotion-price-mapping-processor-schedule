package main

import (
	"log"
	"unisun/api/promotion-price-mapping-processor-schedule/src"
	config "unisun/api/promotion-price-mapping-processor-schedule/src/configs"
	"unisun/api/promotion-price-mapping-processor-schedule/src/constants"
)

func main() {
	envService := config.New(constants.ENV_FILE, constants.ENV_PATH)
	if err := envService.ConfigENV(); err != nil {
		log.Panic(err)
	}

	src.App()
}

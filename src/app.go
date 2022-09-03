package src

import (
	config "unisun/api/promotion-price-mapping-processor-schedule/src/configs"
	"unisun/api/promotion-price-mapping-processor-schedule/src/crons"

	"github.com/spf13/viper"
)

func App() {
	config.ConnectDatabase()
	crons.NewPromotionPriceCronAdapter(viper.GetInt("cron.every"), viper.GetString("cron.at")).ProcessSchedule()
}

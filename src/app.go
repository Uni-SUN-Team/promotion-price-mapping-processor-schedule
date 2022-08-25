package src

import (
	"unisun/api/class-room-price-mapping-processor-schedule/src/crons"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func App() *gin.Engine {
	crons.NewClassRoomPriceCronAdapter(viper.GetInt("cron.every"), viper.GetString("cron.at"))
	crons.NewPromotionPriceCronAdapter(viper.GetInt("cron.every"), viper.GetString("cron.at"))
	r := gin.Default()
	return r
}

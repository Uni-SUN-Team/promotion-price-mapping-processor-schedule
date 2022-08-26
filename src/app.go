package src

import (
	"unisun/api/class-room-price-mapping-processor-schedule/src/controllers"
	"unisun/api/class-room-price-mapping-processor-schedule/src/crons"
	"unisun/api/class-room-price-mapping-processor-schedule/src/repositories"
	"unisun/api/class-room-price-mapping-processor-schedule/src/routes"
	"unisun/api/class-room-price-mapping-processor-schedule/src/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func App() *gin.Engine {
	crons.NewClassRoomPriceCronAdapter(viper.GetInt("cron.every"), viper.GetString("cron.at"))
	crons.NewPromotionPriceCronAdapter(viper.GetInt("cron.every"), viper.GetString("cron.at"))
	r := gin.Default()
	g := r.Group(viper.GetString("app.context_path") + viper.GetString("app.root_path") + "/v1")
	{
		handleClassRoomPrice().ClassRoomPriceRoute(g)
	}
	return r
}

func handleClassRoomPrice() *routes.ClassRoomPriceRouteAdapter {
	repo := repositories.NewClassRoomPriceRepositoriesAdapter()
	service := services.NewClassRoomPriceGetFromRepoAdapter(repo)
	controller := controllers.NewClassRoomPriceControllerAdapter(service)
	router := routes.NewClassRoomPriceRouteAdapter(controller)
	return router
}

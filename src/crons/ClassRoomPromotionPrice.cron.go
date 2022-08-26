package crons

import (
	"strings"
	"time"
	"unisun/api/class-room-price-mapping-processor-schedule/src/components"
	"unisun/api/class-room-price-mapping-processor-schedule/src/components/client"
	"unisun/api/class-room-price-mapping-processor-schedule/src/constants"
	"unisun/api/class-room-price-mapping-processor-schedule/src/repositories"
	"unisun/api/class-room-price-mapping-processor-schedule/src/services"
	"unisun/api/class-room-price-mapping-processor-schedule/src/utils"

	"github.com/go-co-op/gocron"
	"github.com/spf13/viper"
)

type PromotionPriceCronAdapter struct {
	Every int
	At    string
}

func NewPromotionPriceCronAdapter(every int, at string) *PromotionPriceCronAdapter {
	return &PromotionPriceCronAdapter{
		Every: every,
		At:    at,
	}
}

func (srv *PromotionPriceCronAdapter) ProcessSchedule() {
	s := gocron.NewScheduler(time.UTC)
	s.Every("5m").Do(func() {
		httpRequestPort := utils.New()
		classRoomHttpRequestAdapter := client.NewClassRoomHttpRequestAdapter(httpRequestPort)
		path := strings.Join([]string{viper.GetString("endpoint.promotion.path"), viper.GetString("endpoint.promotion.query")}, "")
		method := constants.GET
		mappingValueRequestPayloadAdapte := components.NewMappingValueRequestPayloadAdapte(path, method, nil)
		classRoomPriceRepositoriesAdapter := repositories.NewClassRoomPriceRepositoriesAdapter()
		classRoomPriceAdapter := services.NewPromotionPriceAdapter(classRoomHttpRequestAdapter, mappingValueRequestPayloadAdapte, classRoomPriceRepositoriesAdapter)
		classRoomPriceAdapter.ManagePromotion()
	})
	s.StartAsync()
	s.StartBlocking()
}

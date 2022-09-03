package services

import (
	"encoding/json"
	"log"
	"strconv"
	"unisun/api/promotion-price-mapping-processor-schedule/src/models/promotion"
	"unisun/api/promotion-price-mapping-processor-schedule/src/ports/component"
	"unisun/api/promotion-price-mapping-processor-schedule/src/ports/repository"

	"github.com/dariubs/percent"
)

type PromotionPriceAdapter struct {
	ClassRoomHttpRequest       component.ClassRoomHttpRequestPort
	MappingValueRequestPayload component.MappingValueRequestPayloadPort
	ClassRoomPriceGorm         repository.ClassRoomPricePort
}

func NewPromotionPriceAdapter(classRoomHttpRequest component.ClassRoomHttpRequestPort,
	mappingValueRequestPayload component.MappingValueRequestPayloadPort,
	ClassRoomPrice repository.ClassRoomPricePort) *PromotionPriceAdapter {
	return &PromotionPriceAdapter{
		ClassRoomHttpRequest:       classRoomHttpRequest,
		MappingValueRequestPayload: mappingValueRequestPayload,
		ClassRoomPriceGorm:         ClassRoomPrice,
	}
}

func (srv *PromotionPriceAdapter) ManagePromotion() {
	promotions := promotion.Promotions{}
	payload := srv.MappingValueRequestPayload.MappingPayload()
	data, err := srv.ClassRoomHttpRequest.GetInformationFormStrapi(*payload)
	if err != nil {
		log.Panic("Have problem! Call get value from stapi is error.")
	}
	if err := json.Unmarshal([]byte(data), &promotions); err != nil {
		log.Panic("Have problem! Decode json is error.")
	}
	for _, promotion := range promotions.Data {
		switch promotion.Group[0].Component {
		case "classroom.subject-list":
			classrooms := srv.ClassRoomPriceGorm.GetByClassRoomId(promotion.Group[0].Classrooms[0].Id)
			if classrooms.ClassRoomId != 0 {
				payload := *classrooms
				payload.SpecialPrice = calculatePrice(classrooms.RegularPrice, promotion.Discount[0].Discount, promotion.Discount[0].Component)
				srv.ClassRoomPriceGorm.Update(payload)
			}
		case "classroom.advisors-list":
			classrooms := srv.ClassRoomPriceGorm.GetByAdvisor(strconv.Itoa(promotion.Group[0].Advisors[0].Id))
			for _, adv := range *classrooms {
				adv.SpecialPrice = calculatePrice(adv.RegularPrice, promotion.Discount[0].Discount, promotion.Discount[0].Component)
				srv.ClassRoomPriceGorm.Update(adv)
			}
		case "classroom.categories-list":
			classrooms := srv.ClassRoomPriceGorm.GetByCategories(strconv.Itoa(promotion.Group[0].Categories[0].Id))
			for _, cat := range *classrooms {
				cat.SpecialPrice = calculatePrice(cat.RegularPrice, promotion.Discount[0].Discount, promotion.Discount[0].Component)
				srv.ClassRoomPriceGorm.Update(cat)
			}
		}
	}
}

func calculatePrice(regularPrice float64, discount float64, typeDiscount string) float64 {
	switch typeDiscount {
	case "promotion.reduce-by-difference":
		return regularPrice - discount
	case "promotion.reduce-by-percentage":
		return percent.PercentFloat(discount, regularPrice)
	case "promotion.reduce-equal-to-the-number":
		return discount
	}
	return 0.0
}

package services

import (
	"encoding/json"
	"log"
	"strconv"
	"strings"
	"unisun/api/class-room-price-mapping-processor-schedule/src/entitys"
	classroom "unisun/api/class-room-price-mapping-processor-schedule/src/models/class-room"
	classroomprice "unisun/api/class-room-price-mapping-processor-schedule/src/models/class-room-price"
	"unisun/api/class-room-price-mapping-processor-schedule/src/ports/component"
	"unisun/api/class-room-price-mapping-processor-schedule/src/ports/repository"
)

type ClassRoomPriceAdapter struct {
	ClassRoomHttpRequest       component.ClassRoomHttpRequestPort
	MappingValueRequestPayload component.MappingValueRequestPayloadPort
	ClassRoomPriceGorm         repository.ClassRoomPricePort
}

type ClassRoomPriceGetFromRepoAdapter struct {
	ClassRoomPriceGorm repository.ClassRoomPricePort
}

func NewClassRoomPriceAdapter(classRoomHttpRequest component.ClassRoomHttpRequestPort,
	mappingValueRequestPayload component.MappingValueRequestPayloadPort,
	ClassRoomPrice repository.ClassRoomPricePort) *ClassRoomPriceAdapter {
	return &ClassRoomPriceAdapter{
		ClassRoomHttpRequest:       classRoomHttpRequest,
		MappingValueRequestPayload: mappingValueRequestPayload,
		ClassRoomPriceGorm:         ClassRoomPrice,
	}
}

func NewClassRoomPriceGetFromRepoAdapter(classRoomPriceRepo repository.ClassRoomPricePort) *ClassRoomPriceGetFromRepoAdapter {
	return &ClassRoomPriceGetFromRepoAdapter{
		ClassRoomPriceGorm: classRoomPriceRepo,
	}
}

func (srv *ClassRoomPriceAdapter) ManagePrice() {
	classRooms := classroom.ClassRooms{}
	payload := srv.MappingValueRequestPayload.MappingPayload()
	data, err := srv.ClassRoomHttpRequest.GetInformationFormStrapi(*payload)
	if err != nil {
		log.Panic("Have problem! Call get value from stapi is error.")
	}
	if err := json.Unmarshal([]byte(data), &classRooms); err != nil {
		log.Panic("Have problem! Decode json is error.")
	}
	for _, classRoom := range classRooms.Data {
		result := srv.ClassRoomPriceGorm.GetByClassRoomId(classRoom.Id)
		if result.ClassRoomId != 0 {
			if result.RegularPrice != classRoom.Price {
				body := *result
				body.RegularPrice = classRoom.Price
				if len(classRoom.Advisors) > 1 {
					var sum []string
					for _, id := range classRoom.Advisors {
						sum = append(sum, strconv.Itoa(id.Id))
					}
					body.Advisors = strings.Join(sum, ",")
				} else {
					body.Advisors = strconv.Itoa(classRoom.Advisors[0].Id)
				}

				if len(classRoom.Categories) > 1 {
					var sum []string
					for _, id := range classRoom.Categories {
						sum = append(sum, strconv.Itoa(id.Id))
					}
					body.Categories = strings.Join(sum, ",")
				} else {
					body.Categories = strconv.Itoa(classRoom.Categories[0].Id)
				}
				srv.ClassRoomPriceGorm.Update(body)
			}

		} else {
			body := entitys.ClassRoomPriceEntity{}
			body.ClassRoomId = classRoom.Id
			body.RegularPrice = classRoom.Price
			if len(classRoom.Advisors) > 1 {
				var sum []string
				for _, id := range classRoom.Advisors {
					sum = append(sum, strconv.Itoa(id.Id))
				}
				body.Advisors = strings.Join(sum, ",")
			} else {
				body.Advisors = strconv.Itoa(classRoom.Advisors[0].Id)
			}

			if len(classRoom.Categories) > 1 {
				var sum []string
				for _, id := range classRoom.Categories {
					sum = append(sum, strconv.Itoa(id.Id))
				}
				body.Categories = strings.Join(sum, ",")
			} else {
				body.Categories = strconv.Itoa(classRoom.Categories[0].Id)
			}
			srv.ClassRoomPriceGorm.Save(body)
		}
	}
}

func (srv *ClassRoomPriceGetFromRepoAdapter) GetClassRoomPrice(id int) *classroomprice.ClassRoomPriceEntity {
	result := srv.ClassRoomPriceGorm.GetById(id)
	return &classroomprice.ClassRoomPriceEntity{
		Id:           result.Id,
		ClassRoomId:  result.ClassRoomId,
		RegularPrice: result.RegularPrice,
		SpecialPrice: result.SpecialPrice,
		Advisors:     result.Advisors,
		Categories:   result.Categories,
	}
}

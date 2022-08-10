package services

import (
	"encoding/json"
	"log"
	"unisun/api/class-room-price-mapping-processor-schedule/src/entitys"
	classroom "unisun/api/class-room-price-mapping-processor-schedule/src/models/class-room"
	"unisun/api/class-room-price-mapping-processor-schedule/src/ports/component"
	"unisun/api/class-room-price-mapping-processor-schedule/src/ports/repository"
)

type ClassRoomPriceAdapter struct {
	ClassRoomHttpRequest       component.ClassRoomHttpRequestPort
	MappingValueRequestPayload component.MappingValueRequestPayloadPort
	ClassRoomPriceGorm         repository.ClassRoomPricePort
}

func New(classRoomHttpRequest component.ClassRoomHttpRequestPort,
	mappingValueRequestPayload component.MappingValueRequestPayloadPort,
	ClassRoomPrice repository.ClassRoomPricePort) *ClassRoomPriceAdapter {
	return &ClassRoomPriceAdapter{
		ClassRoomHttpRequest:       classRoomHttpRequest,
		MappingValueRequestPayload: mappingValueRequestPayload,
		ClassRoomPriceGorm:         ClassRoomPrice,
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
				srv.ClassRoomPriceGorm.Update(body)
			}
			
		} else {
			body := entitys.ClassRoomPriceEntity{}
			body.ClassRoomId = classRoom.Id
			body.RegularPrice = classRoom.Price
			srv.ClassRoomPriceGorm.Save(body)
		}
	}
}

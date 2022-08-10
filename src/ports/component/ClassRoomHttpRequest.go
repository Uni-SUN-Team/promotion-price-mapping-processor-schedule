package component

import "unisun/api/class-room-price-mapping-processor-schedule/src/models"

type ClassRoomHttpRequestPort interface {
	GetInformationFormStrapi(payloadRequest models.ServiceIncomeRequest) (string, error)
}

package component

import "unisun/api/promotion-price-mapping-processor-schedule/src/models"

type ClassRoomHttpRequestPort interface {
	GetInformationFormStrapi(payloadRequest models.ServiceIncomeRequest) (string, error)
}

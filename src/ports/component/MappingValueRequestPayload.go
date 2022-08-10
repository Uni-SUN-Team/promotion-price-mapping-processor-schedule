package component

import "unisun/api/class-room-price-mapping-processor-schedule/src/models"

type MappingValueRequestPayloadPort interface {
	MappingPayload() *models.ServiceIncomeRequest
}

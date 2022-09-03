package component

import "unisun/api/promotion-price-mapping-processor-schedule/src/models"

type MappingValueRequestPayloadPort interface {
	MappingPayload() *models.ServiceIncomeRequest
}

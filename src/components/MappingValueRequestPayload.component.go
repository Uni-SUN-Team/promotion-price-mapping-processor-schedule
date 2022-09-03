package components

import "unisun/api/promotion-price-mapping-processor-schedule/src/models"

type MappingValueRequestPayloadAdapte struct {
	Payload models.ServiceIncomeRequest
}

func NewMappingValueRequestPayloadAdapte(path string, method string, body []byte) *MappingValueRequestPayloadAdapte {
	return &MappingValueRequestPayloadAdapte{
		Payload: models.ServiceIncomeRequest{
			Path:   path,
			Method: method,
			Body:   body,
		},
	}
}

func (srv *MappingValueRequestPayloadAdapte) MappingPayload() *models.ServiceIncomeRequest {
	payload := models.ServiceIncomeRequest{}
	payload.Path = srv.Payload.Path
	payload.Method = srv.Payload.Method
	payload.Body = srv.Payload.Body
	return &payload
}

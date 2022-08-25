package client

import (
	"encoding/json"
	"io/ioutil"
	"unisun/api/class-room-price-mapping-processor-schedule/src/constants"
	"unisun/api/class-room-price-mapping-processor-schedule/src/models"
	"unisun/api/class-room-price-mapping-processor-schedule/src/ports/util"

	"github.com/spf13/viper"
)

type ClassRoomHttpRequestAdapter struct {
	HttpRequest util.HttpRequestPort
}

func NewClassRoomHttpRequestAdapter(httpRequestPort util.HttpRequestPort) *ClassRoomHttpRequestAdapter {
	return &ClassRoomHttpRequestAdapter{
		HttpRequest: httpRequestPort,
	}
}

func (svr *ClassRoomHttpRequestAdapter) GetInformationFormStrapi(payloadRequest models.ServiceIncomeRequest) (string, error) {
	var serviceIncomeResponse = models.ServiceIncomeResponse{}
	url := viper.GetString("endpoint.strapi-information-gateway.host") + viper.GetString("endpoint.strapi-information-gateway.path")
	payload, err := json.Marshal(payloadRequest)
	if err != nil {
		return "", err
	}
	response, err := svr.HttpRequest.HTTPRequest(url, constants.POST, payload)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	err = json.Unmarshal([]byte(body), &serviceIncomeResponse)
	if err != nil {
		return "", err
	}
	return serviceIncomeResponse.Payload, nil
}

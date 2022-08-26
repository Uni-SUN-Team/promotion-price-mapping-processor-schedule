package controllers

import (
	"net/http"
	"strconv"
	classroomprice "unisun/api/class-room-price-mapping-processor-schedule/src/models/class-room-price"
	"unisun/api/class-room-price-mapping-processor-schedule/src/ports/service"

	"github.com/gin-gonic/gin"
)

type ClassRoomPriceControllerAdapter struct {
	ClassRoomPriceService service.ClassRoomPriceRepo
}

func NewClassRoomPriceControllerAdapter(classRoomPriceService service.ClassRoomPriceRepo) *ClassRoomPriceControllerAdapter {
	return &ClassRoomPriceControllerAdapter{
		ClassRoomPriceService: classRoomPriceService,
	}
}

func (srv *ClassRoomPriceControllerAdapter) GetClassRoomPriceById(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, struct {
			Error string `json:"error"`
			Code  int    `json:"code"`
		}{
			Error: err.Error(),
			Code:  http.StatusInternalServerError,
		})
	}
	service := srv.ClassRoomPriceService.GetClassRoomPrice(id)
	c.AbortWithStatusJSON(http.StatusOK, struct {
		Data classroomprice.ClassRoomPriceEntity `json:"data"`
	}{
		Data: *service,
	})
}

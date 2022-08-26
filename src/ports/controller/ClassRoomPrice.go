package controller

import "github.com/gin-gonic/gin"

type ClassRoomPriceControllerPort interface {
	GetClassRoomPriceById(c *gin.Context)
}

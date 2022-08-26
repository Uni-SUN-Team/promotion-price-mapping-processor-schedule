package routes

import (
	"unisun/api/class-room-price-mapping-processor-schedule/src/ports/controller"

	"github.com/gin-gonic/gin"
)

type ClassRoomPriceRouteAdapter struct {
	ClassRoomPriceController controller.ClassRoomPriceControllerPort
}

func NewClassRoomPriceRouteAdapter(classRoomPriceController controller.ClassRoomPriceControllerPort) *ClassRoomPriceRouteAdapter {
	return &ClassRoomPriceRouteAdapter{
		ClassRoomPriceController: classRoomPriceController,
	}
}

func (srv *ClassRoomPriceRouteAdapter) ClassRoomPriceRoute(g *gin.RouterGroup) {
	g.GET("/class-room-price/:id", srv.ClassRoomPriceController.GetClassRoomPriceById)
}

package repositories

import (
	config "unisun/api/class-room-price-mapping-processor-schedule/src/configs"
	"unisun/api/class-room-price-mapping-processor-schedule/src/entitys"

	"gorm.io/gorm"
)

type ClassRoomPriceAdapter struct {
	Context *gorm.DB
}

func New() *ClassRoomPriceAdapter {
	return &ClassRoomPriceAdapter{
		Context: config.DB,
	}
}

func (srv *ClassRoomPriceAdapter) GetById(id int) *entitys.ClassRoomPriceEntity {
	result := entitys.ClassRoomPriceEntity{}
	srv.Context.Find(&result, "id = ?", id)
	return &result
}

func (srv *ClassRoomPriceAdapter) GetByClassRoomId(id int) *entitys.ClassRoomPriceEntity {
	result := entitys.ClassRoomPriceEntity{}
	srv.Context.Find(&result, "classroomid = ?", id)
	return &result
}

func (srv *ClassRoomPriceAdapter) Save(classRoom entitys.ClassRoomPriceEntity) {
	srv.Context.Create(&classRoom)
}

func (srv *ClassRoomPriceAdapter) Update(classRoom entitys.ClassRoomPriceEntity) {
	srv.Context.Model(&classRoom).Where("id", classRoom.Id).Update("classroomid", classRoom.ClassRoomId).Update("regularprice", classRoom.RegularPrice).Update("specialprice", classRoom.SpecialPrice)
}

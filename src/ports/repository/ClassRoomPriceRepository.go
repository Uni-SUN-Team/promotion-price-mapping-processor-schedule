package repository

import (
	"unisun/api/class-room-price-mapping-processor-schedule/src/entitys"
)

type ClassRoomPricePort interface {
	GetById(id int) *entitys.ClassRoomPriceEntity
	GetByClassRoomId(id int) *entitys.ClassRoomPriceEntity
	Save(classRoom entitys.ClassRoomPriceEntity)
	Update(classRoom entitys.ClassRoomPriceEntity)
	GetByAdvisor(id string) *[]entitys.ClassRoomPriceEntity
	GetByCategories(id string) *[]entitys.ClassRoomPriceEntity
}

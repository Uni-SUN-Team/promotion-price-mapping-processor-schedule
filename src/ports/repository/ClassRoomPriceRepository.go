package repository

import "unisun/api/promotion-price-mapping-processor-schedule/src/entitys"

type ClassRoomPricePort interface {
	GetById(id int) *entitys.ClassRoomPrice
	GetByClassRoomId(id int) *entitys.ClassRoomPrice
	Save(classRoom entitys.ClassRoomPrice)
	Update(classRoom entitys.ClassRoomPrice)
	GetByAdvisor(id string) *[]entitys.ClassRoomPrice
	GetByCategories(id string) *[]entitys.ClassRoomPrice
}

package repositories

import (
	"strings"
	config "unisun/api/promotion-price-mapping-processor-schedule/src/configs"
	"unisun/api/promotion-price-mapping-processor-schedule/src/entitys"

	"gorm.io/gorm"
)

type ClassRoomPriceRepositoriesAdapter struct {
	Context *gorm.DB
}

func NewClassRoomPriceRepositoriesAdapter() *ClassRoomPriceRepositoriesAdapter {
	return &ClassRoomPriceRepositoriesAdapter{
		Context: config.DB,
	}
}

func (srv *ClassRoomPriceRepositoriesAdapter) GetById(id int) *entitys.ClassRoomPrice {
	result := entitys.ClassRoomPrice{}
	srv.Context.Find(&result, "id = ?", id)
	return &result
}

func (srv *ClassRoomPriceRepositoriesAdapter) GetByClassRoomId(id int) *entitys.ClassRoomPrice {
	result := entitys.ClassRoomPrice{}
	srv.Context.Find(&result, "class_room_id = ?", id)
	return &result
}

func (srv *ClassRoomPriceRepositoriesAdapter) GetByAdvisor(id string) *[]entitys.ClassRoomPrice {
	result := []entitys.ClassRoomPrice{}
	srv.Context.Find(&result, "advisors like ?", strings.Join([]string{"%", id, "%"}, ""))
	return &result
}

func (srv *ClassRoomPriceRepositoriesAdapter) GetByCategories(id string) *[]entitys.ClassRoomPrice {
	result := []entitys.ClassRoomPrice{}
	srv.Context.Find(&result, "categories like ?", strings.Join([]string{"%", id, "%"}, ""))
	return &result
}

func (srv *ClassRoomPriceRepositoriesAdapter) Save(classRoom entitys.ClassRoomPrice) {
	srv.Context.Create(&classRoom)
}

func (srv *ClassRoomPriceRepositoriesAdapter) Update(classRoom entitys.ClassRoomPrice) {
	srv.Context.Model(&classRoom).Where("id", classRoom.Id).Update("class_room_id", classRoom.ClassRoomId).Update("regular_price", classRoom.RegularPrice).Update("special_price", classRoom.SpecialPrice)
}

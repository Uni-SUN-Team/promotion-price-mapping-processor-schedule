package service

import classroomprice "unisun/api/class-room-price-mapping-processor-schedule/src/models/class-room-price"

type ClassRoomPrice interface {
	ManagePrice()
}

type ClassRoomPriceRepo interface {
	GetClassRoomPrice(id int) *classroomprice.ClassRoomPriceEntity
}

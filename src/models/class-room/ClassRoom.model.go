package classroom

type ClassRooms struct {
	Data []ClassRoomBody `json:"data"`
	Meta pagination      `json:"meta"`
}

type ClassRoom struct {
	Data ClassRoomBody `json:"data"`
	Meta pagination    `json:"meta"`
}

type ClassRoomBody struct {
	Id    int     `json:"id"`
	Price float64 `json:"price"`
}

type pagination struct {
	Pagination paginationContent `json:"pagination"`
}

type paginationContent struct {
	Page      int64 `json:"page"`
	PageSize  int64 `json:"pageSize"`
	PageCount int64 `json:"pageCount"`
	Total     int64 `json:"total"`
}

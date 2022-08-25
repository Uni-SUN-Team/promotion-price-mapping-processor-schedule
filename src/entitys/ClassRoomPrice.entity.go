package entitys

type ClassRoomPriceEntity struct {
	Id           int
	ClassRoomId  int
	RegularPrice float64
	SpecialPrice float64
	Advisors     string
	Categories   string
}

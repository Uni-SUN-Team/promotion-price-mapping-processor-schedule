package promotion

import "time"

type Promotions struct {
	Data []PromotionBody `json:"data"`
	Meta pagination      `json:"meta"`
}

type PromotionBody struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	PublishedAt time.Time `json:"publishedAt"`
	Locale      string    `json:"locale"`
	Group       []struct {
		Id         int    `json:"id"`
		Component  string `json:"__component"`
		Categories []struct {
			Id int `json:"id"`
		} `json:"categories"`
		Advisors []struct {
			Id int `json:"id"`
		} `json:"advisors"`
		Classrooms []struct {
			Id int `json:"id"`
		} `json:"class_rooms"`
	} `json:"Group"`
	Discount []struct {
		Id        int     `json:"id"`
		Component string  `json:"__component"`
		Discount  float64 `json:"discount"`
	} `json:"Discount"`
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

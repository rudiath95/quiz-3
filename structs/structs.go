package structs

import "time"

type Category struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type Books struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Image       string    `json:"image_url"`
	ReleaseYear int64     `json:"release_year"`
	Price       string    `json:"price"`
	TotalPage   int64     `json:"total_page"`
	Thickness   string    `json:"thickness"`
	Created_at  time.Time `json:"created_at"`
	Updated_at  time.Time `json:"updated_at"`
	CategoryID  int64     `json:"category_id"`
}

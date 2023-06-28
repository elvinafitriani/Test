package model

type Product struct {
	ID    int     `json:"id" binding:"required"`
	Name  string  `json:"name" binding:"required"`
	Price float64 `json:"price" binding:"required"`
}

// soal 3
type User struct {
	ID    string
	Name  string
	Email string
	Age   int
}

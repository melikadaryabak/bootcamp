package entity

type Bootcamp struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
    Description string `json:"description"`
    Category Category `json:"category"`
}
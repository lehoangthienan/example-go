package domain

type Book struct {
	Model
	Name        string `json:"name"`
	Category_id UUID   `sql:",type:uuid" json:"category_id"`
	Author      string `json:"author"`
	Description string `json:"description"`
}

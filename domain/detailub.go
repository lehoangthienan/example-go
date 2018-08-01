package domain

import "time"

// User describe user in systenm
type Detailub struct {
	Model
	Book_id UUID      `sql:",type:uuid" json:"book_id"`
	User_id UUID      `sql:",type:uuid" json:"user_id"`
	From    time.Time `json:"from"`
	To      time.Time `json:"to"`
}

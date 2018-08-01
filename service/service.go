package service

import "github.com/lehoangthienan/example-go/service/user"
import "github.com/lehoangthienan/example-go/service/category"
import "github.com/lehoangthienan/example-go/service/book"

// Service define list of all services in projects
type Service struct {
	UserService user.Service

	CategoryService category.Service

	BookService book.Service
}

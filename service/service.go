package service

import "github.com/lehoangthienan/example-go/service/user"
import "github.com/lehoangthienan/example-go/service/category"
import "github.com/lehoangthienan/example-go/service/book"
import "github.com/lehoangthienan/example-go/service/detailub"

// Service define list of all services in projects
type Service struct {
	UserService user.Service

	CategoryService category.Service

	BookService book.Service

	DetailubService detailub.Service
}

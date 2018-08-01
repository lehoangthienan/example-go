package book

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/lehoangthienan/example-go/domain"
)

// pgService implmenter for Book serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for Book service
func (s *pgService) Create(_ context.Context, p *domain.Book) error {
	return s.db.Create(p).Error
}

// Update implement Update for Book service
func (s *pgService) Update(_ context.Context, p *domain.Book) (*domain.Book, error) {
	find := domain.Book{}
	if err := s.db.Find(&find, "name =?", p.Name).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			old := domain.Book{Model: domain.Model{ID: p.ID}}
			if err := s.db.Find(&old).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return nil, ErrNotFound
				}
				return nil, err
			}

			if p.Name != "" {
				old.Name = p.Name
			}
			if p.Author != "" {
				old.Author = p.Author
			}
			if p.Description != "" {
				old.Description = p.Description
			}
			if !p.Category_id.IsZero() {
				old.Category_id = p.Category_id
			}

			return &old, s.db.Save(&old).Error
		}
		return nil, err
	}

	return nil, ErrNameSake
	// old := domain.Book{Model: domain.Model{ID: p.ID}}
	// if err := s.db.Find(&old).Error; err != nil {
	// 	if err == gorm.ErrRecordNotFound {
	// 		return nil, ErrNotFound
	// 	}
	// 	return nil, err
	// }

	// // old.Name = p.Name
	// // old.Author = p.Author
	// // old.Category_id = p.Category_id
	// // old.Description = p.Description

	// if p.Name != "" {
	// 	old.Name = p.Name
	// }
	// if p.Author != "" {
	// 	old.Author = p.Author
	// }
	// if p.Description != "" {
	// 	old.Description = p.Description
	// }
	// if !p.Category_id.IsZero() {
	// 	old.Category_id = p.Category_id
	// }

	// return &old, s.db.Save(&old).Error
}

// Find implement Find for Book service
func (s *pgService) Find(_ context.Context, p *domain.Book) (*domain.Book, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Book service
func (s *pgService) FindAll(_ context.Context) ([]domain.Book, error) {
	res := []domain.Book{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Book service
func (s *pgService) Delete(_ context.Context, p *domain.Book) error {
	old := domain.Book{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}

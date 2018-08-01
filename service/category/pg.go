package category

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/lehoangthienan/example-go/domain"
)

// pgService implmenter for User serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for User service
func (s *pgService) Create(_ context.Context, p *domain.Category) error {
	return s.db.Create(p).Error
}

// Update implement Update for User service
func (s *pgService) Update(_ context.Context, p *domain.Category) (*domain.Category, error) {
	find := domain.Category{}
	if err := s.db.Find(&find, "name =?", p.Name).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			old := domain.Category{Model: domain.Model{ID: p.ID}}
			if err := s.db.Find(&old).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return nil, ErrNotFound
				}
				return nil, err
			}

			old.Name = p.Name

			return &old, s.db.Save(&old).Error
		}
		return nil, err
	}

	return nil, ErrNameExisted

}

// Find implement Find for User service
func (s *pgService) Find(_ context.Context, p *domain.Category) (*domain.Category, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for User service
func (s *pgService) FindAll(_ context.Context) ([]domain.Category, error) {
	res := []domain.Category{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for User service
func (s *pgService) Delete(_ context.Context, p *domain.Category) error {
	fBook := domain.Book{}
	if err := s.db.Find(&fBook, "category_id =?", p.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
	}
	if err := s.db.Delete(fBook).Error; err != nil {
		return err
	}

	old := domain.Category{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}

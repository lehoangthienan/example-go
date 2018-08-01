package detailub

import (
	"context"

	"github.com/jinzhu/gorm"

	"github.com/lehoangthienan/example-go/domain"
)

// pgService implmenter for Detailub serivce in postgres
type pgService struct {
	db *gorm.DB
}

// NewPGService create new PGService
func NewPGService(db *gorm.DB) Service {
	return &pgService{
		db: db,
	}
}

// Create implement Create for Detailub service
func (s *pgService) Create(_ context.Context, p *domain.Detailub) error {
	return s.db.Create(p).Error
}

// Update implement Update for Detailub service
func (s *pgService) Update(_ context.Context, p *domain.Detailub) (*domain.Detailub, error) {
	find := domain.Detailub{}
	if err := s.db.Find(&find, "book_id =?", p.Book_id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			old := domain.Detailub{Model: domain.Model{ID: p.ID}}
			if err := s.db.Find(&old).Error; err != nil {
				if err == gorm.ErrRecordNotFound {
					return nil, ErrNotFound
				}
				return nil, err
			}

			return &old, s.db.Save(&old).Error
		}
		return nil, err
	}
	return nil, ErrBookHanded

}

// Find implement Find for Detailub service
func (s *pgService) Find(_ context.Context, p *domain.Detailub) (*domain.Detailub, error) {
	res := p
	if err := s.db.Find(&res).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, ErrNotFound
		}
		return nil, err
	}

	return res, nil
}

// FindAll implement FindAll for Detailub service
func (s *pgService) FindAll(_ context.Context) ([]domain.Detailub, error) {
	res := []domain.Detailub{}
	return res, s.db.Find(&res).Error
}

// Delete implement Delete for Detailub service
func (s *pgService) Delete(_ context.Context, p *domain.Detailub) error {
	old := domain.Detailub{Model: domain.Model{ID: p.ID}}
	if err := s.db.Find(&old).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ErrNotFound
		}
		return err
	}
	return s.db.Delete(old).Error
}

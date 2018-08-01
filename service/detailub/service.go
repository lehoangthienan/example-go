package detailub

import (
	"context"

	"github.com/lehoangthienan/example-go/domain"
)

// Service interface for project service
type Service interface {
	Create(ctx context.Context, p *domain.Detailub) error
	Update(ctx context.Context, p *domain.Detailub) (*domain.Detailub, error)
	Find(ctx context.Context, p *domain.Detailub) (*domain.Detailub, error)
	FindAll(ctx context.Context) ([]domain.Detailub, error)
	Delete(ctx context.Context, p *domain.Detailub) error
}

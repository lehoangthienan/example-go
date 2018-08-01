package detailub

import (
	"context"

	"github.com/lehoangthienan/example-go/domain"
)

// Declare Regex
const (
	emailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

type validationMiddleware struct {
	Service
}

// ValidationMiddleware ...
func ValidationMiddleware() func(Service) Service {
	return func(next Service) Service {
		return &validationMiddleware{
			Service: next,
		}
	}
}

func (mw validationMiddleware) Create(ctx context.Context, detailub *domain.Detailub) (err error) {
	return mw.Service.Create(ctx, detailub)
}
func (mw validationMiddleware) FindAll(ctx context.Context) ([]domain.Detailub, error) {
	return mw.Service.FindAll(ctx)
}
func (mw validationMiddleware) Find(ctx context.Context, detailub *domain.Detailub) (*domain.Detailub, error) {
	return mw.Service.Find(ctx, detailub)
}

func (mw validationMiddleware) Update(ctx context.Context, detailub *domain.Detailub) (*domain.Detailub, error) {

	return mw.Service.Update(ctx, detailub)
}
func (mw validationMiddleware) Delete(ctx context.Context, detailub *domain.Detailub) error {
	return mw.Service.Delete(ctx, detailub)
}

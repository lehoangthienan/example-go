package detailub

import (
	"context"
	"net/http"
	"time"

	"github.com/go-kit/kit/endpoint"

	"github.com/lehoangthienan/example-go/domain"
	"github.com/lehoangthienan/example-go/service"
)

// CreateData data for CreateDetailub
type CreateData struct {
	Book_id domain.UUID `sql:",type:uuid" json:"book_id"`
	User_id domain.UUID `sql:",type:uuid" json:"user_id"`
	From    time.Time   `json:"from"`
	To      time.Time   `json:"to"`
}

// CreateRequest request struct for CreateDetailub
type CreateRequest struct {
	Detailub CreateData `json:"detailub"`
}

// CreateResponse response struct for CreateDetailub
type CreateResponse struct {
	Detailub domain.Detailub `json:"detailub"`
}

// StatusCode customstatus code for success create Detailub
func (CreateResponse) StatusCode() int {
	return http.StatusCreated
}

// MakeCreateEndpoint make endpoint for create a Detailub
func MakeCreateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(CreateRequest)
			detailub = &domain.Detailub{
				Book_id: req.Detailub.Book_id,
				User_id: req.Detailub.User_id,
				From:    req.Detailub.From,
				To:      req.Detailub.To,
			}
		)

		err := s.DetailubService.Create(ctx, detailub)
		if err != nil {
			return nil, err
		}

		return CreateResponse{Detailub: *detailub}, nil
	}
}

// FindRequest request struct for Find a Detailub
type FindRequest struct {
	DetailubID domain.UUID
}

// FindResponse response struct for Find a Detailub
type FindResponse struct {
	Detailub *domain.Detailub `json:"detailub"`
}

// MakeFindEndPoint make endpoint for find Detailub
func MakeFindEndPoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var detailubFind domain.Detailub
		req := request.(FindRequest)
		detailubFind.ID = req.DetailubID

		detailub, err := s.DetailubService.Find(ctx, &detailubFind)
		if err != nil {
			return nil, err
		}
		return FindResponse{Detailub: detailub}, nil
	}
}

// FindAllRequest request struct for FindAll Detailub
type FindAllRequest struct{}

// FindAllResponse request struct for find all Detailub
type FindAllResponse struct {
	Detailubs []domain.Detailub `json:"detailubs"`
}

// MakeFindAllEndpoint make endpoint for find all Detailub
func MakeFindAllEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(FindAllRequest)
		detailubs, err := s.DetailubService.FindAll(ctx)
		if err != nil {
			return nil, err
		}
		return FindAllResponse{Detailubs: detailubs}, nil
	}
}

// UpdateData data for Create
type UpdateData struct {
	ID      domain.UUID `json:"-"`
	Book_id domain.UUID `sql:",type:uuid" json:"book_id"`
	User_id domain.UUID `sql:",type:uuid" json:"user_id"`
	From    time.Time   `json:"from"`
	To      time.Time   `json:"to"`
}

// UpdateRequest request struct for update
type UpdateRequest struct {
	Detailub UpdateData `json:"detailub"`
}

// UpdateResponse response struct for Create
type UpdateResponse struct {
	Detailub domain.Detailub `json:"detailub"`
}

// MakeUpdateEndpoint make endpoint for update a Detailub
func MakeUpdateEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			req      = request.(UpdateRequest)
			detailub = domain.Detailub{
				Model: domain.Model{ID: req.Detailub.ID},
				// Name:  req.Detailub.Name,
				// Email: req.Detailub.Email,
				Book_id: req.Detailub.Book_id,
				User_id: req.Detailub.User_id,
				From:    req.Detailub.From,
				To:      req.Detailub.To,
			}
		)

		res, err := s.DetailubService.Update(ctx, &detailub)
		if err != nil {
			return nil, err
		}

		return UpdateResponse{Detailub: *res}, nil
	}
}

// DeleteRequest request struct for delete a Detailub
type DeleteRequest struct {
	DetailubID domain.UUID
}

// DeleteResponse response struct for Find a Detailub
type DeleteResponse struct {
	Status string `json:"status"`
}

// MakeDeleteEndpoint make endpoint for update a Detailub
func MakeDeleteEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		var (
			detailubFind = domain.Detailub{}
			req          = request.(DeleteRequest)
		)
		detailubFind.ID = req.DetailubID

		err := s.DetailubService.Delete(ctx, &detailubFind)
		if err != nil {
			return nil, err
		}

		return DeleteResponse{"success"}, nil
	}
}

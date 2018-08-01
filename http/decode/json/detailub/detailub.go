package detailub

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/lehoangthienan/example-go/domain"
	detailubEndpoint "github.com/lehoangthienan/example-go/endpoints/detailub"
)

// FindRequest .
func FindRequest(_ context.Context, r *http.Request) (interface{}, error) {
	detailubID, err := domain.UUIDFromString(chi.URLParam(r, "detailub_id"))
	if err != nil {
		return nil, err
	}
	return detailubEndpoint.FindRequest{DetailubID: detailubID}, nil
}

// FindAllRequest .
func FindAllRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return detailubEndpoint.FindAllRequest{}, nil
}

// CreateRequest .
func CreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var req detailubEndpoint.CreateRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

// UpdateRequest .
func UpdateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	detailubID, err := domain.UUIDFromString(chi.URLParam(r, "detailub_id"))
	if err != nil {
		return nil, err
	}

	var req detailubEndpoint.UpdateRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		return nil, err
	}

	req.Detailub.ID = detailubID

	return req, nil
}

// DeleteRequest .
func DeleteRequest(_ context.Context, r *http.Request) (interface{}, error) {
	detailubID, err := domain.UUIDFromString(chi.URLParam(r, "detailub_id"))
	if err != nil {
		return nil, err
	}
	return detailubEndpoint.DeleteRequest{DetailubID: detailubID}, nil
}

package http

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/lehoangthienan/example-go/endpoints"
	userDecode "github.com/lehoangthienan/example-go/http/decode/json/user"

	categoryDecode "github.com/lehoangthienan/example-go/http/decode/json/category"

	bookDecode "github.com/lehoangthienan/example-go/http/decode/json/book"

	detailubDecode "github.com/lehoangthienan/example-go/http/decode/json/detailub"
)

// NewHTTPHandler ...
func NewHTTPHandler(endpoints endpoints.Endpoints,
	logger log.Logger,
	useCORS bool) http.Handler {
	r := chi.NewRouter()

	// if running on local (using `make dev`), include cors middleware
	if useCORS {
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
		})
		r.Use(cors.Handler)
	}

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Get("/_warm", httptransport.NewServer(
		endpoint.Nop,
		httptransport.NopRequestDecoder,
		httptransport.EncodeJSONResponse,
		options...,
	).ServeHTTP)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.FindAllUser,
			userDecode.FindAllRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/{user_id}", httptransport.NewServer(
			endpoints.FindUser,
			userDecode.FindRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/", httptransport.NewServer(
			endpoints.CreateUser,
			userDecode.CreateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{user_id}", httptransport.NewServer(
			endpoints.UpdateUser,
			userDecode.UpdateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Delete("/{user_id}", httptransport.NewServer(
			endpoints.DeleteUser,
			userDecode.DeleteRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})

	r.Route("/categories", func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.FindAllCategory,
			categoryDecode.FindAllRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/{category_id}", httptransport.NewServer(
			endpoints.FindCategory,
			categoryDecode.FindRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/", httptransport.NewServer(
			endpoints.CreateCategory,
			categoryDecode.CreateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{category_id}", httptransport.NewServer(
			endpoints.UpdateCategory,
			categoryDecode.UpdateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Delete("/{category_id}", httptransport.NewServer(
			endpoints.DeleteCategory,
			categoryDecode.DeleteRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})

	r.Route("/books", func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.FindAllBook,
			bookDecode.FindAllRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/{book_id}", httptransport.NewServer(
			endpoints.FindBook,
			bookDecode.FindRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/", httptransport.NewServer(
			endpoints.CreateBook,
			bookDecode.CreateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{book_id}", httptransport.NewServer(
			endpoints.UpdateBook,
			bookDecode.UpdateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Delete("/{book_id}", httptransport.NewServer(
			endpoints.DeleteBook,
			bookDecode.DeleteRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})

	r.Route("/detailubs", func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.FindAllDetailub,
			detailubDecode.FindAllRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/{detailub_id}", httptransport.NewServer(
			endpoints.FindDetailub,
			detailubDecode.FindRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/", httptransport.NewServer(
			endpoints.CreateDetailub,
			detailubDecode.CreateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{detailub_id}", httptransport.NewServer(
			endpoints.UpdateDetailub,
			detailubDecode.UpdateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Delete("/{detailub_id}", httptransport.NewServer(
			endpoints.DeleteDetailub,
			detailubDecode.DeleteRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})

	return r
}

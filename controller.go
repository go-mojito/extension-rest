package rest

import "github.com/go-mojito/mojito"

// Controller defines the interface of a valid REST controller
type Controller[T any] interface {

	/// Request Handlers

	// Create is called when the POST endpoint is called with a valid object as its body
	Create(ctx mojito.Context, newObj *T) (*T, error)

	// Delete is called when the DELETE endpoint is called with a unique resource identifier
	Delete(ctx mojito.Context, obj *T) (*T, error)

	// Get is called when the rest GET endpoint is called with a unique resource identifier
	Get(ctx mojito.Context, obj *T) (*T, error)

	// List is called when the rest GET endpoint is called without a unique resource identifier
	// Optionally, Go-Mojito will parse filtering options like pagination into the filter object.
	List(ctx mojito.Context, filter ListFilter) ([]T, error)

	// Update is called when the PUT endpoint is called with a valid object as its body and with
	// a valid unique resource identifier
	Update(ctx mojito.Context, obj *T, updatedObj *T) (*T, error)

	/// Internal Methods

	// ResolveIdentifier defines how the rest controller can retrieve a unique value, based on the
	// resource identifier
	ResolveIdentifier(ctx mojito.Context, identifier string) (*T, error)
}

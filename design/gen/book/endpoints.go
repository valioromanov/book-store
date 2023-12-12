// Code generated by goa v3.14.1, DO NOT EDIT.
//
// book endpoints
//
// Command:
// $ goa gen book-store/design

package book

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Endpoints wraps the "book" service endpoints.
type Endpoints struct {
	GetBook  goa.Endpoint
	PostBook goa.Endpoint
}

// NewEndpoints wraps the methods of the "book" service with endpoints.
func NewEndpoints(s Service) *Endpoints {
	return &Endpoints{
		GetBook:  NewGetBookEndpoint(s),
		PostBook: NewPostBookEndpoint(s),
	}
}

// Use applies the given middleware to all the "book" service endpoints.
func (e *Endpoints) Use(m func(goa.Endpoint) goa.Endpoint) {
	e.GetBook = m(e.GetBook)
	e.PostBook = m(e.PostBook)
}

// NewGetBookEndpoint returns an endpoint function that calls the method
// "getBook" of service "book".
func NewGetBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*GetBookPayload)
		res, err := s.GetBook(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedBookResult(res, "default")
		return vres, nil
	}
}

// NewPostBookEndpoint returns an endpoint function that calls the method
// "postBook" of service "book".
func NewPostBookEndpoint(s Service) goa.Endpoint {
	return func(ctx context.Context, req any) (any, error) {
		p := req.(*BookResult)
		res, err := s.PostBook(ctx, p)
		if err != nil {
			return nil, err
		}
		vres := NewViewedBookResult(res, "resultOperation")
		return vres, nil
	}
}

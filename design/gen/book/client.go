// Code generated by goa v3.14.1, DO NOT EDIT.
//
// book client
//
// Command:
// $ goa gen book-store/design

package book

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "book" service client.
type Client struct {
	GetBookEndpoint    goa.Endpoint
	PostBookEndpoint   goa.Endpoint
	PatchBookEndpoint  goa.Endpoint
	DeleteBookEndpoint goa.Endpoint
}

// NewClient initializes a "book" service client given the endpoints.
func NewClient(getBook, postBook, patchBook, deleteBook goa.Endpoint) *Client {
	return &Client{
		GetBookEndpoint:    getBook,
		PostBookEndpoint:   postBook,
		PatchBookEndpoint:  patchBook,
		DeleteBookEndpoint: deleteBook,
	}
}

// GetBook calls the "getBook" endpoint of the "book" service.
// GetBook may return the following errors:
//   - "NotFound" (type *goa.ServiceError)
//   - "BadRequest" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) GetBook(ctx context.Context, p *GetBookPayload) (res *BookResult, err error) {
	var ires any
	ires, err = c.GetBookEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*BookResult), nil
}

// PostBook calls the "postBook" endpoint of the "book" service.
// PostBook may return the following errors:
//   - "BadRequest" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) PostBook(ctx context.Context, p *BookReq) (res *BookResult, err error) {
	var ires any
	ires, err = c.PostBookEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*BookResult), nil
}

// PatchBook calls the "patchBook" endpoint of the "book" service.
// PatchBook may return the following errors:
//   - "NotFound" (type *goa.ServiceError)
//   - "BadRequest" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) PatchBook(ctx context.Context, p *BookPathcReq) (err error) {
	_, err = c.PatchBookEndpoint(ctx, p)
	return
}

// DeleteBook calls the "deleteBook" endpoint of the "book" service.
// DeleteBook may return the following errors:
//   - "NotFound" (type *goa.ServiceError)
//   - "BadRequest" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) DeleteBook(ctx context.Context, p *DeleteBookPayload) (err error) {
	_, err = c.DeleteBookEndpoint(ctx, p)
	return
}

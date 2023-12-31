// Code generated by goa v3.14.1, DO NOT EDIT.
//
// book HTTP client encoders and decoders
//
// Command:
// $ goa gen book-store/design

package client

import (
	book "book-store/design/gen/book"
	bookviews "book-store/design/gen/book/views"
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	goahttp "goa.design/goa/v3/http"
)

// BuildGetBookRequest instantiates a HTTP request object with method and path
// set to call the "book" service "getBook" endpoint
func (c *Client) BuildGetBookRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		bookID int
	)
	{
		p, ok := v.(*book.GetBookPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("book", "getBook", "*book.GetBookPayload", v)
		}
		bookID = p.BookID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetBookBookPath(bookID)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("book", "getBook", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetBookResponse returns a decoder for responses returned by the book
// getBook endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetBookResponse may return the following errors:
//   - "NotFound" (type *goa.ServiceError): http.StatusNotFound
//   - "BadRequest" (type *goa.ServiceError): http.StatusBadRequest
//   - error: internal error
func DecodeGetBookResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetBookResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("book", "getBook", err)
			}
			p := NewGetBookBookResultOK(&body)
			view := "default"
			vres := &bookviews.BookResult{Projected: p, View: view}
			if err = bookviews.ValidateBookResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("book", "getBook", err)
			}
			res := book.NewBookResult(vres)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetBookNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("book", "getBook", err)
			}
			err = ValidateGetBookNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("book", "getBook", err)
			}
			return nil, NewGetBookNotFound(&body)
		case http.StatusBadRequest:
			var (
				body GetBookBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("book", "getBook", err)
			}
			err = ValidateGetBookBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("book", "getBook", err)
			}
			return nil, NewGetBookBadRequest(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("book", "getBook", resp.StatusCode, string(body))
		}
	}
}

// BuildPostBookRequest instantiates a HTTP request object with method and path
// set to call the "book" service "postBook" endpoint
func (c *Client) BuildPostBookRequest(ctx context.Context, v any) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PostBookBookPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("book", "postBook", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePostBookRequest returns an encoder for requests sent to the book
// postBook server.
func EncodePostBookRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*book.BookReq)
		if !ok {
			return goahttp.ErrInvalidType("book", "postBook", "*book.BookReq", v)
		}
		body := NewPostBookRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("book", "postBook", err)
		}
		return nil
	}
}

// DecodePostBookResponse returns a decoder for responses returned by the book
// postBook endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodePostBookResponse may return the following errors:
//   - "BadRequest" (type *goa.ServiceError): http.StatusBadRequest
//   - error: internal error
func DecodePostBookResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body PostBookResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("book", "postBook", err)
			}
			p := NewPostBookBookResultOK(&body)
			view := "resultOperation"
			vres := &bookviews.BookResult{Projected: p, View: view}
			if err = bookviews.ValidateBookResult(vres); err != nil {
				return nil, goahttp.ErrValidationError("book", "postBook", err)
			}
			res := book.NewBookResult(vres)
			return res, nil
		case http.StatusBadRequest:
			var (
				body PostBookBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("book", "postBook", err)
			}
			err = ValidatePostBookBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("book", "postBook", err)
			}
			return nil, NewPostBookBadRequest(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("book", "postBook", resp.StatusCode, string(body))
		}
	}
}

// BuildPatchBookRequest instantiates a HTTP request object with method and
// path set to call the "book" service "patchBook" endpoint
func (c *Client) BuildPatchBookRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*book.BookPathcReq)
		if !ok {
			return nil, goahttp.ErrInvalidType("book", "patchBook", "*book.BookPathcReq", v)
		}
		if p.ID != nil {
			id = *p.ID
		}
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: PatchBookBookPath(id)}
	req, err := http.NewRequest("PATCH", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("book", "patchBook", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodePatchBookRequest returns an encoder for requests sent to the book
// patchBook server.
func EncodePatchBookRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, any) error {
	return func(req *http.Request, v any) error {
		p, ok := v.(*book.BookPathcReq)
		if !ok {
			return goahttp.ErrInvalidType("book", "patchBook", "*book.BookPathcReq", v)
		}
		body := NewPatchBookRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("book", "patchBook", err)
		}
		return nil
	}
}

// DecodePatchBookResponse returns a decoder for responses returned by the book
// patchBook endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodePatchBookResponse may return the following errors:
//   - "BadRequest" (type *goa.ServiceError): http.StatusBadRequest
//   - "NotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodePatchBookResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body PatchBookBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("book", "patchBook", err)
			}
			err = ValidatePatchBookBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("book", "patchBook", err)
			}
			return nil, NewPatchBookBadRequest(&body)
		case http.StatusNotFound:
			var (
				body PatchBookNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("book", "patchBook", err)
			}
			err = ValidatePatchBookNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("book", "patchBook", err)
			}
			return nil, NewPatchBookNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("book", "patchBook", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteBookRequest instantiates a HTTP request object with method and
// path set to call the "book" service "deleteBook" endpoint
func (c *Client) BuildDeleteBookRequest(ctx context.Context, v any) (*http.Request, error) {
	var (
		bookID int
	)
	{
		p, ok := v.(*book.DeleteBookPayload)
		if !ok {
			return nil, goahttp.ErrInvalidType("book", "deleteBook", "*book.DeleteBookPayload", v)
		}
		bookID = p.BookID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteBookBookPath(bookID)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("book", "deleteBook", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteBookResponse returns a decoder for responses returned by the
// book deleteBook endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDeleteBookResponse may return the following errors:
//   - "BadRequest" (type *goa.ServiceError): http.StatusBadRequest
//   - "NotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeDeleteBookResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (any, error) {
	return func(resp *http.Response) (any, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusBadRequest:
			var (
				body DeleteBookBadRequestResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("book", "deleteBook", err)
			}
			err = ValidateDeleteBookBadRequestResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("book", "deleteBook", err)
			}
			return nil, NewDeleteBookBadRequest(&body)
		case http.StatusNotFound:
			var (
				body DeleteBookNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("book", "deleteBook", err)
			}
			err = ValidateDeleteBookNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("book", "deleteBook", err)
			}
			return nil, NewDeleteBookNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("book", "deleteBook", resp.StatusCode, string(body))
		}
	}
}

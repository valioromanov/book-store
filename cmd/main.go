package main

import (
	"book-store/cmd/internal"
	"book-store/design/gen/book"
	"book-store/design/gen/http/book/server"
	"book-store/pkg/repository"
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

func main() {
	repository := repository.NewBookRepo()
	controller := internal.NewController(repository)
	endpoints := book.NewEndpoints(controller)
	mux := goahttp.NewMuxer()
	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	svr := server.New(endpoints, mux, dec, enc, nil, nil, nil)
	server.Mount(mux, svr)
	httpsvr := &http.Server{
		Addr:    "localhost:8081",
		Handler: mux,
	}
	if err := httpsvr.ListenAndServe(); err != nil {
		panic(err)
	}
}

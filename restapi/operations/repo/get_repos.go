// Code generated by go-swagger; DO NOT EDIT.

package repo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetReposHandlerFunc turns a function with the right signature into a get repos handler
type GetReposHandlerFunc func(GetReposParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetReposHandlerFunc) Handle(params GetReposParams) middleware.Responder {
	return fn(params)
}

// GetReposHandler interface for that can handle valid get repos params
type GetReposHandler interface {
	Handle(GetReposParams) middleware.Responder
}

// NewGetRepos creates a new http.Handler for the get repos operation
func NewGetRepos(ctx *middleware.Context, handler GetReposHandler) *GetRepos {
	return &GetRepos{Context: ctx, Handler: handler}
}

/*GetRepos swagger:route GET /repos repo getRepos

Gets some repos

Returns a list containing all repos.

*/
type GetRepos struct {
	Context *middleware.Context
	Handler GetReposHandler
}

func (o *GetRepos) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetReposParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
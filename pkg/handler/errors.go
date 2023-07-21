package handler

//--
// File that holds the Error response model & renderers that will render said error model in JSON in response to the client's request using Chi.Render.
//--

import (
	"net/http"

	"github.com/go-chi/render"
)

// the base error response model
type ErrorResponse struct {
	Err        error  `json:"-"`           // low-level runtime error
	StatusCode int    `json:"-"`           // http response status code
	StatusText string `json:"status_text"` // http response status message
	Message    string `json:"message"`     // error message for user that usually has more detail.
}

// Various instances of ErrorResponse representing the generic HTTP errors prevalent in the app.
// A struct is like a POJO class. To create an instance of the struct, you basically set up a pointer to the struct of a given type with its elements as literals.
var (
	ErrorNotFound         = &ErrorResponse{StatusCode: 404, Message: "resource not found"}
	ErrorMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "method not allowed"}
)

// Declare the Render method on the *ErrorResponse type so now *ErrorResponse implements the Render interface
func (e *ErrorResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.StatusCode)
	return nil
}

// Generic business logic error response instance to be used across route handlers
func RenderServerError(err error) render.Renderer {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 500,
		StatusText: "internal server error",
		Message:    err.Error(),
	}
}

// Generic error response for invalid/bad requests
func RenderInvalidRequestError(err error) render.Renderer {
	return &ErrorResponse{
		Err:        err,
		StatusCode: 400,
		StatusText: "invalid Request",
		Message:    err.Error(),
	}
}

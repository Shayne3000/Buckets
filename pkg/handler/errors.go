package handler

//--
// Error response payloads & renderers
//--

// the base error response model
type ErrorResponse struct {
	Err        error  `json:"-"`           // low-level runtime error
	StatusCode int    `json:"-"`           // http response status code
	StatusText string `json:"status_text"` // http response status message
	Message    string `json:"message"`     // error message for user that usually has more detail.
}

// Various ErrorResponse instances representing the kinds of errors prevalent in the app.
// A struct is like a POJO class to create an instance of the struct, you basically set up a pointer to the value of the struct of type "something"
var (
	ErrorNotFound         = &ErrorResponse{StatusCode: 404, Message: "Resource not found."}
	ErrorMethodNotAllowed = &ErrorResponse{StatusCode: 405, Message: "Method not allowed."}
	ErrorBadRequest       = &ErrorResponse{StatusCode: 400, Message: "Bad request."}
)

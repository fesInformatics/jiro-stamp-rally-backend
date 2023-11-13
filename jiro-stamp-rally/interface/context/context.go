package context

import (
	"encoding/json"
	"io"
	"net/http"
)

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func (c *Context) JSON(code int, v any) {
	c.w.Header().Set("Contetnt-Type", "application/json; charset=utf-8")
	c.w.WriteHeader(code)
	if v != nil {
		if err := json.NewEncoder(c.w).Encode(v); err != nil {
			c.InternalServerError(err)
		}
	}
}

func (c *Context) BadRequest(err error) {
	c.JSON(http.StatusBadRequest, ErrorResponse{StatusCode: http.StatusBadRequest, Message: err.Error()})
}

func (c *Context) MethodNotAllowed(err error) {
	c.JSON(http.StatusMethodNotAllowed, ErrorResponse{StatusCode: http.StatusMethodNotAllowed, Message: err.Error()})
}

func (c *Context) InternalServerError(err error) {
	c.JSON(http.StatusInternalServerError, ErrorResponse{StatusCode: http.StatusInternalServerError, Message: err.Error()})
}

func (c Context) GetHttpMethod() string {
	return c.r.Method
}

func (c Context) GetBody() io.ReadCloser {
	return c.r.Body
}

func NewContext(w http.ResponseWriter, r *http.Request) Context {
	return Context{
		w: w,
		r: r,
	}
}

type ErrorResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

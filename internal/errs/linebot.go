package errs

import (
	"bytes"
	"fmt"
)

type errorResponseDetail struct {
	Message  string `json:"message"`
	Property string `json:"property"`
}

// LinebotErrorResponse type
type LinebotErrorResponse struct {
	Message          string                `json:"message"`
	Details          []errorResponseDetail `json:"details"`
	Error            string                `json:"error"`
	ErrorDescription string                `json:"error_description"`
}

// LinebotError type
type LinebotError struct {
	Code     int
	Response *LinebotErrorResponse
}

// NewLinebotError method
func NewLinebotError(
	code int,
	resp *LinebotErrorResponse,
) error {
	return &LinebotError{
		Code:     code,
		Response: resp,
	}
}

// Error method
func (e *LinebotError) Error() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "LinebotError %d ", e.Code)
	if e.Response != nil {
		fmt.Fprintf(&buf, "%s", e.Response.Message)
		for _, d := range e.Response.Details {
			fmt.Fprintf(&buf, "\n[%s] %s", d.Property, d.Message)
		}
	}
	return buf.String()
}

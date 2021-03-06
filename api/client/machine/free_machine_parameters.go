// Code generated by go-swagger; DO NOT EDIT.

package machine

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewFreeMachineParams creates a new FreeMachineParams object
// with the default values initialized.
func NewFreeMachineParams() *FreeMachineParams {
	var ()
	return &FreeMachineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewFreeMachineParamsWithTimeout creates a new FreeMachineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewFreeMachineParamsWithTimeout(timeout time.Duration) *FreeMachineParams {
	var ()
	return &FreeMachineParams{

		timeout: timeout,
	}
}

// NewFreeMachineParamsWithContext creates a new FreeMachineParams object
// with the default values initialized, and the ability to set a context for a request
func NewFreeMachineParamsWithContext(ctx context.Context) *FreeMachineParams {
	var ()
	return &FreeMachineParams{

		Context: ctx,
	}
}

// NewFreeMachineParamsWithHTTPClient creates a new FreeMachineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewFreeMachineParamsWithHTTPClient(client *http.Client) *FreeMachineParams {
	var ()
	return &FreeMachineParams{
		HTTPClient: client,
	}
}

/*FreeMachineParams contains all the parameters to send to the API endpoint
for the free machine operation typically these are written to a http.Request
*/
type FreeMachineParams struct {

	/*ID
	  identifier of the machine

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the free machine params
func (o *FreeMachineParams) WithTimeout(timeout time.Duration) *FreeMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the free machine params
func (o *FreeMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the free machine params
func (o *FreeMachineParams) WithContext(ctx context.Context) *FreeMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the free machine params
func (o *FreeMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the free machine params
func (o *FreeMachineParams) WithHTTPClient(client *http.Client) *FreeMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the free machine params
func (o *FreeMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the free machine params
func (o *FreeMachineParams) WithID(id string) *FreeMachineParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the free machine params
func (o *FreeMachineParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *FreeMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

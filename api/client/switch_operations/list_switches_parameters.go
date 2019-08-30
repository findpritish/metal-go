// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListSwitchesParams creates a new ListSwitchesParams object
// with the default values initialized.
func NewListSwitchesParams() *ListSwitchesParams {

	return &ListSwitchesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSwitchesParamsWithTimeout creates a new ListSwitchesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSwitchesParamsWithTimeout(timeout time.Duration) *ListSwitchesParams {

	return &ListSwitchesParams{

		timeout: timeout,
	}
}

// NewListSwitchesParamsWithContext creates a new ListSwitchesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSwitchesParamsWithContext(ctx context.Context) *ListSwitchesParams {

	return &ListSwitchesParams{

		Context: ctx,
	}
}

// NewListSwitchesParamsWithHTTPClient creates a new ListSwitchesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSwitchesParamsWithHTTPClient(client *http.Client) *ListSwitchesParams {

	return &ListSwitchesParams{
		HTTPClient: client,
	}
}

/*ListSwitchesParams contains all the parameters to send to the API endpoint
for the list switches operation typically these are written to a http.Request
*/
type ListSwitchesParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list switches params
func (o *ListSwitchesParams) WithTimeout(timeout time.Duration) *ListSwitchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list switches params
func (o *ListSwitchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list switches params
func (o *ListSwitchesParams) WithContext(ctx context.Context) *ListSwitchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list switches params
func (o *ListSwitchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list switches params
func (o *ListSwitchesParams) WithHTTPClient(client *http.Client) *ListSwitchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list switches params
func (o *ListSwitchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListSwitchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

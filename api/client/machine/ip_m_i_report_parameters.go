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

	models "github.com/metal-stack/metal-go/api/models"
)

// NewIPMIReportParams creates a new IPMIReportParams object
// with the default values initialized.
func NewIPMIReportParams() *IPMIReportParams {
	var ()
	return &IPMIReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPMIReportParamsWithTimeout creates a new IPMIReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPMIReportParamsWithTimeout(timeout time.Duration) *IPMIReportParams {
	var ()
	return &IPMIReportParams{

		timeout: timeout,
	}
}

// NewIPMIReportParamsWithContext creates a new IPMIReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPMIReportParamsWithContext(ctx context.Context) *IPMIReportParams {
	var ()
	return &IPMIReportParams{

		Context: ctx,
	}
}

// NewIPMIReportParamsWithHTTPClient creates a new IPMIReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPMIReportParamsWithHTTPClient(client *http.Client) *IPMIReportParams {
	var ()
	return &IPMIReportParams{
		HTTPClient: client,
	}
}

/*IPMIReportParams contains all the parameters to send to the API endpoint
for the ipmi report operation typically these are written to a http.Request
*/
type IPMIReportParams struct {

	/*Body*/
	Body *models.V1MachineIPMIReport

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipmi report params
func (o *IPMIReportParams) WithTimeout(timeout time.Duration) *IPMIReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipmi report params
func (o *IPMIReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipmi report params
func (o *IPMIReportParams) WithContext(ctx context.Context) *IPMIReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipmi report params
func (o *IPMIReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipmi report params
func (o *IPMIReportParams) WithHTTPClient(client *http.Client) *IPMIReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipmi report params
func (o *IPMIReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ipmi report params
func (o *IPMIReportParams) WithBody(body *models.V1MachineIPMIReport) *IPMIReportParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ipmi report params
func (o *IPMIReportParams) SetBody(body *models.V1MachineIPMIReport) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *IPMIReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

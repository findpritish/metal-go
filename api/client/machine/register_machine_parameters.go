// Code generated by go-swagger; DO NOT EDIT.

package machine

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

	models "github.com/metal-pod/metal-go/api/models"
)

// NewRegisterMachineParams creates a new RegisterMachineParams object
// with the default values initialized.
func NewRegisterMachineParams() *RegisterMachineParams {
	var ()
	return &RegisterMachineParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRegisterMachineParamsWithTimeout creates a new RegisterMachineParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRegisterMachineParamsWithTimeout(timeout time.Duration) *RegisterMachineParams {
	var ()
	return &RegisterMachineParams{

		timeout: timeout,
	}
}

// NewRegisterMachineParamsWithContext creates a new RegisterMachineParams object
// with the default values initialized, and the ability to set a context for a request
func NewRegisterMachineParamsWithContext(ctx context.Context) *RegisterMachineParams {
	var ()
	return &RegisterMachineParams{

		Context: ctx,
	}
}

// NewRegisterMachineParamsWithHTTPClient creates a new RegisterMachineParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRegisterMachineParamsWithHTTPClient(client *http.Client) *RegisterMachineParams {
	var ()
	return &RegisterMachineParams{
		HTTPClient: client,
	}
}

/*RegisterMachineParams contains all the parameters to send to the API endpoint
for the register machine operation typically these are written to a http.Request
*/
type RegisterMachineParams struct {

	/*Body*/
	Body *models.V1MachineRegisterRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the register machine params
func (o *RegisterMachineParams) WithTimeout(timeout time.Duration) *RegisterMachineParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the register machine params
func (o *RegisterMachineParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the register machine params
func (o *RegisterMachineParams) WithContext(ctx context.Context) *RegisterMachineParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the register machine params
func (o *RegisterMachineParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the register machine params
func (o *RegisterMachineParams) WithHTTPClient(client *http.Client) *RegisterMachineParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the register machine params
func (o *RegisterMachineParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the register machine params
func (o *RegisterMachineParams) WithBody(body *models.V1MachineRegisterRequest) *RegisterMachineParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the register machine params
func (o *RegisterMachineParams) SetBody(body *models.V1MachineRegisterRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *RegisterMachineParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

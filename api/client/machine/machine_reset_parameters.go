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

// NewMachineResetParams creates a new MachineResetParams object
// with the default values initialized.
func NewMachineResetParams() *MachineResetParams {
	var ()
	return &MachineResetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMachineResetParamsWithTimeout creates a new MachineResetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMachineResetParamsWithTimeout(timeout time.Duration) *MachineResetParams {
	var ()
	return &MachineResetParams{

		timeout: timeout,
	}
}

// NewMachineResetParamsWithContext creates a new MachineResetParams object
// with the default values initialized, and the ability to set a context for a request
func NewMachineResetParamsWithContext(ctx context.Context) *MachineResetParams {
	var ()
	return &MachineResetParams{

		Context: ctx,
	}
}

// NewMachineResetParamsWithHTTPClient creates a new MachineResetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMachineResetParamsWithHTTPClient(client *http.Client) *MachineResetParams {
	var ()
	return &MachineResetParams{
		HTTPClient: client,
	}
}

/*MachineResetParams contains all the parameters to send to the API endpoint
for the machine reset operation typically these are written to a http.Request
*/
type MachineResetParams struct {

	/*Body*/
	Body models.V1EmptyBody
	/*ID
	  identifier of the machine

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the machine reset params
func (o *MachineResetParams) WithTimeout(timeout time.Duration) *MachineResetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the machine reset params
func (o *MachineResetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the machine reset params
func (o *MachineResetParams) WithContext(ctx context.Context) *MachineResetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the machine reset params
func (o *MachineResetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the machine reset params
func (o *MachineResetParams) WithHTTPClient(client *http.Client) *MachineResetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the machine reset params
func (o *MachineResetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the machine reset params
func (o *MachineResetParams) WithBody(body models.V1EmptyBody) *MachineResetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the machine reset params
func (o *MachineResetParams) SetBody(body models.V1EmptyBody) {
	o.Body = body
}

// WithID adds the id to the machine reset params
func (o *MachineResetParams) WithID(id string) *MachineResetParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the machine reset params
func (o *MachineResetParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *MachineResetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

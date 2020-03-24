// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/metal-stack/metal-go/api/models"
)

// HealthReader is a Reader for the Health structure.
type HealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewHealthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHealthOK creates a HealthOK with default headers values
func NewHealthOK() *HealthOK {
	return &HealthOK{}
}

/*HealthOK handles this case with default header values.

OK
*/
type HealthOK struct {
	Payload *models.RestStatus
}

func (o *HealthOK) Error() string {
	return fmt.Sprintf("[GET /v1/health][%d] healthOK  %+v", 200, o.Payload)
}

func (o *HealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewHealthInternalServerError creates a HealthInternalServerError with default headers values
func NewHealthInternalServerError() *HealthInternalServerError {
	return &HealthInternalServerError{}
}

/*HealthInternalServerError handles this case with default header values.

Unhealthy
*/
type HealthInternalServerError struct {
	Payload *models.RestStatus
}

func (o *HealthInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/health][%d] healthInternalServerError  %+v", 500, o.Payload)
}

func (o *HealthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

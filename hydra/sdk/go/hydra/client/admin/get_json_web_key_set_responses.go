// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/justpark/auth/hydra/sdk/go/hydra/models"
)

// GetJSONWebKeySetReader is a Reader for the GetJSONWebKeySet structure.
type GetJSONWebKeySetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJSONWebKeySetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetJSONWebKeySetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetJSONWebKeySetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetJSONWebKeySetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetJSONWebKeySetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetJSONWebKeySetOK creates a GetJSONWebKeySetOK with default headers values
func NewGetJSONWebKeySetOK() *GetJSONWebKeySetOK {
	return &GetJSONWebKeySetOK{}
}

/*GetJSONWebKeySetOK handles this case with default header values.

JSONWebKeySet
*/
type GetJSONWebKeySetOK struct {
	Payload *models.SwaggerJSONWebKeySet
}

func (o *GetJSONWebKeySetOK) Error() string {
	return fmt.Sprintf("[GET /keys/{set}][%d] getJsonWebKeySetOK  %+v", 200, o.Payload)
}

func (o *GetJSONWebKeySetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SwaggerJSONWebKeySet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJSONWebKeySetUnauthorized creates a GetJSONWebKeySetUnauthorized with default headers values
func NewGetJSONWebKeySetUnauthorized() *GetJSONWebKeySetUnauthorized {
	return &GetJSONWebKeySetUnauthorized{}
}

/*GetJSONWebKeySetUnauthorized handles this case with default header values.

genericError
*/
type GetJSONWebKeySetUnauthorized struct {
	Payload *models.GenericError
}

func (o *GetJSONWebKeySetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /keys/{set}][%d] getJsonWebKeySetUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJSONWebKeySetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJSONWebKeySetForbidden creates a GetJSONWebKeySetForbidden with default headers values
func NewGetJSONWebKeySetForbidden() *GetJSONWebKeySetForbidden {
	return &GetJSONWebKeySetForbidden{}
}

/*GetJSONWebKeySetForbidden handles this case with default header values.

genericError
*/
type GetJSONWebKeySetForbidden struct {
	Payload *models.GenericError
}

func (o *GetJSONWebKeySetForbidden) Error() string {
	return fmt.Sprintf("[GET /keys/{set}][%d] getJsonWebKeySetForbidden  %+v", 403, o.Payload)
}

func (o *GetJSONWebKeySetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJSONWebKeySetInternalServerError creates a GetJSONWebKeySetInternalServerError with default headers values
func NewGetJSONWebKeySetInternalServerError() *GetJSONWebKeySetInternalServerError {
	return &GetJSONWebKeySetInternalServerError{}
}

/*GetJSONWebKeySetInternalServerError handles this case with default header values.

genericError
*/
type GetJSONWebKeySetInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetJSONWebKeySetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /keys/{set}][%d] getJsonWebKeySetInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJSONWebKeySetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

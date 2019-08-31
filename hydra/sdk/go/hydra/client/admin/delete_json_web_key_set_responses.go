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

// DeleteJSONWebKeySetReader is a Reader for the DeleteJSONWebKeySet structure.
type DeleteJSONWebKeySetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteJSONWebKeySetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteJSONWebKeySetNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteJSONWebKeySetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteJSONWebKeySetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteJSONWebKeySetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteJSONWebKeySetNoContent creates a DeleteJSONWebKeySetNoContent with default headers values
func NewDeleteJSONWebKeySetNoContent() *DeleteJSONWebKeySetNoContent {
	return &DeleteJSONWebKeySetNoContent{}
}

/*DeleteJSONWebKeySetNoContent handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type DeleteJSONWebKeySetNoContent struct {
}

func (o *DeleteJSONWebKeySetNoContent) Error() string {
	return fmt.Sprintf("[DELETE /keys/{set}][%d] deleteJsonWebKeySetNoContent ", 204)
}

func (o *DeleteJSONWebKeySetNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteJSONWebKeySetUnauthorized creates a DeleteJSONWebKeySetUnauthorized with default headers values
func NewDeleteJSONWebKeySetUnauthorized() *DeleteJSONWebKeySetUnauthorized {
	return &DeleteJSONWebKeySetUnauthorized{}
}

/*DeleteJSONWebKeySetUnauthorized handles this case with default header values.

genericError
*/
type DeleteJSONWebKeySetUnauthorized struct {
	Payload *models.GenericError
}

func (o *DeleteJSONWebKeySetUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /keys/{set}][%d] deleteJsonWebKeySetUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteJSONWebKeySetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJSONWebKeySetForbidden creates a DeleteJSONWebKeySetForbidden with default headers values
func NewDeleteJSONWebKeySetForbidden() *DeleteJSONWebKeySetForbidden {
	return &DeleteJSONWebKeySetForbidden{}
}

/*DeleteJSONWebKeySetForbidden handles this case with default header values.

genericError
*/
type DeleteJSONWebKeySetForbidden struct {
	Payload *models.GenericError
}

func (o *DeleteJSONWebKeySetForbidden) Error() string {
	return fmt.Sprintf("[DELETE /keys/{set}][%d] deleteJsonWebKeySetForbidden  %+v", 403, o.Payload)
}

func (o *DeleteJSONWebKeySetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteJSONWebKeySetInternalServerError creates a DeleteJSONWebKeySetInternalServerError with default headers values
func NewDeleteJSONWebKeySetInternalServerError() *DeleteJSONWebKeySetInternalServerError {
	return &DeleteJSONWebKeySetInternalServerError{}
}

/*DeleteJSONWebKeySetInternalServerError handles this case with default header values.

genericError
*/
type DeleteJSONWebKeySetInternalServerError struct {
	Payload *models.GenericError
}

func (o *DeleteJSONWebKeySetInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /keys/{set}][%d] deleteJsonWebKeySetInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteJSONWebKeySetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

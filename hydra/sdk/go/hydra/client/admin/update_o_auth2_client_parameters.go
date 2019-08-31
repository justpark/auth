// Code generated by go-swagger; DO NOT EDIT.

package admin

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

	models "github.com/justpark/auth/hydra/sdk/go/hydra/models"
)

// NewUpdateOAuth2ClientParams creates a new UpdateOAuth2ClientParams object
// with the default values initialized.
func NewUpdateOAuth2ClientParams() *UpdateOAuth2ClientParams {
	var ()
	return &UpdateOAuth2ClientParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOAuth2ClientParamsWithTimeout creates a new UpdateOAuth2ClientParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOAuth2ClientParamsWithTimeout(timeout time.Duration) *UpdateOAuth2ClientParams {
	var ()
	return &UpdateOAuth2ClientParams{

		timeout: timeout,
	}
}

// NewUpdateOAuth2ClientParamsWithContext creates a new UpdateOAuth2ClientParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOAuth2ClientParamsWithContext(ctx context.Context) *UpdateOAuth2ClientParams {
	var ()
	return &UpdateOAuth2ClientParams{

		Context: ctx,
	}
}

// NewUpdateOAuth2ClientParamsWithHTTPClient creates a new UpdateOAuth2ClientParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOAuth2ClientParamsWithHTTPClient(client *http.Client) *UpdateOAuth2ClientParams {
	var ()
	return &UpdateOAuth2ClientParams{
		HTTPClient: client,
	}
}

/*UpdateOAuth2ClientParams contains all the parameters to send to the API endpoint
for the update o auth2 client operation typically these are written to a http.Request
*/
type UpdateOAuth2ClientParams struct {

	/*Body*/
	Body *models.Client
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update o auth2 client params
func (o *UpdateOAuth2ClientParams) WithTimeout(timeout time.Duration) *UpdateOAuth2ClientParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update o auth2 client params
func (o *UpdateOAuth2ClientParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update o auth2 client params
func (o *UpdateOAuth2ClientParams) WithContext(ctx context.Context) *UpdateOAuth2ClientParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update o auth2 client params
func (o *UpdateOAuth2ClientParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update o auth2 client params
func (o *UpdateOAuth2ClientParams) WithHTTPClient(client *http.Client) *UpdateOAuth2ClientParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update o auth2 client params
func (o *UpdateOAuth2ClientParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update o auth2 client params
func (o *UpdateOAuth2ClientParams) WithBody(body *models.Client) *UpdateOAuth2ClientParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update o auth2 client params
func (o *UpdateOAuth2ClientParams) SetBody(body *models.Client) {
	o.Body = body
}

// WithID adds the id to the update o auth2 client params
func (o *UpdateOAuth2ClientParams) WithID(id string) *UpdateOAuth2ClientParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update o auth2 client params
func (o *UpdateOAuth2ClientParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOAuth2ClientParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

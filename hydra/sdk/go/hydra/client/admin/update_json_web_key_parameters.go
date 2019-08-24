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

	models "github.com/ory/hydra/sdk/go/hydra/models"
)

// NewUpdateJSONWebKeyParams creates a new UpdateJSONWebKeyParams object
// with the default values initialized.
func NewUpdateJSONWebKeyParams() *UpdateJSONWebKeyParams {
	var ()
	return &UpdateJSONWebKeyParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateJSONWebKeyParamsWithTimeout creates a new UpdateJSONWebKeyParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateJSONWebKeyParamsWithTimeout(timeout time.Duration) *UpdateJSONWebKeyParams {
	var ()
	return &UpdateJSONWebKeyParams{

		timeout: timeout,
	}
}

// NewUpdateJSONWebKeyParamsWithContext creates a new UpdateJSONWebKeyParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateJSONWebKeyParamsWithContext(ctx context.Context) *UpdateJSONWebKeyParams {
	var ()
	return &UpdateJSONWebKeyParams{

		Context: ctx,
	}
}

// NewUpdateJSONWebKeyParamsWithHTTPClient creates a new UpdateJSONWebKeyParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateJSONWebKeyParamsWithHTTPClient(client *http.Client) *UpdateJSONWebKeyParams {
	var ()
	return &UpdateJSONWebKeyParams{
		HTTPClient: client,
	}
}

/*UpdateJSONWebKeyParams contains all the parameters to send to the API endpoint
for the update Json web key operation typically these are written to a http.Request
*/
type UpdateJSONWebKeyParams struct {

	/*Body*/
	Body *models.SwaggerJSONWebKey
	/*Kid
	  The kid of the desired key

	*/
	Kid string
	/*Set
	  The set

	*/
	Set string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update Json web key params
func (o *UpdateJSONWebKeyParams) WithTimeout(timeout time.Duration) *UpdateJSONWebKeyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Json web key params
func (o *UpdateJSONWebKeyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Json web key params
func (o *UpdateJSONWebKeyParams) WithContext(ctx context.Context) *UpdateJSONWebKeyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Json web key params
func (o *UpdateJSONWebKeyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Json web key params
func (o *UpdateJSONWebKeyParams) WithHTTPClient(client *http.Client) *UpdateJSONWebKeyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Json web key params
func (o *UpdateJSONWebKeyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update Json web key params
func (o *UpdateJSONWebKeyParams) WithBody(body *models.SwaggerJSONWebKey) *UpdateJSONWebKeyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update Json web key params
func (o *UpdateJSONWebKeyParams) SetBody(body *models.SwaggerJSONWebKey) {
	o.Body = body
}

// WithKid adds the kid to the update Json web key params
func (o *UpdateJSONWebKeyParams) WithKid(kid string) *UpdateJSONWebKeyParams {
	o.SetKid(kid)
	return o
}

// SetKid adds the kid to the update Json web key params
func (o *UpdateJSONWebKeyParams) SetKid(kid string) {
	o.Kid = kid
}

// WithSet adds the set to the update Json web key params
func (o *UpdateJSONWebKeyParams) WithSet(set string) *UpdateJSONWebKeyParams {
	o.SetSet(set)
	return o
}

// SetSet adds the set to the update Json web key params
func (o *UpdateJSONWebKeyParams) SetSet(set string) {
	o.Set = set
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateJSONWebKeyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param kid
	if err := r.SetPathParam("kid", o.Kid); err != nil {
		return err
	}

	// path param set
	if err := r.SetPathParam("set", o.Set); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

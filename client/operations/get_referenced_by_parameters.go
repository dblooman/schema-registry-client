// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetReferencedByParams creates a new GetReferencedByParams object
// with the default values initialized.
func NewGetReferencedByParams() *GetReferencedByParams {
	var ()
	return &GetReferencedByParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetReferencedByParamsWithTimeout creates a new GetReferencedByParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetReferencedByParamsWithTimeout(timeout time.Duration) *GetReferencedByParams {
	var ()
	return &GetReferencedByParams{

		timeout: timeout,
	}
}

// NewGetReferencedByParamsWithContext creates a new GetReferencedByParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetReferencedByParamsWithContext(ctx context.Context) *GetReferencedByParams {
	var ()
	return &GetReferencedByParams{

		Context: ctx,
	}
}

// NewGetReferencedByParamsWithHTTPClient creates a new GetReferencedByParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetReferencedByParamsWithHTTPClient(client *http.Client) *GetReferencedByParams {
	var ()
	return &GetReferencedByParams{
		HTTPClient: client,
	}
}

/*GetReferencedByParams contains all the parameters to send to the API endpoint
for the get referenced by operation typically these are written to a http.Request
*/
type GetReferencedByParams struct {

	/*Subject
	  Name of the Subject

	*/
	Subject string
	/*Version
	  Version of the schema to be returned. Valid values for versionId are between [1,2^31-1] or the string "latest". "latest" returns the last registered schema under the specified subject. Note that there may be a new latest schema that gets registered right after this request is served.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get referenced by params
func (o *GetReferencedByParams) WithTimeout(timeout time.Duration) *GetReferencedByParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get referenced by params
func (o *GetReferencedByParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get referenced by params
func (o *GetReferencedByParams) WithContext(ctx context.Context) *GetReferencedByParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get referenced by params
func (o *GetReferencedByParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get referenced by params
func (o *GetReferencedByParams) WithHTTPClient(client *http.Client) *GetReferencedByParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get referenced by params
func (o *GetReferencedByParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubject adds the subject to the get referenced by params
func (o *GetReferencedByParams) WithSubject(subject string) *GetReferencedByParams {
	o.SetSubject(subject)
	return o
}

// SetSubject adds the subject to the get referenced by params
func (o *GetReferencedByParams) SetSubject(subject string) {
	o.Subject = subject
}

// WithVersion adds the version to the get referenced by params
func (o *GetReferencedByParams) WithVersion(version string) *GetReferencedByParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the get referenced by params
func (o *GetReferencedByParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *GetReferencedByParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param subject
	if err := r.SetPathParam("subject", o.Subject); err != nil {
		return err
	}

	// path param version
	if err := r.SetPathParam("version", o.Version); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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
	"github.com/go-openapi/swag"
)

// NewGetSubjectsParams creates a new GetSubjectsParams object
// with the default values initialized.
func NewGetSubjectsParams() *GetSubjectsParams {
	var ()
	return &GetSubjectsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetSubjectsParamsWithTimeout creates a new GetSubjectsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetSubjectsParamsWithTimeout(timeout time.Duration) *GetSubjectsParams {
	var ()
	return &GetSubjectsParams{

		timeout: timeout,
	}
}

// NewGetSubjectsParamsWithContext creates a new GetSubjectsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetSubjectsParamsWithContext(ctx context.Context) *GetSubjectsParams {
	var ()
	return &GetSubjectsParams{

		Context: ctx,
	}
}

// NewGetSubjectsParamsWithHTTPClient creates a new GetSubjectsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetSubjectsParamsWithHTTPClient(client *http.Client) *GetSubjectsParams {
	var ()
	return &GetSubjectsParams{
		HTTPClient: client,
	}
}

/*GetSubjectsParams contains all the parameters to send to the API endpoint
for the get subjects operation typically these are written to a http.Request
*/
type GetSubjectsParams struct {

	/*Deleted*/
	Deleted *bool
	/*ID
	  Globally unique identifier of the schema

	*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get subjects params
func (o *GetSubjectsParams) WithTimeout(timeout time.Duration) *GetSubjectsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get subjects params
func (o *GetSubjectsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get subjects params
func (o *GetSubjectsParams) WithContext(ctx context.Context) *GetSubjectsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get subjects params
func (o *GetSubjectsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get subjects params
func (o *GetSubjectsParams) WithHTTPClient(client *http.Client) *GetSubjectsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get subjects params
func (o *GetSubjectsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeleted adds the deleted to the get subjects params
func (o *GetSubjectsParams) WithDeleted(deleted *bool) *GetSubjectsParams {
	o.SetDeleted(deleted)
	return o
}

// SetDeleted adds the deleted to the get subjects params
func (o *GetSubjectsParams) SetDeleted(deleted *bool) {
	o.Deleted = deleted
}

// WithID adds the id to the get subjects params
func (o *GetSubjectsParams) WithID(id int32) *GetSubjectsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get subjects params
func (o *GetSubjectsParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetSubjectsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Deleted != nil {

		// query param deleted
		var qrDeleted bool
		if o.Deleted != nil {
			qrDeleted = *o.Deleted
		}
		qDeleted := swag.FormatBool(qrDeleted)
		if qDeleted != "" {
			if err := r.SetQueryParam("deleted", qDeleted); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

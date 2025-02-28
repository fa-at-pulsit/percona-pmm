// Code generated by go-swagger; DO NOT EDIT.

package server

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

// NewVersionParams creates a new VersionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVersionParams() *VersionParams {
	return &VersionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVersionParamsWithTimeout creates a new VersionParams object
// with the ability to set a timeout on a request.
func NewVersionParamsWithTimeout(timeout time.Duration) *VersionParams {
	return &VersionParams{
		timeout: timeout,
	}
}

// NewVersionParamsWithContext creates a new VersionParams object
// with the ability to set a context for a request.
func NewVersionParamsWithContext(ctx context.Context) *VersionParams {
	return &VersionParams{
		Context: ctx,
	}
}

// NewVersionParamsWithHTTPClient creates a new VersionParams object
// with the ability to set a custom HTTPClient for a request.
func NewVersionParamsWithHTTPClient(client *http.Client) *VersionParams {
	return &VersionParams{
		HTTPClient: client,
	}
}

/*
VersionParams contains all the parameters to send to the API endpoint

	for the version operation.

	Typically these are written to a http.Request.
*/
type VersionParams struct {
	/* Dummy.

	   Dummy parameter for internal testing. Do not use.
	*/
	Dummy *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersionParams) WithDefaults() *VersionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the version params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VersionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the version params
func (o *VersionParams) WithTimeout(timeout time.Duration) *VersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the version params
func (o *VersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the version params
func (o *VersionParams) WithContext(ctx context.Context) *VersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the version params
func (o *VersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the version params
func (o *VersionParams) WithHTTPClient(client *http.Client) *VersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the version params
func (o *VersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDummy adds the dummy to the version params
func (o *VersionParams) WithDummy(dummy *string) *VersionParams {
	o.SetDummy(dummy)
	return o
}

// SetDummy adds the dummy to the version params
func (o *VersionParams) SetDummy(dummy *string) {
	o.Dummy = dummy
}

// WriteToRequest writes these params to a swagger request
func (o *VersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Dummy != nil {

		// query param dummy
		var qrDummy string

		if o.Dummy != nil {
			qrDummy = *o.Dummy
		}
		qDummy := qrDummy
		if qDummy != "" {
			if err := r.SetQueryParam("dummy", qDummy); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

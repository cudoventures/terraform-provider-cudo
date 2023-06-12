// Code generated by go-swagger; DO NOT EDIT.

package virtual_machines

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

// NewCountVMsParams creates a new CountVMsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCountVMsParams() *CountVMsParams {
	return &CountVMsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCountVMsParamsWithTimeout creates a new CountVMsParams object
// with the ability to set a timeout on a request.
func NewCountVMsParamsWithTimeout(timeout time.Duration) *CountVMsParams {
	return &CountVMsParams{
		timeout: timeout,
	}
}

// NewCountVMsParamsWithContext creates a new CountVMsParams object
// with the ability to set a context for a request.
func NewCountVMsParamsWithContext(ctx context.Context) *CountVMsParams {
	return &CountVMsParams{
		Context: ctx,
	}
}

// NewCountVMsParamsWithHTTPClient creates a new CountVMsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCountVMsParamsWithHTTPClient(client *http.Client) *CountVMsParams {
	return &CountVMsParams{
		HTTPClient: client,
	}
}

/*
CountVMsParams contains all the parameters to send to the API endpoint

	for the count v ms operation.

	Typically these are written to a http.Request.
*/
type CountVMsParams struct {

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the count v ms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CountVMsParams) WithDefaults() *CountVMsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the count v ms params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CountVMsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the count v ms params
func (o *CountVMsParams) WithTimeout(timeout time.Duration) *CountVMsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the count v ms params
func (o *CountVMsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the count v ms params
func (o *CountVMsParams) WithContext(ctx context.Context) *CountVMsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the count v ms params
func (o *CountVMsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the count v ms params
func (o *CountVMsParams) WithHTTPClient(client *http.Client) *CountVMsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the count v ms params
func (o *CountVMsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the count v ms params
func (o *CountVMsParams) WithProjectID(projectID string) *CountVMsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the count v ms params
func (o *CountVMsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *CountVMsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewCreatePrivateVMImageParams creates a new CreatePrivateVMImageParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreatePrivateVMImageParams() *CreatePrivateVMImageParams {
	return &CreatePrivateVMImageParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreatePrivateVMImageParamsWithTimeout creates a new CreatePrivateVMImageParams object
// with the ability to set a timeout on a request.
func NewCreatePrivateVMImageParamsWithTimeout(timeout time.Duration) *CreatePrivateVMImageParams {
	return &CreatePrivateVMImageParams{
		timeout: timeout,
	}
}

// NewCreatePrivateVMImageParamsWithContext creates a new CreatePrivateVMImageParams object
// with the ability to set a context for a request.
func NewCreatePrivateVMImageParamsWithContext(ctx context.Context) *CreatePrivateVMImageParams {
	return &CreatePrivateVMImageParams{
		Context: ctx,
	}
}

// NewCreatePrivateVMImageParamsWithHTTPClient creates a new CreatePrivateVMImageParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreatePrivateVMImageParamsWithHTTPClient(client *http.Client) *CreatePrivateVMImageParams {
	return &CreatePrivateVMImageParams{
		HTTPClient: client,
	}
}

/*
CreatePrivateVMImageParams contains all the parameters to send to the API endpoint

	for the create private VM image operation.

	Typically these are written to a http.Request.
*/
type CreatePrivateVMImageParams struct {

	// Description.
	Description *string

	// ID.
	ID string

	// ProjectID.
	ProjectID string

	// VMID.
	VMID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create private VM image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePrivateVMImageParams) WithDefaults() *CreatePrivateVMImageParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create private VM image params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreatePrivateVMImageParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create private VM image params
func (o *CreatePrivateVMImageParams) WithTimeout(timeout time.Duration) *CreatePrivateVMImageParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create private VM image params
func (o *CreatePrivateVMImageParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create private VM image params
func (o *CreatePrivateVMImageParams) WithContext(ctx context.Context) *CreatePrivateVMImageParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create private VM image params
func (o *CreatePrivateVMImageParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create private VM image params
func (o *CreatePrivateVMImageParams) WithHTTPClient(client *http.Client) *CreatePrivateVMImageParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create private VM image params
func (o *CreatePrivateVMImageParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDescription adds the description to the create private VM image params
func (o *CreatePrivateVMImageParams) WithDescription(description *string) *CreatePrivateVMImageParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the create private VM image params
func (o *CreatePrivateVMImageParams) SetDescription(description *string) {
	o.Description = description
}

// WithID adds the id to the create private VM image params
func (o *CreatePrivateVMImageParams) WithID(id string) *CreatePrivateVMImageParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the create private VM image params
func (o *CreatePrivateVMImageParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the create private VM image params
func (o *CreatePrivateVMImageParams) WithProjectID(projectID string) *CreatePrivateVMImageParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the create private VM image params
func (o *CreatePrivateVMImageParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithVMID adds the vMID to the create private VM image params
func (o *CreatePrivateVMImageParams) WithVMID(vMID string) *CreatePrivateVMImageParams {
	o.SetVMID(vMID)
	return o
}

// SetVMID adds the vmId to the create private VM image params
func (o *CreatePrivateVMImageParams) SetVMID(vMID string) {
	o.VMID = vMID
}

// WriteToRequest writes these params to a swagger request
func (o *CreatePrivateVMImageParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	// query param id
	qrID := o.ID
	qID := qrID
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	// query param vmId
	qrVMID := o.VMID
	qVMID := qrVMID
	if qVMID != "" {

		if err := r.SetQueryParam("vmId", qVMID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
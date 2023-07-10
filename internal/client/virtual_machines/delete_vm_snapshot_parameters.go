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

// NewDeleteVMSnapshotParams creates a new DeleteVMSnapshotParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteVMSnapshotParams() *DeleteVMSnapshotParams {
	return &DeleteVMSnapshotParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteVMSnapshotParamsWithTimeout creates a new DeleteVMSnapshotParams object
// with the ability to set a timeout on a request.
func NewDeleteVMSnapshotParamsWithTimeout(timeout time.Duration) *DeleteVMSnapshotParams {
	return &DeleteVMSnapshotParams{
		timeout: timeout,
	}
}

// NewDeleteVMSnapshotParamsWithContext creates a new DeleteVMSnapshotParams object
// with the ability to set a context for a request.
func NewDeleteVMSnapshotParamsWithContext(ctx context.Context) *DeleteVMSnapshotParams {
	return &DeleteVMSnapshotParams{
		Context: ctx,
	}
}

// NewDeleteVMSnapshotParamsWithHTTPClient creates a new DeleteVMSnapshotParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteVMSnapshotParamsWithHTTPClient(client *http.Client) *DeleteVMSnapshotParams {
	return &DeleteVMSnapshotParams{
		HTTPClient: client,
	}
}

/*
DeleteVMSnapshotParams contains all the parameters to send to the API endpoint

	for the delete VM snapshot operation.

	Typically these are written to a http.Request.
*/
type DeleteVMSnapshotParams struct {

	// ID.
	ID string

	// ProjectID.
	ProjectID string

	// SnapshotID.
	SnapshotID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete VM snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVMSnapshotParams) WithDefaults() *DeleteVMSnapshotParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete VM snapshot params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteVMSnapshotParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete VM snapshot params
func (o *DeleteVMSnapshotParams) WithTimeout(timeout time.Duration) *DeleteVMSnapshotParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete VM snapshot params
func (o *DeleteVMSnapshotParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete VM snapshot params
func (o *DeleteVMSnapshotParams) WithContext(ctx context.Context) *DeleteVMSnapshotParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete VM snapshot params
func (o *DeleteVMSnapshotParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete VM snapshot params
func (o *DeleteVMSnapshotParams) WithHTTPClient(client *http.Client) *DeleteVMSnapshotParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete VM snapshot params
func (o *DeleteVMSnapshotParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete VM snapshot params
func (o *DeleteVMSnapshotParams) WithID(id string) *DeleteVMSnapshotParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete VM snapshot params
func (o *DeleteVMSnapshotParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the delete VM snapshot params
func (o *DeleteVMSnapshotParams) WithProjectID(projectID string) *DeleteVMSnapshotParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete VM snapshot params
func (o *DeleteVMSnapshotParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithSnapshotID adds the snapshotID to the delete VM snapshot params
func (o *DeleteVMSnapshotParams) WithSnapshotID(snapshotID string) *DeleteVMSnapshotParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the delete VM snapshot params
func (o *DeleteVMSnapshotParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteVMSnapshotParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	// query param snapshotId
	qrSnapshotID := o.SnapshotID
	qSnapshotID := qrSnapshotID
	if qSnapshotID != "" {

		if err := r.SetQueryParam("snapshotId", qSnapshotID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

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

// NewRevertVMParams creates a new RevertVMParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRevertVMParams() *RevertVMParams {
	return &RevertVMParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRevertVMParamsWithTimeout creates a new RevertVMParams object
// with the ability to set a timeout on a request.
func NewRevertVMParamsWithTimeout(timeout time.Duration) *RevertVMParams {
	return &RevertVMParams{
		timeout: timeout,
	}
}

// NewRevertVMParamsWithContext creates a new RevertVMParams object
// with the ability to set a context for a request.
func NewRevertVMParamsWithContext(ctx context.Context) *RevertVMParams {
	return &RevertVMParams{
		Context: ctx,
	}
}

// NewRevertVMParamsWithHTTPClient creates a new RevertVMParams object
// with the ability to set a custom HTTPClient for a request.
func NewRevertVMParamsWithHTTPClient(client *http.Client) *RevertVMParams {
	return &RevertVMParams{
		HTTPClient: client,
	}
}

/*
RevertVMParams contains all the parameters to send to the API endpoint

	for the revert VM operation.

	Typically these are written to a http.Request.
*/
type RevertVMParams struct {

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

// WithDefaults hydrates default values in the revert VM params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevertVMParams) WithDefaults() *RevertVMParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the revert VM params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RevertVMParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the revert VM params
func (o *RevertVMParams) WithTimeout(timeout time.Duration) *RevertVMParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the revert VM params
func (o *RevertVMParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the revert VM params
func (o *RevertVMParams) WithContext(ctx context.Context) *RevertVMParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the revert VM params
func (o *RevertVMParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the revert VM params
func (o *RevertVMParams) WithHTTPClient(client *http.Client) *RevertVMParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the revert VM params
func (o *RevertVMParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the revert VM params
func (o *RevertVMParams) WithID(id string) *RevertVMParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the revert VM params
func (o *RevertVMParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the revert VM params
func (o *RevertVMParams) WithProjectID(projectID string) *RevertVMParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the revert VM params
func (o *RevertVMParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithSnapshotID adds the snapshotID to the revert VM params
func (o *RevertVMParams) WithSnapshotID(snapshotID string) *RevertVMParams {
	o.SetSnapshotID(snapshotID)
	return o
}

// SetSnapshotID adds the snapshotId to the revert VM params
func (o *RevertVMParams) SetSnapshotID(snapshotID string) {
	o.SnapshotID = snapshotID
}

// WriteToRequest writes these params to a swagger request
func (o *RevertVMParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

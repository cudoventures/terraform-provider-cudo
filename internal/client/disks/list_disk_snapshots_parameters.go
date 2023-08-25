// Code generated by go-swagger; DO NOT EDIT.

package disks

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

// NewListDiskSnapshotsParams creates a new ListDiskSnapshotsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListDiskSnapshotsParams() *ListDiskSnapshotsParams {
	return &ListDiskSnapshotsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListDiskSnapshotsParamsWithTimeout creates a new ListDiskSnapshotsParams object
// with the ability to set a timeout on a request.
func NewListDiskSnapshotsParamsWithTimeout(timeout time.Duration) *ListDiskSnapshotsParams {
	return &ListDiskSnapshotsParams{
		timeout: timeout,
	}
}

// NewListDiskSnapshotsParamsWithContext creates a new ListDiskSnapshotsParams object
// with the ability to set a context for a request.
func NewListDiskSnapshotsParamsWithContext(ctx context.Context) *ListDiskSnapshotsParams {
	return &ListDiskSnapshotsParams{
		Context: ctx,
	}
}

// NewListDiskSnapshotsParamsWithHTTPClient creates a new ListDiskSnapshotsParams object
// with the ability to set a custom HTTPClient for a request.
func NewListDiskSnapshotsParamsWithHTTPClient(client *http.Client) *ListDiskSnapshotsParams {
	return &ListDiskSnapshotsParams{
		HTTPClient: client,
	}
}

/*
ListDiskSnapshotsParams contains all the parameters to send to the API endpoint

	for the list disk snapshots operation.

	Typically these are written to a http.Request.
*/
type ListDiskSnapshotsParams struct {

	// ID.
	ID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list disk snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDiskSnapshotsParams) WithDefaults() *ListDiskSnapshotsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list disk snapshots params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListDiskSnapshotsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list disk snapshots params
func (o *ListDiskSnapshotsParams) WithTimeout(timeout time.Duration) *ListDiskSnapshotsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list disk snapshots params
func (o *ListDiskSnapshotsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list disk snapshots params
func (o *ListDiskSnapshotsParams) WithContext(ctx context.Context) *ListDiskSnapshotsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list disk snapshots params
func (o *ListDiskSnapshotsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list disk snapshots params
func (o *ListDiskSnapshotsParams) WithHTTPClient(client *http.Client) *ListDiskSnapshotsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list disk snapshots params
func (o *ListDiskSnapshotsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the list disk snapshots params
func (o *ListDiskSnapshotsParams) WithID(id string) *ListDiskSnapshotsParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the list disk snapshots params
func (o *ListDiskSnapshotsParams) SetID(id string) {
	o.ID = id
}

// WithProjectID adds the projectID to the list disk snapshots params
func (o *ListDiskSnapshotsParams) WithProjectID(projectID string) *ListDiskSnapshotsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list disk snapshots params
func (o *ListDiskSnapshotsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListDiskSnapshotsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

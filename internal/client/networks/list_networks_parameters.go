// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewListNetworksParams creates a new ListNetworksParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListNetworksParams() *ListNetworksParams {
	return &ListNetworksParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListNetworksParamsWithTimeout creates a new ListNetworksParams object
// with the ability to set a timeout on a request.
func NewListNetworksParamsWithTimeout(timeout time.Duration) *ListNetworksParams {
	return &ListNetworksParams{
		timeout: timeout,
	}
}

// NewListNetworksParamsWithContext creates a new ListNetworksParams object
// with the ability to set a context for a request.
func NewListNetworksParamsWithContext(ctx context.Context) *ListNetworksParams {
	return &ListNetworksParams{
		Context: ctx,
	}
}

// NewListNetworksParamsWithHTTPClient creates a new ListNetworksParams object
// with the ability to set a custom HTTPClient for a request.
func NewListNetworksParamsWithHTTPClient(client *http.Client) *ListNetworksParams {
	return &ListNetworksParams{
		HTTPClient: client,
	}
}

/*
ListNetworksParams contains all the parameters to send to the API endpoint

	for the list networks operation.

	Typically these are written to a http.Request.
*/
type ListNetworksParams struct {

	// PageNumber.
	//
	// Format: int32
	PageNumber *int32

	// PageSize.
	//
	// Format: int32
	PageSize *int32

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNetworksParams) WithDefaults() *ListNetworksParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list networks params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListNetworksParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list networks params
func (o *ListNetworksParams) WithTimeout(timeout time.Duration) *ListNetworksParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list networks params
func (o *ListNetworksParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list networks params
func (o *ListNetworksParams) WithContext(ctx context.Context) *ListNetworksParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list networks params
func (o *ListNetworksParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list networks params
func (o *ListNetworksParams) WithHTTPClient(client *http.Client) *ListNetworksParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list networks params
func (o *ListNetworksParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageNumber adds the pageNumber to the list networks params
func (o *ListNetworksParams) WithPageNumber(pageNumber *int32) *ListNetworksParams {
	o.SetPageNumber(pageNumber)
	return o
}

// SetPageNumber adds the pageNumber to the list networks params
func (o *ListNetworksParams) SetPageNumber(pageNumber *int32) {
	o.PageNumber = pageNumber
}

// WithPageSize adds the pageSize to the list networks params
func (o *ListNetworksParams) WithPageSize(pageSize *int32) *ListNetworksParams {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the list networks params
func (o *ListNetworksParams) SetPageSize(pageSize *int32) {
	o.PageSize = pageSize
}

// WithProjectID adds the projectID to the list networks params
func (o *ListNetworksParams) WithProjectID(projectID string) *ListNetworksParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the list networks params
func (o *ListNetworksParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *ListNetworksParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PageNumber != nil {

		// query param pageNumber
		var qrPageNumber int32

		if o.PageNumber != nil {
			qrPageNumber = *o.PageNumber
		}
		qPageNumber := swag.FormatInt32(qrPageNumber)
		if qPageNumber != "" {

			if err := r.SetQueryParam("pageNumber", qPageNumber); err != nil {
				return err
			}
		}
	}

	if o.PageSize != nil {

		// query param pageSize
		var qrPageSize int32

		if o.PageSize != nil {
			qrPageSize = *o.PageSize
		}
		qPageSize := swag.FormatInt32(qrPageSize)
		if qPageSize != "" {

			if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
				return err
			}
		}
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
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
)

// NewUpdateSecurityGroupParams creates a new UpdateSecurityGroupParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateSecurityGroupParams() *UpdateSecurityGroupParams {
	return &UpdateSecurityGroupParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateSecurityGroupParamsWithTimeout creates a new UpdateSecurityGroupParams object
// with the ability to set a timeout on a request.
func NewUpdateSecurityGroupParamsWithTimeout(timeout time.Duration) *UpdateSecurityGroupParams {
	return &UpdateSecurityGroupParams{
		timeout: timeout,
	}
}

// NewUpdateSecurityGroupParamsWithContext creates a new UpdateSecurityGroupParams object
// with the ability to set a context for a request.
func NewUpdateSecurityGroupParamsWithContext(ctx context.Context) *UpdateSecurityGroupParams {
	return &UpdateSecurityGroupParams{
		Context: ctx,
	}
}

// NewUpdateSecurityGroupParamsWithHTTPClient creates a new UpdateSecurityGroupParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateSecurityGroupParamsWithHTTPClient(client *http.Client) *UpdateSecurityGroupParams {
	return &UpdateSecurityGroupParams{
		HTTPClient: client,
	}
}

/*
UpdateSecurityGroupParams contains all the parameters to send to the API endpoint

	for the update security group operation.

	Typically these are written to a http.Request.
*/
type UpdateSecurityGroupParams struct {

	// SecurityGroup.
	SecurityGroup UpdateSecurityGroupBody

	// SecurityGroupID.
	SecurityGroupID string

	// SecurityGroupProjectID.
	SecurityGroupProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update security group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSecurityGroupParams) WithDefaults() *UpdateSecurityGroupParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update security group params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateSecurityGroupParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update security group params
func (o *UpdateSecurityGroupParams) WithTimeout(timeout time.Duration) *UpdateSecurityGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update security group params
func (o *UpdateSecurityGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update security group params
func (o *UpdateSecurityGroupParams) WithContext(ctx context.Context) *UpdateSecurityGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update security group params
func (o *UpdateSecurityGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update security group params
func (o *UpdateSecurityGroupParams) WithHTTPClient(client *http.Client) *UpdateSecurityGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update security group params
func (o *UpdateSecurityGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSecurityGroup adds the securityGroup to the update security group params
func (o *UpdateSecurityGroupParams) WithSecurityGroup(securityGroup UpdateSecurityGroupBody) *UpdateSecurityGroupParams {
	o.SetSecurityGroup(securityGroup)
	return o
}

// SetSecurityGroup adds the securityGroup to the update security group params
func (o *UpdateSecurityGroupParams) SetSecurityGroup(securityGroup UpdateSecurityGroupBody) {
	o.SecurityGroup = securityGroup
}

// WithSecurityGroupID adds the securityGroupID to the update security group params
func (o *UpdateSecurityGroupParams) WithSecurityGroupID(securityGroupID string) *UpdateSecurityGroupParams {
	o.SetSecurityGroupID(securityGroupID)
	return o
}

// SetSecurityGroupID adds the securityGroupId to the update security group params
func (o *UpdateSecurityGroupParams) SetSecurityGroupID(securityGroupID string) {
	o.SecurityGroupID = securityGroupID
}

// WithSecurityGroupProjectID adds the securityGroupProjectID to the update security group params
func (o *UpdateSecurityGroupParams) WithSecurityGroupProjectID(securityGroupProjectID string) *UpdateSecurityGroupParams {
	o.SetSecurityGroupProjectID(securityGroupProjectID)
	return o
}

// SetSecurityGroupProjectID adds the securityGroupProjectId to the update security group params
func (o *UpdateSecurityGroupParams) SetSecurityGroupProjectID(securityGroupProjectID string) {
	o.SecurityGroupProjectID = securityGroupProjectID
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateSecurityGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.SecurityGroup); err != nil {
		return err
	}

	// path param securityGroup.id
	if err := r.SetPathParam("securityGroup.id", o.SecurityGroupID); err != nil {
		return err
	}

	// path param securityGroup.projectId
	if err := r.SetPathParam("securityGroup.projectId", o.SecurityGroupProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
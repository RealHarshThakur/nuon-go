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

	"github.com/nuonco/nuon-go/models"
)

// NewCreateKubernetesManifestComponentConfigParams creates a new CreateKubernetesManifestComponentConfigParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateKubernetesManifestComponentConfigParams() *CreateKubernetesManifestComponentConfigParams {
	return &CreateKubernetesManifestComponentConfigParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateKubernetesManifestComponentConfigParamsWithTimeout creates a new CreateKubernetesManifestComponentConfigParams object
// with the ability to set a timeout on a request.
func NewCreateKubernetesManifestComponentConfigParamsWithTimeout(timeout time.Duration) *CreateKubernetesManifestComponentConfigParams {
	return &CreateKubernetesManifestComponentConfigParams{
		timeout: timeout,
	}
}

// NewCreateKubernetesManifestComponentConfigParamsWithContext creates a new CreateKubernetesManifestComponentConfigParams object
// with the ability to set a context for a request.
func NewCreateKubernetesManifestComponentConfigParamsWithContext(ctx context.Context) *CreateKubernetesManifestComponentConfigParams {
	return &CreateKubernetesManifestComponentConfigParams{
		Context: ctx,
	}
}

// NewCreateKubernetesManifestComponentConfigParamsWithHTTPClient creates a new CreateKubernetesManifestComponentConfigParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateKubernetesManifestComponentConfigParamsWithHTTPClient(client *http.Client) *CreateKubernetesManifestComponentConfigParams {
	return &CreateKubernetesManifestComponentConfigParams{
		HTTPClient: client,
	}
}

/*
CreateKubernetesManifestComponentConfigParams contains all the parameters to send to the API endpoint

	for the create kubernetes manifest component config operation.

	Typically these are written to a http.Request.
*/
type CreateKubernetesManifestComponentConfigParams struct {

	/* ComponentID.

	   component ID
	*/
	ComponentID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateKubernetesManifestComponentConfigRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create kubernetes manifest component config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateKubernetesManifestComponentConfigParams) WithDefaults() *CreateKubernetesManifestComponentConfigParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create kubernetes manifest component config params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateKubernetesManifestComponentConfigParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create kubernetes manifest component config params
func (o *CreateKubernetesManifestComponentConfigParams) WithTimeout(timeout time.Duration) *CreateKubernetesManifestComponentConfigParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create kubernetes manifest component config params
func (o *CreateKubernetesManifestComponentConfigParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create kubernetes manifest component config params
func (o *CreateKubernetesManifestComponentConfigParams) WithContext(ctx context.Context) *CreateKubernetesManifestComponentConfigParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create kubernetes manifest component config params
func (o *CreateKubernetesManifestComponentConfigParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create kubernetes manifest component config params
func (o *CreateKubernetesManifestComponentConfigParams) WithHTTPClient(client *http.Client) *CreateKubernetesManifestComponentConfigParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create kubernetes manifest component config params
func (o *CreateKubernetesManifestComponentConfigParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComponentID adds the componentID to the create kubernetes manifest component config params
func (o *CreateKubernetesManifestComponentConfigParams) WithComponentID(componentID string) *CreateKubernetesManifestComponentConfigParams {
	o.SetComponentID(componentID)
	return o
}

// SetComponentID adds the componentId to the create kubernetes manifest component config params
func (o *CreateKubernetesManifestComponentConfigParams) SetComponentID(componentID string) {
	o.ComponentID = componentID
}

// WithReq adds the req to the create kubernetes manifest component config params
func (o *CreateKubernetesManifestComponentConfigParams) WithReq(req *models.ServiceCreateKubernetesManifestComponentConfigRequest) *CreateKubernetesManifestComponentConfigParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create kubernetes manifest component config params
func (o *CreateKubernetesManifestComponentConfigParams) SetReq(req *models.ServiceCreateKubernetesManifestComponentConfigRequest) {
	o.Req = req
}

// WriteToRequest writes these params to a swagger request
func (o *CreateKubernetesManifestComponentConfigParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param component_id
	if err := r.SetPathParam("component_id", o.ComponentID); err != nil {
		return err
	}
	if o.Req != nil {
		if err := r.SetBodyParam(o.Req); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

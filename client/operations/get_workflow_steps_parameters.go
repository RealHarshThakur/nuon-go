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
)

// NewGetWorkflowStepsParams creates a new GetWorkflowStepsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetWorkflowStepsParams() *GetWorkflowStepsParams {
	return &GetWorkflowStepsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetWorkflowStepsParamsWithTimeout creates a new GetWorkflowStepsParams object
// with the ability to set a timeout on a request.
func NewGetWorkflowStepsParamsWithTimeout(timeout time.Duration) *GetWorkflowStepsParams {
	return &GetWorkflowStepsParams{
		timeout: timeout,
	}
}

// NewGetWorkflowStepsParamsWithContext creates a new GetWorkflowStepsParams object
// with the ability to set a context for a request.
func NewGetWorkflowStepsParamsWithContext(ctx context.Context) *GetWorkflowStepsParams {
	return &GetWorkflowStepsParams{
		Context: ctx,
	}
}

// NewGetWorkflowStepsParamsWithHTTPClient creates a new GetWorkflowStepsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetWorkflowStepsParamsWithHTTPClient(client *http.Client) *GetWorkflowStepsParams {
	return &GetWorkflowStepsParams{
		HTTPClient: client,
	}
}

/*
GetWorkflowStepsParams contains all the parameters to send to the API endpoint

	for the get workflow steps operation.

	Typically these are written to a http.Request.
*/
type GetWorkflowStepsParams struct {

	/* WorkflowID.

	   workflow ID
	*/
	WorkflowID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get workflow steps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowStepsParams) WithDefaults() *GetWorkflowStepsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get workflow steps params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetWorkflowStepsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get workflow steps params
func (o *GetWorkflowStepsParams) WithTimeout(timeout time.Duration) *GetWorkflowStepsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get workflow steps params
func (o *GetWorkflowStepsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get workflow steps params
func (o *GetWorkflowStepsParams) WithContext(ctx context.Context) *GetWorkflowStepsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get workflow steps params
func (o *GetWorkflowStepsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get workflow steps params
func (o *GetWorkflowStepsParams) WithHTTPClient(client *http.Client) *GetWorkflowStepsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get workflow steps params
func (o *GetWorkflowStepsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithWorkflowID adds the workflowID to the get workflow steps params
func (o *GetWorkflowStepsParams) WithWorkflowID(workflowID string) *GetWorkflowStepsParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the get workflow steps params
func (o *GetWorkflowStepsParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowStepsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param workflow_id
	if err := r.SetPathParam("workflow_id", o.WorkflowID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

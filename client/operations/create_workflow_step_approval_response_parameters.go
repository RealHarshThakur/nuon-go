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

// NewCreateWorkflowStepApprovalResponseParams creates a new CreateWorkflowStepApprovalResponseParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateWorkflowStepApprovalResponseParams() *CreateWorkflowStepApprovalResponseParams {
	return &CreateWorkflowStepApprovalResponseParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateWorkflowStepApprovalResponseParamsWithTimeout creates a new CreateWorkflowStepApprovalResponseParams object
// with the ability to set a timeout on a request.
func NewCreateWorkflowStepApprovalResponseParamsWithTimeout(timeout time.Duration) *CreateWorkflowStepApprovalResponseParams {
	return &CreateWorkflowStepApprovalResponseParams{
		timeout: timeout,
	}
}

// NewCreateWorkflowStepApprovalResponseParamsWithContext creates a new CreateWorkflowStepApprovalResponseParams object
// with the ability to set a context for a request.
func NewCreateWorkflowStepApprovalResponseParamsWithContext(ctx context.Context) *CreateWorkflowStepApprovalResponseParams {
	return &CreateWorkflowStepApprovalResponseParams{
		Context: ctx,
	}
}

// NewCreateWorkflowStepApprovalResponseParamsWithHTTPClient creates a new CreateWorkflowStepApprovalResponseParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateWorkflowStepApprovalResponseParamsWithHTTPClient(client *http.Client) *CreateWorkflowStepApprovalResponseParams {
	return &CreateWorkflowStepApprovalResponseParams{
		HTTPClient: client,
	}
}

/*
CreateWorkflowStepApprovalResponseParams contains all the parameters to send to the API endpoint

	for the create workflow step approval response operation.

	Typically these are written to a http.Request.
*/
type CreateWorkflowStepApprovalResponseParams struct {

	/* ApprovalID.

	   approval id
	*/
	ApprovalID string

	/* Req.

	   Input
	*/
	Req *models.ServiceCreateWorkflowStepApprovalResponseRequest

	/* WorkflowID.

	   workflow id
	*/
	WorkflowID string

	/* WorkflowStepID.

	   step id
	*/
	WorkflowStepID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create workflow step approval response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateWorkflowStepApprovalResponseParams) WithDefaults() *CreateWorkflowStepApprovalResponseParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create workflow step approval response params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateWorkflowStepApprovalResponseParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) WithTimeout(timeout time.Duration) *CreateWorkflowStepApprovalResponseParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) WithContext(ctx context.Context) *CreateWorkflowStepApprovalResponseParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) WithHTTPClient(client *http.Client) *CreateWorkflowStepApprovalResponseParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApprovalID adds the approvalID to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) WithApprovalID(approvalID string) *CreateWorkflowStepApprovalResponseParams {
	o.SetApprovalID(approvalID)
	return o
}

// SetApprovalID adds the approvalId to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) SetApprovalID(approvalID string) {
	o.ApprovalID = approvalID
}

// WithReq adds the req to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) WithReq(req *models.ServiceCreateWorkflowStepApprovalResponseRequest) *CreateWorkflowStepApprovalResponseParams {
	o.SetReq(req)
	return o
}

// SetReq adds the req to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) SetReq(req *models.ServiceCreateWorkflowStepApprovalResponseRequest) {
	o.Req = req
}

// WithWorkflowID adds the workflowID to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) WithWorkflowID(workflowID string) *CreateWorkflowStepApprovalResponseParams {
	o.SetWorkflowID(workflowID)
	return o
}

// SetWorkflowID adds the workflowId to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) SetWorkflowID(workflowID string) {
	o.WorkflowID = workflowID
}

// WithWorkflowStepID adds the workflowStepID to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) WithWorkflowStepID(workflowStepID string) *CreateWorkflowStepApprovalResponseParams {
	o.SetWorkflowStepID(workflowStepID)
	return o
}

// SetWorkflowStepID adds the workflowStepId to the create workflow step approval response params
func (o *CreateWorkflowStepApprovalResponseParams) SetWorkflowStepID(workflowStepID string) {
	o.WorkflowStepID = workflowStepID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateWorkflowStepApprovalResponseParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param approval_id
	if err := r.SetPathParam("approval_id", o.ApprovalID); err != nil {
		return err
	}
	if o.Req != nil {
		if err := r.SetBodyParam(o.Req); err != nil {
			return err
		}
	}

	// path param workflow_id
	if err := r.SetPathParam("workflow_id", o.WorkflowID); err != nil {
		return err
	}

	// path param workflow_step_id
	if err := r.SetPathParam("workflow_step_id", o.WorkflowStepID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

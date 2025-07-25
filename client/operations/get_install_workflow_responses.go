// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nuonco/nuon-go/models"
)

// GetInstallWorkflowReader is a Reader for the GetInstallWorkflow structure.
type GetInstallWorkflowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstallWorkflowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetInstallWorkflowOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetInstallWorkflowBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetInstallWorkflowUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetInstallWorkflowForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetInstallWorkflowNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetInstallWorkflowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /v1/install-workflows/{install_workflow_id}] GetInstallWorkflow", response, response.Code())
	}
}

// NewGetInstallWorkflowOK creates a GetInstallWorkflowOK with default headers values
func NewGetInstallWorkflowOK() *GetInstallWorkflowOK {
	return &GetInstallWorkflowOK{}
}

/*
GetInstallWorkflowOK describes a response with status code 200, with default header values.

OK
*/
type GetInstallWorkflowOK struct {
	Payload *models.AppWorkflow
}

// IsSuccess returns true when this get install workflow o k response has a 2xx status code
func (o *GetInstallWorkflowOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get install workflow o k response has a 3xx status code
func (o *GetInstallWorkflowOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install workflow o k response has a 4xx status code
func (o *GetInstallWorkflowOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install workflow o k response has a 5xx status code
func (o *GetInstallWorkflowOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get install workflow o k response a status code equal to that given
func (o *GetInstallWorkflowOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get install workflow o k response
func (o *GetInstallWorkflowOK) Code() int {
	return 200
}

func (o *GetInstallWorkflowOK) Error() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}][%d] getInstallWorkflowOK  %+v", 200, o.Payload)
}

func (o *GetInstallWorkflowOK) String() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}][%d] getInstallWorkflowOK  %+v", 200, o.Payload)
}

func (o *GetInstallWorkflowOK) GetPayload() *models.AppWorkflow {
	return o.Payload
}

func (o *GetInstallWorkflowOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AppWorkflow)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallWorkflowBadRequest creates a GetInstallWorkflowBadRequest with default headers values
func NewGetInstallWorkflowBadRequest() *GetInstallWorkflowBadRequest {
	return &GetInstallWorkflowBadRequest{}
}

/*
GetInstallWorkflowBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetInstallWorkflowBadRequest struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install workflow bad request response has a 2xx status code
func (o *GetInstallWorkflowBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install workflow bad request response has a 3xx status code
func (o *GetInstallWorkflowBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install workflow bad request response has a 4xx status code
func (o *GetInstallWorkflowBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install workflow bad request response has a 5xx status code
func (o *GetInstallWorkflowBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get install workflow bad request response a status code equal to that given
func (o *GetInstallWorkflowBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get install workflow bad request response
func (o *GetInstallWorkflowBadRequest) Code() int {
	return 400
}

func (o *GetInstallWorkflowBadRequest) Error() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}][%d] getInstallWorkflowBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallWorkflowBadRequest) String() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}][%d] getInstallWorkflowBadRequest  %+v", 400, o.Payload)
}

func (o *GetInstallWorkflowBadRequest) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallWorkflowBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallWorkflowUnauthorized creates a GetInstallWorkflowUnauthorized with default headers values
func NewGetInstallWorkflowUnauthorized() *GetInstallWorkflowUnauthorized {
	return &GetInstallWorkflowUnauthorized{}
}

/*
GetInstallWorkflowUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetInstallWorkflowUnauthorized struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install workflow unauthorized response has a 2xx status code
func (o *GetInstallWorkflowUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install workflow unauthorized response has a 3xx status code
func (o *GetInstallWorkflowUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install workflow unauthorized response has a 4xx status code
func (o *GetInstallWorkflowUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install workflow unauthorized response has a 5xx status code
func (o *GetInstallWorkflowUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get install workflow unauthorized response a status code equal to that given
func (o *GetInstallWorkflowUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get install workflow unauthorized response
func (o *GetInstallWorkflowUnauthorized) Code() int {
	return 401
}

func (o *GetInstallWorkflowUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}][%d] getInstallWorkflowUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallWorkflowUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}][%d] getInstallWorkflowUnauthorized  %+v", 401, o.Payload)
}

func (o *GetInstallWorkflowUnauthorized) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallWorkflowUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallWorkflowForbidden creates a GetInstallWorkflowForbidden with default headers values
func NewGetInstallWorkflowForbidden() *GetInstallWorkflowForbidden {
	return &GetInstallWorkflowForbidden{}
}

/*
GetInstallWorkflowForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetInstallWorkflowForbidden struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install workflow forbidden response has a 2xx status code
func (o *GetInstallWorkflowForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install workflow forbidden response has a 3xx status code
func (o *GetInstallWorkflowForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install workflow forbidden response has a 4xx status code
func (o *GetInstallWorkflowForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install workflow forbidden response has a 5xx status code
func (o *GetInstallWorkflowForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get install workflow forbidden response a status code equal to that given
func (o *GetInstallWorkflowForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get install workflow forbidden response
func (o *GetInstallWorkflowForbidden) Code() int {
	return 403
}

func (o *GetInstallWorkflowForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}][%d] getInstallWorkflowForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallWorkflowForbidden) String() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}][%d] getInstallWorkflowForbidden  %+v", 403, o.Payload)
}

func (o *GetInstallWorkflowForbidden) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallWorkflowForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallWorkflowNotFound creates a GetInstallWorkflowNotFound with default headers values
func NewGetInstallWorkflowNotFound() *GetInstallWorkflowNotFound {
	return &GetInstallWorkflowNotFound{}
}

/*
GetInstallWorkflowNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetInstallWorkflowNotFound struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install workflow not found response has a 2xx status code
func (o *GetInstallWorkflowNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install workflow not found response has a 3xx status code
func (o *GetInstallWorkflowNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install workflow not found response has a 4xx status code
func (o *GetInstallWorkflowNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get install workflow not found response has a 5xx status code
func (o *GetInstallWorkflowNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get install workflow not found response a status code equal to that given
func (o *GetInstallWorkflowNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get install workflow not found response
func (o *GetInstallWorkflowNotFound) Code() int {
	return 404
}

func (o *GetInstallWorkflowNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}][%d] getInstallWorkflowNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallWorkflowNotFound) String() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}][%d] getInstallWorkflowNotFound  %+v", 404, o.Payload)
}

func (o *GetInstallWorkflowNotFound) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallWorkflowNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstallWorkflowInternalServerError creates a GetInstallWorkflowInternalServerError with default headers values
func NewGetInstallWorkflowInternalServerError() *GetInstallWorkflowInternalServerError {
	return &GetInstallWorkflowInternalServerError{}
}

/*
GetInstallWorkflowInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetInstallWorkflowInternalServerError struct {
	Payload *models.StderrErrResponse
}

// IsSuccess returns true when this get install workflow internal server error response has a 2xx status code
func (o *GetInstallWorkflowInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get install workflow internal server error response has a 3xx status code
func (o *GetInstallWorkflowInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get install workflow internal server error response has a 4xx status code
func (o *GetInstallWorkflowInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get install workflow internal server error response has a 5xx status code
func (o *GetInstallWorkflowInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get install workflow internal server error response a status code equal to that given
func (o *GetInstallWorkflowInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get install workflow internal server error response
func (o *GetInstallWorkflowInternalServerError) Code() int {
	return 500
}

func (o *GetInstallWorkflowInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}][%d] getInstallWorkflowInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallWorkflowInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/install-workflows/{install_workflow_id}][%d] getInstallWorkflowInternalServerError  %+v", 500, o.Payload)
}

func (o *GetInstallWorkflowInternalServerError) GetPayload() *models.StderrErrResponse {
	return o.Payload
}

func (o *GetInstallWorkflowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StderrErrResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

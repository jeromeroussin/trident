// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// S3AuditEventSelectorModifyReader is a Reader for the S3AuditEventSelectorModify structure.
type S3AuditEventSelectorModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3AuditEventSelectorModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewS3AuditEventSelectorModifyAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3AuditEventSelectorModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3AuditEventSelectorModifyAccepted creates a S3AuditEventSelectorModifyAccepted with default headers values
func NewS3AuditEventSelectorModifyAccepted() *S3AuditEventSelectorModifyAccepted {
	return &S3AuditEventSelectorModifyAccepted{}
}

/* S3AuditEventSelectorModifyAccepted describes a response with status code 202, with default header values.

Accepted
*/
type S3AuditEventSelectorModifyAccepted struct {
	Payload *models.S3AuditEventSelector
}

func (o *S3AuditEventSelectorModifyAccepted) Error() string {
	return fmt.Sprintf("[PATCH /protocols/event-selectors/{svm.uuid}/{bucket}][%d] s3AuditEventSelectorModifyAccepted  %+v", 202, o.Payload)
}
func (o *S3AuditEventSelectorModifyAccepted) GetPayload() *models.S3AuditEventSelector {
	return o.Payload
}

func (o *S3AuditEventSelectorModifyAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3AuditEventSelector)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3AuditEventSelectorModifyDefault creates a S3AuditEventSelectorModifyDefault with default headers values
func NewS3AuditEventSelectorModifyDefault(code int) *S3AuditEventSelectorModifyDefault {
	return &S3AuditEventSelectorModifyDefault{
		_statusCode: code,
	}
}

/* S3AuditEventSelectorModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 140902570 | S3 audit event selector configuration is not available. |
| 140902571 | S3 audit configuration was not created for the SVM. |
| 140902572 | S3 audit event selector operation failed. |
| 140902573 | Not all S3 audit event selector have been remove for the SVM. |

*/
type S3AuditEventSelectorModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 audit event selector modify default response
func (o *S3AuditEventSelectorModifyDefault) Code() int {
	return o._statusCode
}

func (o *S3AuditEventSelectorModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/event-selectors/{svm.uuid}/{bucket}][%d] s3_audit_event_selector_modify default  %+v", o._statusCode, o.Payload)
}
func (o *S3AuditEventSelectorModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3AuditEventSelectorModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

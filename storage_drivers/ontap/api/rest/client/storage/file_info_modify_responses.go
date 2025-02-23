// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FileInfoModifyReader is a Reader for the FileInfoModify structure.
type FileInfoModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileInfoModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFileInfoModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileInfoModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileInfoModifyOK creates a FileInfoModifyOK with default headers values
func NewFileInfoModifyOK() *FileInfoModifyOK {
	return &FileInfoModifyOK{}
}

/* FileInfoModifyOK describes a response with status code 200, with default header values.

OK
*/
type FileInfoModifyOK struct {
}

func (o *FileInfoModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /storage/volumes/{volume.uuid}/files/{path}][%d] fileInfoModifyOK ", 200)
}

func (o *FileInfoModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFileInfoModifyDefault creates a FileInfoModifyDefault with default headers values
func NewFileInfoModifyDefault(code int) *FileInfoModifyDefault {
	return &FileInfoModifyDefault{
		_statusCode: code,
	}
}

/* FileInfoModifyDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 918235 | A volume with UUID {volume.uuid} was not found. |
| 6488081 | The {field} field is not supported for PATCH operations. |
| 6488082 | Failed to rename {path}. |
| 6488083 | Failed to rename {path} to {path} because a directory named {path} already exists. |

*/
type FileInfoModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the file info modify default response
func (o *FileInfoModifyDefault) Code() int {
	return o._statusCode
}

func (o *FileInfoModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /storage/volumes/{volume.uuid}/files/{path}][%d] file_info_modify default  %+v", o._statusCode, o.Payload)
}
func (o *FileInfoModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileInfoModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

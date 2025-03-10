// Code generated by go-swagger; DO NOT EDIT.

package svm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SvmDeleteReader is a Reader for the SvmDelete structure.
type SvmDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewSvmDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmDeleteOK creates a SvmDeleteOK with default headers values
func NewSvmDeleteOK() *SvmDeleteOK {
	return &SvmDeleteOK{}
}

/*
SvmDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SvmDeleteOK struct {
	Payload *models.SvmJobLinkResponse
}

// IsSuccess returns true when this svm delete o k response has a 2xx status code
func (o *SvmDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm delete o k response has a 3xx status code
func (o *SvmDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm delete o k response has a 4xx status code
func (o *SvmDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm delete o k response has a 5xx status code
func (o *SvmDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm delete o k response a status code equal to that given
func (o *SvmDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the svm delete o k response
func (o *SvmDeleteOK) Code() int {
	return 200
}

func (o *SvmDeleteOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/svms/{uuid}][%d] svmDeleteOK %s", 200, payload)
}

func (o *SvmDeleteOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/svms/{uuid}][%d] svmDeleteOK %s", 200, payload)
}

func (o *SvmDeleteOK) GetPayload() *models.SvmJobLinkResponse {
	return o.Payload
}

func (o *SvmDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SvmJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmDeleteAccepted creates a SvmDeleteAccepted with default headers values
func NewSvmDeleteAccepted() *SvmDeleteAccepted {
	return &SvmDeleteAccepted{}
}

/*
SvmDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SvmDeleteAccepted struct {
	Payload *models.SvmJobLinkResponse
}

// IsSuccess returns true when this svm delete accepted response has a 2xx status code
func (o *SvmDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm delete accepted response has a 3xx status code
func (o *SvmDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm delete accepted response has a 4xx status code
func (o *SvmDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm delete accepted response has a 5xx status code
func (o *SvmDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this svm delete accepted response a status code equal to that given
func (o *SvmDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

// Code gets the status code for the svm delete accepted response
func (o *SvmDeleteAccepted) Code() int {
	return 202
}

func (o *SvmDeleteAccepted) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/svms/{uuid}][%d] svmDeleteAccepted %s", 202, payload)
}

func (o *SvmDeleteAccepted) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/svms/{uuid}][%d] svmDeleteAccepted %s", 202, payload)
}

func (o *SvmDeleteAccepted) GetPayload() *models.SvmJobLinkResponse {
	return o.Payload
}

func (o *SvmDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SvmJobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmDeleteDefault creates a SvmDeleteDefault with default headers values
func NewSvmDeleteDefault(code int) *SvmDeleteDefault {
	return &SvmDeleteDefault{
		_statusCode: code,
	}
}

/*
	SvmDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

<br/>
```
| Error codes | Description |
| ----------- | ----------- |
| 13434894    | Maximum allowed SVM jobs exceeded. Wait and retry. |
| 2621525     | SVM cannot be deleted as it is associated with an Active Directory configured CIFS server. Delete the CIFS server using "cifs delete" and retry the operation. |
```
<br/>
*/
type SvmDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this svm delete default response has a 2xx status code
func (o *SvmDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm delete default response has a 3xx status code
func (o *SvmDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm delete default response has a 4xx status code
func (o *SvmDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm delete default response has a 5xx status code
func (o *SvmDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm delete default response a status code equal to that given
func (o *SvmDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the svm delete default response
func (o *SvmDeleteDefault) Code() int {
	return o._statusCode
}

func (o *SvmDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/svms/{uuid}][%d] svm_delete default %s", o._statusCode, payload)
}

func (o *SvmDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /svm/svms/{uuid}][%d] svm_delete default %s", o._statusCode, payload)
}

func (o *SvmDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

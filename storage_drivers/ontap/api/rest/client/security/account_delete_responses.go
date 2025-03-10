// Code generated by go-swagger; DO NOT EDIT.

package security

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

// AccountDeleteReader is a Reader for the AccountDelete structure.
type AccountDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountDeleteOK creates a AccountDeleteOK with default headers values
func NewAccountDeleteOK() *AccountDeleteOK {
	return &AccountDeleteOK{}
}

/*
AccountDeleteOK describes a response with status code 200, with default header values.

OK
*/
type AccountDeleteOK struct {
}

// IsSuccess returns true when this account delete o k response has a 2xx status code
func (o *AccountDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account delete o k response has a 3xx status code
func (o *AccountDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account delete o k response has a 4xx status code
func (o *AccountDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account delete o k response has a 5xx status code
func (o *AccountDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account delete o k response a status code equal to that given
func (o *AccountDeleteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account delete o k response
func (o *AccountDeleteOK) Code() int {
	return 200
}

func (o *AccountDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /security/accounts/{owner.uuid}/{name}][%d] accountDeleteOK", 200)
}

func (o *AccountDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /security/accounts/{owner.uuid}/{name}][%d] accountDeleteOK", 200)
}

func (o *AccountDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewAccountDeleteDefault creates a AccountDeleteDefault with default headers values
func NewAccountDeleteDefault(code int) *AccountDeleteDefault {
	return &AccountDeleteDefault{
		_statusCode: code,
	}
}

/*
	AccountDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5636098 | Last unlocked account that has an admin role cannot be deleted. |
| 5636125 | The operation is not supported on system accounts. |
| 5636146 | Cannot delete the last console account with admin role. |
Also see the table of common errors in the <a href="#Response_body">Response body</a> overview section of this documentation.
*/
type AccountDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this account delete default response has a 2xx status code
func (o *AccountDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this account delete default response has a 3xx status code
func (o *AccountDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this account delete default response has a 4xx status code
func (o *AccountDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this account delete default response has a 5xx status code
func (o *AccountDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this account delete default response a status code equal to that given
func (o *AccountDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the account delete default response
func (o *AccountDeleteDefault) Code() int {
	return o._statusCode
}

func (o *AccountDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/accounts/{owner.uuid}/{name}][%d] account_delete default %s", o._statusCode, payload)
}

func (o *AccountDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /security/accounts/{owner.uuid}/{name}][%d] account_delete default %s", o._statusCode, payload)
}

func (o *AccountDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *AccountDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

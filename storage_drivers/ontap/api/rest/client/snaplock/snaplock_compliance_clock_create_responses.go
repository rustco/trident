// Code generated by go-swagger; DO NOT EDIT.

package snaplock

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

// SnaplockComplianceClockCreateReader is a Reader for the SnaplockComplianceClockCreate structure.
type SnaplockComplianceClockCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockComplianceClockCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSnaplockComplianceClockCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockComplianceClockCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockComplianceClockCreateCreated creates a SnaplockComplianceClockCreateCreated with default headers values
func NewSnaplockComplianceClockCreateCreated() *SnaplockComplianceClockCreateCreated {
	return &SnaplockComplianceClockCreateCreated{}
}

/*
SnaplockComplianceClockCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SnaplockComplianceClockCreateCreated struct {

	/* Useful for tracking the resource location
	 */
	Location string
}

// IsSuccess returns true when this snaplock compliance clock create created response has a 2xx status code
func (o *SnaplockComplianceClockCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock compliance clock create created response has a 3xx status code
func (o *SnaplockComplianceClockCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock compliance clock create created response has a 4xx status code
func (o *SnaplockComplianceClockCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock compliance clock create created response has a 5xx status code
func (o *SnaplockComplianceClockCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock compliance clock create created response a status code equal to that given
func (o *SnaplockComplianceClockCreateCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the snaplock compliance clock create created response
func (o *SnaplockComplianceClockCreateCreated) Code() int {
	return 201
}

func (o *SnaplockComplianceClockCreateCreated) Error() string {
	return fmt.Sprintf("[POST /storage/snaplock/compliance-clocks][%d] snaplockComplianceClockCreateCreated", 201)
}

func (o *SnaplockComplianceClockCreateCreated) String() string {
	return fmt.Sprintf("[POST /storage/snaplock/compliance-clocks][%d] snaplockComplianceClockCreateCreated", 201)
}

func (o *SnaplockComplianceClockCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	return nil
}

// NewSnaplockComplianceClockCreateDefault creates a SnaplockComplianceClockCreateDefault with default headers values
func NewSnaplockComplianceClockCreateDefault(code int) *SnaplockComplianceClockCreateDefault {
	return &SnaplockComplianceClockCreateDefault{
		_statusCode: code,
	}
}

/*
	SnaplockComplianceClockCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
|  13763062   | Compliance Clock already initialized  |
|  13763084   | ComplianceClock re-initialization failed as it requires that all nodes in the cluster are healthy, all volumes are in online state, no volumes are present in the volume recovery queue and there are no SnapLock volumes or volumes with \"snapshot-locking-enabled\" set to true or S3 buckets with object locking enabled.  |
|  13763085   | ComplianceClock re-initialization supported on effective cluster version of ONTAP 9.14.1 or later. |
|  14090240   | Node with the specified UUID does not exist |
|  14090241   | The specified node name and node UUID refer to different nodes |
|  14090343   | Invalid field |
*/
type SnaplockComplianceClockCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// IsSuccess returns true when this snaplock compliance clock create default response has a 2xx status code
func (o *SnaplockComplianceClockCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock compliance clock create default response has a 3xx status code
func (o *SnaplockComplianceClockCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock compliance clock create default response has a 4xx status code
func (o *SnaplockComplianceClockCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock compliance clock create default response has a 5xx status code
func (o *SnaplockComplianceClockCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock compliance clock create default response a status code equal to that given
func (o *SnaplockComplianceClockCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the snaplock compliance clock create default response
func (o *SnaplockComplianceClockCreateDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockComplianceClockCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/snaplock/compliance-clocks][%d] snaplock_compliance_clock_create default %s", o._statusCode, payload)
}

func (o *SnaplockComplianceClockCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /storage/snaplock/compliance-clocks][%d] snaplock_compliance_clock_create default %s", o._statusCode, payload)
}

func (o *SnaplockComplianceClockCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockComplianceClockCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

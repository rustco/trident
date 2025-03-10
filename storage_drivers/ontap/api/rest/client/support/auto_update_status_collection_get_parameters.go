// Code generated by go-swagger; DO NOT EDIT.

package support

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
	"github.com/go-openapi/swag"
)

// NewAutoUpdateStatusCollectionGetParams creates a new AutoUpdateStatusCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAutoUpdateStatusCollectionGetParams() *AutoUpdateStatusCollectionGetParams {
	return &AutoUpdateStatusCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAutoUpdateStatusCollectionGetParamsWithTimeout creates a new AutoUpdateStatusCollectionGetParams object
// with the ability to set a timeout on a request.
func NewAutoUpdateStatusCollectionGetParamsWithTimeout(timeout time.Duration) *AutoUpdateStatusCollectionGetParams {
	return &AutoUpdateStatusCollectionGetParams{
		timeout: timeout,
	}
}

// NewAutoUpdateStatusCollectionGetParamsWithContext creates a new AutoUpdateStatusCollectionGetParams object
// with the ability to set a context for a request.
func NewAutoUpdateStatusCollectionGetParamsWithContext(ctx context.Context) *AutoUpdateStatusCollectionGetParams {
	return &AutoUpdateStatusCollectionGetParams{
		Context: ctx,
	}
}

// NewAutoUpdateStatusCollectionGetParamsWithHTTPClient creates a new AutoUpdateStatusCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAutoUpdateStatusCollectionGetParamsWithHTTPClient(client *http.Client) *AutoUpdateStatusCollectionGetParams {
	return &AutoUpdateStatusCollectionGetParams{
		HTTPClient: client,
	}
}

/*
AutoUpdateStatusCollectionGetParams contains all the parameters to send to the API endpoint

	for the auto update status collection get operation.

	Typically these are written to a http.Request.
*/
type AutoUpdateStatusCollectionGetParams struct {

	/* ContentCategory.

	   Filter by content_category
	*/
	ContentCategory *string

	/* ContentType.

	   Filter by content_type
	*/
	ContentType *string

	/* CreationTime.

	   Filter by creation_time
	*/
	CreationTime *string

	/* Description.

	   Filter by description
	*/
	Description *string

	/* EndTime.

	   Filter by end_time
	*/
	EndTime *string

	/* ExpiryTime.

	   Filter by expiry_time
	*/
	ExpiryTime *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* LastStateChangeTime.

	   Filter by last_state_change_time
	*/
	LastStateChangeTime *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* PackageID.

	   Filter by package_id
	*/
	PackageID *string

	/* PercentComplete.

	   Filter by percent_complete
	*/
	PercentComplete *int64

	/* RemainingTime.

	   Filter by remaining_time
	*/
	RemainingTime *string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* ScheduledTime.

	   Filter by scheduled_time
	*/
	ScheduledTime *string

	/* StartTime.

	   Filter by start_time
	*/
	StartTime *string

	/* State.

	   Filter by state
	*/
	State *string

	/* StatusArgumentsCode.

	   Filter by status.arguments.code
	*/
	StatusArgumentsCode *string

	/* StatusArgumentsMessage.

	   Filter by status.arguments.message
	*/
	StatusArgumentsMessage *string

	/* StatusCode.

	   Filter by status.code
	*/
	StatusCode *string

	/* StatusMessage.

	   Filter by status.message
	*/
	StatusMessage *string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auto update status collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AutoUpdateStatusCollectionGetParams) WithDefaults() *AutoUpdateStatusCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auto update status collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AutoUpdateStatusCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := AutoUpdateStatusCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithTimeout(timeout time.Duration) *AutoUpdateStatusCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithContext(ctx context.Context) *AutoUpdateStatusCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithHTTPClient(client *http.Client) *AutoUpdateStatusCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContentCategory adds the contentCategory to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithContentCategory(contentCategory *string) *AutoUpdateStatusCollectionGetParams {
	o.SetContentCategory(contentCategory)
	return o
}

// SetContentCategory adds the contentCategory to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetContentCategory(contentCategory *string) {
	o.ContentCategory = contentCategory
}

// WithContentType adds the contentType to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithContentType(contentType *string) *AutoUpdateStatusCollectionGetParams {
	o.SetContentType(contentType)
	return o
}

// SetContentType adds the contentType to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetContentType(contentType *string) {
	o.ContentType = contentType
}

// WithCreationTime adds the creationTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithCreationTime(creationTime *string) *AutoUpdateStatusCollectionGetParams {
	o.SetCreationTime(creationTime)
	return o
}

// SetCreationTime adds the creationTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetCreationTime(creationTime *string) {
	o.CreationTime = creationTime
}

// WithDescription adds the description to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithDescription(description *string) *AutoUpdateStatusCollectionGetParams {
	o.SetDescription(description)
	return o
}

// SetDescription adds the description to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetDescription(description *string) {
	o.Description = description
}

// WithEndTime adds the endTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithEndTime(endTime *string) *AutoUpdateStatusCollectionGetParams {
	o.SetEndTime(endTime)
	return o
}

// SetEndTime adds the endTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetEndTime(endTime *string) {
	o.EndTime = endTime
}

// WithExpiryTime adds the expiryTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithExpiryTime(expiryTime *string) *AutoUpdateStatusCollectionGetParams {
	o.SetExpiryTime(expiryTime)
	return o
}

// SetExpiryTime adds the expiryTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetExpiryTime(expiryTime *string) {
	o.ExpiryTime = expiryTime
}

// WithFields adds the fields to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithFields(fields []string) *AutoUpdateStatusCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithLastStateChangeTime adds the lastStateChangeTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithLastStateChangeTime(lastStateChangeTime *string) *AutoUpdateStatusCollectionGetParams {
	o.SetLastStateChangeTime(lastStateChangeTime)
	return o
}

// SetLastStateChangeTime adds the lastStateChangeTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetLastStateChangeTime(lastStateChangeTime *string) {
	o.LastStateChangeTime = lastStateChangeTime
}

// WithMaxRecords adds the maxRecords to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithMaxRecords(maxRecords *int64) *AutoUpdateStatusCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithOrderBy(orderBy []string) *AutoUpdateStatusCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithPackageID adds the packageID to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithPackageID(packageID *string) *AutoUpdateStatusCollectionGetParams {
	o.SetPackageID(packageID)
	return o
}

// SetPackageID adds the packageId to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetPackageID(packageID *string) {
	o.PackageID = packageID
}

// WithPercentComplete adds the percentComplete to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithPercentComplete(percentComplete *int64) *AutoUpdateStatusCollectionGetParams {
	o.SetPercentComplete(percentComplete)
	return o
}

// SetPercentComplete adds the percentComplete to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetPercentComplete(percentComplete *int64) {
	o.PercentComplete = percentComplete
}

// WithRemainingTime adds the remainingTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithRemainingTime(remainingTime *string) *AutoUpdateStatusCollectionGetParams {
	o.SetRemainingTime(remainingTime)
	return o
}

// SetRemainingTime adds the remainingTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetRemainingTime(remainingTime *string) {
	o.RemainingTime = remainingTime
}

// WithReturnRecords adds the returnRecords to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithReturnRecords(returnRecords *bool) *AutoUpdateStatusCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *AutoUpdateStatusCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithScheduledTime adds the scheduledTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithScheduledTime(scheduledTime *string) *AutoUpdateStatusCollectionGetParams {
	o.SetScheduledTime(scheduledTime)
	return o
}

// SetScheduledTime adds the scheduledTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetScheduledTime(scheduledTime *string) {
	o.ScheduledTime = scheduledTime
}

// WithStartTime adds the startTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithStartTime(startTime *string) *AutoUpdateStatusCollectionGetParams {
	o.SetStartTime(startTime)
	return o
}

// SetStartTime adds the startTime to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetStartTime(startTime *string) {
	o.StartTime = startTime
}

// WithState adds the state to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithState(state *string) *AutoUpdateStatusCollectionGetParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetState(state *string) {
	o.State = state
}

// WithStatusArgumentsCode adds the statusArgumentsCode to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithStatusArgumentsCode(statusArgumentsCode *string) *AutoUpdateStatusCollectionGetParams {
	o.SetStatusArgumentsCode(statusArgumentsCode)
	return o
}

// SetStatusArgumentsCode adds the statusArgumentsCode to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetStatusArgumentsCode(statusArgumentsCode *string) {
	o.StatusArgumentsCode = statusArgumentsCode
}

// WithStatusArgumentsMessage adds the statusArgumentsMessage to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithStatusArgumentsMessage(statusArgumentsMessage *string) *AutoUpdateStatusCollectionGetParams {
	o.SetStatusArgumentsMessage(statusArgumentsMessage)
	return o
}

// SetStatusArgumentsMessage adds the statusArgumentsMessage to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetStatusArgumentsMessage(statusArgumentsMessage *string) {
	o.StatusArgumentsMessage = statusArgumentsMessage
}

// WithStatusCode adds the statusCode to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithStatusCode(statusCode *string) *AutoUpdateStatusCollectionGetParams {
	o.SetStatusCode(statusCode)
	return o
}

// SetStatusCode adds the statusCode to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetStatusCode(statusCode *string) {
	o.StatusCode = statusCode
}

// WithStatusMessage adds the statusMessage to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithStatusMessage(statusMessage *string) *AutoUpdateStatusCollectionGetParams {
	o.SetStatusMessage(statusMessage)
	return o
}

// SetStatusMessage adds the statusMessage to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetStatusMessage(statusMessage *string) {
	o.StatusMessage = statusMessage
}

// WithUUID adds the uuid to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) WithUUID(uuid *string) *AutoUpdateStatusCollectionGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the auto update status collection get params
func (o *AutoUpdateStatusCollectionGetParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *AutoUpdateStatusCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ContentCategory != nil {

		// query param content_category
		var qrContentCategory string

		if o.ContentCategory != nil {
			qrContentCategory = *o.ContentCategory
		}
		qContentCategory := qrContentCategory
		if qContentCategory != "" {

			if err := r.SetQueryParam("content_category", qContentCategory); err != nil {
				return err
			}
		}
	}

	if o.ContentType != nil {

		// query param content_type
		var qrContentType string

		if o.ContentType != nil {
			qrContentType = *o.ContentType
		}
		qContentType := qrContentType
		if qContentType != "" {

			if err := r.SetQueryParam("content_type", qContentType); err != nil {
				return err
			}
		}
	}

	if o.CreationTime != nil {

		// query param creation_time
		var qrCreationTime string

		if o.CreationTime != nil {
			qrCreationTime = *o.CreationTime
		}
		qCreationTime := qrCreationTime
		if qCreationTime != "" {

			if err := r.SetQueryParam("creation_time", qCreationTime); err != nil {
				return err
			}
		}
	}

	if o.Description != nil {

		// query param description
		var qrDescription string

		if o.Description != nil {
			qrDescription = *o.Description
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.EndTime != nil {

		// query param end_time
		var qrEndTime string

		if o.EndTime != nil {
			qrEndTime = *o.EndTime
		}
		qEndTime := qrEndTime
		if qEndTime != "" {

			if err := r.SetQueryParam("end_time", qEndTime); err != nil {
				return err
			}
		}
	}

	if o.ExpiryTime != nil {

		// query param expiry_time
		var qrExpiryTime string

		if o.ExpiryTime != nil {
			qrExpiryTime = *o.ExpiryTime
		}
		qExpiryTime := qrExpiryTime
		if qExpiryTime != "" {

			if err := r.SetQueryParam("expiry_time", qExpiryTime); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.LastStateChangeTime != nil {

		// query param last_state_change_time
		var qrLastStateChangeTime string

		if o.LastStateChangeTime != nil {
			qrLastStateChangeTime = *o.LastStateChangeTime
		}
		qLastStateChangeTime := qrLastStateChangeTime
		if qLastStateChangeTime != "" {

			if err := r.SetQueryParam("last_state_change_time", qLastStateChangeTime); err != nil {
				return err
			}
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.PackageID != nil {

		// query param package_id
		var qrPackageID string

		if o.PackageID != nil {
			qrPackageID = *o.PackageID
		}
		qPackageID := qrPackageID
		if qPackageID != "" {

			if err := r.SetQueryParam("package_id", qPackageID); err != nil {
				return err
			}
		}
	}

	if o.PercentComplete != nil {

		// query param percent_complete
		var qrPercentComplete int64

		if o.PercentComplete != nil {
			qrPercentComplete = *o.PercentComplete
		}
		qPercentComplete := swag.FormatInt64(qrPercentComplete)
		if qPercentComplete != "" {

			if err := r.SetQueryParam("percent_complete", qPercentComplete); err != nil {
				return err
			}
		}
	}

	if o.RemainingTime != nil {

		// query param remaining_time
		var qrRemainingTime string

		if o.RemainingTime != nil {
			qrRemainingTime = *o.RemainingTime
		}
		qRemainingTime := qrRemainingTime
		if qRemainingTime != "" {

			if err := r.SetQueryParam("remaining_time", qRemainingTime); err != nil {
				return err
			}
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.ScheduledTime != nil {

		// query param scheduled_time
		var qrScheduledTime string

		if o.ScheduledTime != nil {
			qrScheduledTime = *o.ScheduledTime
		}
		qScheduledTime := qrScheduledTime
		if qScheduledTime != "" {

			if err := r.SetQueryParam("scheduled_time", qScheduledTime); err != nil {
				return err
			}
		}
	}

	if o.StartTime != nil {

		// query param start_time
		var qrStartTime string

		if o.StartTime != nil {
			qrStartTime = *o.StartTime
		}
		qStartTime := qrStartTime
		if qStartTime != "" {

			if err := r.SetQueryParam("start_time", qStartTime); err != nil {
				return err
			}
		}
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if o.StatusArgumentsCode != nil {

		// query param status.arguments.code
		var qrStatusArgumentsCode string

		if o.StatusArgumentsCode != nil {
			qrStatusArgumentsCode = *o.StatusArgumentsCode
		}
		qStatusArgumentsCode := qrStatusArgumentsCode
		if qStatusArgumentsCode != "" {

			if err := r.SetQueryParam("status.arguments.code", qStatusArgumentsCode); err != nil {
				return err
			}
		}
	}

	if o.StatusArgumentsMessage != nil {

		// query param status.arguments.message
		var qrStatusArgumentsMessage string

		if o.StatusArgumentsMessage != nil {
			qrStatusArgumentsMessage = *o.StatusArgumentsMessage
		}
		qStatusArgumentsMessage := qrStatusArgumentsMessage
		if qStatusArgumentsMessage != "" {

			if err := r.SetQueryParam("status.arguments.message", qStatusArgumentsMessage); err != nil {
				return err
			}
		}
	}

	if o.StatusCode != nil {

		// query param status.code
		var qrStatusCode string

		if o.StatusCode != nil {
			qrStatusCode = *o.StatusCode
		}
		qStatusCode := qrStatusCode
		if qStatusCode != "" {

			if err := r.SetQueryParam("status.code", qStatusCode); err != nil {
				return err
			}
		}
	}

	if o.StatusMessage != nil {

		// query param status.message
		var qrStatusMessage string

		if o.StatusMessage != nil {
			qrStatusMessage = *o.StatusMessage
		}
		qStatusMessage := qrStatusMessage
		if qStatusMessage != "" {

			if err := r.SetQueryParam("status.message", qStatusMessage); err != nil {
				return err
			}
		}
	}

	if o.UUID != nil {

		// query param uuid
		var qrUUID string

		if o.UUID != nil {
			qrUUID = *o.UUID
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamAutoUpdateStatusCollectionGet binds the parameter fields
func (o *AutoUpdateStatusCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamAutoUpdateStatusCollectionGet binds the parameter order_by
func (o *AutoUpdateStatusCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}

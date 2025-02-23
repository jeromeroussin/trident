// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewSecurityKeyManagerGetParams creates a new SecurityKeyManagerGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityKeyManagerGetParams() *SecurityKeyManagerGetParams {
	return &SecurityKeyManagerGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityKeyManagerGetParamsWithTimeout creates a new SecurityKeyManagerGetParams object
// with the ability to set a timeout on a request.
func NewSecurityKeyManagerGetParamsWithTimeout(timeout time.Duration) *SecurityKeyManagerGetParams {
	return &SecurityKeyManagerGetParams{
		timeout: timeout,
	}
}

// NewSecurityKeyManagerGetParamsWithContext creates a new SecurityKeyManagerGetParams object
// with the ability to set a context for a request.
func NewSecurityKeyManagerGetParamsWithContext(ctx context.Context) *SecurityKeyManagerGetParams {
	return &SecurityKeyManagerGetParams{
		Context: ctx,
	}
}

// NewSecurityKeyManagerGetParamsWithHTTPClient creates a new SecurityKeyManagerGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityKeyManagerGetParamsWithHTTPClient(client *http.Client) *SecurityKeyManagerGetParams {
	return &SecurityKeyManagerGetParams{
		HTTPClient: client,
	}
}

/* SecurityKeyManagerGetParams contains all the parameters to send to the API endpoint
   for the security key manager get operation.

   Typically these are written to a http.Request.
*/
type SecurityKeyManagerGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* UUID.

	   Key manager UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security key manager get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerGetParams) WithDefaults() *SecurityKeyManagerGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security key manager get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityKeyManagerGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security key manager get params
func (o *SecurityKeyManagerGetParams) WithTimeout(timeout time.Duration) *SecurityKeyManagerGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security key manager get params
func (o *SecurityKeyManagerGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security key manager get params
func (o *SecurityKeyManagerGetParams) WithContext(ctx context.Context) *SecurityKeyManagerGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security key manager get params
func (o *SecurityKeyManagerGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security key manager get params
func (o *SecurityKeyManagerGetParams) WithHTTPClient(client *http.Client) *SecurityKeyManagerGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security key manager get params
func (o *SecurityKeyManagerGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the security key manager get params
func (o *SecurityKeyManagerGetParams) WithFieldsQueryParameter(fields []string) *SecurityKeyManagerGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the security key manager get params
func (o *SecurityKeyManagerGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithUUIDPathParameter adds the uuid to the security key manager get params
func (o *SecurityKeyManagerGetParams) WithUUIDPathParameter(uuid string) *SecurityKeyManagerGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the security key manager get params
func (o *SecurityKeyManagerGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityKeyManagerGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSecurityKeyManagerGet binds the parameter fields
func (o *SecurityKeyManagerGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

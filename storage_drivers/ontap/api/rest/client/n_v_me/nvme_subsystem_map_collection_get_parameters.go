// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NewNvmeSubsystemMapCollectionGetParams creates a new NvmeSubsystemMapCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNvmeSubsystemMapCollectionGetParams() *NvmeSubsystemMapCollectionGetParams {
	return &NvmeSubsystemMapCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNvmeSubsystemMapCollectionGetParamsWithTimeout creates a new NvmeSubsystemMapCollectionGetParams object
// with the ability to set a timeout on a request.
func NewNvmeSubsystemMapCollectionGetParamsWithTimeout(timeout time.Duration) *NvmeSubsystemMapCollectionGetParams {
	return &NvmeSubsystemMapCollectionGetParams{
		timeout: timeout,
	}
}

// NewNvmeSubsystemMapCollectionGetParamsWithContext creates a new NvmeSubsystemMapCollectionGetParams object
// with the ability to set a context for a request.
func NewNvmeSubsystemMapCollectionGetParamsWithContext(ctx context.Context) *NvmeSubsystemMapCollectionGetParams {
	return &NvmeSubsystemMapCollectionGetParams{
		Context: ctx,
	}
}

// NewNvmeSubsystemMapCollectionGetParamsWithHTTPClient creates a new NvmeSubsystemMapCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNvmeSubsystemMapCollectionGetParamsWithHTTPClient(client *http.Client) *NvmeSubsystemMapCollectionGetParams {
	return &NvmeSubsystemMapCollectionGetParams{
		HTTPClient: client,
	}
}

/* NvmeSubsystemMapCollectionGetParams contains all the parameters to send to the API endpoint
   for the nvme subsystem map collection get operation.

   Typically these are written to a http.Request.
*/
type NvmeSubsystemMapCollectionGetParams struct {

	/* Anagrpid.

	   Filter by anagrpid
	*/
	AnagrpIDQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* NamespaceName.

	   Filter by namespace.name
	*/
	NamespaceNameQueryParameter *string

	/* NamespaceNodeName.

	   Filter by namespace.node.name
	*/
	NamespaceNodeNameQueryParameter *string

	/* NamespaceNodeUUID.

	   Filter by namespace.node.uuid
	*/
	NamespaceNodeUUIDQueryParameter *string

	/* NamespaceUUID.

	   Filter by namespace.uuid
	*/
	NamespaceUUIDQueryParameter *string

	/* Nsid.

	   Filter by nsid
	*/
	NsIDQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* SubsystemName.

	   Filter by subsystem.name
	*/
	SubsystemNameQueryParameter *string

	/* SubsystemUUID.

	   Filter by subsystem.uuid
	*/
	SubsystemUUIDQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nvme subsystem map collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemMapCollectionGetParams) WithDefaults() *NvmeSubsystemMapCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nvme subsystem map collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemMapCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := NvmeSubsystemMapCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithTimeout(timeout time.Duration) *NvmeSubsystemMapCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithContext(ctx context.Context) *NvmeSubsystemMapCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithHTTPClient(client *http.Client) *NvmeSubsystemMapCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAnagrpIDQueryParameter adds the anagrpid to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithAnagrpIDQueryParameter(anagrpid *string) *NvmeSubsystemMapCollectionGetParams {
	o.SetAnagrpIDQueryParameter(anagrpid)
	return o
}

// SetAnagrpIDQueryParameter adds the anagrpid to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetAnagrpIDQueryParameter(anagrpid *string) {
	o.AnagrpIDQueryParameter = anagrpid
}

// WithFieldsQueryParameter adds the fields to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithFieldsQueryParameter(fields []string) *NvmeSubsystemMapCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *NvmeSubsystemMapCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNamespaceNameQueryParameter adds the namespaceName to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithNamespaceNameQueryParameter(namespaceName *string) *NvmeSubsystemMapCollectionGetParams {
	o.SetNamespaceNameQueryParameter(namespaceName)
	return o
}

// SetNamespaceNameQueryParameter adds the namespaceName to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetNamespaceNameQueryParameter(namespaceName *string) {
	o.NamespaceNameQueryParameter = namespaceName
}

// WithNamespaceNodeNameQueryParameter adds the namespaceNodeName to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithNamespaceNodeNameQueryParameter(namespaceNodeName *string) *NvmeSubsystemMapCollectionGetParams {
	o.SetNamespaceNodeNameQueryParameter(namespaceNodeName)
	return o
}

// SetNamespaceNodeNameQueryParameter adds the namespaceNodeName to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetNamespaceNodeNameQueryParameter(namespaceNodeName *string) {
	o.NamespaceNodeNameQueryParameter = namespaceNodeName
}

// WithNamespaceNodeUUIDQueryParameter adds the namespaceNodeUUID to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithNamespaceNodeUUIDQueryParameter(namespaceNodeUUID *string) *NvmeSubsystemMapCollectionGetParams {
	o.SetNamespaceNodeUUIDQueryParameter(namespaceNodeUUID)
	return o
}

// SetNamespaceNodeUUIDQueryParameter adds the namespaceNodeUuid to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetNamespaceNodeUUIDQueryParameter(namespaceNodeUUID *string) {
	o.NamespaceNodeUUIDQueryParameter = namespaceNodeUUID
}

// WithNamespaceUUIDQueryParameter adds the namespaceUUID to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithNamespaceUUIDQueryParameter(namespaceUUID *string) *NvmeSubsystemMapCollectionGetParams {
	o.SetNamespaceUUIDQueryParameter(namespaceUUID)
	return o
}

// SetNamespaceUUIDQueryParameter adds the namespaceUuid to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetNamespaceUUIDQueryParameter(namespaceUUID *string) {
	o.NamespaceUUIDQueryParameter = namespaceUUID
}

// WithNsIDQueryParameter adds the nsid to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithNsIDQueryParameter(nsid *string) *NvmeSubsystemMapCollectionGetParams {
	o.SetNsIDQueryParameter(nsid)
	return o
}

// SetNsIDQueryParameter adds the nsid to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetNsIDQueryParameter(nsid *string) {
	o.NsIDQueryParameter = nsid
}

// WithOrderByQueryParameter adds the orderBy to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *NvmeSubsystemMapCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *NvmeSubsystemMapCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *NvmeSubsystemMapCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSubsystemNameQueryParameter adds the subsystemName to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithSubsystemNameQueryParameter(subsystemName *string) *NvmeSubsystemMapCollectionGetParams {
	o.SetSubsystemNameQueryParameter(subsystemName)
	return o
}

// SetSubsystemNameQueryParameter adds the subsystemName to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetSubsystemNameQueryParameter(subsystemName *string) {
	o.SubsystemNameQueryParameter = subsystemName
}

// WithSubsystemUUIDQueryParameter adds the subsystemUUID to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithSubsystemUUIDQueryParameter(subsystemUUID *string) *NvmeSubsystemMapCollectionGetParams {
	o.SetSubsystemUUIDQueryParameter(subsystemUUID)
	return o
}

// SetSubsystemUUIDQueryParameter adds the subsystemUuid to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetSubsystemUUIDQueryParameter(subsystemUUID *string) {
	o.SubsystemUUIDQueryParameter = subsystemUUID
}

// WithSVMNameQueryParameter adds the svmName to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *NvmeSubsystemMapCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *NvmeSubsystemMapCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the nvme subsystem map collection get params
func (o *NvmeSubsystemMapCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NvmeSubsystemMapCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AnagrpIDQueryParameter != nil {

		// query param anagrpid
		var qrAnagrpid string

		if o.AnagrpIDQueryParameter != nil {
			qrAnagrpid = *o.AnagrpIDQueryParameter
		}
		qAnagrpid := qrAnagrpid
		if qAnagrpid != "" {

			if err := r.SetQueryParam("anagrpid", qAnagrpid); err != nil {
				return err
			}
		}
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NamespaceNameQueryParameter != nil {

		// query param namespace.name
		var qrNamespaceName string

		if o.NamespaceNameQueryParameter != nil {
			qrNamespaceName = *o.NamespaceNameQueryParameter
		}
		qNamespaceName := qrNamespaceName
		if qNamespaceName != "" {

			if err := r.SetQueryParam("namespace.name", qNamespaceName); err != nil {
				return err
			}
		}
	}

	if o.NamespaceNodeNameQueryParameter != nil {

		// query param namespace.node.name
		var qrNamespaceNodeName string

		if o.NamespaceNodeNameQueryParameter != nil {
			qrNamespaceNodeName = *o.NamespaceNodeNameQueryParameter
		}
		qNamespaceNodeName := qrNamespaceNodeName
		if qNamespaceNodeName != "" {

			if err := r.SetQueryParam("namespace.node.name", qNamespaceNodeName); err != nil {
				return err
			}
		}
	}

	if o.NamespaceNodeUUIDQueryParameter != nil {

		// query param namespace.node.uuid
		var qrNamespaceNodeUUID string

		if o.NamespaceNodeUUIDQueryParameter != nil {
			qrNamespaceNodeUUID = *o.NamespaceNodeUUIDQueryParameter
		}
		qNamespaceNodeUUID := qrNamespaceNodeUUID
		if qNamespaceNodeUUID != "" {

			if err := r.SetQueryParam("namespace.node.uuid", qNamespaceNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.NamespaceUUIDQueryParameter != nil {

		// query param namespace.uuid
		var qrNamespaceUUID string

		if o.NamespaceUUIDQueryParameter != nil {
			qrNamespaceUUID = *o.NamespaceUUIDQueryParameter
		}
		qNamespaceUUID := qrNamespaceUUID
		if qNamespaceUUID != "" {

			if err := r.SetQueryParam("namespace.uuid", qNamespaceUUID); err != nil {
				return err
			}
		}
	}

	if o.NsIDQueryParameter != nil {

		// query param nsid
		var qrNsid string

		if o.NsIDQueryParameter != nil {
			qrNsid = *o.NsIDQueryParameter
		}
		qNsid := qrNsid
		if qNsid != "" {

			if err := r.SetQueryParam("nsid", qNsid); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.SubsystemNameQueryParameter != nil {

		// query param subsystem.name
		var qrSubsystemName string

		if o.SubsystemNameQueryParameter != nil {
			qrSubsystemName = *o.SubsystemNameQueryParameter
		}
		qSubsystemName := qrSubsystemName
		if qSubsystemName != "" {

			if err := r.SetQueryParam("subsystem.name", qSubsystemName); err != nil {
				return err
			}
		}
	}

	if o.SubsystemUUIDQueryParameter != nil {

		// query param subsystem.uuid
		var qrSubsystemUUID string

		if o.SubsystemUUIDQueryParameter != nil {
			qrSubsystemUUID = *o.SubsystemUUIDQueryParameter
		}
		qSubsystemUUID := qrSubsystemUUID
		if qSubsystemUUID != "" {

			if err := r.SetQueryParam("subsystem.uuid", qSubsystemUUID); err != nil {
				return err
			}
		}
	}

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNvmeSubsystemMapCollectionGet binds the parameter fields
func (o *NvmeSubsystemMapCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNvmeSubsystemMapCollectionGet binds the parameter order_by
func (o *NvmeSubsystemMapCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}

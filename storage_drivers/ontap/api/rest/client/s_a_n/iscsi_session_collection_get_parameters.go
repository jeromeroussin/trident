// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewIscsiSessionCollectionGetParams creates a new IscsiSessionCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIscsiSessionCollectionGetParams() *IscsiSessionCollectionGetParams {
	return &IscsiSessionCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIscsiSessionCollectionGetParamsWithTimeout creates a new IscsiSessionCollectionGetParams object
// with the ability to set a timeout on a request.
func NewIscsiSessionCollectionGetParamsWithTimeout(timeout time.Duration) *IscsiSessionCollectionGetParams {
	return &IscsiSessionCollectionGetParams{
		timeout: timeout,
	}
}

// NewIscsiSessionCollectionGetParamsWithContext creates a new IscsiSessionCollectionGetParams object
// with the ability to set a context for a request.
func NewIscsiSessionCollectionGetParamsWithContext(ctx context.Context) *IscsiSessionCollectionGetParams {
	return &IscsiSessionCollectionGetParams{
		Context: ctx,
	}
}

// NewIscsiSessionCollectionGetParamsWithHTTPClient creates a new IscsiSessionCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewIscsiSessionCollectionGetParamsWithHTTPClient(client *http.Client) *IscsiSessionCollectionGetParams {
	return &IscsiSessionCollectionGetParams{
		HTTPClient: client,
	}
}

/* IscsiSessionCollectionGetParams contains all the parameters to send to the API endpoint
   for the iscsi session collection get operation.

   Typically these are written to a http.Request.
*/
type IscsiSessionCollectionGetParams struct {

	/* ConnectionsAuthenticationType.

	   Filter by connections.authentication_type
	*/
	ConnectionsAuthenticationTypeQueryParameter *string

	/* ConnectionsCid.

	   Filter by connections.cid
	*/
	ConnectionsCIDQueryParameter *int64

	/* ConnectionsInitiatorAddressAddress.

	   Filter by connections.initiator_address.address
	*/
	ConnectionsInitiatorAddressAddressQueryParameter *string

	/* ConnectionsInitiatorAddressPort.

	   Filter by connections.initiator_address.port
	*/
	ConnectionsInitiatorAddressPortQueryParameter *int64

	/* ConnectionsInterfaceIPAddress.

	   Filter by connections.interface.ip.address
	*/
	ConnectionsInterfaceIPAddressQueryParameter *string

	/* ConnectionsInterfaceIPPort.

	   Filter by connections.interface.ip.port
	*/
	ConnectionsInterfaceIPPortQueryParameter *int64

	/* ConnectionsInterfaceName.

	   Filter by connections.interface.name
	*/
	ConnectionsInterfaceNameQueryParameter *string

	/* ConnectionsInterfaceUUID.

	   Filter by connections.interface.uuid
	*/
	ConnectionsInterfaceUUIDQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* IgroupsName.

	   Filter by igroups.name
	*/
	IgroupsNameQueryParameter *string

	/* IgroupsUUID.

	   Filter by igroups.uuid
	*/
	IgroupsUUIDQueryParameter *string

	/* InitiatorAlias.

	   Filter by initiator.alias
	*/
	InitiatorAliasQueryParameter *string

	/* InitiatorComment.

	   Filter by initiator.comment
	*/
	InitiatorCommentQueryParameter *string

	/* InitiatorName.

	   Filter by initiator.name
	*/
	InitiatorNameQueryParameter *string

	/* Isid.

	   Filter by isid
	*/
	IsIDQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

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

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* TargetPortalGroup.

	   Filter by target_portal_group
	*/
	TargetPortalGroupQueryParameter *string

	/* TargetPortalGroupTag.

	   Filter by target_portal_group_tag
	*/
	TargetPortalGroupTagQueryParameter *int64

	/* Tsih.

	   Filter by tsih
	*/
	TsihQueryParameter *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the iscsi session collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiSessionCollectionGetParams) WithDefaults() *IscsiSessionCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the iscsi session collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IscsiSessionCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := IscsiSessionCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithTimeout(timeout time.Duration) *IscsiSessionCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithContext(ctx context.Context) *IscsiSessionCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithHTTPClient(client *http.Client) *IscsiSessionCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionsAuthenticationTypeQueryParameter adds the connectionsAuthenticationType to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithConnectionsAuthenticationTypeQueryParameter(connectionsAuthenticationType *string) *IscsiSessionCollectionGetParams {
	o.SetConnectionsAuthenticationTypeQueryParameter(connectionsAuthenticationType)
	return o
}

// SetConnectionsAuthenticationTypeQueryParameter adds the connectionsAuthenticationType to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetConnectionsAuthenticationTypeQueryParameter(connectionsAuthenticationType *string) {
	o.ConnectionsAuthenticationTypeQueryParameter = connectionsAuthenticationType
}

// WithConnectionsCIDQueryParameter adds the connectionsCid to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithConnectionsCIDQueryParameter(connectionsCid *int64) *IscsiSessionCollectionGetParams {
	o.SetConnectionsCIDQueryParameter(connectionsCid)
	return o
}

// SetConnectionsCIDQueryParameter adds the connectionsCid to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetConnectionsCIDQueryParameter(connectionsCid *int64) {
	o.ConnectionsCIDQueryParameter = connectionsCid
}

// WithConnectionsInitiatorAddressAddressQueryParameter adds the connectionsInitiatorAddressAddress to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithConnectionsInitiatorAddressAddressQueryParameter(connectionsInitiatorAddressAddress *string) *IscsiSessionCollectionGetParams {
	o.SetConnectionsInitiatorAddressAddressQueryParameter(connectionsInitiatorAddressAddress)
	return o
}

// SetConnectionsInitiatorAddressAddressQueryParameter adds the connectionsInitiatorAddressAddress to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetConnectionsInitiatorAddressAddressQueryParameter(connectionsInitiatorAddressAddress *string) {
	o.ConnectionsInitiatorAddressAddressQueryParameter = connectionsInitiatorAddressAddress
}

// WithConnectionsInitiatorAddressPortQueryParameter adds the connectionsInitiatorAddressPort to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithConnectionsInitiatorAddressPortQueryParameter(connectionsInitiatorAddressPort *int64) *IscsiSessionCollectionGetParams {
	o.SetConnectionsInitiatorAddressPortQueryParameter(connectionsInitiatorAddressPort)
	return o
}

// SetConnectionsInitiatorAddressPortQueryParameter adds the connectionsInitiatorAddressPort to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetConnectionsInitiatorAddressPortQueryParameter(connectionsInitiatorAddressPort *int64) {
	o.ConnectionsInitiatorAddressPortQueryParameter = connectionsInitiatorAddressPort
}

// WithConnectionsInterfaceIPAddressQueryParameter adds the connectionsInterfaceIPAddress to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithConnectionsInterfaceIPAddressQueryParameter(connectionsInterfaceIPAddress *string) *IscsiSessionCollectionGetParams {
	o.SetConnectionsInterfaceIPAddressQueryParameter(connectionsInterfaceIPAddress)
	return o
}

// SetConnectionsInterfaceIPAddressQueryParameter adds the connectionsInterfaceIpAddress to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetConnectionsInterfaceIPAddressQueryParameter(connectionsInterfaceIPAddress *string) {
	o.ConnectionsInterfaceIPAddressQueryParameter = connectionsInterfaceIPAddress
}

// WithConnectionsInterfaceIPPortQueryParameter adds the connectionsInterfaceIPPort to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithConnectionsInterfaceIPPortQueryParameter(connectionsInterfaceIPPort *int64) *IscsiSessionCollectionGetParams {
	o.SetConnectionsInterfaceIPPortQueryParameter(connectionsInterfaceIPPort)
	return o
}

// SetConnectionsInterfaceIPPortQueryParameter adds the connectionsInterfaceIpPort to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetConnectionsInterfaceIPPortQueryParameter(connectionsInterfaceIPPort *int64) {
	o.ConnectionsInterfaceIPPortQueryParameter = connectionsInterfaceIPPort
}

// WithConnectionsInterfaceNameQueryParameter adds the connectionsInterfaceName to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithConnectionsInterfaceNameQueryParameter(connectionsInterfaceName *string) *IscsiSessionCollectionGetParams {
	o.SetConnectionsInterfaceNameQueryParameter(connectionsInterfaceName)
	return o
}

// SetConnectionsInterfaceNameQueryParameter adds the connectionsInterfaceName to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetConnectionsInterfaceNameQueryParameter(connectionsInterfaceName *string) {
	o.ConnectionsInterfaceNameQueryParameter = connectionsInterfaceName
}

// WithConnectionsInterfaceUUIDQueryParameter adds the connectionsInterfaceUUID to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithConnectionsInterfaceUUIDQueryParameter(connectionsInterfaceUUID *string) *IscsiSessionCollectionGetParams {
	o.SetConnectionsInterfaceUUIDQueryParameter(connectionsInterfaceUUID)
	return o
}

// SetConnectionsInterfaceUUIDQueryParameter adds the connectionsInterfaceUuid to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetConnectionsInterfaceUUIDQueryParameter(connectionsInterfaceUUID *string) {
	o.ConnectionsInterfaceUUIDQueryParameter = connectionsInterfaceUUID
}

// WithFieldsQueryParameter adds the fields to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithFieldsQueryParameter(fields []string) *IscsiSessionCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIgroupsNameQueryParameter adds the igroupsName to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithIgroupsNameQueryParameter(igroupsName *string) *IscsiSessionCollectionGetParams {
	o.SetIgroupsNameQueryParameter(igroupsName)
	return o
}

// SetIgroupsNameQueryParameter adds the igroupsName to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetIgroupsNameQueryParameter(igroupsName *string) {
	o.IgroupsNameQueryParameter = igroupsName
}

// WithIgroupsUUIDQueryParameter adds the igroupsUUID to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithIgroupsUUIDQueryParameter(igroupsUUID *string) *IscsiSessionCollectionGetParams {
	o.SetIgroupsUUIDQueryParameter(igroupsUUID)
	return o
}

// SetIgroupsUUIDQueryParameter adds the igroupsUuid to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetIgroupsUUIDQueryParameter(igroupsUUID *string) {
	o.IgroupsUUIDQueryParameter = igroupsUUID
}

// WithInitiatorAliasQueryParameter adds the initiatorAlias to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithInitiatorAliasQueryParameter(initiatorAlias *string) *IscsiSessionCollectionGetParams {
	o.SetInitiatorAliasQueryParameter(initiatorAlias)
	return o
}

// SetInitiatorAliasQueryParameter adds the initiatorAlias to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetInitiatorAliasQueryParameter(initiatorAlias *string) {
	o.InitiatorAliasQueryParameter = initiatorAlias
}

// WithInitiatorCommentQueryParameter adds the initiatorComment to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithInitiatorCommentQueryParameter(initiatorComment *string) *IscsiSessionCollectionGetParams {
	o.SetInitiatorCommentQueryParameter(initiatorComment)
	return o
}

// SetInitiatorCommentQueryParameter adds the initiatorComment to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetInitiatorCommentQueryParameter(initiatorComment *string) {
	o.InitiatorCommentQueryParameter = initiatorComment
}

// WithInitiatorNameQueryParameter adds the initiatorName to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithInitiatorNameQueryParameter(initiatorName *string) *IscsiSessionCollectionGetParams {
	o.SetInitiatorNameQueryParameter(initiatorName)
	return o
}

// SetInitiatorNameQueryParameter adds the initiatorName to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetInitiatorNameQueryParameter(initiatorName *string) {
	o.InitiatorNameQueryParameter = initiatorName
}

// WithIsIDQueryParameter adds the isid to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithIsIDQueryParameter(isid *string) *IscsiSessionCollectionGetParams {
	o.SetIsIDQueryParameter(isid)
	return o
}

// SetIsIDQueryParameter adds the isid to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetIsIDQueryParameter(isid *string) {
	o.IsIDQueryParameter = isid
}

// WithMaxRecordsQueryParameter adds the maxRecords to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *IscsiSessionCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *IscsiSessionCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *IscsiSessionCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *IscsiSessionCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *IscsiSessionCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *IscsiSessionCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithTargetPortalGroupQueryParameter adds the targetPortalGroup to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithTargetPortalGroupQueryParameter(targetPortalGroup *string) *IscsiSessionCollectionGetParams {
	o.SetTargetPortalGroupQueryParameter(targetPortalGroup)
	return o
}

// SetTargetPortalGroupQueryParameter adds the targetPortalGroup to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetTargetPortalGroupQueryParameter(targetPortalGroup *string) {
	o.TargetPortalGroupQueryParameter = targetPortalGroup
}

// WithTargetPortalGroupTagQueryParameter adds the targetPortalGroupTag to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithTargetPortalGroupTagQueryParameter(targetPortalGroupTag *int64) *IscsiSessionCollectionGetParams {
	o.SetTargetPortalGroupTagQueryParameter(targetPortalGroupTag)
	return o
}

// SetTargetPortalGroupTagQueryParameter adds the targetPortalGroupTag to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetTargetPortalGroupTagQueryParameter(targetPortalGroupTag *int64) {
	o.TargetPortalGroupTagQueryParameter = targetPortalGroupTag
}

// WithTsihQueryParameter adds the tsih to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) WithTsihQueryParameter(tsih *int64) *IscsiSessionCollectionGetParams {
	o.SetTsihQueryParameter(tsih)
	return o
}

// SetTsihQueryParameter adds the tsih to the iscsi session collection get params
func (o *IscsiSessionCollectionGetParams) SetTsihQueryParameter(tsih *int64) {
	o.TsihQueryParameter = tsih
}

// WriteToRequest writes these params to a swagger request
func (o *IscsiSessionCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConnectionsAuthenticationTypeQueryParameter != nil {

		// query param connections.authentication_type
		var qrConnectionsAuthenticationType string

		if o.ConnectionsAuthenticationTypeQueryParameter != nil {
			qrConnectionsAuthenticationType = *o.ConnectionsAuthenticationTypeQueryParameter
		}
		qConnectionsAuthenticationType := qrConnectionsAuthenticationType
		if qConnectionsAuthenticationType != "" {

			if err := r.SetQueryParam("connections.authentication_type", qConnectionsAuthenticationType); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsCIDQueryParameter != nil {

		// query param connections.cid
		var qrConnectionsCid int64

		if o.ConnectionsCIDQueryParameter != nil {
			qrConnectionsCid = *o.ConnectionsCIDQueryParameter
		}
		qConnectionsCid := swag.FormatInt64(qrConnectionsCid)
		if qConnectionsCid != "" {

			if err := r.SetQueryParam("connections.cid", qConnectionsCid); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsInitiatorAddressAddressQueryParameter != nil {

		// query param connections.initiator_address.address
		var qrConnectionsInitiatorAddressAddress string

		if o.ConnectionsInitiatorAddressAddressQueryParameter != nil {
			qrConnectionsInitiatorAddressAddress = *o.ConnectionsInitiatorAddressAddressQueryParameter
		}
		qConnectionsInitiatorAddressAddress := qrConnectionsInitiatorAddressAddress
		if qConnectionsInitiatorAddressAddress != "" {

			if err := r.SetQueryParam("connections.initiator_address.address", qConnectionsInitiatorAddressAddress); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsInitiatorAddressPortQueryParameter != nil {

		// query param connections.initiator_address.port
		var qrConnectionsInitiatorAddressPort int64

		if o.ConnectionsInitiatorAddressPortQueryParameter != nil {
			qrConnectionsInitiatorAddressPort = *o.ConnectionsInitiatorAddressPortQueryParameter
		}
		qConnectionsInitiatorAddressPort := swag.FormatInt64(qrConnectionsInitiatorAddressPort)
		if qConnectionsInitiatorAddressPort != "" {

			if err := r.SetQueryParam("connections.initiator_address.port", qConnectionsInitiatorAddressPort); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsInterfaceIPAddressQueryParameter != nil {

		// query param connections.interface.ip.address
		var qrConnectionsInterfaceIPAddress string

		if o.ConnectionsInterfaceIPAddressQueryParameter != nil {
			qrConnectionsInterfaceIPAddress = *o.ConnectionsInterfaceIPAddressQueryParameter
		}
		qConnectionsInterfaceIPAddress := qrConnectionsInterfaceIPAddress
		if qConnectionsInterfaceIPAddress != "" {

			if err := r.SetQueryParam("connections.interface.ip.address", qConnectionsInterfaceIPAddress); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsInterfaceIPPortQueryParameter != nil {

		// query param connections.interface.ip.port
		var qrConnectionsInterfaceIPPort int64

		if o.ConnectionsInterfaceIPPortQueryParameter != nil {
			qrConnectionsInterfaceIPPort = *o.ConnectionsInterfaceIPPortQueryParameter
		}
		qConnectionsInterfaceIPPort := swag.FormatInt64(qrConnectionsInterfaceIPPort)
		if qConnectionsInterfaceIPPort != "" {

			if err := r.SetQueryParam("connections.interface.ip.port", qConnectionsInterfaceIPPort); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsInterfaceNameQueryParameter != nil {

		// query param connections.interface.name
		var qrConnectionsInterfaceName string

		if o.ConnectionsInterfaceNameQueryParameter != nil {
			qrConnectionsInterfaceName = *o.ConnectionsInterfaceNameQueryParameter
		}
		qConnectionsInterfaceName := qrConnectionsInterfaceName
		if qConnectionsInterfaceName != "" {

			if err := r.SetQueryParam("connections.interface.name", qConnectionsInterfaceName); err != nil {
				return err
			}
		}
	}

	if o.ConnectionsInterfaceUUIDQueryParameter != nil {

		// query param connections.interface.uuid
		var qrConnectionsInterfaceUUID string

		if o.ConnectionsInterfaceUUIDQueryParameter != nil {
			qrConnectionsInterfaceUUID = *o.ConnectionsInterfaceUUIDQueryParameter
		}
		qConnectionsInterfaceUUID := qrConnectionsInterfaceUUID
		if qConnectionsInterfaceUUID != "" {

			if err := r.SetQueryParam("connections.interface.uuid", qConnectionsInterfaceUUID); err != nil {
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

	if o.IgroupsNameQueryParameter != nil {

		// query param igroups.name
		var qrIgroupsName string

		if o.IgroupsNameQueryParameter != nil {
			qrIgroupsName = *o.IgroupsNameQueryParameter
		}
		qIgroupsName := qrIgroupsName
		if qIgroupsName != "" {

			if err := r.SetQueryParam("igroups.name", qIgroupsName); err != nil {
				return err
			}
		}
	}

	if o.IgroupsUUIDQueryParameter != nil {

		// query param igroups.uuid
		var qrIgroupsUUID string

		if o.IgroupsUUIDQueryParameter != nil {
			qrIgroupsUUID = *o.IgroupsUUIDQueryParameter
		}
		qIgroupsUUID := qrIgroupsUUID
		if qIgroupsUUID != "" {

			if err := r.SetQueryParam("igroups.uuid", qIgroupsUUID); err != nil {
				return err
			}
		}
	}

	if o.InitiatorAliasQueryParameter != nil {

		// query param initiator.alias
		var qrInitiatorAlias string

		if o.InitiatorAliasQueryParameter != nil {
			qrInitiatorAlias = *o.InitiatorAliasQueryParameter
		}
		qInitiatorAlias := qrInitiatorAlias
		if qInitiatorAlias != "" {

			if err := r.SetQueryParam("initiator.alias", qInitiatorAlias); err != nil {
				return err
			}
		}
	}

	if o.InitiatorCommentQueryParameter != nil {

		// query param initiator.comment
		var qrInitiatorComment string

		if o.InitiatorCommentQueryParameter != nil {
			qrInitiatorComment = *o.InitiatorCommentQueryParameter
		}
		qInitiatorComment := qrInitiatorComment
		if qInitiatorComment != "" {

			if err := r.SetQueryParam("initiator.comment", qInitiatorComment); err != nil {
				return err
			}
		}
	}

	if o.InitiatorNameQueryParameter != nil {

		// query param initiator.name
		var qrInitiatorName string

		if o.InitiatorNameQueryParameter != nil {
			qrInitiatorName = *o.InitiatorNameQueryParameter
		}
		qInitiatorName := qrInitiatorName
		if qInitiatorName != "" {

			if err := r.SetQueryParam("initiator.name", qInitiatorName); err != nil {
				return err
			}
		}
	}

	if o.IsIDQueryParameter != nil {

		// query param isid
		var qrIsid string

		if o.IsIDQueryParameter != nil {
			qrIsid = *o.IsIDQueryParameter
		}
		qIsid := qrIsid
		if qIsid != "" {

			if err := r.SetQueryParam("isid", qIsid); err != nil {
				return err
			}
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

	if o.TargetPortalGroupQueryParameter != nil {

		// query param target_portal_group
		var qrTargetPortalGroup string

		if o.TargetPortalGroupQueryParameter != nil {
			qrTargetPortalGroup = *o.TargetPortalGroupQueryParameter
		}
		qTargetPortalGroup := qrTargetPortalGroup
		if qTargetPortalGroup != "" {

			if err := r.SetQueryParam("target_portal_group", qTargetPortalGroup); err != nil {
				return err
			}
		}
	}

	if o.TargetPortalGroupTagQueryParameter != nil {

		// query param target_portal_group_tag
		var qrTargetPortalGroupTag int64

		if o.TargetPortalGroupTagQueryParameter != nil {
			qrTargetPortalGroupTag = *o.TargetPortalGroupTagQueryParameter
		}
		qTargetPortalGroupTag := swag.FormatInt64(qrTargetPortalGroupTag)
		if qTargetPortalGroupTag != "" {

			if err := r.SetQueryParam("target_portal_group_tag", qTargetPortalGroupTag); err != nil {
				return err
			}
		}
	}

	if o.TsihQueryParameter != nil {

		// query param tsih
		var qrTsih int64

		if o.TsihQueryParameter != nil {
			qrTsih = *o.TsihQueryParameter
		}
		qTsih := swag.FormatInt64(qrTsih)
		if qTsih != "" {

			if err := r.SetQueryParam("tsih", qTsih); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamIscsiSessionCollectionGet binds the parameter fields
func (o *IscsiSessionCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamIscsiSessionCollectionGet binds the parameter order_by
func (o *IscsiSessionCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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

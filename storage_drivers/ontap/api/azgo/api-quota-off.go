// Code generated automatically. DO NOT EDIT.
package azgo

import (
	"encoding/xml"
	"reflect"

	log "github.com/sirupsen/logrus"
)

// QuotaOffRequest is a structure to represent a quota-off Request ZAPI object
type QuotaOffRequest struct {
	XMLName   xml.Name `xml:"quota-off"`
	VolumePtr *string  `xml:"volume"`
}

// QuotaOffResponse is a structure to represent a quota-off Response ZAPI object
type QuotaOffResponse struct {
	XMLName         xml.Name               `xml:"netapp"`
	ResponseVersion string                 `xml:"version,attr"`
	ResponseXmlns   string                 `xml:"xmlns,attr"`
	Result          QuotaOffResponseResult `xml:"results"`
}

// NewQuotaOffResponse is a factory method for creating new instances of QuotaOffResponse objects
func NewQuotaOffResponse() *QuotaOffResponse {
	return &QuotaOffResponse{}
}

// String returns a string representation of this object's fields and implements the Stringer interface
func (o QuotaOffResponse) String() string {
	return ToString(reflect.ValueOf(o))
}

// ToXML converts this object into an xml string representation
func (o *QuotaOffResponse) ToXML() (string, error) {
	output, err := xml.MarshalIndent(o, " ", "    ")
	if err != nil {
		log.Errorf("error: %v", err)
	}
	return string(output), err
}

// QuotaOffResponseResult is a structure to represent a quota-off Response Result ZAPI object
type QuotaOffResponseResult struct {
	XMLName               xml.Name `xml:"results"`
	ResultStatusAttr      string   `xml:"status,attr"`
	ResultReasonAttr      string   `xml:"reason,attr"`
	ResultErrnoAttr       string   `xml:"errno,attr"`
	ResultErrorCodePtr    *int     `xml:"result-error-code"`
	ResultErrorMessagePtr *string  `xml:"result-error-message"`
	ResultJobidPtr        *int     `xml:"result-jobid"`
	ResultStatusPtr       *string  `xml:"result-status"`
}

// NewQuotaOffRequest is a factory method for creating new instances of QuotaOffRequest objects
func NewQuotaOffRequest() *QuotaOffRequest {
	return &QuotaOffRequest{}
}

// NewQuotaOffResponseResult is a factory method for creating new instances of QuotaOffResponseResult objects
func NewQuotaOffResponseResult() *QuotaOffResponseResult {
	return &QuotaOffResponseResult{}
}

// ToXML converts this object into an xml string representation
func (o *QuotaOffRequest) ToXML() (string, error) {
	output, err := xml.MarshalIndent(o, " ", "    ")
	if err != nil {
		log.Errorf("error: %v", err)
	}
	return string(output), err
}

// ToXML converts this object into an xml string representation
func (o *QuotaOffResponseResult) ToXML() (string, error) {
	output, err := xml.MarshalIndent(o, " ", "    ")
	if err != nil {
		log.Errorf("error: %v", err)
	}
	return string(output), err
}

// String returns a string representation of this object's fields and implements the Stringer interface
func (o QuotaOffRequest) String() string {
	return ToString(reflect.ValueOf(o))
}

// String returns a string representation of this object's fields and implements the Stringer interface
func (o QuotaOffResponseResult) String() string {
	return ToString(reflect.ValueOf(o))
}

// ExecuteUsing converts this object to a ZAPI XML representation and uses the supplied ZapiRunner to send to a filer

func (o *QuotaOffRequest) ExecuteUsing(zr *ZapiRunner) (*QuotaOffResponse, error) {
	return o.executeWithoutIteration(zr)
}

// executeWithoutIteration converts this object to a ZAPI XML representation and uses the supplied ZapiRunner to send to a filer

func (o *QuotaOffRequest) executeWithoutIteration(zr *ZapiRunner) (*QuotaOffResponse, error) {
	result, err := zr.ExecuteUsing(o, "QuotaOffRequest", NewQuotaOffResponse())
	if result == nil {
		return nil, err
	}
	return result.(*QuotaOffResponse), err
}

// Volume is a 'getter' method
func (o *QuotaOffRequest) Volume() string {
	r := *o.VolumePtr
	return r
}

// SetVolume is a fluent style 'setter' method that can be chained
func (o *QuotaOffRequest) SetVolume(newValue string) *QuotaOffRequest {
	o.VolumePtr = &newValue
	return o
}

// ResultErrorCode is a 'getter' method
func (o *QuotaOffResponseResult) ResultErrorCode() int {
	r := *o.ResultErrorCodePtr
	return r
}

// SetResultErrorCode is a fluent style 'setter' method that can be chained
func (o *QuotaOffResponseResult) SetResultErrorCode(newValue int) *QuotaOffResponseResult {
	o.ResultErrorCodePtr = &newValue
	return o
}

// ResultErrorMessage is a 'getter' method
func (o *QuotaOffResponseResult) ResultErrorMessage() string {
	r := *o.ResultErrorMessagePtr
	return r
}

// SetResultErrorMessage is a fluent style 'setter' method that can be chained
func (o *QuotaOffResponseResult) SetResultErrorMessage(newValue string) *QuotaOffResponseResult {
	o.ResultErrorMessagePtr = &newValue
	return o
}

// ResultJobid is a 'getter' method
func (o *QuotaOffResponseResult) ResultJobid() int {
	r := *o.ResultJobidPtr
	return r
}

// SetResultJobid is a fluent style 'setter' method that can be chained
func (o *QuotaOffResponseResult) SetResultJobid(newValue int) *QuotaOffResponseResult {
	o.ResultJobidPtr = &newValue
	return o
}

// ResultStatus is a 'getter' method
func (o *QuotaOffResponseResult) ResultStatus() string {
	r := *o.ResultStatusPtr
	return r
}

// SetResultStatus is a fluent style 'setter' method that can be chained
func (o *QuotaOffResponseResult) SetResultStatus(newValue string) *QuotaOffResponseResult {
	o.ResultStatusPtr = &newValue
	return o
}

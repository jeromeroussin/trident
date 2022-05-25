// Code generated automatically. DO NOT EDIT.
// Copyright 2017 NetApp, Inc. All Rights Reserved.

package azgo

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

// LunMapGetIterRequest is a structure to represent a lun-map-get-iter ZAPI request object
type LunMapGetIterRequest struct {
	XMLName xml.Name `xml:"lun-map-get-iter"`

	DesiredAttributesPtr *LunMapInfoType `xml:"desired-attributes>lun-map-info"`
	MaxRecordsPtr        *int            `xml:"max-records"`
	QueryPtr             *LunMapInfoType `xml:"query>lun-map-info"`
	TagPtr               *string         `xml:"tag"`
}

// ToXML converts this object into an xml string representation
func (o *LunMapGetIterRequest) ToXML() (string, error) {
	output, err := xml.MarshalIndent(o, " ", "    ")
	//if err != nil { log.Errorf("error: %v\n", err) }
	return string(output), err
}

// NewLunMapGetIterRequest is a factory method for creating new instances of LunMapGetIterRequest objects
func NewLunMapGetIterRequest() *LunMapGetIterRequest { return &LunMapGetIterRequest{} }

// ExecuteUsing converts this object to a ZAPI XML representation and uses the supplied ZapiRunner to send to a filer
func (o *LunMapGetIterRequest) ExecuteUsing(zr *ZapiRunner) (LunMapGetIterResponse, error) {

	if zr.DebugTraceFlags["method"] {
		fields := log.Fields{"Method": "ExecuteUsing", "Type": "LunMapGetIterRequest"}
		log.WithFields(fields).Debug(">>>> ExecuteUsing")
		defer log.WithFields(fields).Debug("<<<< ExecuteUsing")
	}

	combined := NewLunMapGetIterResponse()
	var nextTagPtr *string
	done := false
	for !done {

		resp, err := zr.SendZapi(o)
		if err != nil {
			log.Errorf("API invocation failed. %v", err.Error())
			return *combined, err
		}
		defer resp.Body.Close()
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			log.Errorf("Error reading response body. %v", readErr.Error())
			return *combined, readErr
		}
		if zr.DebugTraceFlags["api"] {
			log.Debugf("response Body:\n%s", string(body))
		}

		var n LunMapGetIterResponse
		unmarshalErr := xml.Unmarshal(body, &n)
		if unmarshalErr != nil {
			log.WithField("body", string(body)).Warnf("Error unmarshaling response body. %v", unmarshalErr.Error())
			//return *combined, unmarshalErr
		}
		if zr.DebugTraceFlags["api"] {
			log.Debugf("lun-map-get-iter result:\n%s", n.Result)
		}

		if err == nil {
			nextTagPtr = n.Result.NextTagPtr
			if nextTagPtr == nil {
				done = true
			} else {
				o.SetTag(*nextTagPtr)
			}

			if n.Result.NumRecordsPtr == nil {
				done = true
			} else {
				recordsRead := n.Result.NumRecords()
				if recordsRead == 0 {
					done = true
				}
			}

			if n.Result.AttributesListPtr != nil {
				combined.Result.SetAttributesList(append(combined.Result.AttributesList(), n.Result.AttributesList()...))
			}

			if done {
				combined.Result.ResultErrnoAttr = n.Result.ResultErrnoAttr
				combined.Result.ResultReasonAttr = n.Result.ResultReasonAttr
				combined.Result.ResultStatusAttr = n.Result.ResultStatusAttr
				combined.Result.SetNumRecords(len(combined.Result.AttributesList()))
			}
		}
	}

	return *combined, nil
}

// String returns a string representation of this object's fields and implements the Stringer interface
func (o LunMapGetIterRequest) String() string {
	var buffer bytes.Buffer
	if o.DesiredAttributesPtr != nil {
		buffer.WriteString(fmt.Sprintf("%s: %v\n", "desired-attributes", *o.DesiredAttributesPtr))
	} else {
		buffer.WriteString("desired-attributes: nil\n")
	}
	if o.MaxRecordsPtr != nil {
		buffer.WriteString(fmt.Sprintf("%s: %v\n", "max-records", *o.MaxRecordsPtr))
	} else {
		buffer.WriteString("max-records: nil\n")
	}
	if o.QueryPtr != nil {
		buffer.WriteString(fmt.Sprintf("%s: %v\n", "query", *o.QueryPtr))
	} else {
		buffer.WriteString("query: nil\n")
	}
	if o.TagPtr != nil {
		buffer.WriteString(fmt.Sprintf("%s: %v\n", "tag", *o.TagPtr))
	} else {
		buffer.WriteString("tag: nil\n")
	}
	return buffer.String()
}

// DesiredAttributes is a fluent style 'getter' method that can be chained
func (o *LunMapGetIterRequest) DesiredAttributes() LunMapInfoType {
	r := *o.DesiredAttributesPtr
	return r
}

// SetDesiredAttributes is a fluent style 'setter' method that can be chained
func (o *LunMapGetIterRequest) SetDesiredAttributes(newValue LunMapInfoType) *LunMapGetIterRequest {
	o.DesiredAttributesPtr = &newValue
	return o
}

// MaxRecords is a fluent style 'getter' method that can be chained
func (o *LunMapGetIterRequest) MaxRecords() int {
	r := *o.MaxRecordsPtr
	return r
}

// SetMaxRecords is a fluent style 'setter' method that can be chained
func (o *LunMapGetIterRequest) SetMaxRecords(newValue int) *LunMapGetIterRequest {
	o.MaxRecordsPtr = &newValue
	return o
}

// Query is a fluent style 'getter' method that can be chained
func (o *LunMapGetIterRequest) Query() LunMapInfoType {
	r := *o.QueryPtr
	return r
}

// SetQuery is a fluent style 'setter' method that can be chained
func (o *LunMapGetIterRequest) SetQuery(newValue LunMapInfoType) *LunMapGetIterRequest {
	o.QueryPtr = &newValue
	return o
}

// Tag is a fluent style 'getter' method that can be chained
func (o *LunMapGetIterRequest) Tag() string {
	r := *o.TagPtr
	return r
}

// SetTag is a fluent style 'setter' method that can be chained
func (o *LunMapGetIterRequest) SetTag(newValue string) *LunMapGetIterRequest {
	o.TagPtr = &newValue
	return o
}

// LunMapGetIterResponse is a structure to represent a lun-map-get-iter ZAPI response object
type LunMapGetIterResponse struct {
	XMLName xml.Name `xml:"netapp"`

	ResponseVersion string `xml:"version,attr"`
	ResponseXmlns   string `xml:"xmlns,attr"`

	Result LunMapGetIterResponseResult `xml:"results"`
}

// String returns a string representation of this object's fields and implements the Stringer interface
func (o LunMapGetIterResponse) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s: %s\n", "version", o.ResponseVersion))
	buffer.WriteString(fmt.Sprintf("%s: %s\n", "xmlns", o.ResponseXmlns))
	buffer.WriteString(fmt.Sprintf("%s: %s\n", "results", o.Result))
	return buffer.String()
}

// LunMapGetIterResponseResult is a structure to represent a lun-map-get-iter ZAPI object's result
type LunMapGetIterResponseResult struct {
	XMLName xml.Name `xml:"results"`

	ResultStatusAttr  string            `xml:"status,attr"`
	ResultReasonAttr  string            `xml:"reason,attr"`
	ResultErrnoAttr   string            `xml:"errno,attr"`
	AttributesListPtr []LunMapInfoType  `xml:"attributes-list>lun-map-info"`
	NextTagPtr        *string           `xml:"next-tag"`
	NumRecordsPtr     *int              `xml:"num-records"`
	VolumeErrorsPtr   []VolumeErrorType `xml:"volume-errors>volume-error"`
}

// ToXML converts this object into an xml string representation
func (o *LunMapGetIterResponse) ToXML() (string, error) {
	output, err := xml.MarshalIndent(o, " ", "    ")
	//if err != nil { log.Debugf("error: %v", err) }
	return string(output), err
}

// NewLunMapGetIterResponse is a factory method for creating new instances of LunMapGetIterResponse objects
func NewLunMapGetIterResponse() *LunMapGetIterResponse { return &LunMapGetIterResponse{} }

// String returns a string representation of this object's fields and implements the Stringer interface
func (o LunMapGetIterResponseResult) String() string {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s: %s\n", "resultStatusAttr", o.ResultStatusAttr))
	buffer.WriteString(fmt.Sprintf("%s: %s\n", "resultReasonAttr", o.ResultReasonAttr))
	buffer.WriteString(fmt.Sprintf("%s: %s\n", "resultErrnoAttr", o.ResultErrnoAttr))
	if o.AttributesListPtr != nil {
		buffer.WriteString(fmt.Sprintf("%s: %v\n", "attributes-list", o.AttributesListPtr))
	} else {
		buffer.WriteString("attributes-list: nil\n")
	}
	if o.NextTagPtr != nil {
		buffer.WriteString(fmt.Sprintf("%s: %v\n", "next-tag", *o.NextTagPtr))
	} else {
		buffer.WriteString("next-tag: nil\n")
	}
	if o.NumRecordsPtr != nil {
		buffer.WriteString(fmt.Sprintf("%s: %v\n", "num-records", *o.NumRecordsPtr))
	} else {
		buffer.WriteString("num-records: nil\n")
	}
	if o.VolumeErrorsPtr != nil {
		buffer.WriteString(fmt.Sprintf("%s: %v\n", "volume-errors", o.VolumeErrorsPtr))
	} else {
		buffer.WriteString("volume-errors: nil\n")
	}
	return buffer.String()
}

// AttributesList is a fluent style 'getter' method that can be chained
func (o *LunMapGetIterResponseResult) AttributesList() []LunMapInfoType {
	r := o.AttributesListPtr
	return r
}

// SetAttributesList is a fluent style 'setter' method that can be chained
func (o *LunMapGetIterResponseResult) SetAttributesList(newValue []LunMapInfoType) *LunMapGetIterResponseResult {
	newSlice := make([]LunMapInfoType, len(newValue))
	copy(newSlice, newValue)
	o.AttributesListPtr = newSlice
	return o
}

// NextTag is a fluent style 'getter' method that can be chained
func (o *LunMapGetIterResponseResult) NextTag() string {
	r := *o.NextTagPtr
	return r
}

// SetNextTag is a fluent style 'setter' method that can be chained
func (o *LunMapGetIterResponseResult) SetNextTag(newValue string) *LunMapGetIterResponseResult {
	o.NextTagPtr = &newValue
	return o
}

// NumRecords is a fluent style 'getter' method that can be chained
func (o *LunMapGetIterResponseResult) NumRecords() int {
	r := *o.NumRecordsPtr
	return r
}

// SetNumRecords is a fluent style 'setter' method that can be chained
func (o *LunMapGetIterResponseResult) SetNumRecords(newValue int) *LunMapGetIterResponseResult {
	o.NumRecordsPtr = &newValue
	return o
}

// VolumeErrors is a fluent style 'getter' method that can be chained
func (o *LunMapGetIterResponseResult) VolumeErrors() []VolumeErrorType {
	r := o.VolumeErrorsPtr
	return r
}

// SetVolumeErrors is a fluent style 'setter' method that can be chained
func (o *LunMapGetIterResponseResult) SetVolumeErrors(newValue []VolumeErrorType) *LunMapGetIterResponseResult {
	newSlice := make([]VolumeErrorType, len(newValue))
	copy(newSlice, newValue)
	o.VolumeErrorsPtr = newSlice
	return o
}

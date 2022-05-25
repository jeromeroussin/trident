// Code generated automatically. DO NOT EDIT.
package azgo

import (
	"encoding/xml"
	"reflect"

	log "github.com/sirupsen/logrus"
)

// VserverPeerInfoType is a structure to represent a vserver-peer-info ZAPI object
type VserverPeerInfoType struct {
	XMLName         xml.Name                         `xml:"vserver-peer-info"`
	ApplicationsPtr *VserverPeerInfoTypeApplications `xml:"applications"`
	// work in progress
	PeerClusterPtr       *string               `xml:"peer-cluster"`
	PeerStatePtr         *VserverPeerStateType `xml:"peer-state"`
	PeerVserverPtr       *string               `xml:"peer-vserver"`
	PeerVserverUuidPtr   *UuidType             `xml:"peer-vserver-uuid"`
	RemoteVserverNamePtr *string               `xml:"remote-vserver-name"`
	VserverPtr           *string               `xml:"vserver"`
	VserverUuidPtr       *UuidType             `xml:"vserver-uuid"`
}

// NewVserverPeerInfoType is a factory method for creating new instances of VserverPeerInfoType objects
func NewVserverPeerInfoType() *VserverPeerInfoType {
	return &VserverPeerInfoType{}
}

// ToXML converts this object into an xml string representation
func (o *VserverPeerInfoType) ToXML() (string, error) {
	output, err := xml.MarshalIndent(o, " ", "    ")
	if err != nil {
		log.Errorf("error: %v", err)
	}
	return string(output), err
}

// String returns a string representation of this object's fields and implements the Stringer interface
func (o VserverPeerInfoType) String() string {
	return ToString(reflect.ValueOf(o))
}

// VserverPeerInfoTypeApplications is a wrapper
type VserverPeerInfoTypeApplications struct {
	XMLName                   xml.Name                     `xml:"applications"`
	VserverPeerApplicationPtr []VserverPeerApplicationType `xml:"vserver-peer-application"`
}

// VserverPeerApplication is a 'getter' method
func (o *VserverPeerInfoTypeApplications) VserverPeerApplication() []VserverPeerApplicationType {
	r := o.VserverPeerApplicationPtr
	return r
}

// SetVserverPeerApplication is a fluent style 'setter' method that can be chained
func (o *VserverPeerInfoTypeApplications) SetVserverPeerApplication(newValue []VserverPeerApplicationType) *VserverPeerInfoTypeApplications {
	newSlice := make([]VserverPeerApplicationType, len(newValue))
	copy(newSlice, newValue)
	o.VserverPeerApplicationPtr = newSlice
	return o
}

// Applications is a 'getter' method
func (o *VserverPeerInfoType) Applications() VserverPeerInfoTypeApplications {
	r := *o.ApplicationsPtr
	return r
}

// SetApplications is a fluent style 'setter' method that can be chained
func (o *VserverPeerInfoType) SetApplications(newValue VserverPeerInfoTypeApplications) *VserverPeerInfoType {
	o.ApplicationsPtr = &newValue
	return o
}

// PeerCluster is a 'getter' method
func (o *VserverPeerInfoType) PeerCluster() string {
	r := *o.PeerClusterPtr
	return r
}

// SetPeerCluster is a fluent style 'setter' method that can be chained
func (o *VserverPeerInfoType) SetPeerCluster(newValue string) *VserverPeerInfoType {
	o.PeerClusterPtr = &newValue
	return o
}

// PeerState is a 'getter' method
func (o *VserverPeerInfoType) PeerState() VserverPeerStateType {
	r := *o.PeerStatePtr
	return r
}

// SetPeerState is a fluent style 'setter' method that can be chained
func (o *VserverPeerInfoType) SetPeerState(newValue VserverPeerStateType) *VserverPeerInfoType {
	o.PeerStatePtr = &newValue
	return o
}

// PeerVserver is a 'getter' method
func (o *VserverPeerInfoType) PeerVserver() string {
	r := *o.PeerVserverPtr
	return r
}

// SetPeerVserver is a fluent style 'setter' method that can be chained
func (o *VserverPeerInfoType) SetPeerVserver(newValue string) *VserverPeerInfoType {
	o.PeerVserverPtr = &newValue
	return o
}

// PeerVserverUuid is a 'getter' method
func (o *VserverPeerInfoType) PeerVserverUuid() UuidType {
	r := *o.PeerVserverUuidPtr
	return r
}

// SetPeerVserverUuid is a fluent style 'setter' method that can be chained
func (o *VserverPeerInfoType) SetPeerVserverUuid(newValue UuidType) *VserverPeerInfoType {
	o.PeerVserverUuidPtr = &newValue
	return o
}

// RemoteVserverName is a 'getter' method
func (o *VserverPeerInfoType) RemoteVserverName() string {
	r := *o.RemoteVserverNamePtr
	return r
}

// SetRemoteVserverName is a fluent style 'setter' method that can be chained
func (o *VserverPeerInfoType) SetRemoteVserverName(newValue string) *VserverPeerInfoType {
	o.RemoteVserverNamePtr = &newValue
	return o
}

// Vserver is a 'getter' method
func (o *VserverPeerInfoType) Vserver() string {
	r := *o.VserverPtr
	return r
}

// SetVserver is a fluent style 'setter' method that can be chained
func (o *VserverPeerInfoType) SetVserver(newValue string) *VserverPeerInfoType {
	o.VserverPtr = &newValue
	return o
}

// VserverUuid is a 'getter' method
func (o *VserverPeerInfoType) VserverUuid() UuidType {
	r := *o.VserverUuidPtr
	return r
}

// SetVserverUuid is a fluent style 'setter' method that can be chained
func (o *VserverPeerInfoType) SetVserverUuid(newValue UuidType) *VserverPeerInfoType {
	o.VserverUuidPtr = &newValue
	return o
}

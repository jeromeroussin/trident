// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NdmpAuthType The different interface roles.
//
// swagger:model ndmp_auth_type
type NdmpAuthType string

const (

	// NdmpAuthTypePlaintext captures enum value "plaintext"
	NdmpAuthTypePlaintext NdmpAuthType = "plaintext"

	// NdmpAuthTypeChallenge captures enum value "challenge"
	NdmpAuthTypeChallenge NdmpAuthType = "challenge"

	// NdmpAuthTypePlaintextSso captures enum value "plaintext_sso"
	NdmpAuthTypePlaintextSso NdmpAuthType = "plaintext_sso"
)

// for schema
var ndmpAuthTypeEnum []interface{}

func init() {
	var res []NdmpAuthType
	if err := json.Unmarshal([]byte(`["plaintext","challenge","plaintext_sso"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		ndmpAuthTypeEnum = append(ndmpAuthTypeEnum, v)
	}
}

func (m NdmpAuthType) validateNdmpAuthTypeEnum(path, location string, value NdmpAuthType) error {
	if err := validate.EnumCase(path, location, value, ndmpAuthTypeEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this ndmp auth type
func (m NdmpAuthType) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateNdmpAuthTypeEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this ndmp auth type based on context it is used
func (m NdmpAuthType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

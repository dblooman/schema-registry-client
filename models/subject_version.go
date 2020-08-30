// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SubjectVersion subject version
//
// swagger:model SubjectVersion
type SubjectVersion struct {

	// subject
	Subject string `json:"subject,omitempty"`

	// version
	Version int32 `json:"version,omitempty"`
}

// Validate validates this subject version
func (m *SubjectVersion) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SubjectVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SubjectVersion) UnmarshalBinary(b []byte) error {
	var res SubjectVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

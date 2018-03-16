// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixCreateJobResponse openpitrix create job response
// swagger:model openpitrixCreateJobResponse
type OpenpitrixCreateJobResponse struct {

	// app id
	AppID string `json:"app_id,omitempty"`

	// cluster id
	ClusterID string `json:"cluster_id,omitempty"`

	// job id
	JobID string `json:"job_id,omitempty"`

	// version id
	VersionID string `json:"version_id,omitempty"`
}

// Validate validates this openpitrix create job response
func (m *OpenpitrixCreateJobResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixCreateJobResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixCreateJobResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixCreateJobResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

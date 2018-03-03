// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// OpenpitrixGetTaskStatusResponse openpitrix get task status response
// swagger:model openpitrixGetTaskStatusResponse
type OpenpitrixGetTaskStatusResponse struct {

	// task status set
	TaskStatusSet OpenpitrixGetTaskStatusResponseTaskStatusSet `json:"task_status_set"`

	// total count
	TotalCount int64 `json:"total_count,omitempty"`
}

// Validate validates this openpitrix get task status response
func (m *OpenpitrixGetTaskStatusResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *OpenpitrixGetTaskStatusResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *OpenpitrixGetTaskStatusResponse) UnmarshalBinary(b []byte) error {
	var res OpenpitrixGetTaskStatusResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

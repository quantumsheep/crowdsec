// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// MetricsMeta MetricsMeta
//
// swagger:model MetricsMeta
type MetricsMeta struct {

	// UTC timestamp of the current time
	UtcNowTimestamp int64 `json:"utc_now_timestamp,omitempty"`

	// UTC timestamp of the startup of the software
	UtcStartupTimestamp int64 `json:"utc_startup_timestamp,omitempty"`

	// Size, in seconds, of the window used to compute the metric
	WindowSizeSeconds int64 `json:"window_size_seconds,omitempty"`
}

// Validate validates this metrics meta
func (m *MetricsMeta) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metrics meta based on context it is used
func (m *MetricsMeta) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *MetricsMeta) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetricsMeta) UnmarshalBinary(b []byte) error {
	var res MetricsMeta
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
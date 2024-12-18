// Code generated by go-swagger; DO NOT EDIT.

package modelscapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// AddSignalsRequestItem Signal
//
// swagger:model AddSignalsRequestItem
type AddSignalsRequestItem struct {

	// alert id
	AlertID int64 `json:"alert_id,omitempty"`

	// context
	Context []*AddSignalsRequestItemContextItems0 `json:"context"`

	// created at
	CreatedAt string `json:"created_at,omitempty"`

	// decisions
	Decisions AddSignalsRequestItemDecisions `json:"decisions,omitempty"`

	// machine id
	MachineID string `json:"machine_id,omitempty"`

	// a human readable message
	// Required: true
	Message *string `json:"message"`

	// scenario
	// Required: true
	Scenario *string `json:"scenario"`

	// scenario hash
	// Required: true
	ScenarioHash *string `json:"scenario_hash"`

	// scenario trust
	ScenarioTrust string `json:"scenario_trust,omitempty"`

	// scenario version
	// Required: true
	ScenarioVersion *string `json:"scenario_version"`

	// source
	// Required: true
	Source *AddSignalsRequestItemSource `json:"source"`

	// start at
	// Required: true
	StartAt *string `json:"start_at"`

	// stop at
	// Required: true
	StopAt *string `json:"stop_at"`

	// UUID of the alert
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this add signals request item
func (m *AddSignalsRequestItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDecisions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScenario(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScenarioHash(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScenarioVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSource(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStopAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddSignalsRequestItem) validateContext(formats strfmt.Registry) error {
	if swag.IsZero(m.Context) { // not required
		return nil
	}

	for i := 0; i < len(m.Context); i++ {
		if swag.IsZero(m.Context[i]) { // not required
			continue
		}

		if m.Context[i] != nil {
			if err := m.Context[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("context" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("context" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AddSignalsRequestItem) validateDecisions(formats strfmt.Registry) error {
	if swag.IsZero(m.Decisions) { // not required
		return nil
	}

	if err := m.Decisions.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("decisions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("decisions")
		}
		return err
	}

	return nil
}

func (m *AddSignalsRequestItem) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

func (m *AddSignalsRequestItem) validateScenario(formats strfmt.Registry) error {

	if err := validate.Required("scenario", "body", m.Scenario); err != nil {
		return err
	}

	return nil
}

func (m *AddSignalsRequestItem) validateScenarioHash(formats strfmt.Registry) error {

	if err := validate.Required("scenario_hash", "body", m.ScenarioHash); err != nil {
		return err
	}

	return nil
}

func (m *AddSignalsRequestItem) validateScenarioVersion(formats strfmt.Registry) error {

	if err := validate.Required("scenario_version", "body", m.ScenarioVersion); err != nil {
		return err
	}

	return nil
}

func (m *AddSignalsRequestItem) validateSource(formats strfmt.Registry) error {

	if err := validate.Required("source", "body", m.Source); err != nil {
		return err
	}

	if m.Source != nil {
		if err := m.Source.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

func (m *AddSignalsRequestItem) validateStartAt(formats strfmt.Registry) error {

	if err := validate.Required("start_at", "body", m.StartAt); err != nil {
		return err
	}

	return nil
}

func (m *AddSignalsRequestItem) validateStopAt(formats strfmt.Registry) error {

	if err := validate.Required("stop_at", "body", m.StopAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this add signals request item based on the context it is used
func (m *AddSignalsRequestItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDecisions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSource(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddSignalsRequestItem) contextValidateContext(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Context); i++ {

		if m.Context[i] != nil {

			if swag.IsZero(m.Context[i]) { // not required
				return nil
			}

			if err := m.Context[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("context" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("context" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *AddSignalsRequestItem) contextValidateDecisions(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Decisions.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("decisions")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("decisions")
		}
		return err
	}

	return nil
}

func (m *AddSignalsRequestItem) contextValidateSource(ctx context.Context, formats strfmt.Registry) error {

	if m.Source != nil {

		if err := m.Source.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("source")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("source")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddSignalsRequestItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddSignalsRequestItem) UnmarshalBinary(b []byte) error {
	var res AddSignalsRequestItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// AddSignalsRequestItemContextItems0 add signals request item context items0
//
// swagger:model AddSignalsRequestItemContextItems0
type AddSignalsRequestItemContextItems0 struct {

	// key
	Key string `json:"key,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this add signals request item context items0
func (m *AddSignalsRequestItemContextItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this add signals request item context items0 based on context it is used
func (m *AddSignalsRequestItemContextItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AddSignalsRequestItemContextItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddSignalsRequestItemContextItems0) UnmarshalBinary(b []byte) error {
	var res AddSignalsRequestItemContextItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

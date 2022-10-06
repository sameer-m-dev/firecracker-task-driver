// Code generated by go-swagger; DO NOT EDIT.

// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Logger Describes the configuration option for the logging capability.
// swagger:model Logger
type Logger struct {

	// Set the level. The possible values are case-insensitive.
	// Enum: [Error Warning Info Debug]
	Level *string `json:"level,omitempty"`

	// Path to the named pipe or file for the human readable log output.
	// Required: true
	LogPath *string `json:"log_path"`

	// Whether or not to output the level in the logs.
	ShowLevel *bool `json:"show_level,omitempty"`

	// Whether or not to include the file path and line number of the log's origin.
	ShowLogOrigin *bool `json:"show_log_origin,omitempty"`
}

// Validate validates this logger
func (m *Logger) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLevel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLogPath(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var loggerTypeLevelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Error","Warning","Info","Debug"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		loggerTypeLevelPropEnum = append(loggerTypeLevelPropEnum, v)
	}
}

const (

	// LoggerLevelError captures enum value "Error"
	LoggerLevelError string = "Error"

	// LoggerLevelWarning captures enum value "Warning"
	LoggerLevelWarning string = "Warning"

	// LoggerLevelInfo captures enum value "Info"
	LoggerLevelInfo string = "Info"

	// LoggerLevelDebug captures enum value "Debug"
	LoggerLevelDebug string = "Debug"
)

// prop value enum
func (m *Logger) validateLevelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, loggerTypeLevelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Logger) validateLevel(formats strfmt.Registry) error {

	if swag.IsZero(m.Level) { // not required
		return nil
	}

	// value enum
	if err := m.validateLevelEnum("level", "body", *m.Level); err != nil {
		return err
	}

	return nil
}

func (m *Logger) validateLogPath(formats strfmt.Registry) error {

	if err := validate.Required("log_path", "body", m.LogPath); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Logger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Logger) UnmarshalBinary(b []byte) error {
	var res Logger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

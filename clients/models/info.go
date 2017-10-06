package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/xml"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// Info Collection of general purpose informational attributes which can be included in different other elements
// swagger:model Info
type Info struct {
	XMLName xml.Name `xml:"http://www.sap.com/lmsl/slp Info"`

	// description
	Description string `xml:"description,omitempty"`

	// display name
	DisplayName string `xml:"displayName,omitempty"`

	// external info
	ExternalInfo strfmt.URI `xml:"externalInfo,omitempty"`
}

// Validate validates this info
func (m *Info) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
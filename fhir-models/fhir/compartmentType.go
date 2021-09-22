// Copyright 2019 Daybreak Health
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package fhir

import (
	"encoding/json"
	"fmt"
	"strings"
)

// THIS FILE IS GENERATED BY https://github.com/daybreakhealth/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// CompartmentType is documented here http://hl7.org/fhir/ValueSet/compartment-type
type CompartmentType int

const (
	CompartmentTypePatient CompartmentType = iota
	CompartmentTypeEncounter
	CompartmentTypeRelatedPerson
	CompartmentTypePractitioner
	CompartmentTypeDevice
)

func (code CompartmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *CompartmentType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "Patient":
		*code = CompartmentTypePatient
	case "Encounter":
		*code = CompartmentTypeEncounter
	case "RelatedPerson":
		*code = CompartmentTypeRelatedPerson
	case "Practitioner":
		*code = CompartmentTypePractitioner
	case "Device":
		*code = CompartmentTypeDevice
	default:
		return fmt.Errorf("unknown CompartmentType code `%s`", s)
	}
	return nil
}
func (code CompartmentType) String() string {
	return code.Code()
}
func (code CompartmentType) Code() string {
	switch code {
	case CompartmentTypePatient:
		return "Patient"
	case CompartmentTypeEncounter:
		return "Encounter"
	case CompartmentTypeRelatedPerson:
		return "RelatedPerson"
	case CompartmentTypePractitioner:
		return "Practitioner"
	case CompartmentTypeDevice:
		return "Device"
	}
	return "<unknown>"
}
func (code CompartmentType) Display() string {
	switch code {
	case CompartmentTypePatient:
		return "Patient"
	case CompartmentTypeEncounter:
		return "Encounter"
	case CompartmentTypeRelatedPerson:
		return "RelatedPerson"
	case CompartmentTypePractitioner:
		return "Practitioner"
	case CompartmentTypeDevice:
		return "Device"
	}
	return "<unknown>"
}
func (code CompartmentType) Definition() string {
	switch code {
	case CompartmentTypePatient:
		return "The compartment definition is for the patient compartment."
	case CompartmentTypeEncounter:
		return "The compartment definition is for the encounter compartment."
	case CompartmentTypeRelatedPerson:
		return "The compartment definition is for the related-person compartment."
	case CompartmentTypePractitioner:
		return "The compartment definition is for the practitioner compartment."
	case CompartmentTypeDevice:
		return "The compartment definition is for the device compartment."
	}
	return "<unknown>"
}

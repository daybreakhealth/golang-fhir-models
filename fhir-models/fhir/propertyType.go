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

// PropertyType is documented here http://hl7.org/fhir/ValueSet/concept-property-type
type PropertyType int

const (
	PropertyTypeCode PropertyType = iota
	PropertyTypeCoding
	PropertyTypeString
	PropertyTypeInteger
	PropertyTypeBoolean
	PropertyTypeDateTime
	PropertyTypeDecimal
)

func (code PropertyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *PropertyType) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "code":
		*code = PropertyTypeCode
	case "Coding":
		*code = PropertyTypeCoding
	case "string":
		*code = PropertyTypeString
	case "integer":
		*code = PropertyTypeInteger
	case "boolean":
		*code = PropertyTypeBoolean
	case "dateTime":
		*code = PropertyTypeDateTime
	case "decimal":
		*code = PropertyTypeDecimal
	default:
		return fmt.Errorf("unknown PropertyType code `%s`", s)
	}
	return nil
}
func (code PropertyType) String() string {
	return code.Code()
}
func (code PropertyType) Code() string {
	switch code {
	case PropertyTypeCode:
		return "code"
	case PropertyTypeCoding:
		return "Coding"
	case PropertyTypeString:
		return "string"
	case PropertyTypeInteger:
		return "integer"
	case PropertyTypeBoolean:
		return "boolean"
	case PropertyTypeDateTime:
		return "dateTime"
	case PropertyTypeDecimal:
		return "decimal"
	}
	return "<unknown>"
}
func (code PropertyType) Display() string {
	switch code {
	case PropertyTypeCode:
		return "code (internal reference)"
	case PropertyTypeCoding:
		return "Coding (external reference)"
	case PropertyTypeString:
		return "string"
	case PropertyTypeInteger:
		return "integer"
	case PropertyTypeBoolean:
		return "boolean"
	case PropertyTypeDateTime:
		return "dateTime"
	case PropertyTypeDecimal:
		return "decimal"
	}
	return "<unknown>"
}
func (code PropertyType) Definition() string {
	switch code {
	case PropertyTypeCode:
		return "The property value is a code that identifies a concept defined in the code system."
	case PropertyTypeCoding:
		return "The property  value is a code defined in an external code system. This may be used for translations, but is not the intent."
	case PropertyTypeString:
		return "The property value is a string."
	case PropertyTypeInteger:
		return "The property value is a string (often used to assign ranking values to concepts for supporting score assessments)."
	case PropertyTypeBoolean:
		return "The property value is a boolean true | false."
	case PropertyTypeDateTime:
		return "The property is a date or a date + time."
	case PropertyTypeDecimal:
		return "The property value is a decimal number."
	}
	return "<unknown>"
}

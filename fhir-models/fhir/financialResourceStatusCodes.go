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

// FinancialResourceStatusCodes is documented here http://hl7.org/fhir/ValueSet/fm-status
type FinancialResourceStatusCodes int

const (
	FinancialResourceStatusCodesActive FinancialResourceStatusCodes = iota
	FinancialResourceStatusCodesCancelled
	FinancialResourceStatusCodesDraft
	FinancialResourceStatusCodesEnteredInError
)

func (code FinancialResourceStatusCodes) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *FinancialResourceStatusCodes) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "active":
		*code = FinancialResourceStatusCodesActive
	case "cancelled":
		*code = FinancialResourceStatusCodesCancelled
	case "draft":
		*code = FinancialResourceStatusCodesDraft
	case "entered-in-error":
		*code = FinancialResourceStatusCodesEnteredInError
	default:
		return fmt.Errorf("unknown FinancialResourceStatusCodes code `%s`", s)
	}
	return nil
}
func (code FinancialResourceStatusCodes) String() string {
	return code.Code()
}
func (code FinancialResourceStatusCodes) Code() string {
	switch code {
	case FinancialResourceStatusCodesActive:
		return "active"
	case FinancialResourceStatusCodesCancelled:
		return "cancelled"
	case FinancialResourceStatusCodesDraft:
		return "draft"
	case FinancialResourceStatusCodesEnteredInError:
		return "entered-in-error"
	}
	return "<unknown>"
}
func (code FinancialResourceStatusCodes) Display() string {
	switch code {
	case FinancialResourceStatusCodesActive:
		return "Active"
	case FinancialResourceStatusCodesCancelled:
		return "Cancelled"
	case FinancialResourceStatusCodesDraft:
		return "Draft"
	case FinancialResourceStatusCodesEnteredInError:
		return "Entered in Error"
	}
	return "<unknown>"
}
func (code FinancialResourceStatusCodes) Definition() string {
	switch code {
	case FinancialResourceStatusCodesActive:
		return "The instance is currently in-force."
	case FinancialResourceStatusCodesCancelled:
		return "The instance is withdrawn, rescinded or reversed."
	case FinancialResourceStatusCodesDraft:
		return "A new instance the contents of which is not complete."
	case FinancialResourceStatusCodesEnteredInError:
		return "The instance was entered in error."
	}
	return "<unknown>"
}

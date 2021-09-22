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

// ImagingStudyStatus is documented here http://hl7.org/fhir/ValueSet/imagingstudy-status
type ImagingStudyStatus int

const (
	ImagingStudyStatusRegistered ImagingStudyStatus = iota
	ImagingStudyStatusAvailable
	ImagingStudyStatusCancelled
	ImagingStudyStatusEnteredInError
	ImagingStudyStatusUnknown
)

func (code ImagingStudyStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(code.Code())
}
func (code *ImagingStudyStatus) UnmarshalJSON(json []byte) error {
	s := strings.Trim(string(json), "\"")
	switch s {
	case "registered":
		*code = ImagingStudyStatusRegistered
	case "available":
		*code = ImagingStudyStatusAvailable
	case "cancelled":
		*code = ImagingStudyStatusCancelled
	case "entered-in-error":
		*code = ImagingStudyStatusEnteredInError
	case "unknown":
		*code = ImagingStudyStatusUnknown
	default:
		return fmt.Errorf("unknown ImagingStudyStatus code `%s`", s)
	}
	return nil
}
func (code ImagingStudyStatus) String() string {
	return code.Code()
}
func (code ImagingStudyStatus) Code() string {
	switch code {
	case ImagingStudyStatusRegistered:
		return "registered"
	case ImagingStudyStatusAvailable:
		return "available"
	case ImagingStudyStatusCancelled:
		return "cancelled"
	case ImagingStudyStatusEnteredInError:
		return "entered-in-error"
	case ImagingStudyStatusUnknown:
		return "unknown"
	}
	return "<unknown>"
}
func (code ImagingStudyStatus) Display() string {
	switch code {
	case ImagingStudyStatusRegistered:
		return "Registered"
	case ImagingStudyStatusAvailable:
		return "Available"
	case ImagingStudyStatusCancelled:
		return "Cancelled"
	case ImagingStudyStatusEnteredInError:
		return "Entered in Error"
	case ImagingStudyStatusUnknown:
		return "Unknown"
	}
	return "<unknown>"
}
func (code ImagingStudyStatus) Definition() string {
	switch code {
	case ImagingStudyStatusRegistered:
		return "The existence of the imaging study is registered, but there is nothing yet available."
	case ImagingStudyStatusAvailable:
		return "At least one instance has been associated with this imaging study."
	case ImagingStudyStatusCancelled:
		return "The imaging study is unavailable because the imaging study was not started or not completed (also sometimes called \"aborted\")."
	case ImagingStudyStatusEnteredInError:
		return "The imaging study has been withdrawn following a previous final release.  This electronic record should never have existed, though it is possible that real-world decisions were based on it. (If real-world activity has occurred, the status should be \"cancelled\" rather than \"entered-in-error\".)."
	case ImagingStudyStatusUnknown:
		return "The system does not know which of the status values currently applies for this request. Note: This concept is not to be used for \"other\" - one of the listed statuses is presumed to apply, it's just not known which one."
	}
	return "<unknown>"
}

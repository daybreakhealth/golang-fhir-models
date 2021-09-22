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

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/daybreakhealth/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// FamilyMemberHistory is documented here http://hl7.org/fhir/StructureDefinition/FamilyMemberHistory
type FamilyMemberHistory struct {
	Id                    *string                        `bson:"id,omitempty" json:"id,omitempty"`
	Meta                  *Meta                          `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules         *string                        `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language              *string                        `bson:"language,omitempty" json:"language,omitempty"`
	Text                  *Narrative                     `bson:"text,omitempty" json:"text,omitempty"`
	Extension             []Extension                    `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension     []Extension                    `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier            []Identifier                   `bson:"identifier,omitempty" json:"identifier,omitempty"`
	InstantiatesCanonical []string                       `bson:"instantiatesCanonical,omitempty" json:"instantiatesCanonical,omitempty"`
	InstantiatesUri       []string                       `bson:"instantiatesUri,omitempty" json:"instantiatesUri,omitempty"`
	Status                FamilyHistoryStatus            `bson:"status" json:"status"`
	DataAbsentReason      *CodeableConcept               `bson:"dataAbsentReason,omitempty" json:"dataAbsentReason,omitempty"`
	Patient               Reference                      `bson:"patient" json:"patient"`
	Date                  *string                        `bson:"date,omitempty" json:"date,omitempty"`
	Name                  *string                        `bson:"name,omitempty" json:"name,omitempty"`
	Relationship          CodeableConcept                `bson:"relationship" json:"relationship"`
	Sex                   *CodeableConcept               `bson:"sex,omitempty" json:"sex,omitempty"`
	EstimatedAge          *bool                          `bson:"estimatedAge,omitempty" json:"estimatedAge,omitempty"`
	ReasonCode            []CodeableConcept              `bson:"reasonCode,omitempty" json:"reasonCode,omitempty"`
	ReasonReference       []Reference                    `bson:"reasonReference,omitempty" json:"reasonReference,omitempty"`
	Note                  []Annotation                   `bson:"note,omitempty" json:"note,omitempty"`
	Condition             []FamilyMemberHistoryCondition `bson:"condition,omitempty" json:"condition,omitempty"`
}
type FamilyMemberHistoryCondition struct {
	Id                 *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension          []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension  []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code               CodeableConcept  `bson:"code" json:"code"`
	Outcome            *CodeableConcept `bson:"outcome,omitempty" json:"outcome,omitempty"`
	ContributedToDeath *bool            `bson:"contributedToDeath,omitempty" json:"contributedToDeath,omitempty"`
	Note               []Annotation     `bson:"note,omitempty" json:"note,omitempty"`
}
type OtherFamilyMemberHistory FamilyMemberHistory

// MarshalJSON marshals the given FamilyMemberHistory as JSON into a byte slice
func (r FamilyMemberHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherFamilyMemberHistory
		ResourceType string `json:"resourceType"`
	}{
		OtherFamilyMemberHistory: OtherFamilyMemberHistory(r),
		ResourceType:             "FamilyMemberHistory",
	})
}

// UnmarshalFamilyMemberHistory unmarshals a FamilyMemberHistory.
func UnmarshalFamilyMemberHistory(b []byte) (FamilyMemberHistory, error) {
	var familyMemberHistory FamilyMemberHistory
	if err := json.Unmarshal(b, &familyMemberHistory); err != nil {
		return familyMemberHistory, err
	}
	return familyMemberHistory, nil
}

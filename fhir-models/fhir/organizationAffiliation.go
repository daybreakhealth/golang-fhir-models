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

// OrganizationAffiliation is documented here http://hl7.org/fhir/StructureDefinition/OrganizationAffiliation
type OrganizationAffiliation struct {
	Id                        *string           `bson:"id,omitempty" json:"id,omitempty"`
	Meta                      *Meta             `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules             *string           `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language                  *string           `bson:"language,omitempty" json:"language,omitempty"`
	Text                      *Narrative        `bson:"text,omitempty" json:"text,omitempty"`
	Extension                 []Extension       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension         []Extension       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier                []Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Active                    *bool             `bson:"active,omitempty" json:"active,omitempty"`
	Period                    *Period           `bson:"period,omitempty" json:"period,omitempty"`
	Organization              *Reference        `bson:"organization,omitempty" json:"organization,omitempty"`
	ParticipatingOrganization *Reference        `bson:"participatingOrganization,omitempty" json:"participatingOrganization,omitempty"`
	Network                   []Reference       `bson:"network,omitempty" json:"network,omitempty"`
	Code                      []CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Specialty                 []CodeableConcept `bson:"specialty,omitempty" json:"specialty,omitempty"`
	Location                  []Reference       `bson:"location,omitempty" json:"location,omitempty"`
	HealthcareService         []Reference       `bson:"healthcareService,omitempty" json:"healthcareService,omitempty"`
	Telecom                   []ContactPoint    `bson:"telecom,omitempty" json:"telecom,omitempty"`
	Endpoint                  []Reference       `bson:"endpoint,omitempty" json:"endpoint,omitempty"`
}
type OtherOrganizationAffiliation OrganizationAffiliation

// MarshalJSON marshals the given OrganizationAffiliation as JSON into a byte slice
func (r OrganizationAffiliation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherOrganizationAffiliation
		ResourceType string `json:"resourceType"`
	}{
		OtherOrganizationAffiliation: OtherOrganizationAffiliation(r),
		ResourceType:                 "OrganizationAffiliation",
	})
}

// UnmarshalOrganizationAffiliation unmarshals a OrganizationAffiliation.
func UnmarshalOrganizationAffiliation(b []byte) (OrganizationAffiliation, error) {
	var organizationAffiliation OrganizationAffiliation
	if err := json.Unmarshal(b, &organizationAffiliation); err != nil {
		return organizationAffiliation, err
	}
	return organizationAffiliation, nil
}

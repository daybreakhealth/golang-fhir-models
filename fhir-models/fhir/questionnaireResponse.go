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

// QuestionnaireResponse is documented here http://hl7.org/fhir/StructureDefinition/QuestionnaireResponse
type QuestionnaireResponse struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                       `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                     `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                     `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                  `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	BasedOn           []Reference                 `bson:"basedOn,omitempty" json:"basedOn,omitempty"`
	PartOf            []Reference                 `bson:"partOf,omitempty" json:"partOf,omitempty"`
	Questionnaire     *string                     `bson:"questionnaire,omitempty" json:"questionnaire,omitempty"`
	Status            QuestionnaireResponseStatus `bson:"status" json:"status"`
	Subject           *Reference                  `bson:"subject,omitempty" json:"subject,omitempty"`
	Encounter         *Reference                  `bson:"encounter,omitempty" json:"encounter,omitempty"`
	Authored          *string                     `bson:"authored,omitempty" json:"authored,omitempty"`
	Author            *Reference                  `bson:"author,omitempty" json:"author,omitempty"`
	Source            *Reference                  `bson:"source,omitempty" json:"source,omitempty"`
	Item              []QuestionnaireResponseItem `bson:"item,omitempty" json:"item,omitempty"`
}
type QuestionnaireResponseItem struct {
	Id                *string                           `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                       `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                       `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	LinkId            string                            `bson:"linkId" json:"linkId"`
	Definition        *string                           `bson:"definition,omitempty" json:"definition,omitempty"`
	Text              *string                           `bson:"text,omitempty" json:"text,omitempty"`
	Answer            []QuestionnaireResponseItemAnswer `bson:"answer,omitempty" json:"answer,omitempty"`
	Item              []QuestionnaireResponseItem       `bson:"item,omitempty" json:"item,omitempty"`
}
type QuestionnaireResponseItemAnswer struct {
	Id                *string                     `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                 `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                 `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Item              []QuestionnaireResponseItem `bson:"item,omitempty" json:"item,omitempty"`
}
type OtherQuestionnaireResponse QuestionnaireResponse

// MarshalJSON marshals the given QuestionnaireResponse as JSON into a byte slice
func (r QuestionnaireResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherQuestionnaireResponse
		ResourceType string `json:"resourceType"`
	}{
		OtherQuestionnaireResponse: OtherQuestionnaireResponse(r),
		ResourceType:               "QuestionnaireResponse",
	})
}

// UnmarshalQuestionnaireResponse unmarshals a QuestionnaireResponse.
func UnmarshalQuestionnaireResponse(b []byte) (QuestionnaireResponse, error) {
	var questionnaireResponse QuestionnaireResponse
	if err := json.Unmarshal(b, &questionnaireResponse); err != nil {
		return questionnaireResponse, err
	}
	return questionnaireResponse, nil
}

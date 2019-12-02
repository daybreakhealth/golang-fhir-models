// Copyright 2019 The Samply Development Community
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

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// DataRequirement is documented here http://hl7.org/fhir/StructureDefinition/DataRequirement
type DataRequirement struct {
	Id          *string                     `json:"id,omitempty"`
	Extension   []Extension                 `json:"extension,omitempty"`
	Type        string                      `json:"type"`
	Profile     []string                    `json:"profile,omitempty"`
	MustSupport []string                    `json:"mustSupport,omitempty"`
	CodeFilter  []DataRequirementCodeFilter `json:"codeFilter,omitempty"`
	DateFilter  []DataRequirementDateFilter `json:"dateFilter,omitempty"`
	Limit       *int                        `json:"limit,omitempty"`
	Sort        []DataRequirementSort       `json:"sort,omitempty"`
}
type DataRequirementCodeFilter struct {
	Id          *string     `json:"id,omitempty"`
	Extension   []Extension `json:"extension,omitempty"`
	Path        *string     `json:"path,omitempty"`
	SearchParam *string     `json:"searchParam,omitempty"`
	ValueSet    *string     `json:"valueSet,omitempty"`
	Code        []Coding    `json:"code,omitempty"`
}
type DataRequirementDateFilter struct {
	Id          *string     `json:"id,omitempty"`
	Extension   []Extension `json:"extension,omitempty"`
	Path        *string     `json:"path,omitempty"`
	SearchParam *string     `json:"searchParam,omitempty"`
}
type DataRequirementSort struct {
	Id        *string       `json:"id,omitempty"`
	Extension []Extension   `json:"extension,omitempty"`
	Path      string        `json:"path"`
	Direction SortDirection `json:"direction"`
}
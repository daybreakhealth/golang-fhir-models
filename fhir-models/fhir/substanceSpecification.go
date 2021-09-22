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

// SubstanceSpecification is documented here http://hl7.org/fhir/StructureDefinition/SubstanceSpecification
type SubstanceSpecification struct {
	Id                   *string                                                 `bson:"id,omitempty" json:"id,omitempty"`
	Meta                 *Meta                                                   `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules        *string                                                 `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language             *string                                                 `bson:"language,omitempty" json:"language,omitempty"`
	Text                 *Narrative                                              `bson:"text,omitempty" json:"text,omitempty"`
	Extension            []Extension                                             `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension    []Extension                                             `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier           *Identifier                                             `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Type                 *CodeableConcept                                        `bson:"type,omitempty" json:"type,omitempty"`
	Status               *CodeableConcept                                        `bson:"status,omitempty" json:"status,omitempty"`
	Domain               *CodeableConcept                                        `bson:"domain,omitempty" json:"domain,omitempty"`
	Description          *string                                                 `bson:"description,omitempty" json:"description,omitempty"`
	Source               []Reference                                             `bson:"source,omitempty" json:"source,omitempty"`
	Comment              *string                                                 `bson:"comment,omitempty" json:"comment,omitempty"`
	Moiety               []SubstanceSpecificationMoiety                          `bson:"moiety,omitempty" json:"moiety,omitempty"`
	Property             []SubstanceSpecificationProperty                        `bson:"property,omitempty" json:"property,omitempty"`
	ReferenceInformation *Reference                                              `bson:"referenceInformation,omitempty" json:"referenceInformation,omitempty"`
	Structure            *SubstanceSpecificationStructure                        `bson:"structure,omitempty" json:"structure,omitempty"`
	Code                 []SubstanceSpecificationCode                            `bson:"code,omitempty" json:"code,omitempty"`
	Name                 []SubstanceSpecificationName                            `bson:"name,omitempty" json:"name,omitempty"`
	MolecularWeight      []SubstanceSpecificationStructureIsotopeMolecularWeight `bson:"molecularWeight,omitempty" json:"molecularWeight,omitempty"`
	Relationship         []SubstanceSpecificationRelationship                    `bson:"relationship,omitempty" json:"relationship,omitempty"`
	NucleicAcid          *Reference                                              `bson:"nucleicAcid,omitempty" json:"nucleicAcid,omitempty"`
	Polymer              *Reference                                              `bson:"polymer,omitempty" json:"polymer,omitempty"`
	Protein              *Reference                                              `bson:"protein,omitempty" json:"protein,omitempty"`
	SourceMaterial       *Reference                                              `bson:"sourceMaterial,omitempty" json:"sourceMaterial,omitempty"`
}
type SubstanceSpecificationMoiety struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Role              *CodeableConcept `bson:"role,omitempty" json:"role,omitempty"`
	Identifier        *Identifier      `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name              *string          `bson:"name,omitempty" json:"name,omitempty"`
	Stereochemistry   *CodeableConcept `bson:"stereochemistry,omitempty" json:"stereochemistry,omitempty"`
	OpticalActivity   *CodeableConcept `bson:"opticalActivity,omitempty" json:"opticalActivity,omitempty"`
	MolecularFormula  *string          `bson:"molecularFormula,omitempty" json:"molecularFormula,omitempty"`
}
type SubstanceSpecificationProperty struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Category          *CodeableConcept `bson:"category,omitempty" json:"category,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Parameters        *string          `bson:"parameters,omitempty" json:"parameters,omitempty"`
}
type SubstanceSpecificationStructure struct {
	Id                       *string                                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension                []Extension                                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension        []Extension                                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Stereochemistry          *CodeableConcept                                       `bson:"stereochemistry,omitempty" json:"stereochemistry,omitempty"`
	OpticalActivity          *CodeableConcept                                       `bson:"opticalActivity,omitempty" json:"opticalActivity,omitempty"`
	MolecularFormula         *string                                                `bson:"molecularFormula,omitempty" json:"molecularFormula,omitempty"`
	MolecularFormulaByMoiety *string                                                `bson:"molecularFormulaByMoiety,omitempty" json:"molecularFormulaByMoiety,omitempty"`
	Isotope                  []SubstanceSpecificationStructureIsotope               `bson:"isotope,omitempty" json:"isotope,omitempty"`
	MolecularWeight          *SubstanceSpecificationStructureIsotopeMolecularWeight `bson:"molecularWeight,omitempty" json:"molecularWeight,omitempty"`
	Source                   []Reference                                            `bson:"source,omitempty" json:"source,omitempty"`
	Representation           []SubstanceSpecificationStructureRepresentation        `bson:"representation,omitempty" json:"representation,omitempty"`
}
type SubstanceSpecificationStructureIsotope struct {
	Id                *string                                                `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                                            `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                                            `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        *Identifier                                            `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Name              *CodeableConcept                                       `bson:"name,omitempty" json:"name,omitempty"`
	Substitution      *CodeableConcept                                       `bson:"substitution,omitempty" json:"substitution,omitempty"`
	HalfLife          *Quantity                                              `bson:"halfLife,omitempty" json:"halfLife,omitempty"`
	MolecularWeight   *SubstanceSpecificationStructureIsotopeMolecularWeight `bson:"molecularWeight,omitempty" json:"molecularWeight,omitempty"`
}
type SubstanceSpecificationStructureIsotopeMolecularWeight struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Method            *CodeableConcept `bson:"method,omitempty" json:"method,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Amount            *Quantity        `bson:"amount,omitempty" json:"amount,omitempty"`
}
type SubstanceSpecificationStructureRepresentation struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Type              *CodeableConcept `bson:"type,omitempty" json:"type,omitempty"`
	Representation    *string          `bson:"representation,omitempty" json:"representation,omitempty"`
	Attachment        *Attachment      `bson:"attachment,omitempty" json:"attachment,omitempty"`
}
type SubstanceSpecificationCode struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Code              *CodeableConcept `bson:"code,omitempty" json:"code,omitempty"`
	Status            *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	StatusDate        *string          `bson:"statusDate,omitempty" json:"statusDate,omitempty"`
	Comment           *string          `bson:"comment,omitempty" json:"comment,omitempty"`
	Source            []Reference      `bson:"source,omitempty" json:"source,omitempty"`
}
type SubstanceSpecificationName struct {
	Id                *string                              `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension                          `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                          `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Name              string                               `bson:"name" json:"name"`
	Type              *CodeableConcept                     `bson:"type,omitempty" json:"type,omitempty"`
	Status            *CodeableConcept                     `bson:"status,omitempty" json:"status,omitempty"`
	Preferred         *bool                                `bson:"preferred,omitempty" json:"preferred,omitempty"`
	Language          []CodeableConcept                    `bson:"language,omitempty" json:"language,omitempty"`
	Domain            []CodeableConcept                    `bson:"domain,omitempty" json:"domain,omitempty"`
	Jurisdiction      []CodeableConcept                    `bson:"jurisdiction,omitempty" json:"jurisdiction,omitempty"`
	Synonym           []SubstanceSpecificationName         `bson:"synonym,omitempty" json:"synonym,omitempty"`
	Translation       []SubstanceSpecificationName         `bson:"translation,omitempty" json:"translation,omitempty"`
	Official          []SubstanceSpecificationNameOfficial `bson:"official,omitempty" json:"official,omitempty"`
	Source            []Reference                          `bson:"source,omitempty" json:"source,omitempty"`
}
type SubstanceSpecificationNameOfficial struct {
	Id                *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension         []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Authority         *CodeableConcept `bson:"authority,omitempty" json:"authority,omitempty"`
	Status            *CodeableConcept `bson:"status,omitempty" json:"status,omitempty"`
	Date              *string          `bson:"date,omitempty" json:"date,omitempty"`
}
type SubstanceSpecificationRelationship struct {
	Id                  *string          `bson:"id,omitempty" json:"id,omitempty"`
	Extension           []Extension      `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension   []Extension      `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Relationship        *CodeableConcept `bson:"relationship,omitempty" json:"relationship,omitempty"`
	IsDefining          *bool            `bson:"isDefining,omitempty" json:"isDefining,omitempty"`
	AmountRatioLowLimit *Ratio           `bson:"amountRatioLowLimit,omitempty" json:"amountRatioLowLimit,omitempty"`
	AmountType          *CodeableConcept `bson:"amountType,omitempty" json:"amountType,omitempty"`
	Source              []Reference      `bson:"source,omitempty" json:"source,omitempty"`
}
type OtherSubstanceSpecification SubstanceSpecification

// MarshalJSON marshals the given SubstanceSpecification as JSON into a byte slice
func (r SubstanceSpecification) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSubstanceSpecification
		ResourceType string `json:"resourceType"`
	}{
		OtherSubstanceSpecification: OtherSubstanceSpecification(r),
		ResourceType:                "SubstanceSpecification",
	})
}

// UnmarshalSubstanceSpecification unmarshals a SubstanceSpecification.
func UnmarshalSubstanceSpecification(b []byte) (SubstanceSpecification, error) {
	var substanceSpecification SubstanceSpecification
	if err := json.Unmarshal(b, &substanceSpecification); err != nil {
		return substanceSpecification, err
	}
	return substanceSpecification, nil
}

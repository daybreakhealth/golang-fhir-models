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

// PaymentNotice is documented here http://hl7.org/fhir/StructureDefinition/PaymentNotice
type PaymentNotice struct {
	Id                *string                      `bson:"id,omitempty" json:"id,omitempty"`
	Meta              *Meta                        `bson:"meta,omitempty" json:"meta,omitempty"`
	ImplicitRules     *string                      `bson:"implicitRules,omitempty" json:"implicitRules,omitempty"`
	Language          *string                      `bson:"language,omitempty" json:"language,omitempty"`
	Text              *Narrative                   `bson:"text,omitempty" json:"text,omitempty"`
	Extension         []Extension                  `bson:"extension,omitempty" json:"extension,omitempty"`
	ModifierExtension []Extension                  `bson:"modifierExtension,omitempty" json:"modifierExtension,omitempty"`
	Identifier        []Identifier                 `bson:"identifier,omitempty" json:"identifier,omitempty"`
	Status            FinancialResourceStatusCodes `bson:"status" json:"status"`
	Request           *Reference                   `bson:"request,omitempty" json:"request,omitempty"`
	Response          *Reference                   `bson:"response,omitempty" json:"response,omitempty"`
	Created           string                       `bson:"created" json:"created"`
	Provider          *Reference                   `bson:"provider,omitempty" json:"provider,omitempty"`
	Payment           Reference                    `bson:"payment" json:"payment"`
	PaymentDate       *string                      `bson:"paymentDate,omitempty" json:"paymentDate,omitempty"`
	Payee             *Reference                   `bson:"payee,omitempty" json:"payee,omitempty"`
	Recipient         Reference                    `bson:"recipient" json:"recipient"`
	Amount            Money                        `bson:"amount" json:"amount"`
	PaymentStatus     *CodeableConcept             `bson:"paymentStatus,omitempty" json:"paymentStatus,omitempty"`
}
type OtherPaymentNotice PaymentNotice

// MarshalJSON marshals the given PaymentNotice as JSON into a byte slice
func (r PaymentNotice) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherPaymentNotice
		ResourceType string `json:"resourceType"`
	}{
		OtherPaymentNotice: OtherPaymentNotice(r),
		ResourceType:       "PaymentNotice",
	})
}

// UnmarshalPaymentNotice unmarshals a PaymentNotice.
func UnmarshalPaymentNotice(b []byte) (PaymentNotice, error) {
	var paymentNotice PaymentNotice
	if err := json.Unmarshal(b, &paymentNotice); err != nil {
		return paymentNotice, err
	}
	return paymentNotice, nil
}

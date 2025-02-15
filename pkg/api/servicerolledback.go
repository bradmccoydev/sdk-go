/*
Copyright 2022 The CDEvents Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"time"
)

const (
	// ServiceRolledback event
	ServiceRolledbackEventV1 CDEventType = "dev.cdevents.service.rolledback.v1"
)

type ServiceRolledbackSubjectContent struct{}

type ServiceRolledbackSubject struct {
	SubjectBase
	Content ServiceRolledbackSubjectContent `json:"content"`
}

func (sc ServiceRolledbackSubject) GetEventType() CDEventType {
	return ServiceRolledbackEventV1
}

func (sc ServiceRolledbackSubject) GetSubjectType() SubjectType {
	return ServiceSubjectType
}

type ServiceRolledbackEvent struct {
	Context Context                  `json:"context"`
	Subject ServiceRolledbackSubject `json:"subject"`
}

// CDEventsReader implementation

func (e ServiceRolledbackEvent) GetType() CDEventType {
	return ServiceRolledbackEventV1
}

func (e ServiceRolledbackEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e ServiceRolledbackEvent) GetId() string {
	return e.Context.Id
}

func (e ServiceRolledbackEvent) GetSource() string {
	return e.Context.Source
}

func (e ServiceRolledbackEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e ServiceRolledbackEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e ServiceRolledbackEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e ServiceRolledbackEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *ServiceRolledbackEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *ServiceRolledbackEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *ServiceRolledbackEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *ServiceRolledbackEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *ServiceRolledbackEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func newServiceRolledbackEvent() CDEvent {
	return &ServiceRolledbackEvent{
		Context: Context{
			Type:    ServiceRolledbackEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: ServiceRolledbackSubject{},
	}
}

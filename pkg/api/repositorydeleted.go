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
	// RepositoryDeleted event
	RepositoryDeletedEventV1 CDEventType = "dev.cdevents.repository.deleted.v1"
)

type RepositoryDeletedSubjectContent struct{}

type RepositoryDeletedSubject struct {
	SubjectBase
	Content RepositoryDeletedSubjectContent `json:"content"`
}

func (sc RepositoryDeletedSubject) GetEventType() CDEventType {
	return RepositoryDeletedEventV1
}

func (sc RepositoryDeletedSubject) GetSubjectType() SubjectType {
	return RepositorySubjectType
}

type RepositoryDeletedEvent struct {
	Context Context                  `json:"context"`
	Subject RepositoryDeletedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e RepositoryDeletedEvent) GetType() CDEventType {
	return RepositoryDeletedEventV1
}

func (e RepositoryDeletedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e RepositoryDeletedEvent) GetId() string {
	return e.Context.Id
}

func (e RepositoryDeletedEvent) GetSource() string {
	return e.Context.Source
}

func (e RepositoryDeletedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e RepositoryDeletedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e RepositoryDeletedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e RepositoryDeletedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *RepositoryDeletedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *RepositoryDeletedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *RepositoryDeletedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *RepositoryDeletedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *RepositoryDeletedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func newRepositoryDeletedEvent() CDEvent {
	return &RepositoryDeletedEvent{
		Context: Context{
			Type:    RepositoryDeletedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: RepositoryDeletedSubject{},
	}
}

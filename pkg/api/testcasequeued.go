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
	// TestCaseQueued event
	TestCaseQueuedEventV1 CDEventType = "dev.cdevents.testcase.queued.v1"
)

type TestCaseQueuedSubjectContent struct{}

type TestCaseQueuedSubject struct {
	SubjectBase
	Content TestCaseQueuedSubjectContent `json:"content"`
}

func (sc TestCaseQueuedSubject) GetEventType() CDEventType {
	return TestCaseQueuedEventV1
}

func (sc TestCaseQueuedSubject) GetSubjectType() SubjectType {
	return TestCaseSubjectType
}

type TestCaseQueuedEvent struct {
	Context Context               `json:"context"`
	Subject TestCaseQueuedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e TestCaseQueuedEvent) GetType() CDEventType {
	return TestCaseQueuedEventV1
}

func (e TestCaseQueuedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e TestCaseQueuedEvent) GetId() string {
	return e.Context.Id
}

func (e TestCaseQueuedEvent) GetSource() string {
	return e.Context.Source
}

func (e TestCaseQueuedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e TestCaseQueuedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e TestCaseQueuedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e TestCaseQueuedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *TestCaseQueuedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *TestCaseQueuedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *TestCaseQueuedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *TestCaseQueuedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *TestCaseQueuedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func newTestCaseQueuedEvent() CDEvent {
	return &TestCaseQueuedEvent{
		Context: Context{
			Type:    TestCaseQueuedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: TestCaseQueuedSubject{},
	}
}

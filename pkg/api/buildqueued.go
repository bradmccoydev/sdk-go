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
	// BuildQueued event
	BuildQueuedEventV1 CDEventType = "dev.cdevents.build.queued.v1"
)

type BuildQueuedSubjectContent struct{}

type BuildQueuedSubject struct {
	SubjectBase
	Content BuildQueuedSubjectContent `json:"content"`
}

func (sc BuildQueuedSubject) GetEventType() CDEventType {
	return BuildQueuedEventV1
}

func (sc BuildQueuedSubject) GetSubjectType() SubjectType {
	return BuildSubjectType
}

type BuildQueuedEvent struct {
	Context Context            `json:"context"`
	Subject BuildQueuedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e BuildQueuedEvent) GetType() CDEventType {
	return BuildQueuedEventV1
}

func (e BuildQueuedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e BuildQueuedEvent) GetId() string {
	return e.Context.Id
}

func (e BuildQueuedEvent) GetSource() string {
	return e.Context.Source
}

func (e BuildQueuedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e BuildQueuedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e BuildQueuedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e BuildQueuedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *BuildQueuedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *BuildQueuedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *BuildQueuedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *BuildQueuedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *BuildQueuedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func newBuildQueuedEvent() CDEvent {
	return &BuildQueuedEvent{
		Context: Context{
			Type:    BuildQueuedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: BuildQueuedSubject{},
	}
}

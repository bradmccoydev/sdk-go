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
	// BuildFinished event
	BuildFinishedEventV1 CDEventType = "dev.cdevents.build.finished.v1"
)

type BuildFinishedSubjectContent struct{}

type BuildFinishedSubject struct {
	SubjectBase
	Content BuildFinishedSubjectContent `json:"content"`
}

func (sc BuildFinishedSubject) GetEventType() CDEventType {
	return BuildFinishedEventV1
}

func (sc BuildFinishedSubject) GetSubjectType() SubjectType {
	return BuildSubjectType
}

type BuildFinishedEvent struct {
	Context Context              `json:"context"`
	Subject BuildFinishedSubject `json:"subject"`
}

// CDEventsReader implementation

func (e BuildFinishedEvent) GetType() CDEventType {
	return BuildFinishedEventV1
}

func (e BuildFinishedEvent) GetVersion() string {
	return CDEventsSpecVersion
}

func (e BuildFinishedEvent) GetId() string {
	return e.Context.Id
}

func (e BuildFinishedEvent) GetSource() string {
	return e.Context.Source
}

func (e BuildFinishedEvent) GetTimestamp() time.Time {
	return e.Context.Timestamp
}

func (e BuildFinishedEvent) GetSubjectId() string {
	return e.Subject.Id
}

func (e BuildFinishedEvent) GetSubjectSource() string {
	return e.Subject.Source
}

func (e BuildFinishedEvent) GetSubject() Subject {
	return e.Subject
}

// CDEventsWriter implementation

func (e *BuildFinishedEvent) SetId(id string) {
	e.Context.Id = id
}

func (e *BuildFinishedEvent) SetSource(source string) {
	e.Context.Source = source
	// Default the subject source to the event source
	if e.Subject.Source == "" {
		e.Subject.Source = source
	}
}

func (e *BuildFinishedEvent) SetTimestamp(timestamp time.Time) {
	e.Context.Timestamp = timestamp
}

func (e *BuildFinishedEvent) SetSubjectId(subjectId string) {
	e.Subject.Id = subjectId
}

func (e *BuildFinishedEvent) SetSubjectSource(subjectSource string) {
	e.Subject.Source = subjectSource
}

func newBuildFinishedEvent() CDEvent {
	return &BuildFinishedEvent{
		Context: Context{
			Type:    BuildFinishedEventV1,
			Version: CDEventsSpecVersion,
		},
		Subject: BuildFinishedSubject{},
	}
}

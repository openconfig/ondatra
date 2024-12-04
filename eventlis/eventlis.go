// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package eventlis provides an API to attach listeners that are called at
// events during the test execution.
//
// Code can register a callback to be invoked after the reservation is complete
// but before any tests are executed by calling
// [EventListener.AddBeforeTestsCallback]:
//
//	ondatra.EventListener().AddBeforeTestsCallback(func (e *eventlis.BeforeTestsEvent) {
//		handleBeforeTestsEvent(e)
//	})
//
// Code can register a callback to be invoked after all the tests are complete
// but before the reservation is released by calling
// [EventListener.AddAfterTestsCallback]:
//
//	ondatra.EventListener().AddAfterTestsCallback(func (e *eventlis.AfterTestsEvent) {
//		handleAfterTestsEvent(e)
//	})
//
// These calls can be made from the test itself, from the binding, or from a
// helper library that the test imports.
package eventlis

import (
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/internal/events"
)

// EventListener is an event listener API.
type EventListener struct{}

// BeforeTestsEvent occurs after the reservation is complete and before tests
// start executing.
type BeforeTestsEvent struct {
	Reservation *binding.Reservation
}

// AfterTestsEvent occurs after the tests are finished executing and before
// the reservation is released.
type AfterTestsEvent struct {
	ExitCode *int
}

// AddBeforeTestsCallback adds a callback to run after the reservation is
// complete and before tests start executing.
func (el *EventListener) AddBeforeTestsCallback(cb func(*BeforeTestsEvent) error) {
	events.AddBeforeTests(func(res *binding.Reservation) error {
		if err := cb(&BeforeTestsEvent{Reservation: res}); err != nil {
			return events.CallbackErr(err, cb)
		}
		return nil
	})
}

// AddAfterTestsCallback adds a callback to run after the tests are finished
// executing and before the reservation is released.
func (el *EventListener) AddAfterTestsCallback(cb func(*AfterTestsEvent) error) {
	events.AddAfterTests(func(exitCode *int) error {
		if err := cb(&AfterTestsEvent{ExitCode: exitCode}); err != nil {
			return events.CallbackErr(err, cb)
		}
		return nil
	})
}

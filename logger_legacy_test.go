// Copyright © 2021 Kris Nóva <kris@nivenly.com>
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
//
// ----------------------------------------------------------------------------
//
// Legacy Regression tests for 0.1.0
//
// Note: @christopherhein added these tests and I am unsure if Amazon
// is still using the old io.Writer append() convention.
// Either way, I am keeping legacy (deprecated) support for it.
//
// All of these regression tests should still pass.
//
// ----------------------------------------------------------------------------

package logger

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"
)

const (
	format          = "%v, %v, %v, all eyes on me!"
	formatExp       = `^\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}.* \[%s\]  \d, \d, \d, all eyes on me!`
	formatWOTimeExp = `\[%s\]  \d, \d, \d, all eyes on me!`
)

var (
	a = []interface{}{1, 2, 3}
)

func TestMain(m *testing.M) {
	TestMode = true
	m.Run()
}

func TestAlwaysLegacy(t *testing.T) {
	e, err := regexp.Compile(fmt.Sprintf(formatExp, PreAlways))
	g := captureLoggerOutput(Always, format, a)

	if err != nil {
		t.Fatalf("Failed to compile regexp '%v': %v", e.String(), err)
	}

	if !e.MatchString(g) {
		t.Fatalf("Always should produce a pattern '%v' but produces: %v", e.String(), g)
	}
}

func TestCriticalLegacy(t *testing.T) {
	Level = 1

	e, err := regexp.Compile(fmt.Sprintf(formatExp, PreCritical))
	g := captureLoggerOutput(Critical, format, a)

	if err != nil {
		t.Fatalf("Failed to compile regexp '%v': %v", e.String(), err)
	}

	if !e.MatchString(g) {
		t.Fatalf("Critical should produce a pattern '%v' but produces: %v", e.String(), g)
	}
}

func TestInfoLegacy(t *testing.T) {
	Level = 3

	e, err := regexp.Compile(fmt.Sprintf(formatExp, PreInfo))
	g := captureLoggerOutput(Info, format, a)

	if err != nil {
		t.Fatalf("Failed to compile regexp '%v': %v", e.String(), err)
	}

	if !e.MatchString(g) {
		t.Fatalf("Info should produce a pattern '%v' but produces: %v", e.String(), g)
	}
}

func TestSuccessLegacy(t *testing.T) {
	Level = 3

	e, err := regexp.Compile(fmt.Sprintf(formatExp, PreSuccess))
	g := captureLoggerOutput(Success, format, a)

	if err != nil {
		t.Fatalf("Failed to compile regexp '%v': %v", e.String(), err)
	}

	if !e.MatchString(g) {
		t.Fatalf("Success should produce a pattern '%v' but produces: %v", e.String(), g)
	}
}

func TestDebugLegacy(t *testing.T) {
	Level = 4

	e, err := regexp.Compile(fmt.Sprintf(formatExp, PreDebug))
	g := captureLoggerOutput(Debug, format, a)

	if err != nil {
		t.Fatalf("Failed to compile regexp '%v': %v", e.String(), err)
	}

	if !e.MatchString(g) {
		t.Fatalf("Info should produce a pattern '%v' but produces: %v", e.String(), g)
	}
}

func TestWarningLegacy(t *testing.T) {
	Level = 2

	e, err := regexp.Compile(fmt.Sprintf(formatExp, PreWarning))
	g := captureLoggerOutput(Warning, format, a)

	if err != nil {
		t.Fatalf("Failed to compile regexp '%v': %v", e.String(), err)
	}

	if !e.MatchString(g) {
		t.Fatalf("Info should produce a pattern '%v' but produces: %v", e.String(), g)
	}
}

func TestWithoutTimestampsLegacy(t *testing.T) {
	Timestamps = false
	e, err := regexp.Compile(fmt.Sprintf(formatWOTimeExp, PreAlways))
	g := captureLoggerOutput(Always, format, a)

	if err != nil {
		t.Fatalf("Failed to compile regexp '%v': %v", e.String(), err)
	}

	if !e.MatchString(g) {
		t.Fatalf("Always should produce a pattern '%v' but produces: %v", e.String(), g)
	}
}

// captureLoggerOutput is used to test the log functions
func captureLoggerOutput(l LoggerFunc, format string, a []interface{}) string {
	b := new(bytes.Buffer)
	l(format, append(a, b)...)
	return b.String()
}

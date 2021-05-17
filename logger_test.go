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

package logger

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/fatih/color"
)

const (
	TestFormat = "%s %s %s money, success, fame glamour"
)

var (
	testA = []interface{}{"1", "2", "3"}
)

func TestAlwaysOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogAlways
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreAlways, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestSuccessOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogSuccess
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreSuccess, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestDebugOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogDebug
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreDebug, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestInfoOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogInfo
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreInfo, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestWarningOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogWarning
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreWarning, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestCriticalOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogCritical
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreCritical, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestDeprecatedOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogDeprecated
	buffer := new(bytes.Buffer)
	Writer = buffer
	expected := Line(PreDeprecated, TestFormat, testA...)
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestEverything(t *testing.T) {
	Level = -1
	BitwiseLevel = LogEverything
	buffer := new(bytes.Buffer)
	Writer = buffer
	cases := []string{PreDeprecated, PreDebug, PreInfo, PreWarning, PreCritical, PreSuccess, PreAlways}
	expected := ""
	for _, c := range cases {
		expected = fmt.Sprintf("%s%s", expected, Line(c, TestFormat, testA...))
	}
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

func TestAlwaysCriticalDebugOnly(t *testing.T) {
	Level = -1
	BitwiseLevel = LogAlways | LogCritical | LogDebug
	buffer := new(bytes.Buffer)
	Writer = buffer
	cases := []string{PreDebug, PreCritical, PreAlways}
	expected := ""
	for _, c := range cases {
		expected = fmt.Sprintf("%s%s", expected, Line(c, TestFormat, testA...))
	}
	Deprecated(TestFormat, testA...)
	Debug(TestFormat, testA...)
	Info(TestFormat, testA...)
	Warning(TestFormat, testA...)
	Critical(TestFormat, testA...)
	Success(TestFormat, testA...)
	Always(TestFormat, testA...)
	actual := buffer.String()
	if expected != actual {
		t.Errorf("Expected (%s) Actual (%s)", expected, actual)
	}
}

// Note: Taken from the use case in eksctl
// https://github.com/michaelbeaumont/eksctl/blob/a1564e4cf825be9fd8936794e230d3b5c2d267ea/cmd/eksctl/main.go#L78-L79
func TestLineOverride(t *testing.T) {
	Level = -1
	BitwiseLevel = LogAlways
	Line = func(prefix, format string, a ...interface{}) string {
		if !strings.Contains(format, "\n") {
			format = fmt.Sprintf("%s%s", format, "\n")
		}
		if Timestamps {
			now := time.Now()
			fNow := now.Format(Layout)
			prefix = fmt.Sprintf("%s [%s]", fNow, prefix)
		} else {
			prefix = fmt.Sprintf("[%s]", prefix)
		}
		out := fmt.Sprintf("%s  %s", prefix, fmt.Sprintf(format, a...))
		out = color.GreenString(out)
		return out
	}
	buffer := new(bytes.Buffer)
	Writer = buffer
	Always(TestFormat, testA...)
	actual := buffer.String()
	expected := Line(PreAlways, TestFormat, testA...)
	if actual != expected {
		t.Errorf("actual(%s) != expected(%s)", actual, expected)
		t.FailNow()
	}
}

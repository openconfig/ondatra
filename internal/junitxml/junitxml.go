// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package junitxml provides a mechanism to convert a streamed test log to JUnit XML.
package junitxml

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/jstemmer/go-junit-report/v2/junit"
	"github.com/jstemmer/go-junit-report/v2/parser/gotest"
)

var (
	convSingle *converter // Converter singleton.
	timeNowFn  = time.Now // To be stubbed out by unit tests.
)

// StartConverting starts converting a Go test log in stdout to a JUnit XML
// file written at the specified path. Stop _must_ be called on the returned
// Converter to ensure the XML is fully written.
func StartConverting(xmlPath string) error {
	xmlFile, err := os.Create(xmlPath)
	if err != nil {
		return fmt.Errorf("failed to create JUnit XML file: %w", err)
	}
	conv, err := startConverter(os.Stdout, xmlFile)
	if err != nil {
		return err
	}
	os.Stdout = conv.file
	convSingle = conv
	return err
}

// AddProperty adds a name-value property associated with the given test name
// to the JUnit XML testsuite that is generated.
func AddProperty(test, name, value string) {
	if convSingle != nil {
		convSingle.addProperty(test, name, value)
	}
}

// StopConverting stops converting and waits for the XML to be fully written.
func StopConverting() error {
	if convSingle != nil {
		return convSingle.Stop()
	}
	return nil
}

func startConverter(src, dst io.Writer) (*converter, error) {
	// Bytes written to the piped writer are read from the piped reader,
	// translated to XML, and that XML written to the destination writer.
	pr, pw, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("unable to create file pipe: %w", err)
	}
	tr := io.TeeReader(pr, src)
	errCh := make(chan error)
	conv := &converter{file: pw, errCh: errCh}
	go func() {
		errCh <- conv.logToXML(tr, dst)
	}()
	return conv, nil
}

type converter struct {
	file  *os.File
	errCh <-chan error
	props []junit.Property
}

func (c *converter) addProperty(test, name, value string) {
	c.props = append(c.props, encodeProperty(test, name, value))
}

func (c *converter) Stop() error {
	if err := c.file.Close(); err != nil {
		return err
	}
	return <-c.errCh
}

func (c *converter) logToXML(r io.Reader, w io.Writer) error {
	parser := gotest.NewParser(gotest.TimestampFunc(timeNowFn))
	report, err := parser.Parse(r)
	if err != nil {
		return fmt.Errorf("error parsing test log: %w", err)
	}
	if len(report.Packages) != 1 {
		return fmt.Errorf("expecting 1 generated package but got: %v", report.Packages)
	}
	for _, p := range c.props {
		report.Packages[0].AddProperty(p.Name, p.Value)
	}

	testsuites := junit.CreateFromReport(report, "")
	if err := writeXML(&testsuites, w); err != nil {
		return fmt.Errorf("error writing XML output: %w", err)
	}
	return nil
}

func writeXML(testsuites *junit.Testsuites, w io.Writer) error {
	_, err := fmt.Fprintf(w, xml.Header)
	if err != nil {
		return err
	}
	enc := xml.NewEncoder(w)
	enc.Indent("", "\t")
	if err := enc.Encode(testsuites); err != nil {
		return err
	}
	if err := enc.Flush(); err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "\n")
	return err
}

// encodeProperty encodes a test name and property name-value in a JUnit Property.
func encodeProperty(test, name, value string) junit.Property {
	name = strings.ReplaceAll(name, "/", "//")
	if test != "" {
		name = fmt.Sprintf("%s/%s", test, name)
	}
	return junit.Property{Name: name, Value: value}
}

// DecodeProperty decodes a test name and property name-value from a JUnit Property.
func DecodeProperty(prop junit.Property) (string, string, string) {
	var lastSlashIdx int
	if doubleSlashIdx := strings.Index(prop.Name, "//"); doubleSlashIdx == -1 {
		lastSlashIdx = strings.LastIndex(prop.Name, "/")
	} else {
		lastSlashIdx = strings.LastIndex(prop.Name[0:doubleSlashIdx], "/")
	}
	test := ""
	name := prop.Name
	if lastSlashIdx != -1 {
		test = prop.Name[:lastSlashIdx]
		name = prop.Name[lastSlashIdx+1:]
	}
	name = strings.ReplaceAll(name, "//", "/")
	return test, name, prop.Value
}

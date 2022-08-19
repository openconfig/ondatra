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

// Package junitxml provides a mechanism to covert a streamed test log to JUnit XML.
package junitxml

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jstemmer/go-junit-report/v2/junit"
	"github.com/jstemmer/go-junit-report/v2/parser/gotest"
)

// To be stubbed out by unit tests.
var timeNowFn = time.Now

// StartConverter starts converting a Go test log in stdout to a JUnit XML
// file written at the specified path. Stop _must_ be called on the returned
// Converter to ensure the XML is fully written.
func StartConverter(xmlPath string) (*Converter, error) {
	xmlFile, err := os.Create(xmlPath)
	if err != nil {
		return nil, fmt.Errorf("failed to create JUnit XML file: %w", err)
	}
	conv, err := startConverter(xmlFile)
	if err != nil {
		return nil, err
	}
	os.Stdout = conv.file
	return conv, err
}

func startConverter(dest io.Writer) (*Converter, error) {
	// Bytes written to the piped writer are read from the piped reader,
	// translated to XML, and that XML written to the destination writer.
	pr, pw, err := os.Pipe()
	if err != nil {
		return nil, fmt.Errorf("unable to create file pipe: %w", err)
	}
	errCh := make(chan error)
	go func() {
		errCh <- logToXML(pr, dest)
	}()
	return &Converter{file: pw, errCh: errCh}, nil
}

// Converter converts Go test log format to a JUnit XML file.
type Converter struct {
	file  *os.File
	errCh <-chan error
}

// Stop stops the converter and waits for the XML to be fully written.
func (c *Converter) Stop() error {
	if err := c.file.Close(); err != nil {
		return err
	}
	return <-c.errCh
}

func logToXML(r io.Reader, w io.Writer) error {
	parser := gotest.NewParser(gotest.TimestampFunc(timeNowFn))
	report, err := parser.Parse(r)
	if err != nil {
		return fmt.Errorf("error parsing test log: %w", err)
	}
	// TODO: pass in a hostname?
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

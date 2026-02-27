// Copyright 2026 Google LLC
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

// Package grpclog contains gRPC logging utilities useful to binding implementations.
//
// # Usage:
// Pass the following grpc dial options when creating a gRPC client connection:
//
//	grpcLoggingDialOpts := []grpc.DialOption{
//		grpc.WithChainUnaryInterceptor(grpclog.UnaryClientInterceptor()),
//		grpc.WithChainStreamInterceptor(grpclog.StreamClientInterceptor()),
//	}
//
// # Where do the logs go?
// RPCs made using connections established with these dial options will be logged to:
//  1. Files in the directory specified by the TEST_UNDECLARED_OUTPUTS_DIR environment variable, if set.
//     Otherwise, files in the system temporary directory which is '/tmp' for unix/linux based systems.
//  2. The standard log output at verbosity level 3.
package grpclog

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"golang.org/x/net/context"

	log "github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	grpcTargetRegexp = regexp.MustCompile(`[^a-zA-Z0-9_.-]`)
	testOutputsDir   = os.Getenv("TEST_UNDECLARED_OUTPUTS_DIR")
)

func init() {
	if testOutputsDir == "" {
		testOutputsDir = os.TempDir()
	}
}

// sanitizeTarget makes the target string safe for use as a filename.
func sanitizeTarget(target string) string {
	// Remove common problematic characters for filenames
	return grpcTargetRegexp.ReplaceAllString(target, "_")
}

// writeLog writes the log message inline and to the target specific log file in the test outputs directory.
func writeLog(target string, format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	log.V(3).Infof("[%s] %s", target, msg)

	logFileName := fmt.Sprintf("grpc-target-%s.log", sanitizeTarget(target))
	logFilePath := filepath.Join(testOutputsDir, logFileName)
	logMsg := fmt.Sprintf("[%s] %s\n", time.Now().Format(time.RFC3339Nano), msg)
	if err := appendToFile(logFilePath, logMsg); err != nil {
		log.Errorf("Failed to write log (%s) to file (%s): %v", logMsg, logFilePath, err)
	}
}

// appendToFile appends the msg string to the file at the given path.
func appendToFile(filePath string, msg string) error {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.WriteString(msg)
	return err
}

// UnaryClientInterceptor returns a UnaryClientInterceptor function.
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		target := cc.Target()

		md, _ := metadata.FromOutgoingContext(ctx)

		writeLog(target, "CALL START: Method: %s, Metadata: %+v, Request: %+v", method, md, req)

		startTime := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		duration := time.Since(startTime)

		status := "OK"
		if err != nil {
			status = fmt.Sprintf("ERROR: %v", err)
		}
		writeLog(target, "CALL END: Method: %s, Duration: %s, Status: %s, Reply: %+v", method, duration, status, reply)

		return err
	}
}

// StreamClientInterceptor returns a StreamClientInterceptor function.
func StreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		target := cc.Target()

		md, _ := metadata.FromOutgoingContext(ctx)

		writeLog(target, "CALL START: Method: %s, Metadata: %+v, StreamDesc: %+v", method, md, desc)

		clientStream, err := streamer(ctx, desc, cc, method, opts...)
		if err != nil {
			writeLog(target, "CALL END: Failed to start stream (Method: %s, Metadata: %+v, StreamDesc: %+v): %v", method, md, desc, err)
			return nil, err
		}

		return &wrappedClientStream{
			ClientStream: clientStream,
			method:       method,
			target:       target,
		}, nil
	}
}

// wrappedClientStream wraps grpc.ClientStream to log messages.
type wrappedClientStream struct {
	grpc.ClientStream
	method string
	target string
}

func (w *wrappedClientStream) SendMsg(m any) error {
	err := w.ClientStream.SendMsg(m)
	if err != nil {
		writeLog(w.target, "(Stream %s) SendMsg error: %v", w.method, err)
	} else {
		writeLog(w.target, "(Stream %s) SendMsg success: %+v", w.method, m)
	}
	return err
}

func (w *wrappedClientStream) RecvMsg(m any) error {
	err := w.ClientStream.RecvMsg(m)
	if err != nil {
		writeLog(w.target, "(Stream %s) RecvMsg error: %v", w.method, err)
	} else {
		writeLog(w.target, "(Stream %s) RecvMsg success: %+v", w.method, m)
	}
	return err
}

func (w *wrappedClientStream) Header() (metadata.MD, error) {
	md, err := w.ClientStream.Header()
	if err != nil {
		writeLog(w.target, "(Stream %s) Header error: %v", w.method, err)
	} else {
		writeLog(w.target, "(Stream %s) Header success: %+v", w.method, md)
	}
	return md, err
}

func (w *wrappedClientStream) Trailer() metadata.MD {
	md := w.ClientStream.Trailer()
	writeLog(w.target, "(Stream %s) Trailer: %+v", w.method, md)
	return md
}

// Copyright 2021 Google LLC
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

// Package dnbind provides an Ondatra binding for Drivenets devices.
package dnbind

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/pborman/uuid"
	"golang.org/x/crypto/ssh"

	log "github.com/golang/glog"
	"github.com/openconfig/ondatra/binding"
	opb "github.com/openconfig/ondatra/proto"
)

// New returns a new Drivenets bind instance.
func New(cfg *Config) (*Bind, error) {
	if cfg == nil {
		return nil, fmt.Errorf("config cannot be nil")
	}

	return &Bind{cfg: cfg}, nil
}

// Bind implements the ondatra Binding interface for Drivenets
type Bind struct {
	binding.Binding
	cfg *Config
}

// Reserve implements the binding Reserve method by finding nodes and links in
// the topology specified in the config file that match the requested testbed.
func (b *Bind) Reserve(ctx context.Context, tb *opb.Testbed, runTime time.Duration, waitTime time.Duration, partial map[string]string) (*binding.Reservation, error) {

	res := &binding.Reservation{
		ID:   uuid.New(),
		DUTs: make(map[string]binding.DUT),
		ATEs: make(map[string]binding.ATE),
	}

	for _, dutInfo := range tb.Duts {
		if dutInfo.Vendor != opb.Device_VENDOR_UNSPECIFIED &&
			dutInfo.Vendor != opb.Device_DRIVENETS {
			// not a Drivenets device; ignore
			continue
		}

		node := b.cfg.Lookup(dutInfo.Id)
		if node == nil {
			return nil, fmt.Errorf("missing credetials for dut %s", dutInfo.Id)
		}

		dims := &binding.Dims{
			Name:            node.Hostname,
			Vendor:          opb.Device_DRIVENETS,
			HardwareModel:   node.HardwareModel,
			SoftwareVersion: node.SoftwareVersion,
			Ports:           make(map[string]*binding.Port),
		}
		res.DUTs[dutInfo.Id] = &dnDUT{
			AbstractDUT: &binding.AbstractDUT{dims},
			bind:        b,
		}
	}

	return res, nil
}

// Release is a no-op because there's no need to release Drivenets devices.
func (b *Bind) Release(context.Context) error {
	return nil
}

type dnDUT struct {
	*binding.AbstractDUT
	bind *Bind
	cli  *dnCLI
}

func (dut *dnDUT) DialCLI(ctx context.Context) (binding.CLIClient, error) {
	if dut.cli != nil {
		return dut.cli, nil
	}

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(ctx, time.Minute)
	defer cancel()

	up := dut.bind.cfg.Credentials.Node[dut.Name()]

	var timeout time.Duration = 0
	deadline, ok := ctx.Deadline()
	if ok {
		timeout = time.Until(deadline)
	}

	sshConfig := &ssh.ClientConfig{
		User:            up.Username,
		Auth:            []ssh.AuthMethod{ssh.Password(up.Password)},
		Timeout:         timeout,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := net.JoinHostPort(dut.Name(), "22")

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	dialCli := func() (err error) {
		dut.cli = &dnCLI{}

		if dut.cli.client, err = ssh.Dial("tcp", addr, sshConfig); err != nil {
			return err
		}

		if dut.cli.session, err = dut.cli.client.NewSession(); err != nil {
			return err
		}

		if err = dut.cli.session.RequestPty("xterm", 80, 40, modes); err != nil {
			return err
		}

		if dut.cli.stdin, err = dut.cli.session.StdinPipe(); err != nil {
			return err
		}

		if dut.cli.stdout, err = dut.cli.session.StdoutPipe(); err != nil {
			return err
		}

		if err = dut.cli.session.Shell(); err != nil {
			return err
		}

		return nil
	}

	var err error
	for {
		select {
		case <-ctx.Done():
			return nil, err
		default:
			if err = dialCli(); err != nil {
				dut.cli.Close()
				dut.cli = nil
				log.Errorf("%s: %s; retrying for %s\n", dut.Name(),
					err, time.Until(deadline).Round(time.Millisecond))
				time.Sleep(time.Second * 10)
				break
			}

			// wait for prompt
			if _, err = dut.cli.CommandResult(ctx); err != nil {
				return nil, err
			}

			return dut.cli, nil
		}
	}
}

func (dut *dnDUT) PushConfig(ctx context.Context, config string, reset bool) error {
	if _, err := dut.DialCLI(ctx); err != nil {
		return err
	}

	log.Infof("Pushing config:\n\"%s\"", config)

	check := func(res binding.CommandResult, err error) error {
		if err == nil && len(res.Error()) == 0 {
			return nil
		}
		// revert current changes
		dut.cli.RunCommand(ctx, "rollback 0")
		// exit configure menu
		dut.cli.RunCommand(ctx, "end")
		// propagate error or stderr
		if err != nil {
			return err
		}
		return errors.New(res.Error())
	}

	commands := []string{"configure"}
	if reset {
		commands = append(commands, "load override factory-default")
	}
	commands = append(commands, strings.Split(config, "\n")...)
	commands = append(commands, []string{"commit check", "commit", "end"}...)

	for _, command := range commands {
		if err := check(dut.cli.RunCommand(ctx, command)); err != nil {
			return err
		}
	}

	return nil
}

type dnCLI struct {
	*binding.AbstractCLIClient
	client  *ssh.Client
	session *ssh.Session
	stdin   io.WriteCloser
	stdout  io.Reader
}

func (c *dnCLI) Close() error {
	if c.stdin != nil {
		c.stdin.Close()
	}
	if c.session != nil {
		c.session.Close()
	}
	if c.client != nil {
		c.client.Close()
	}
	return nil
}

func (c *dnCLI) RunCommand(ctx context.Context, cmd string) (binding.CommandResult, error) {
	cmd = strings.Replace(cmd, "\t", " ", -1) + "\n"
	if _, err := c.stdin.Write([]byte(cmd)); err != nil {
		return nil, err
	}

	return c.CommandResult(ctx)
}

func (c *dnCLI) CommandResult(ctx context.Context) (r cmdResult, err error) {
	buffer := make([]byte, 10000000)

	for {
		byteCount, err := c.stdout.Read(buffer)
		if err != nil {
			return r, err
		}
		log.Infof(string(buffer[:byteCount]))
		lines := strings.Split(string(buffer[:byteCount]), "\n")

		for _, line := range lines {
			line = strings.TrimRight(line, " ")
			if len(line) > 0 {
				if line[len(line)-1:] == "#" {
					return r, nil
				}
			}
			r.output += line + "\n"
			if strings.HasPrefix(line, "ERROR:") {
				r.error = r.output
				r.output = ""
			}
		}
	}
}

type cmdResult struct {
	*binding.AbstractCommandResult
	output, error string
}

// Output logs a fatal unimplemented error.
func (r cmdResult) Output() string {
	return r.output
}

// Error logs a fatal unimplemented error.
func (r cmdResult) Error() string {
	return r.error
}

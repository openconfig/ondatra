// Package reservemain provides the logic for the _reserve debug target.
package reservemain

import (
	"golang.org/x/net/context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"flag"
	"github.com/openconfig/ondatra/closer"
	"github.com/openconfig/ondatra/internal/reservation"
	"github.com/openconfig/ondatra/internal/testbed"
)

// NOTE: The ondatra.bzl build rules set these flags to different defaults,
// depending on whether Ondatra is executed as a test or a reserve binary.
var (
	// TestbedPath is a flag for the testbed path.
	TestbedPath = flag.String("testbed", "", "Path to the Ondatra testbed file")
	// RunTime is a flag for the run time duration of the reservation.
	RunTime = flag.Duration("run_time", 0, "Timeout of the test run, excluding the wait time for the testbed to be ready. "+
		" Must be a positive value.")
	// WaitTime is a flag for the wait time duration of the reservation.
	WaitTime = flag.Duration("wait_time", 0, "Maximum amount of time the test should wait until the testbed is ready. "+
		"A zero value lets the binding implementation choose an appropriate wait time. Must be a non-negative value.")
)

// Main can be
func Main() {
	if err := run(context.Background(), *TestbedPath, *RunTime, *WaitTime); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, testbedPath string, runTime, waitTime time.Duration) (rerr error) {
	const addr = "blade:netops-lab-reservation"
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Wait for SIGINT callback.
	waitc := make(chan struct{})
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, os.Interrupt)

	go func() {
		select {
		case <-sigc:
			fmt.Println("\nReservation canceled.")
			cancel()
		case <-waitc:
			select {
			case <-time.After(runTime):
				fmt.Printf("Reservation expired after %v.\n", runTime)
				sigc <- nil
			}
		}
	}()

	fmt.Printf("Waiting for reservation of %s ...\n", testbedPath)
	if err := testbed.Reserve(ctx, testbedPath, runTime, waitTime); err != nil {
		return err
	}
	defer closer.Close(&rerr, func() error {
		return testbed.Release(ctx)
	}, "error releasing testbed")
	waitc <- struct{}{}

	res, err := testbed.Reservation()
	if err != nil {
		return err
	}
	fmt.Printf("Created reservation http://resv/?id=%s for %s\n\n", res.ID, runTime)
	printReservation(res)
	fmt.Println()
	fmt.Println("*** Physical devices are a shared and limited resource")
	fmt.Println("*** Please release the reservation as soon as you can!")
	fmt.Println()
	fmt.Printf("Run blaze test <TEST_TARGET> --test_env=ONDATRA_RESERVE=%s\n", res.ID)
	fmt.Println("Press CTRL-C to Release")

	// Wait for SIGINT.
	<-sigc
	return nil
}

func printReservation(res *reservation.Reservation) {
	fmt.Printf("Reservation Assignment:\n")
	printf := func(id, name string) {
		fmt.Printf("  %-17s %s\n", id+":", name)
	}
	for id, d := range res.DUTs {
		printf(id, d.Name)
		for pid, p := range d.Ports {
			printf(pid, p.Name)
		}
	}
	for id, a := range res.ATEs {
		printf(id, a.Name)
		for pid, p := range a.Ports {
			printf(pid, p.Name)
		}
	}
}

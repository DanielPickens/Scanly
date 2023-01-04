package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"runtime/trace"

	"github.com/urfave/cli"
)

func main() {
	err := cli.NewApp().Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v (try --help)\n", err)
		os.Exit(1)
		log.Fatal()
	}
}

type NewCLI struct {
	args   []string
	config *Config
}

type Config struct {
	// CPU profile
	CPUProfile string
	// Memory profile
	MemProfile string
	// Trace profile
	TraceProfile string
	// Number of goroutines to use
	NumGoroutines int
	// Number of iterations to run
	NumIterations int
	// Number of bytes to allocate
	NumBytes int
}

type CLI struct {
	// Configuration
	config *Config
	// Command line arguments
	args []string
}

func (cli *CLI) parseArgs() error {
	if len(cli.args) < 1 {
		return fmt.Errorf("missing command")
	}
	switch cli.args[0] {
	case "alloc":
		cli.config.NumGoroutines = 1
		cli.config.NumIterations = 1
		cli.config.NumBytes = 1024
	case "alloc-goroutine":
		cli.config.NumGoroutines = 1
		cli.config.NumIterations = 1
		cli.config.NumBytes = 1024
	case "alloc-iteration":
		cli.config.NumGoroutines = 1
		cli.config.NumIterations = 1
		cli.config.NumBytes = 1024
	case "alloc-goroutine-iteration":
		cli.config.NumGoroutines = 1
		cli.config.NumIterations = 1
		cli.config.NumBytes = 1024
	case "alloc-goroutine-iteration-heap":
		cli.config.NumGoroutines = 1
		cli.config.NumIterations = 1
		cli.config.NumBytes = 1024
	default:
		return fmt.Errorf("unknown command: %s", cli.args[0])
	}
	return nil
}

func (cli *CLI) Run() error {
	// Parse the command line arguments
	if err := cli.parseArgs(); err != nil {
		return err
	}

	// Start the CPU profile
	if cli.config.CPUProfile != "" {
		f, err := os.Create(cli.config.CPUProfile)
		if err != nil {
			return err
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	// Start the trace profile
	if cli.config.TraceProfile != "" {
		f, err := os.Create(cli.config.TraceProfile)
		if err != nil {
			return err
		}
		trace.Start(f)
		defer trace.Stop()
	}

	// Write the memory profile
	if cli.config.MemProfile != "" {
		f, err := os.Create(cli.config.MemProfile)
		if err != nil {
			return err
		}
		pprof.WriteHeapProfile(f)
		f.Close()
	}

	return nil
}

func (cli *CLI) runBenchmark() error {
	switch cli.args[0] {
	case "alloc":
		return cli.alloc()
	case "alloc-goroutine":
		return cli.allocGoroutine()
	case "alloc-iteration":
		return cli.allocIteration()
	case "alloc-goroutine-iteration":
		return cli.allocGoroutineIteration()
	case "alloc-goroutine-iteration-heap":
		return cli.allocGoroutineIterationHeap()
	default:
		return fmt.Errorf("unknown command: %s", cli.args[0])
	}
}

func (cli *CLI) alloc() error {
	for i := 0; i < cli.config.NumIterations; i++ {
		_ = make([]byte, cli.config.NumBytes)
	}
	return nil
}

func (cli *CLI) allocGoroutine() error {
	for i := 0; i < cli.config.NumGoroutines; i++ {
		go func() {
			for j := 0; j < cli.config.NumIterations; j++ {
				_ = make([]byte, cli.config.NumBytes)
			}
		}()
	}
	return nil
}

func (cli *CLI) allocIteration() error {
	for i := 0; i < cli.config.NumIterations; i++ {
		for j := 0; j < cli.config.NumGoroutines; j++ {
			_ = make([]byte, cli.config.NumBytes)
		}
	}
	return nil
}

func (cli *CLI) allocGoroutineIteration() error {
	for i := 0; i < cli.config.NumGoroutines; i++ {
		go func() {
			for j := 0; j < cli.config.NumIterations; j++ {
				_ = make([]byte, cli.config.NumBytes)
			}
		}()
	}
	return nil
}

func (cli *CLI) allocGoroutineIterationHeap() error {
	for i := 0; i < cli.config.NumGoroutines; i++ {
		go func() {
			for j := 0; j < cli.config.NumIterations; j++ {
				_ = make([]byte, cli.config.NumBytes)

			}
		}()
	}
	return nil
}

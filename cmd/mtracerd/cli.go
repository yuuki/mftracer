package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/yuuki/mkr-flow-tracer/agent"
	"github.com/yuuki/mkr-flow-tracer/db"
)

const (
	exitCodeOK         = 0
	exitCodeErr        = 10 + iota
	defaultIntervalSec = 30
)

// CLI is the command line object.
type CLI struct {
	// outStream and errStream are the stdout and stderr
	// to write message from the CLI.
	outStream, errStream io.Writer
}

// Run execute the main process.
// It returns exit code.
func (c *CLI) Run(args []string) int {
	log.SetOutput(c.errStream)

	var (
		ver bool
	)
	flags := flag.NewFlagSet("mtracerd", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.Usage = func() {
		fmt.Fprint(c.errStream, helpText)
	}
	flags.BoolVar(&ver, "version", false, "")
	if err := flags.Parse(args[1:]); err != nil {
		return exitCodeErr
	}

	if ver {
		// fmt.Fprintf(c.errStream, "%s version %s, build %s, date %s \n", name, version, commit, date)
		return exitCodeOK
	}

	log.Println("--> Connecting postgres ...")
	db, err := db.New()
	if err != nil {
		log.Printf("postgres initialize error: %v\n", err)
		return exitCodeErr
	}
	log.Println("Connected postgres")

	agent.Start(defaultIntervalSec*time.Second, db)

	return exitCodeOK
}

var helpText = `Usage: mtracerd [options]

  

Options:
  --version, -v	            print version
  --help, -h                print help
`

package main

import (
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/yuuki/mftracer/db"
)

const (
	exitCodeOK  = 0
	exitCodeErr = 10 + iota
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
		ver          bool
		createSchema bool
	)
	flags := flag.NewFlagSet("mftracer", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.Usage = func() {
		fmt.Fprint(c.errStream, helpText)
	}
	flags.BoolVar(&createSchema, "create-schema", false, "")
	flags.BoolVar(&ver, "version", false, "")
	if err := flags.Parse(args[1:]); err != nil {
		return exitCodeErr
	}

	if ver {
		// fmt.Fprintf(c.errStream, "%s version %s, build %s, date %s \n", name, version, commit, date)
		return exitCodeOK
	}

	if createSchema {
		log.Println("Connecting postgres ...")
		db, err := db.New()
		if err != nil {
			log.Printf("postgres initialize error: %v\n", err)
			return exitCodeErr
		}

		log.Println("Creating postgres schema ...")
		if err := db.CreateSchema(); err != nil {
			log.Printf("postgres initialize error: %v\n", err)
			return exitCodeErr
		}
		return exitCodeOK
	}

	return exitCodeOK
}

var helpText = `Usage: mtracer [options]

  

Options:
  --version, -v	            print version
  --help, -h                print help
`

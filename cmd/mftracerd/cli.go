//go:generate go-bindata -pkg main -o credits.go ../../CREDITS
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/yuuki/mftracer/agent"
	"github.com/yuuki/mftracer/db"
)

const (
	exitCodeOK         = 0
	exitCodeErr        = 10 + iota
	defaultIntervalSec = 30
)

var (
	creditsText = string(MustAsset("CREDITS"))
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
		ver     bool
		credits bool

		once   bool
		dbuser string
		dbpass string
		dbhost string
		dbport string
		dbname string
	)
	flags := flag.NewFlagSet("mftracerd", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.Usage = func() {
		fmt.Fprint(c.errStream, helpText)
	}
	flags.BoolVar(&once, "once", false, "")
	flags.StringVar(&dbuser, "dbuser", "", "")
	flags.StringVar(&dbpass, "dbpass", "", "")
	flags.StringVar(&dbhost, "dbhost", "", "")
	flags.StringVar(&dbport, "dbport", "", "")
	flags.StringVar(&dbname, "dbname", "", "")
	flags.BoolVar(&ver, "version", false, "")
	flags.BoolVar(&credits, "credits", false, "")
	if err := flags.Parse(args[1:]); err != nil {
		return exitCodeErr
	}

	if ver {
		// fmt.Fprintf(c.errStream, "%s version %s, build %s, date %s \n", name, version, commit, date)
		return exitCodeOK
	}

	if credits {
		fmt.Fprintln(c.outStream, creditsText)
		return exitCodeOK
	}

	log.Println("--> Connecting postgres ...")
	db, err := db.New(&db.Opt{
		DBName:   dbname,
		User:     dbuser,
		Password: dbpass,
		Host:     dbhost,
		Port:     dbport,
	})
	if err != nil {
		log.Printf("postgres initialize error: %v\n", err)
		return exitCodeErr
	}
	log.Println("Connected postgres")

	if once {
		if err := agent.RunOnce(db); err != nil {
			log.Printf("%+v\n", err)
			return exitCodeErr
		}
	} else {
		agent.Start(defaultIntervalSec*time.Second, db)
	}

	return exitCodeOK
}

var helpText = `Usage: mtracerd [options]

  

Options:
  --once                    run once
  --dbuser                  postgres user
  --dbpass                  postgres user password
  --dbhost                  postgres host
  --dbport                  postgres port
  --dbname                  postgres database name
  --version, -v	            print version
  --help, -h                print help
`

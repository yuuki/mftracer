package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

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
		dbuser       string
		dbpass       string
		dbhost       string
		dbport       string
		dbname       string
		destipv4     string
	)
	flags := flag.NewFlagSet("mftracer", flag.ContinueOnError)
	flags.SetOutput(c.errStream)
	flags.Usage = func() {
		fmt.Fprint(c.errStream, helpText)
	}
	flags.BoolVar(&createSchema, "create-schema", false, "")
	flags.StringVar(&destipv4, "dest-ipv4", "", "")
	flags.StringVar(&dbuser, "dbuser", "", "")
	flags.StringVar(&dbpass, "dbpass", "", "")
	flags.StringVar(&dbhost, "dbhost", "", "")
	flags.StringVar(&dbport, "dbport", "", "")
	flags.StringVar(&dbname, "dbname", "", "")
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

		log.Println("Creating postgres schema ...")
		if err := db.CreateSchema(); err != nil {
			log.Printf("postgres initialize error: %v\n", err)
			return exitCodeErr
		}
		return exitCodeOK
	}

	if destipv4 != "" {
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

		ip := net.ParseIP(destipv4)
		addrports, err := db.FindSourceByDestIPAddr(ip)
		if err != nil {
			log.Printf("postgres find source error: %v\n", err)
			return exitCodeErr
		}
		fmt.Println(ip)
		for _, addrport := range addrports {
			fmt.Print("└<-- ")
			fmt.Println(addrport)
		}
	}

	return exitCodeOK
}

var helpText = `Usage: mtracer [options]

  

Options:
  --create-schema           create mftracer table schema for postgres
  --dbuser                  postgres user
  --dbpass                  postgres user password
  --dbhost                  postgres host
  --dbport                  postgres port
  --dbname                  postgres database name
  --version, -v	            print version
  --help, -h                print help
`

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"github.com/yuuki/mftracer/db"
)

const (
	exitCodeOK    = 0
	exitCodeErr   = 10 + iota
	maxGraphDepth = 4
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
		depth        int
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
	flags.IntVar(&depth, "L", maxGraphDepth, "") // level
	flags.IntVar(&depth, "depth", maxGraphDepth, "")
	flags.BoolVar(&ver, "version", false, "")
	if err := flags.Parse(args[1:]); err != nil {
		return exitCodeErr
	}

	if ver {
		// fmt.Fprintf(c.errStream, "%s version %s, build %s, date %s \n", name, version, commit, date)
		return exitCodeOK
	}

	dbopt := &db.Opt{
		DBName:   dbname,
		User:     dbuser,
		Password: dbpass,
		Host:     dbhost,
		Port:     dbport,
	}

	if depth <= 0 || depth > maxGraphDepth {
		log.Printf("depth must be 0 < depth < %d, but specified %d\n", maxGraphDepth, depth)
		return exitCodeErr
	}

	if createSchema {
		return c.createSchema(dbopt)
	}

	if destipv4 != "" {
		return c.destIPv4(destipv4, depth, dbopt)
	}

	return exitCodeOK
}

func (c *CLI) createSchema(opt *db.Opt) int {
	log.Println("Connecting postgres ...")

	db, err := db.New(opt)
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

func (c *CLI) destIPv4(ipv4 string, depth int, opt *db.Opt) int {
	db, err := db.New(opt)
	if err != nil {
		log.Printf("postgres initialize error: %v\n", err)
		return exitCodeErr
	}
	ip := net.ParseIP(ipv4)
	fmt.Println(ipv4)
	if err := c.printDestIPAddr(db, ip, 1, depth); err != nil {
		return exitCodeErr
	}
	return exitCodeOK
}

func (c *CLI) printDestIPAddr(db *db.DB, ipv4 net.IP, curDepth, depth int) error {
	addrports, err := db.FindSourceByDestIPAddr(ipv4)
	if err != nil {
		return err
	}
	if len(addrports) == 0 {
		return nil
	}
	indent := strings.Repeat("\t", curDepth-1)
	curDepth++
	depth--
	for _, addrport := range addrports {
		fmt.Print(indent)
		fmt.Print("└<-- ")
		fmt.Println(addrport)
		if err := c.printDestIPAddr(db, addrport.IPAddr, curDepth, depth); err != nil {
			return err
		}
	}
	return nil
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

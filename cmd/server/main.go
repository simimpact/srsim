package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/simimpact/srsim/pkg/servermode"
	// gotta manually import stats agg since im not using the nicer srsim executable
	_ "github.com/simimpact/srsim/pkg/statistics/agg/overview"
)

var (
	version string
)

type opts struct {
	host        string
	port        string
	timeout     int
	showVersion bool
}

func main() {
	var opt opts
	flag.StringVar(&opt.host, "host", "localhost", "host to listen to (default: localhost)")
	flag.StringVar(&opt.port, "port", "54321", "port to listen on (default: 54321)")
	flag.IntVar(&opt.timeout, "timeout", 5*60, "how long to run each sim for in seconds before timing out (default: 300s)")
	flag.BoolVar(&opt.showVersion, "version", false, "show currrent version")
	flag.Parse()

	if opt.showVersion {
		fmt.Println("Running version: ", version)
		return
	}

	server, err := servermode.New(
		servermode.WithDefaults(),
		servermode.WithTimeout(time.Duration(opt.timeout)*time.Second),
	)

	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", opt.host, opt.port), server.Router))
}

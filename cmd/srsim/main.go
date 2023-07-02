package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/simimpact/srsim/pkg/logic/gcs/parse"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/pkg/simulator"
)

var (
	sha1ver string // sha1 revision used to build the program
	// buildTime string // when the executable was built
)

type opts struct {
	config string
	script string
}

// command line tool; following options are available:
func main() {
	var opt opts
	var version bool
	flag.BoolVar(&version, "version", false, "check gcsim version (git hash)")
	flag.StringVar(&opt.config, "c", "config.txt", "which config to use (json format)")
	flag.StringVar(&opt.script, "s", "script.txt", "which script to use")

	flag.Parse()

	if version {
		fmt.Println(sha1ver)
		return
	}

	_, err := os.Stat(opt.config)
	usedCLI := false
	flag.Visit(func(f *flag.Flag) {
		usedCLI = true
	})
	if errors.Is(err, os.ErrNotExist) && !usedCLI {
		fmt.Printf("The file %s does not exist.\n", opt.config)
		return
	}

	config, err := ReadConfig(opt.config)
	if err != nil {
		fmt.Println(err)
		return
	}

	script, err := os.ReadFile(opt.script)
	if err != nil {
		fmt.Println(err)
		return
	}
	p := parse.New(string(script))
	list, err := p.Parse()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = simulator.Run(context.Background(), list, config)
	if err != nil {
		log.Println(err)
		return
	}
}

func ReadConfig(path string) (*model.SimConfig, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	res := &model.SimConfig{}
	err = json.Unmarshal(src, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

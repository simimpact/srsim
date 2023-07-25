package main

import (
	"fmt"
	"log"
	"os"

	"github.com/simimpact/srsim/pkg/simulation"
	"github.com/urfave/cli/v2"
)

// TODO: example config, generate when no config found
// TODO: key signing

var globalFlags = []cli.Flag{
	&cli.StringFlag{
		Name:    "outpath",
		Aliases: []string{"output", "out", "o"},
		Usage:   "output path for generated results",
		Value:   "result/",
	},
	&cli.Int64Flag{
		Name:        "seed",
		Usage:       "optional seed for deterministic executions",
		DefaultText: "random",
	},
	&cli.BoolFlag{
		Name:    "keep-serve",
		Aliases: []string{"keep", "ks", "k"},
		Usage:   "keeps the web server running even after its read by UI",
	},
	&cli.BoolFlag{
		Name:    "no-serve",
		Aliases: []string{"ns", "n"},
		Usage:   "skips running the web server and only generates results",
	},
	&cli.StringFlag{
		Name:    "key",
		Usage:   "used to sign the sim for share uploads",
		EnvVars: []string{"SRSIM_SHARE_KEY"},
		Hidden:  true,
	},
}

var runFlags = []cli.Flag{
	&cli.IntFlag{
		Name:    "iterations",
		Aliases: []string{"itrs", "itr", "i"},
		Value:   1000,
		Usage:   "iterations/battles to simulate",
	},
	&cli.IntFlag{
		Name:    "workers",
		Aliases: []string{"w"},
		Value:   10,
		Usage:   "number of workers",
	},
}

var debugFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:    "stdout",
		Aliases: []string{"sout", "std", "s"},
		Usage:   "enable logs also print out to stdout",
	},
}

var version string

func init() {
	// TODO: get version from tag & gh workflow (https://www.forkingbytes.com/blog/dynamic-versioning-your-go-application/)
	// default version commit hash?
	cli.VersionPrinter = func(ctx *cli.Context) {
		fmt.Fprintf(ctx.App.Writer, "srsim version %s (%s)\n", ctx.App.Version, "")
	}
}

func main() {
	app := &cli.App{
		Name:     "srsim",
		Usage:    "a Honkai: Star Rail monte carlo battle simulator",
		Version:  version,
		HideHelp: true,
		Commands: []*cli.Command{
			{
				Name:                   "run",
				Usage:                  "executes a simulation",
				UsageText:              "srsim run [options] <path to config>",
				UseShortOptionHandling: true,
				HideHelpCommand:        true,
				Flags:                  concatMultipleSlices(runFlags, globalFlags),
				Action:                 run,
			},
			{
				Name:                   "debug",
				Usage:                  "debug a simulation",
				UsageText:              "srsim debug [options] <path to config>",
				UseShortOptionHandling: true,
				HideHelpCommand:        true,
				Flags:                  concatMultipleSlices(debugFlags, globalFlags),
				Action:                 run,
			},
			{
				Name:     "version",
				Usage:    "print the version",
				HideHelp: true,
				Action: func(ctx *cli.Context) error {
					cli.ShowVersion(ctx)
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func run(ctx *cli.Context) error {
	var err error

	if ctx.IsSet("no-serve") && ctx.IsSet("keep-serve") {
		return cli.Exit("cannot use both no-serve & keep-serve in same command", 1)
	}

	// TODO: default template print
	if ctx.NArg() != 1 {
		cli.ShowSubcommandHelpAndExit(ctx, 1)
	}

	// validate that input path exists
	path := ctx.Args().Get(0)
	if err := validatePath(path); err != nil {
		return cli.Exit(err, 1)
	}

	// parse path -> SimConfig
	config, err := parseConfig(path)
	if err != nil {
		return cli.Exit(err, 1)
	}

	// TODO: parseLogic, returns ActionList
	list, err := parseLogic(config)
	if err != nil {
		return cli.Exit(err, 1)
	}

	// get or generate seed
	seed, err := seed(ctx)
	if err != nil {
		return cli.Exit(err, 1)
	}

	if err := os.MkdirAll(ctx.String("outpath"), os.ModeDir); err != nil {
		return cli.Exit(err, 1)
	}

	opts := &ExecutionOpts{
		config:     config,
		list:       list,
		configPath: path,
		seed:       seed,
		outpath:    ctx.String("outpath"),
		iterations: ctx.Int("iterations"),
		workers:    ctx.Int("workers"),
		debug:      ctx.Command.Name == "debug",
		silent:     !ctx.Bool("stdout"),
	}

	if err := execute(opts); err != nil {
		return cli.Exit(err, 1)
	}

	if !ctx.Bool("no-serve") {
		serve(&ServeOpts{
			outpath:   ctx.String("outpath"),
			keepAlive: ctx.Bool("keep-serve"),
		})
	}
	return nil
}

func seed(ctx *cli.Context) (int64, error) {
	if !ctx.IsSet("seed") {
		return simulation.RandSeed()
	}
	return ctx.Int64("seed"), nil
}

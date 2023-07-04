package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/simimpact/srsim/pkg/engine/logging"
	"github.com/simimpact/srsim/pkg/logic/gcs"
	"github.com/simimpact/srsim/pkg/logic/gcs/eval"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/pkg/simulation"
)

type ExecutionOpts struct {
	config     *model.SimConfig
	list       *gcs.ActionList
	configPath string
	outpath    string
	seed       int64
	iterations int
	workers    int
	debug      bool
	silent     bool
}

func execute(opts *ExecutionOpts) error {
	title("--- execution settings ---")
	fmt.Printf("%s %s\n", item("config:  "), opts.configPath)
	fmt.Printf("%s %s\n", item("outpath: "), opts.outpath)

	if !opts.debug {
		fmt.Printf("%s %v\n", item("iters:   "), opts.iterations)
		fmt.Printf("%s %v\n", item("workers: "), opts.workers)
	}

	fmt.Printf("%s %v\n", item("debug:   "), opts.debug)
	fmt.Printf("%s %s\n", item("seed:    "), strconv.FormatUint(uint64(opts.seed), 10))
	fmt.Println()

	os.Remove(LogFile(opts.outpath))
	os.Remove(ResultFile(opts.outpath))

	title("--- logged iteration ---")
	if err := executeLogging(opts); err != nil {
		return err
	}
	fmt.Printf("logs generated: %s\n", LogFile(opts.outpath))

	if opts.debug {
		fmt.Println()
		return nil
	}

	fmt.Println()
	title("--- simulation ---")

	result, err := executeSimulation(opts)
	if err != nil {
		return err
	}

	if err := WriteResult(result, opts.outpath); err != nil {
		return err
	}
	fmt.Printf("results generated: %s\n", ResultFile(opts.outpath))
	fmt.Println()

	return nil
}

func executeLogging(opts *ExecutionOpts) error {
	// TODO: need more elegant ways to capture errors and bubble up?
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic occurred: %v\n", err)
		}
	}()

	fileLogger, err := NewGzipLogger(opts.outpath)
	if err != nil {
		return err
	}
	defer fileLogger.Flush()

	loggers := make([]logging.Logger, 0, 2)
	if opts.debug && !opts.silent {
		loggers = append(loggers, fileLogger, ConsoleLogger{})
	} else {
		loggers = append(loggers, fileLogger)
	}

	_, err = simulation.Run(&simulation.RunOpts{
		Config:  opts.config,
		Eval:    eval.New(context.TODO(), opts.list.Program),
		Seed:    opts.seed,
		Loggers: loggers,
	})
	return err
}

func executeSimulation(opts *ExecutionOpts) (*model.SimResult, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	result := simulation.CreateResult(opts.config, opts.seed)
	aggs, err := simulation.InitializeAggregators(opts.iterations, opts.config)
	if err != nil {
		return result, err
	}
	fmt.Printf("sim initialized! starting %v workers...\n", opts.workers)

	p := createPool(ctx, opts)
	fmt.Printf("workers started, executing %v iterations...\n", opts.iterations)

	bar := progressbar.NewOptions(
		opts.iterations,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetDescription("status "),
		progressbar.OptionSetPredictTime(true),
		progressbar.OptionShowCount(),
		progressbar.OptionSetRenderBlankState(true),
		progressbar.OptionShowElapsedTimeOnFinish(),
		progressbar.OptionShowIts(),
		progressbar.OptionThrottle(5*time.Millisecond),
	)

	go p.start(opts)
	for i := 0; i < opts.iterations; i++ {
		p.processWorkerResult(aggs)
		bar.Add(1)
	}
	fmt.Println()

	// TODO: run complete
	result.Statistics = aggs.Flush()
	return result, nil
}

type job struct {
	config *model.SimConfig
	list   *gcs.ActionList
	seed   int64
}

type pool struct {
	ctx      context.Context
	errChan  chan error
	respChan chan *model.IterationResult
	workChan chan job
}

func createPool(ctx context.Context, opts *ExecutionOpts) *pool {
	p := &pool{
		ctx:      ctx,
		errChan:  make(chan error),
		respChan: make(chan *model.IterationResult),
		workChan: make(chan job),
	}

	for i := 0; i < opts.workers; i++ {
		go p.worker()
	}

	return p
}

func (p *pool) start(opts *ExecutionOpts) {
	rand := rand.New(rand.NewSource(opts.seed))
	for i := 0; i < opts.iterations; i++ {
		p.queue(job{
			config: opts.config,
			list:   opts.list,
			seed:   rand.Int63(),
		})
	}
}

func (p *pool) processWorkerResult(aggs simulation.Aggregators) error {
	select {
	case <-p.ctx.Done():
		return p.ctx.Err()
	case err := <-p.errChan:
		return err
	case result := <-p.respChan:
		aggs.Add(result)
		return nil
	}
}

func (p *pool) queue(j job) error {
	select {
	case <-p.ctx.Done():
		return p.ctx.Err()
	default:
		p.workChan <- j
	}
	return nil
}

func (p *pool) worker() {
	for {
		select {
		case <-p.ctx.Done():
			return

		case job := <-p.workChan:
			opts := &simulation.RunOpts{
				Config: job.config,
				Eval:   eval.New(p.ctx, job.list.Program),
				Seed:   job.seed,
			}

			res, err := simulation.Run(opts)
			if err != nil {
				p.errChan <- err
				return
			}

			p.respChan <- res
		}
	}
}

package servermode

import (
	"context"
	"log/slog"
	"math/rand"

	"github.com/simimpact/srsim/pkg/logic/gcs"
	"github.com/simimpact/srsim/pkg/logic/gcs/eval"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/pkg/simulation"
)

type workerpool struct {
	// initial stuff
	id      string
	yamlCfg string
	log     *slog.Logger
	ctx     context.Context
	cancel  chan bool

	// state
	currentCount int              // current count done for logging purposes
	done         bool             // if simulation is done
	result       *model.SimResult // latest result
	err          error            // any errors running sim
}

type job struct {
	config *model.SimConfig
	list   *gcs.ActionList
	seed   int64
}

func (w *workerpool) handleErr(err error) {
	w.err = err
	w.result = nil
}

func (w *workerpool) run(iterations, workerCount, flushInterval int) {
	w.log.Info("worker run started", "id", w.id)
	// handle panic
	defer func() {
		if r := recover(); r != nil {
			w.handleErr(errorRecover(r))
		}
	}()
	// done is true regardless of reason once this function exits
	defer func() {
		w.done = true
	}()

	simcfg, gcsl, err := parseYaml(w.yamlCfg)
	if err != nil {
		w.log.Info("config parsing failed", "id", w.id, "err", err)
		w.handleErr(err)
		return
	}
	w.log.Info("parse ok", "id", w.id)

	debugSeed, err := simulation.RandSeed()
	if err != nil {
		w.log.Warn("could not create debug seed", "id", w.id, "err", err)
		w.handleErr(err)
		return
	}
	w.result = simulation.CreateResult(simcfg, debugSeed)
	aggregators, err := simulation.InitializeAggregators(iterations, simcfg)
	if err != nil {
		w.log.Info("aggregator setup failed", "id", w.id, "err", err)
		w.handleErr(err)
		return
	}
	w.log.Info("aggregators ok", "id", w.id)

	// run jobs
	respCh := make(chan *model.IterationResult)
	errCh := make(chan error)
	workChan := make(chan job)
	for i := 0; i < workerCount; i++ {
		go w.iter(workChan, respCh, errCh)
	}
	w.log.Info("spawned worker", "id", w.id, "count", workChan)
	go func() {
		// make all the seeds
		wip := 0
		for wip < iterations {
			select {
			case <-w.cancel:
				w.log.Info("wip sending ended due to cancel", "id", w.id)
				return
			case workChan <- job{
				config: simcfg,
				list:   gcsl,
				seed:   rand.Int63(),
			}:
				wip++
			}
		}
		close(workChan)
	}()

	lastFlush := 0
iters:
	for w.currentCount < iterations {
		select {
		case result := <-respCh:
			// w.log.Info("got 1 result", "id", w.id, "count", count)
			for _, a := range aggregators {
				a.Add(result)
			}
			w.currentCount++
		case err := <-errCh:
			// error encountered
			w.log.Info("error running sim", "id", w.id, "err", err)
			w.handleErr(err)
			return
		case <-w.cancel:
			w.log.Info("cancel signal received", "id", w.id)
			// expectation is w.cancel is closed causing all go routines to wrap it up
			break iters
		}
		// flush and update results
		if w.currentCount-lastFlush > flushInterval {
			w.log.Debug("flush interval reached, flushing results", "id", w.id, "count", w.currentCount, "flush", lastFlush)
			lastFlush = w.currentCount
			w.result.Statistics = aggregators.Flush()
		}
	}
	w.log.Info("sim done", "id", w.id, "count", w.currentCount, "flush", lastFlush)
	w.result.Statistics = aggregators.Flush()
}

func (w *workerpool) iter(work chan job, respChan chan *model.IterationResult, errChan chan error) {
	for {
		select {
		case <-w.cancel:
			return
		case job, ok := <-work:
			if !ok {
				w.log.Info("work channel closed, iter worker ending", "id", w.id)
				return
			}

			opts := &simulation.RunOpts{
				Config: job.config,
				Eval:   eval.New(w.ctx, job.list.Program),
				Seed:   job.seed,
			}

			res, err := simulation.Run(opts)
			if err != nil {
				errChan <- err
				return
			}

			respChan <- res
		}
	}
}

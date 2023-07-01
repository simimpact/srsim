package workerpool

import (
	"context"

	"github.com/simimpact/srsim/pkg/gcs"
	"github.com/simimpact/srsim/pkg/gcs/eval"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/pkg/simulation"
)

type Pool struct {
	ctx      context.Context
	errChan  chan error
	respChan chan *model.IterationResult
	workChan chan Job
}

type Job struct {
	Script *gcs.ActionList
	Config *model.SimConfig
}

func New(ctx context.Context, workerCount int, respChan chan *model.IterationResult, errChan chan error) *Pool {
	p := &Pool{
		ctx:      ctx,
		errChan:  errChan,
		respChan: respChan,
		workChan: make(chan Job),
	}

	for i := 0; i < workerCount; i++ {
		go p.worker()
	}

	return p
}

// QueueJob attempts to add a new job; blocks until job is sucessfully added
// Return an error if the pool is stopped
func (p *Pool) QueueJob(j Job) error {
	select {
	case <-p.ctx.Done():
		return p.ctx.Err()
	default:
		p.workChan <- j
	}
	return nil
}

func (p *Pool) worker() {
	for {
		select {
		case <-p.ctx.Done():
			return
		case job := <-p.workChan:
			seed, err := simulation.RandSeed()
			if err != nil {
				p.errChan <- err
				return
			}

			res, err := simulation.Run(job.Config, eval.New(p.ctx, job.Script.Program), seed)
			if err != nil {
				p.errChan <- err
				return
			}
			p.respChan <- res
		}
	}
}

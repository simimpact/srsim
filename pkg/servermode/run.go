package servermode

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func errorRecover(r interface{}) error {
	var err error
	switch x := r.(type) {
	case string:
		err = errors.New(x)
	case error:
		err = x
	default:
		err = errors.New("unknown error")
	}
	return err
}

func (s *Server) run() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		s.Log.Info("request to run sim", "id", id)
		var payload struct {
			Config     string `json:"config"`
			Iterations int    `json:"iterations"`
		}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			s.Log.Info("body did not decode to json", "id", id, "err", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		// don't run if already running
		if s.isRunning(id) {
			s.Log.Info("run request failed; already running", "id", id)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("already running!!"))
			return
		}
		s.Log.Info("config decoded ok; running", "id", id)

		ctx, cancel := context.WithTimeout(context.Background(), s.Timeout)
		//nolint:exhaustruct // internal fields don't need to be initialized
		wp := &workerpool{
			id:      id,
			yamlCfg: payload.Config,
			log:     s.Log,
			cancel:  make(chan bool),
			ctx:     ctx,
		}
		s.Lock()
		s.pool[id] = wp
		s.Unlock()

		// sane defaults
		if payload.Iterations <= 0 {
			payload.Iterations = 100
		}

		// start run
		go wp.run(payload.Iterations, s.WorkerCount, s.FlushInterval)

		// add a timeout
		// TODO: i thought simulation run had a context check? i guess not
		go func() {
			defer cancel()
			for {
				select {
				case <-wp.cancel:
					// someone cancelled so we're done
					return
				case <-ctx.Done():
					// context must have timed out
					close(wp.cancel)
					wp.done = true
					if wp.err == nil {
						wp.err = fmt.Errorf("execution timed out after %s", s.Timeout)
					}
					return
				}
			}
		}()

		w.WriteHeader(http.StatusOK)
	}
}

package servermode

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Server) notImplemented() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotImplemented)
	}
}

func (s *Server) ready() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		running := s.isRunning(id)
		if running {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) running() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		running := s.isRunning(id)
		if running {
			w.Write([]byte("true"))
		} else {
			w.Write([]byte("false"))
		}
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) cancel() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		wk, ok := s.pool[id]
		if !ok {
			s.Log.Info("cancel request received; worker does not exist", "id", id)
			w.WriteHeader(http.StatusNotFound)
			return
		}
		if wk.done {
			s.Log.Info("cancel request received; already done", "id", id)
			w.WriteHeader(http.StatusOK)
			return
		}
		s.Log.Info("cancelling run", "id", id)
		close(wk.cancel)
		wk.done = true
		w.WriteHeader(http.StatusOK)
	}
}

func (s *Server) latest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		s.Log.Info("request for latest", "id", id)

		wk, ok := s.pool[id]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		var res struct {
			Result string `json:"result"`
			Hash   string `json:"hash"`
			Done   bool   `json:"done"`
			Error  string `json:"error"`
		}
		res.Done = wk.done

		// regardless of what results looks like, we should delete worker if done
		defer func() {
			if res.Done {
				s.Lock()
				delete(s.pool, id)
				s.Unlock()
			}
		}()

		s.Log.Info("found data", "id", id, "current_count", wk.currentCount, "res.Error", res.Error, "res.Done", res.Done)
		if wk.err == nil {
			// if no error then we expect there to be some kind of result
			if wk.result == nil {
				s.Log.Info("unexpected result is nil", "id", id)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("unexpected result is blank"))
				return
			}
			b, err := wk.result.MarshalJSON()
			if err != nil {
				s.Log.Info("error marshalling result to json", "id", id, "err", err)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
				return
			}
			res.Result = string(b)
		} else {
			res.Error = wk.err.Error()
		}

		msg, err := json.Marshal(res)
		if err != nil {
			s.Log.Info("error marshalling final response to json", "id", id, "err", err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(msg)
	}
}

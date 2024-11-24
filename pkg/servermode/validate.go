package servermode

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func (s *Server) validate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		s.Log.Info("request to run sample", "id", id)
		var payload struct {
			Config string `json:"config"`
		}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			s.Log.Info("body did not decode to json", "id", id, "err", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		simCfg, _, code, err := s.parseYAMLCfg(payload.Config, id)
		if err != nil {
			w.WriteHeader(code)
			w.Write([]byte(err.Error()))
			return
		}

		marshalled, err := json.Marshal(simCfg)
		if err != nil {
			w.Write([]byte(err.Error()))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(marshalled)
		w.WriteHeader(http.StatusOK)
	}
}

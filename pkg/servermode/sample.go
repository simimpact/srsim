package servermode

import (
	"bufio"
	"bytes"
	"compress/flate"
	"compress/gzip"
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/simimpact/srsim/pkg/engine/logging"
	"github.com/simimpact/srsim/pkg/logic/gcs"
	"github.com/simimpact/srsim/pkg/logic/gcs/eval"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/simimpact/srsim/pkg/simulation"
)

func (s *Server) sample() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				w.Write([]byte(errorRecover(r).Error()))
				w.WriteHeader(http.StatusBadRequest)
			}
		}()
		id := chi.URLParam(r, "id")
		s.Log.Info("request to run sample", "id", id)
		var payload struct {
			Config string `json:"config"`
			Seed   uint64 `json:"seed"`
		}
		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			s.Log.Info("body did not decode to json", "id", id, "err", err)
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		simcfg, gcsl, err := parseYaml(payload.Config)
		if err != nil {
			s.Log.Info("config parsing failed", "id", id, "err", err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		d, err := generateLogs(simcfg, gcsl, payload.Seed)
		if err != nil {
			s.Log.Info("generate sample failed", "id", id, "err", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		flush(w, d)
	}
}

func flush(resp http.ResponseWriter, compressed []byte) {
	resp.Header().Set("Content-Type", "application/json")
	resp.Header().Set("Content-Encoding", "gzip")
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.WriteHeader(http.StatusOK)
	resp.Write(compressed)

	if f, ok := resp.(http.Flusher); ok {
		f.Flush()
	}
}

func generateLogs(simcfg *model.SimConfig, list *gcs.ActionList, seed uint64) ([]byte, error) {
	buf := bytes.NewBuffer(nil)

	fileLogger, err := NewGzipLogger(buf)
	if err != nil {
		return nil, err
	}
	defer fileLogger.Flush()

	loggers := make([]logging.Logger, 0, 2)
	loggers = append(loggers, fileLogger)

	_, err = simulation.Run(&simulation.RunOpts{
		Config:  simcfg,
		Eval:    eval.New(context.TODO(), list.Program),
		Seed:    int64(seed),
		Loggers: loggers,
	})
	if err != nil {
		return nil, err
	}

	return io.ReadAll(buf)
}

type GzipLogger struct {
	f  io.Writer
	gf *gzip.Writer
	fw *bufio.Writer
}

func NewGzipLogger(fi io.Writer) (*GzipLogger, error) {
	gf, _ := gzip.NewWriterLevel(fi, flate.BestCompression)
	fw := bufio.NewWriterSize(gf, 32768)
	return &GzipLogger{fi, gf, fw}, nil
}

func (l *GzipLogger) Log(e any) {
	line := logging.Wrap(e)
	res, err := json.Marshal(line)
	if err != nil {
		l.Flush()
	}
	l.fw.Write(res)
	l.fw.WriteString("\n")
}

func (l *GzipLogger) Flush() {
	l.fw.Flush()
	l.gf.Close()
}

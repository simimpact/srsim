package main

import (
	"bufio"
	"compress/flate"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"os"

	"github.com/simimpact/srsim/pkg/engine/logging"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/urfave/cli/v2"
)

var _ logging.Logger = (*GzipLogger)(nil)
var _ logging.Logger = (*ConsoleLogger)(nil)

type ConsoleLogger struct{}

func (l ConsoleLogger) Log(e any) {
	line := logging.Wrap(e)
	res, err := json.Marshal(line)
	if err != nil {
		cli.Exit(err, 1)
	}
	fmt.Println(string(res))
}

type GzipLogger struct {
	f  *os.File
	gf *gzip.Writer
	fw *bufio.Writer
}

func NewGzipLogger(path string) (*GzipLogger, error) {
	fi, err := os.OpenFile(LogFile(path), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return nil, err
	}

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
	l.f.Close()
}

func WriteResult(result *model.SimResult, path string) error {
	fi, err := os.OpenFile(ResultFile(path), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o644)
	if err != nil {
		return err
	}
	defer fi.Close()

	data, err := result.MarshalJSON()
	if err != nil {
		return err
	}

	gf, _ := gzip.NewWriterLevel(fi, flate.BestCompression)
	defer gf.Close()
	gf.Write(data)

	return nil
}

package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/simimpact/srsim/pkg/logic/gcs"
	"github.com/simimpact/srsim/pkg/logic/gcs/parse"
	"github.com/simimpact/srsim/pkg/model"
	"github.com/urfave/cli/v2"
	"sigs.k8s.io/yaml"
)

func validatePath(path string) error {
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		abs, _ := filepath.Abs(path)
		fmt.Printf("File Not Found: no file found at %s\n", abs)
		return cli.Exit("", 1)
	}
	return nil
}

// TODO: support more file formats
func parseConfig(path string) (*model.SimConfig, error) {
	src, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	asJSON, err := yaml.YAMLToJSON(src)
	if err != nil {
		return nil, err
	}

	result := new(model.SimConfig)
	if err := result.UnmarshalJSON(asJSON); err != nil {
		return nil, err
	}

	return result, nil
}

// TODO: should return a logic.Eval supplier rather than the action list
// TODO: move to simulation so can be used by wasm
func parseLogic(config *model.SimConfig) (*gcs.ActionList, error) {
	switch logic := config.Logic.(type) {
	case *model.SimConfig_Gcsl:
		p := parse.New(logic.Gcsl)
		return p.Parse()
	default:
		return nil, fmt.Errorf("unknown logic type: %v", logic)
	}
}

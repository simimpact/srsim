package testyaml

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/simimpact/srsim/pkg/logic/gcs/eval"
	"github.com/simimpact/srsim/pkg/logic/gcs/parse"
	"github.com/simimpact/srsim/pkg/model"
	"sigs.k8s.io/yaml"
)

func ParseConfig(path string) (*model.SimConfig, *eval.Eval, error) {
	curPath, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}
	splitPath := strings.Split(curPath, string(os.PathSeparator))
	idx := 0
	for i, v := range splitPath {
		if v == "tests" {
			idx = i + 1
			break
		}
	}
	if idx == 0 {
		return nil, nil, fmt.Errorf("function is not called from within testframe. Call location: %s", curPath)
	}
	splitPath = splitPath[0:idx]
	splitPath = append(splitPath, "testcfg", "testyaml", path)
	filepath := strings.Join(splitPath, string(os.PathSeparator))

	// parse the SimConfig
	src, err := os.ReadFile(filepath)
	if err != nil {
		return nil, nil, err
	}

	asJSON, err := yaml.YAMLToJSON(src)
	if err != nil {
		return nil, nil, err
	}

	result := new(model.SimConfig)
	if err := result.UnmarshalJSON(asJSON); err != nil {
		return nil, nil, err
	}

	// parse eval function
	var e *eval.Eval
	switch logic := result.Logic.(type) {
	case *model.SimConfig_Gcsl:
		p := parse.New(logic.Gcsl)
		list, err := p.Parse()
		if err != nil {
			return nil, nil, fmt.Errorf("ActionList parse error: %w", err)
		}
		e = eval.New(context.TODO(), list.Program)
	default:
		return nil, nil, fmt.Errorf("unknown logic type: %v", logic)
	}

	return result, e, nil
}

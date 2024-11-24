package servermode

import (
	"fmt"
	"net/http"

	"github.com/simimpact/srsim/pkg/logic/gcs"
	"github.com/simimpact/srsim/pkg/logic/gcs/parse"
	"github.com/simimpact/srsim/pkg/model"
	"sigs.k8s.io/yaml"
)

func parseCfg(jsonCfg []byte) (*model.SimConfig, error) {
	simCfg := new(model.SimConfig)
	if err := simCfg.UnmarshalJSON(jsonCfg); err != nil {
		return nil, err
	}
	return simCfg, nil
}

func parseLogic(simCfg *model.SimConfig) (*gcs.ActionList, error) {
	switch logic := simCfg.Logic.(type) {
	case *model.SimConfig_Gcsl:
		p := parse.New(logic.Gcsl)
		return p.Parse()
	default:
		return nil, fmt.Errorf("unknown logic type: %v", logic)
	}
}

func parseYaml(cfg string) (*model.SimConfig, *gcs.ActionList, error) {
	asJSON, err := yaml.YAMLToJSON([]byte(cfg))
	if err != nil {
		return nil, nil, err
	}
	simCfg, err := parseCfg(asJSON)
	if err != nil {
		return nil, nil, err
	}
	list, err := parseLogic(simCfg)
	if err != nil {
		return nil, nil, err
	}
	return simCfg, list, nil
}

func (s *Server) parseYAMLCfg(cfg, id string) (*model.SimConfig, *gcs.ActionList, int, error) {
	asJSON, err := yaml.YAMLToJSON([]byte(cfg))
	if err != nil {
		s.Log.Warn("could not convert yaml to json", "id", id, "err", err)
		return nil, nil, http.StatusBadRequest, err
	}

	simCfg, err := parseCfg(asJSON)
	if err != nil {
		s.Log.Warn("could not load config json", "id", id, "err", err)
		return nil, nil, http.StatusInternalServerError, err
	}

	list, err := parseLogic(simCfg)
	if err != nil {
		s.Log.Info("could not parse logic", "id", id, "err", err)
		return nil, nil, http.StatusBadRequest, err
	}
	return simCfg, list, http.StatusOK, nil
}

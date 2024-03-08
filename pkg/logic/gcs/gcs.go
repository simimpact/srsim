package gcs

import (
	"encoding/json"
	"log"

	"github.com/simimpact/srsim/pkg/logic/gcs/ast"
	"github.com/simimpact/srsim/pkg/model"
)

// TODO: these structure here should eventually be replaced with protos

type ActionList struct {
	// TODO: this one is a bit trouble some to replace; i think ideally this should be an interface that
	// has an eval method
	Program   *ast.BlockStmt           `json:"-"`
	Settings  *model.SimulatorSettings `json:"settings"`
	Errors    []error                  `json:"-"` // These represents errors preventing ActionList from being executed
	ErrorMsgs []string                 `json:"error_msgs"`
}

func (a *ActionList) Copy() *ActionList {
	r := new(ActionList)
	*r = *a
	r.Program = a.Program.CopyBlock()
	return r
}

func (a *ActionList) PrettyPrint() string {
	prettyJSON, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(prettyJSON)
}

package gcs

import (
	"encoding/json"
	"log"

	"github.com/simimpact/srsim/pkg/gcs/ast"
	"github.com/simimpact/srsim/pkg/model"
)

//TODO: these structure here should eventually be replaced with protos

type ActionList struct {
	//TODO: this one is a bit trouble some to replace; i think ideally this should be an interface that
	//has an eval method
	Program   *ast.BlockStmt           `json:"-"`
	Settings  *model.SimulatorSettings `json:"settings"`
	Errors    []error                  `json:"-"` //These represents errors preventing ActionList from being executed
	ErrorMsgs []string                 `json:"errors"`
}

func (c *ActionList) Copy() *ActionList {
	r := new(ActionList)
	*r = *c
	r.Program = c.Program.CopyBlock()
	return r
}

func (a *ActionList) PrettyPrint() string {
	prettyJson, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	return string(prettyJson)
}

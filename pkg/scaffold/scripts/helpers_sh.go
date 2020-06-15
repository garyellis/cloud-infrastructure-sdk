package scripts

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const helpersShFile = "helpers.sh"

type HelpersSh struct {
	input.Input
}

func (h *HelpersSh) GetInput() (input.Input, error) {
	if h.Path == "" {
		h.Path = filepath.Join(ScriptsDir, helpersShFile)
	}
	h.TemplateBody = helpersShTmpl

	h.IsExec = true

	return h.Input, nil
}

const helpersShTmpl = `#!/bin/bash
#### this file is maintained by cloud-infrastructure-sdk
#### edits to this file will be overwritten the next cloud-infrastructure-sdk runs 

function foo(){
	echo foo
}

eval "$@"
`

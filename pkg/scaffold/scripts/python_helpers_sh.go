package scripts

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const pythonHelpersShFile = "python-helpers.sh"

type PythonHelpersSh struct {
	input.Input
}

func (h *PythonHelpersSh) GetInput() (input.Input, error) {
	if h.Path == "" {
		h.Path = filepath.Join(ScriptsDir, pythonHelpersShFile)
	}
	h.TemplateBody = pythonHelpersShTmpl

	h.IsExec = true

	return h.Input, nil
}

const pythonHelpersShTmpl = `#!/bin/bash
#### this file is maintained by {{.CliName}}-{{.CliVersion}}
#### edits to this file will be overwritten the next time {{.CliName}} runs on this project

function setup_virtualenv(){
  virtualenv .${1}
  source .${1}/bin/activate

  if [ ! -f requirements.txt ]; then
    echo "requirements.txt not found. skipping pip install."
  else
    pip install -r requirements.txt
  fi
}


eval $@
`

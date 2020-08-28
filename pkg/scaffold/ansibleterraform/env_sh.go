package ansibleterraform

import (
	"fmt"
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const envFile = "%s.sh"

type EnvSh struct {
	input.Input
	EnvName             string
	AppName             string
	DCName              string
	TfLiveBaseDir       string
	AnsibleInventoryDir string
}

func (t *EnvSh) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(EnvDir, t.DCName, fmt.Sprintf(envFile, t.EnvName))
	}

	t.TfLiveBaseDir = TfLiveBaseDir
	t.AnsibleInventoryDir = AnsibleInventoryDir

	t.TemplateBody = envTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const envTmpl = `
export IAAS_ENV=./{{.TfLiveBaseDir}}/{{.DCName}}/{{.EnvName}}
export APP_ENV=./{{.AnsibleInventoryDir}}/{{.DCName}}/{{.EnvName}}
`

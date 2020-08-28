package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const ansibleCfgFile = "ansible.cfg"

type AnsibleCfg struct {
	input.Input
}

func (t *AnsibleCfg) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(AnsibleBaseDir, ansibleCfgFile)
	}
	t.TemplateBody = ansibleCfgTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const ansibleCfgTmpl = `[defaults]
hash_behaviour = merge

[ssh_connection]
pipelining = True
`

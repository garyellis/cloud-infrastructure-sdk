package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const varsOverrideYmlFile = "overrides.yml"

type VarsOverrideYml struct {
	input.Input
}

func (t *VarsOverrideYml) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(AnsibleBaseDir, "vars", varsOverrideYmlFile)
	}
	t.TemplateBody = varsOverridesYmlTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const varsOverridesYmlTmpl = `---
### variables in this file are included via ansible playbook command line args
`

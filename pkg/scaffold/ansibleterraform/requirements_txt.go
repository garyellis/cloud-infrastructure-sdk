package ansibleterraform

import (
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const requirementsTxtFile = "requirements.txt"

type RequirementsTxt struct {
	input.Input
	AnsibleVersion string
}

func (t *RequirementsTxt) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = requirementsTxtFile
	}
	t.TemplateBody = requirementsTxtTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const requirementsTxtTmpl = `pip
selinux
pytest
awscli
molecule
docker
hvac
ansible
`

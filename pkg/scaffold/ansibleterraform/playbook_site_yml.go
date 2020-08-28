package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const siteYmlFile = "site.yml"

type SiteYml struct {
	input.Input
}

func (t *SiteYml) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(AnsiblePlaybooksDir, siteYmlFile)
	}
	t.TemplateBody = siteYmlTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const siteYmlTmpl = `---
- import_playbook: os.yml
- import_playbook: middleware.yml
`

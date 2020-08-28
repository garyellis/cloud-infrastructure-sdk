package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const requirementsYmlFile = "ansible.cfg"

type RequirementsYml struct {
	input.Input
}

func (t *RequirementsYml) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(AnsibleBaseDir, requirementsYmlFile)
	}
	t.TemplateBody = requirementsYmlTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const requirementsYmlTmpl = `---
- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/repository.git
  version: v0.1.0

- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/environment.git
  version: v0.1.0

- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/environment.git
  version: v0.1.0

- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/ntp.git
  version: v0.1.0

- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/logging-agent.git
  version: v0.1.0

- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/av.git
  version: v0.1.0

- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/ssh.git
  version: v0.1.0

- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/linux-common.git
  version: v0.1.0

- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/monitoring-agent1.git
  version: v0.1.0

- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/monitoring-agent2.git
  version: v0.1.0

- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/ldap-client.git
  version: v0.1.0

- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/lvm.git
  version: v0.1.0

- src: git+ssh://$GIT_HOSTNAME:7999/$GIT_PROJECT/vault-agent.git
  version: v0.1.0
`

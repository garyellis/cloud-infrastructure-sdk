package ansibleterraform

import (
	"fmt"
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const ansibleInventoryGroupVarsYml = "%s.yml"

type AnsibleInventoryGroupVarsYml struct {
	input.Input
	EnvName string
	AppName string
	DCName  string
}

func (t *AnsibleInventoryGroupVarsYml) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(AnsibleInventoryDir, t.DCName, t.EnvName, "group_vars", fmt.Sprintf(ansibleInventoryGroupVarsYml, t.AppName))
	}
	t.TemplateBody = ansibleInventoryGroupVarsYmlTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const ansibleInventoryGroupVarsYmlTmpl = `---
env_name: "{{.EnvName}}"
dc_name: "{{.DCName}}"
app_name: "{{.AppName}}"


## ssh config
ssh_ca_urls:
  - "https://vault-dev.ews.int/v1/ssh-client-signer/public_key"

# lvm config
lvm_disk: "/dev/nvme1n1"

# vault config
vault_pki_common_name: "{{ "{{ lb_dns }}" }}"
vault_pki_alt_names: "{{ "{{ lb_dns }}" }}"

## vault-agent config
## Set these to false if vault agent is already installed and configured
## it is optional
vault_agent_install: true
vault_agent_configure: true

## zabbix config
zabbix_agent_server: "{{ "{{ corp.zabbix_agent_server }}" }}"
zabbix_agent_serveractive: "{{ "{{ corp.zabbix_agent_serveractive }}" }}"

## appd_machine_agent
appd_controller_environment: corp
appd_machine_agent_application_name: "{{.AppName}}"
appd_machine_agent_machine_path: "EDAP|{{.AppName}}-{{.DCName}}-{{.EnvName}}"
`

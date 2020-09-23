package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const awsUserdataSh = "userdata.sh.tmpl"

type AWSUserdataSh struct {
	input.Input
	AppName string
}

func (t *AWSUserdataSh) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfModuleBaseDir, t.AppName, awsTfModuleDir, awsUserdataSh)
	}
	t.TemplateBody = awsUserdataShTmpl

	t.IfExistsAction = input.Skip
	return t.Input, nil
}

const awsUserdataShTmpl = `#!/bin/bash

exec > /var/log/bootstrap.log 2>&1

VAULT_ADDR=${vault_addr}
VAULT_SSH_CA_PATH=${vault_ssh_ca_path}

trusted_user_ca_keys=/etc/ssh/trusted-user-ca-keys.pem

curl -o $trusted_user_ca_keys $VAULT_ADDR/v1/$VAULT_SSH_CA_PATH/public_key

if grep -q TrustedUserCAKeys /etc/ssh/sshd_config ; then
  echo "TrustedUserCAKeys is already set"
else
  echo "TrustedUserCAKeys $trusted_user_ca_keys" >> /etc/ssh/sshd_config
  systemctl restart sshd
fi
`

package ansibleterraform

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const vmwareUserdataSh = "userdata.sh.tmpl"

type VMwareUserdataSh struct {
	input.Input
	AppName string
}

func (t *VMwareUserdataSh) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(TfModuleBaseDir, t.AppName, vmwareTfModuleDir, vmwareUserdataSh)
	}
	t.TemplateBody = vmwareUserdataShTmpl

	t.IfExistsAction = input.Skip
	return t.Input, nil
}

const vmwareUserdataShTmpl = `#!/bin/bash

exec > /var/log/bootstrap.log 2>&1

#### setup the input variables
export VAULT_ADDR=${vault_addr}
VAULT_SSH_CA_PATH=${vault_ssh_ca_path}
kv_secret_wrapping_token=${kv_secret_wrapping_token}
kv_secret_vas_username_key=${kv_secret_vas_username_key}
kv_secret_vas_password_key=${kv_secrret_vas_password_key}


#### install the vault cli
vault_cli_version=1.3.3
curl -RO https://artifactory.ews.int/artifactory/releases-hashicorp/vault/$vault_cli_version/vault_$${vault_cli_version}_linux_amd64.zip
yum -y install unzip
unzip vault_$${vault_cli_version}_linux_amd64.zip -d /usr/local/bin && chmod 755 /usr/local/bin/vault
rm -f vault_$${vault_cli_version}_linux_amd64.zip


#### read the vault secrets into shell variables
read -s vas_username vas_password <<<$(vault unwrap -format=json $kv_secret_wrapping_token | python -c "import sys, json; data=json.load(sys.stdin)['data']['data']; print data['$kv_secret_vas_username_key'], data['$kv_secret_vas_password_key']")


#### install the vas agent
yum -y install vasclnt vasgp vasutil
        /opt/quest/bin/vastool -u $vas_username -w $vas_password join -f -c "ou=Unix,ou=servers,dc=EWS,dc=INT" ews.int
        /opt/quest/bin/vastool configure vas vasd timesync-interval 0

/opt/quest/bin/vastool -u $vas_username -w $vas_password create -c "ou=Access Groups,ou=Security Groups,ou=Groups,dc=EWS,dc=INT" -t "universal" group ux-hac-$(hostname)-deny
/opt/quest/bin/vastool -u $vas_username -w $vas_password create -c "ou=Access Groups,ou=Security Groups,ou=Groups,dc=EWS,dc=INT" -t "universal" group ux-hac-$(hostname)-allow


#### install the ssh ca certificate
trusted_user_ca_keys=/etc/ssh/trusted-user-ca-keys.pem

curl -o $trusted_user_ca_keys $VAULT_ADDR/v1/$VAULT_SSH_CA_PATH/public_key

if grep -q TrustedUserCAKeys /etc/ssh/sshd_config ; then
  echo "TrustedUserCAKeys is already set"
else
  echo "TrustedUserCAKeys $trusted_user_ca_keys" >> /etc/ssh/sshd_config
  systemctl restart sshd
fi
`

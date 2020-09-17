package ansibleterraform

import (
	"fmt"
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

// env_vmware_sh.go and env_aws_sh.go need to merge into a single file
// aws params, and vault params need to be set by the user after an environment is layed down until
// these opts are parameterized

const envVmwareFile = "%s.sh"

type EnvVmwareSh struct {
	input.Input
	EnvName             string
	AppName             string
	DCName              string
	AWSRegion           string
	VaultAddr           string
	TfLiveBaseDir       string
	AnsibleInventoryDir string
}

func (t *EnvVmwareSh) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(EnvDir, t.DCName, fmt.Sprintf(envVmwareFile, t.EnvName))
	}

	t.TfLiveBaseDir = TfLiveBaseDir
	t.AnsibleInventoryDir = AnsibleInventoryDir

	t.TemplateBody = envVmwareTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const envVmwareTmpl = `
## setup the vault environment
export AWS_CA_BUNDLE=/etc/ssl/certs/ca-bundle.crt
export VAULT_ADDR={{.VaultAddr}}
export VAULT_SSH_CERT_PRINCIPAL=${VAULT_SSH_CERT_PRINCIPAL} # set by user
export VAULT_SSH_CLIENT_SIGNER_PATH=ssh-client-signer/sign/vault-dev_ansible


## set environment for the automation user
export AUTOMATION_USER_VAULT_PATHS=(
  provision/data/trusted-orchestrator/corp
  provision/data/trusted-orchestrator/aws-iam-svc-terraform-dev
)
AUTOMATION_USER=${AUTOMATION_USER:-false}
if [ "$AUTOMATION_USER" == "true" ]; then
  export VAULT_SSH_CERT_PRINCIPAL=provisioner
  for i in ${AUTOMATION_USER_VAULT_PATHS[@]}; do
    $(vault read $i -format=json | jq '.data.data | to_entries| map("export " + .key + "=" + .value)|.[]' -r) || return 1
  done
  # setup the automation user git ssh key
  [ -f /.dockerenv ] && [ ! -z "${id_rsa}" ] && base64 -d <<<"${id_rsa}" > $HOME/.ssh/id_rsa && chmod 600 $HOME/.ssh/id_rsa && ssh-keygen -y -f $HOME/.ssh/id_rsa > $HOME/.ssh/id_rsa.pub
fi


## setup the iaas provider
## iaas provider depends on s3 remote state
export AWS_REGION={{.AWSRegion}}
export ASSUME_ROLE_ARN=$ASSUME_ROLE_ARN # set by the user

export VSPHERE_SERVER=<set-by-user>

if [ "$VSPHERE_USER" == "" ]; then
  echo VSPHERE_USER is not set
  echo -n VSPHERE_USER:
  read VSPHERE_USER
  export VSPHERE_USER
fi

if [ "$VSPHERE_PASSWORD" == "" ]; then
  echo VSPHERE_PASSWORD is not set
  echo -n VSPHERE_PASSWORD:
  read -s VSPHERE_PASSWORD
  export VSPHERE_PASSWORD
fi


## setup the bootstrap ssh user
if [ "$TF_VAR_provisioner_ssh_user" == "" ]; then
  echo TF_VAR_provisioner_ssh_user is not set
  echo -n TF_VAR_provisioner_ssh_user:
  read TF_VAR_provisioner_ssh_user
  export TF_VAR_provisioner_ssh_user
fi

if [ "$TF_VAR_provisioner_ssh_password" == "" ]; then
  echo TF_VAR_provisioner_ssh_password is not set
  echo -n TF_VAR_provisioner_ssh_password:
  read -s TF_VAR_provisioner_ssh_password
  export TF_VAR_provisioner_ssh_password
fi

export IAAS_ENV=./{{.TfLiveBaseDir}}/{{.DCName}}/{{.EnvName}}
export APP_ENV=./{{.AnsibleInventoryDir}}/{{.DCName}}/{{.EnvName}}

## setup a provisioner user hackaround related to ssh agent keys and password
export ANSIBLE_SSH_ARGS="-o PreferredAuthentications=password,publickey"
`

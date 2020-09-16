package ansibleterraform

import (
	"fmt"
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

// env_vmware_sh.go and env_aws_sh.go need to merge into a single file
// aws params, and vault params need to be set by the user after an environment is layed down until
// these opts are parameterized

const envAwsFile = "%s.sh"

type EnvAwsSh struct {
	input.Input
	EnvName             string
	AppName             string
	DCName              string
	AWSRegion           string
	VaultAddr           string
	TfLiveBaseDir       string
	AnsibleInventoryDir string
}

func (t *EnvAwsSh) GetInput() (input.Input, error) {
	if t.Path == "" {
		t.Path = filepath.Join(EnvDir, t.DCName, fmt.Sprintf(envAwsFile, t.EnvName))
	}

	t.TfLiveBaseDir = TfLiveBaseDir
	t.AnsibleInventoryDir = AnsibleInventoryDir

	t.TemplateBody = envAwsTmpl

	t.IfExistsAction = input.Skip

	return t.Input, nil
}

const envAwsTmpl = `
export AWS_CA_BUNDLE=/etc/ssl/certs/ca-bundle.crt
export AWS_REGION={{.AWSRegion}}
export ASSUME_ROLE_ARN=${ASSUME_ROLE_ARN} # set by user
export VAULT_ADDR={{.VaultAddr}}
export VAULT_SSH_CERT_PRINCIPAL=${VAULT_SSH_CERT_PRINCIPAL} # set by user
export VAULT_SSH_CLIENT_SIGNER_PATH=ssh-ca-engine-name/sign/ssh-role-name # set by user


## set environment for the automation user
export AUTOMATION_USER_VAULT_PATHS=(
  provision/data/trusted-orchestrator/corp
  provision/data/trusted-orchestrator/aws-iam-svc-terraform-dev
)
AUTOMATION_USER=${AUTOMATION_USER:-false}
if [ "$AUTOMATION_USER" == "true" ]; then
  for i in ${AUTOMATION_USER_VAULT_PATHS[@]}; do
    $(vault read $i -format=json | jq '.data.data | to_entries| map("export " + .key + "=" + .value)|.[]' -r) || return 1
  done
  # setup the automation user git ssh key
  [ -f /.dockerenv ] && [ ! -z "${id_rsa}" ] && base64 -d <<<"${id_rsa}" > $HOME/.ssh/id_rsa && chmod 600 $HOME/.ssh/id_rsa && ssh-keygen -y -f $HOME/.ssh/id_rsa > $HOME/.ssh/id_rsa.pub
fi

export IAAS_ENV=./{{.TfLiveBaseDir}}/{{.DCName}}/{{.EnvName}}
export APP_ENV=./{{.AnsibleInventoryDir}}/{{.DCName}}/{{.EnvName}}
`

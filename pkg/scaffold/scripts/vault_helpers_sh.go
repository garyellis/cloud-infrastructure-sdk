package scripts

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const vaultHelpersShFile = "vault-helpers.sh"

type VaultHelpersSh struct {
	input.Input
}

func (h *VaultHelpersSh) GetInput() (input.Input, error) {
	if h.Path == "" {
		h.Path = filepath.Join(ScriptsDir, vaultHelpersShFile)
	}
	h.TemplateBody = VaultHelpersShTmpl

	h.IsExec = true

	return h.Input, nil
}

const VaultHelpersShTmpl = `#!/bin/bash
#### this file is maintained by {{.CliName}}-{{.CliVersion}}
#### edits to this file will be overwritten the next time {{.CliName}} runs on this project

# vault-helpers defaults
VAULT_SSH_CERT_PRINCIPAL=${VAULT_SSH_CERT_PRINCIPAL:-provisioner}

# source the project config when it exists
[ -e "./scripts/config.sh" ] && source ./scripts/config.sh
# source the user project config when it exists
[ -e "$HOME/config.sh" ] && source $HOME/config.sh


function vault_get_ssh_cert(){
  if [ -e "$ENV" ]; then
    echo "==> sourcing $ENV"
    source $ENV
  fi
  echo "==> writing signed certificate. signer: $VAULT_SSH_CLIENT_SIGNER_PATH principal: $VAULT_SSH_CERT_PRINCIPAL dest: $HOME/.ssh/id_rsa-cert.pub"
  vault write -field=signed_key $VAULT_SSH_CLIENT_SIGNER_PATH public_key=@$HOME/.ssh/id_rsa.pub valid_principals=$VAULT_SSH_CERT_PRINCIPAL > $HOME/.ssh/id_rsa-cert.pub
  echo "==> ssh cert details"
  ssh-keygen -Lf ~/.ssh/id_rsa-cert.pub
}


eval $@
`

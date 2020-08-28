package scripts

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const awsHelpersShFile = "aws-helpers.sh"

type AwsHelpersSh struct {
	input.Input
}

func (h *AwsHelpersSh) GetInput() (input.Input, error) {
	if h.Path == "" {
		h.Path = filepath.Join(ScriptsDir, awsHelpersShFile)
	}
	h.TemplateBody = awsHelpersShTmpl

	h.IsExec = true

	return h.Input, nil
}

const awsHelpersShTmpl = `#!/bin/bash
#### this file is maintained by {{.CliName}}-{{.CliVersion}}
#### edits to this file will be overwritten the next time {{.CliName}} runs on this project

# aws-helpers config
export ASSUME_ROLE_ARN=${ASSUME_ROLE_ARN}

# source the project config when it exists
[ -e "./scripts/config.sh" ] && source ./scripts/config.sh
# source the user project config when it exists
[ -e "$HOME/config.sh" ] && source $HOME/config.sh

if [ "$AWS_SERIAL_NUMBER" == "" ]; then
  echo AWS_SERIAL_NUMBER is not set
  echo -n AWS_SERIAL_NUMBER: 
  read AWS_SERIAL_NUMBER
  export AWS_SERIAL_NUMBER
fi

if [ "$AWS_DEFAULT_PROFILE" == "" ]; then
  echo AWS_DEFAULT_PROFILE is not set
  echo -n AWS_DEFAULT_PROFILE: 
  read AWS_DEFAULT_PROFILE
  export AWS_DEFAULT_PROFILE
fi

# ensure the AWS_PROFILE environment variable is not also set
unset AWS_PROFILE


function sts_get_session_token(){
  unset AWS_SESSION_TOKEN AWS_ACCESS_KEY_ID AWS_SECRET_ACCESS_KEY
  aws sts get-session-token \
      --serial-number $1 \
      --token-code ${2} |
    jq -r '.Credentials| "AWS_SESSION_TOKEN="+.SessionToken+" AWS_ACCESS_KEY_ID="+.AccessKeyId+" AWS_SECRET_ACCESS_KEY="+.SecretAccessKey'
  
}

function sts_assume_role(){
  aws sts assume-role \
      --role-arn ${1} \
      --role-session-name "${USER}" |
    jq -r '.Credentials| "AWS_SESSION_TOKEN="+.SessionToken+" AWS_ACCESS_KEY_ID="+.AccessKeyId+" AWS_SECRET_ACCESS_KEY="+.SecretAccessKey'
}

function validate_credentials(){
  echo account_alias: $(aws iam list-account-aliases|jq -r '.AccountAliases[0]')
  aws sts get-caller-identity|jq -r '"account: "+ .Account +"\nsession-arn: " + .Arn'
}

function get_temporary_credentials(){
  if [ -e "$1" ]; then
    echo "==> sourcing $1"
    source $1
  fi

  echo mfa serial: $AWS_SERIAL_NUMBER
  echo role_arn: $ASSUME_ROLE_ARN
  echo -n mfa token:
  read -s TOKEN_CODE
  echo
  echo "==> getting session token with mfa credentials"
  export $(sts_get_session_token $AWS_SERIAL_NUMBER $TOKEN_CODE)
  echo "done"
  echo "==> assuming role $ASSUME_ROLE_ARN"
  export $(sts_assume_role $ASSUME_ROLE_ARN)
  echo "done"

  echo "==> validating credentials"
  validate_credentials
  echo "done"
}

eval $@
`

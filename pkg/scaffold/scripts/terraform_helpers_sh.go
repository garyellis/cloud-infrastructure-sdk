package scripts

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const terraformHelpersShFile = "terraform-helpers.sh"

type TerraformHelpersSh struct {
	input.Input
	TerraformVersion      string
	TerraformReleasesURL  string
	TerragruntVersion     string
	TerragruntDownloadURL string
	TfenvRepoURL          string
}

func (h *TerraformHelpersSh) GetInput() (input.Input, error) {
	if h.Path == "" {
		h.Path = filepath.Join(ScriptsDir, terraformHelpersShFile)
	}
	h.TemplateBody = TerraformHelpersShTmpl

	h.IsExec = true

	return h.Input, nil
}

const TerraformHelpersShTmpl = `#!/bin/bash
#### this file is maintained by {{.CliName}}-{{.CliVersion}}
#### edits to this file will be overwritten the next time {{.CliName}} runs on this project

# terraform-helpers config
TERRAFORM_VERSION={{.TerraformVersion}}
TERRAGRUNT_VERSION={{.TerragruntVersion}}
TERRAGRUNT_DOWNLOAD_URL={{.TerragruntDownloadURL}}
TFENV_REPO_URL={{.TfenvRepoURL}}
TERRAFORM_RELEASES_URL={{.TerraformReleasesURL}}


# source the project config when it exists
[ -e "./scripts/config.sh" ] && source ./scripts/config.sh
# source the user project config when it exists
[ -e "~/config.sh" ] && source ~/config.sh


function tfenv(){
  if [ ! $(find ~/.tfenv -name tfenv) ]; then
    git clone $TFENV_REPO_URL ~/.tfenv
  fi
  if ! grep -q 'PATH=.*tfenv/bin:' ~/.bash_profile ; then
    echo 'export PATH=$HOME/.tfenv/bin:$PATH' >> ~/.bash_profile
  fi
  if ! grep -q TERRAFORM_RELEASES_URL ~/.bash_profile ; then
    echo "export TERRAFORM_RELEASES_URL=$TERRAFORM_RELEASES_URL" >> ~/.bash_profile
  fi

  ~/.tfenv/bin/tfenv install $TERRAFORM_VERSION
  ~/.tfenv/bin/tfenv use $TERRAFORM_VERSION
}

function terragrunt(){
  mkdir -p ./bin
  curl -o ./bin/terragrunt -L $TERRAGRUNT_DOWNLOAD_URL
  chmod 755 ./bin/terragrunt
}

function terraform_plugins(){
  if [ ! -f $2 ]; then
    echo "terraform-plugins does not exist. skipping"
  else
    mkdir -p $1
    while read plugin; do
      filename=$(basename $plugin)
      curl -L -RO $plugin
      unzip -o $filename -d $1
      rm -f $filename
    done < $2
   fi
}

function terraform_clean(){
  ( cd $1 && rm -f *.tfstate* ; rm -fr .terraform )
}


eval $@
`

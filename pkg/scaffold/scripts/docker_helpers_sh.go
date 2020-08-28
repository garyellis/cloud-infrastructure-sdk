package scripts

import (
	"path/filepath"

	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const dockerHelpersShFile = "docker-helpers.sh"

type DockerHelpersSh struct {
	input.Input
}

func (h *DockerHelpersSh) GetInput() (input.Input, error) {
	if h.Path == "" {
		h.Path = filepath.Join(ScriptsDir, dockerHelpersShFile)
	}
	h.TemplateBody = dockerHelpersShTmpl

	h.IsExec = true

	return h.Input, nil
}

const dockerHelpersShTmpl = `#!/bin/bash
#### this file is maintained by {{.CliName}}-{{.CliVersion}}
#### edits to this file will be overwritten the next time {{.CliName}} runs on this project

# docker-helpers config
INSTALL_DOCKER_CE_CLI=${INSTALL_DOCKER_CE_CLI:-false}
DOCKER_CE_VERSION=19.03.5-3
DOCKER_CE_CLI_URL=https://artifactory.ews.int/artifactory/docker_centos_2-cache/7/x86_64/stable/Packages/docker-ce-cli-${DOCKER_CE_VERSION}.el7.x86_64.rpm
CI_PIPELINES_DOCKER_IMAGE=dev.registry.ews.int/pegasus/ci-pipelines-infrastructure

# source the project config when it exists
[ -e "./scripts/config.sh" ] && source ./scripts/config.sh
# source the user project config when it exists
[ -e "~/config.sh" ] && source ~/config.sh


function install_docker_ce_cli(){
  [ "$INSTALL_DOCKER_CE_CLI" == "true" ] && yum -y install $DOCKER_CE_CLI_URL || return 0
}

function docker_runner(){
  # workaround to https://jira.atlassian.com/browse/BAM-20553
  # resolved in bamboo 6.10

  set -x
  action=$2
  name=$1
  shift;shift
  args=$@

  # cleanup files created by the container
  if [ "$action" == "cleanup" ]; then
    if ! (docker ps |grep "$name"); then
      echo "container $name is not running"
      return 0
    fi

    docker exec -w /deploy  $name rm -fr version.rendered
    docker exec -w /deploy $name rm -fr .deploy-*
    docker exec -w /deploy $name bash -c "find . -name .terragrunt-cache -type d| xargs rm -fr"
    docker exec -w /deploy $name bash -c "find . -name inventory.yml -delete"

    docker kill $name
    return $?
  fi

  if [ "$action" == "exec" ]; then
    if ! (docker ps |grep $name); then
      docker run \
        --name $name \
        -d \
        --rm \
        -w /deploy \
        -v $HOME/.vault-token:/root/.vault-token \
        -v $PWD:/deploy \
        -e ENV=$ENV \
        -e AUTOMATION_USER=$AUTOMATION_USER \
        $CI_PIPELINES_DOCKER_IMAGE tail -f /dev/null
    fi

    docker exec \
      -w /deploy \
      -e ENV=$ENV \
      -e AUTOMATION_USER=$AUTOMATION_USER \
      $name ${args}
      return $?
  fi
  set +x
}


eval $@
`

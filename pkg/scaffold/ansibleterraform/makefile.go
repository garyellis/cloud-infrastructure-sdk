package ansibleterraform

import (
	"github.com/garyellis/cloud-infrastructure-sdk/pkg/scaffold/input"
)

const makefile = "Makefile"

type Makefile struct {
	input.Input

	AppName string
}

func (r *Makefile) GetInput() (input.Input, error) {
	if r.Path == "" {
		r.Path = makefile
	}
	r.TemplateBody = makefileTmpl

	return r.Input, nil
}

const makefileTmpl = `
#### this file is maintained by {{.CliName}}-{{.CliVersion}}
#### edits to this file will be overwritten the next time {{.CliName}} runs on this project
.PHONY: help version metadata terragrunt-plan package
	.DEFAULT_GOAL := help

export PATH := $(PWD)/bin:$(PATH)

ENABLE_VIRTUALENV := [ -z "$$VIRTUAL_ENV" ] && source .$$(basename $$PWD)/bin/activate
ANSIBLE_ENV :=  source $(ENV) && export ANSIBLE_HOST_KEY_CHECKING=false ANSIBLE_ROLES_PATH=$$PWD/app/ansible/roles ANSIBLE_INVENTORY=$$PWD/$$APP_ENV
TERRAFORM_ENV := source $(ENV)
USER_ENV := [ -e "$$HOME/config.sh" ] && source $$HOME/config.sh

help: ## show this message
	@grep -E '^[a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

clean: ## cleans files created by this project
	rm -fr .deploy-*
	rm -fr version.rendered
	find . -name .terragrunt-cache -type d| xargs rm -fr
	find . -name inventory.yml -delete

dependencies: ## install dependencies and setup local environment
	./scripts/terraform-helpers.sh tfenv
	./scripts/terraform-helpers.sh terraform_plugins ~/.terraform.d/plugins/linux_amd64 terraform-plugins
	./scripts/terraform-helpers.sh terragrunt
	./scripts/python-helpers.sh setup_virtualenv $$(basename $$PWD)

version: ## renders version file
	source ./version && for i in $$(awk -F'=' '/^[A-Z]/ {print $$1}' ./version); do echo $$i=$${!i} || exit 1; done > version.rendered

metadata: version ## creates metadata files
	@./iaas/terraform/live/metadata.yaml.sh

molecule-converge: ## runs molecule converge
	cd ./app/ansible && molecule converge

molecule-destroy: ## runs molecule destroy
	cd ./app/ansible && molecule destroy

molecule-list: ## runs molecule list
	cd ./app/ansible && molecule list

molecule-login: ## runs molecule login on the target host
	cd ./app/ansible && molecule login --host $(HOST)

terragrunt-init: metadata ## runs terragrunt init on the target directory
	source $(ENV) && cd $$IAAS_ENV/{{.AppName}} && TF_INPUT=0 terragrunt init

terragrunt-plan: terragrunt-init ## runs terragrunt plan
	mkdir -p ./reports
	source $(ENV) && (cd $$IAAS_ENV/{{.AppName}}; terragrunt plan) | tee ./reports/$$(basename $$ENV | sed 's/.sh//')-terragrunt-plan.log

terragrunt-apply: terragrunt-init ## runs terragrunt apply
	source $(ENV) && cd $$IAAS_ENV/{{.AppName}} && terragrunt apply -auto-approve

terragrunt-destroy: terragrunt-init ## runs terragrunt destroy
	source $(ENV) && cd $$IAAS_ENV/{{.AppName}} && terragrunt destroy -auto-approve

generate-ansible-inventory: terragrunt-init ## generates ansible inventory from terraform outputs
	$(TERRAFORM_ENV) && echo "$$(cd $$IAAS_ENV/{{.AppName}} && terragrunt output ansible_inventory)" > $$APP_ENV/inventory.yml
	$(ENABLE_VIRTUALENV) ; $(ANSIBLE_ENV) && ansible-inventory --graph

ansible-inventory: ## runs ansible-inventory to list inventory hosts
	$(ENABLE_VIRTUALENV) ; $(ANSIBLE_ENV) && env|grep APP_ENV && ansible-inventory --graph

get-ssh-cert: ## gets an ssh cert from vault
	source $(ENV) && ./scripts/vault-helpers.sh vault_get_ssh_cert

ansible-ping: ## validates ssh connectivity on the target ansible inventory
	$(ENABLE_VIRTUALENV); $(ANSIBLE_ENV); $(USER_ENV); ansible -m ping all -u $$VAULT_SSH_CERT_PRINCIPAL

ansible-shell: ## runs ansible shell module CMD on the target inventory
	$(ENABLE_VIRTUALENV); $(ANSIBLE_ENV); $(USER_ENV); ansible -u $$VAULT_SSH_CERT_PRINCIPAL -m shell -a "$$CMD" all

ansible-galaxy: ## runs ansible-galaxy to install ansible role dependencies
	$(ENABLE_VIRTUALENV) ; ansible-galaxy install -r ./app/ansible/requirements.yml -p ./app/ansible/roles -f

ansible-playbook: ## runs ansible-playbook on the APP_ENV inventory
	$(ENABLE_VIRTUALENV) ; $(ANSIBLE_ENV); $(USER_ENV); cd ./app/ansible && ansible-playbook -u $$VAULT_SSH_CERT_PRINCIPAL -b ./playbooks/site.yml -e "@./vars/global.yml" -e "ansible_password=$$TF_VAR_provisioner_ssh_password"

package: ## packages the project into a tarball for distribution
	./scripts/helpers.sh package
`

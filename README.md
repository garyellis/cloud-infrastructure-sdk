# cloud infrastructure sdk [![Build Status](https://travis-ci.org/garyellis/cloud-infrastructure-sdk.svg?branch=master)](https://travis-ci.org/garyellis/cloud-infrastructure-sdk)
Scaffolding framework cli for cloud infrastructure related projects.

* ansible role (not started)
* ansible-terraform project (in progress)
* terragrunt/terraform live project (in progress)
* terraform module (not started)
* packer build (not started)
* docker image (not started)


## ansible-terraform usage

```bash
create an ansible-terraform project

Usage:
  cloud-infra-sdk ansible-terraform [command]

Available Commands:
  init        creates a new ansible-terraform project

Flags:
      --app-name string         the terraform live live subfolder name (default "my-app")
      --dc-name string          the data center name (default "my-dc")
  -e, --env-name strings        one or more environment names (default [development])
  -h, --help                    help for ansible-terraform
      --infra-provider string   infrastructure provider. Valid providers are aws and vmware (default "aws")
      --project-name string     the teraform live project name (default "my-project")
```

Create an ansible/terraform project my-great-project for application/stack name foo in dc1 with eight aws environments (dev1, dev2, qa1, qa2, uat1, uat2, prod1 and prod2).

```bash
$cloud-infrastructure-sdk ansible-terraform init --app-name foo --project-name my-great-project --dc-name dc1 -e dev1 -e dev2 -e qa1 -e qa2 -e uat1 -e uat2 -e prod1 -e prod2

$ tree my-great-project
my-great-project
├── Makefile
├── README.md
├── app
│   └── ansible
│       ├── ansible.cfg
│       ├── inventory
│       │   └── dc1
│       │       ├── dev1
│       │       │   └── foo.yml
│       │       ├── dev2
│       │       │   └── foo.yml
│       │       ├── prod1
│       │       │   └── foo.yml
│       │       ├── prod2
│       │       │   └── foo.yml
│       │       ├── qa1
│       │       │   └── foo.yml
│       │       ├── qa2
│       │       │   └── foo.yml
│       │       ├── qa3
│       │       │   └── foo.yml
│       │       ├── uat1
│       │       │   └── foo.yml
│       │       └── uat2
│       │           └── foo.yml
│       └── playbooks
│           ├── middleware.yml
│           ├── os.yml
│           └── site.yml
├── env
│   └── dc1
│       ├── dev1.sh
│       ├── dev2.sh
│       ├── prod1.sh
│       ├── prod2.sh
│       ├── qa1.sh
│       ├── qa2.sh
│       ├── qa3.sh
│       ├── uat1.sh
│       └── uat2.sh
├── iaas
│   └── terraform
│       ├── live
│       │   └── dc1
│       │       ├── dev1
│       │       │   ├── foo
│       │       │   │   └── terragrunt.hcl
│       │       │   ├── terragrunt.hcl
│       │       │   └── vars.yml
│       │       ├── dev2
│       │       │   ├── foo
│       │       │   │   └── terragrunt.hcl
│       │       │   ├── terragrunt.hcl
│       │       │   └── vars.yml
│       │       ├── prod1
│       │       │   ├── foo
│       │       │   │   └── terragrunt.hcl
│       │       │   ├── terragrunt.hcl
│       │       │   └── vars.yml
│       │       ├── prod2
│       │       │   ├── foo
│       │       │   │   └── terragrunt.hcl
│       │       │   ├── terragrunt.hcl
│       │       │   └── vars.yml
│       │       ├── qa1
│       │       │   ├── foo
│       │       │   │   └── terragrunt.hcl
│       │       │   ├── terragrunt.hcl
│       │       │   └── vars.yml
│       │       ├── qa2
│       │       │   ├── foo
│       │       │   │   └── terragrunt.hcl
│       │       │   ├── terragrunt.hcl
│       │       │   └── vars.yml
│       │       ├── qa3
│       │       │   ├── foo
│       │       │   │   └── terragrunt.hcl
│       │       │   ├── terragrunt.hcl
│       │       │   └── vars.yml
│       │       ├── uat1
│       │       │   ├── foo
│       │       │   │   └── terragrunt.hcl
│       │       │   ├── terragrunt.hcl
│       │       │   └── vars.yml
│       │       └── uat2
│       │           ├── foo
│       │           │   └── terragrunt.hcl
│       │           ├── terragrunt.hcl
│       │           └── vars.yml
│       └── modules
│           └── foo
│               ├── aws
│               │   ├── ansible_inventory.yml.tmpl
│               │   ├── locals.tf
│               │   ├── main.tf
│               │   ├── outputs.tf
│               │   ├── userdata.sh.tmpl
│               │   └── variables.tf
│               └── vmware
│                   ├── ansible_inventory.yml.tmpl
│                   ├── main.tf
│                   ├── outputs.tf
│                   ├── userdata.sh.tmpl
│                   └── variables.tf
├── scripts
│   ├── aws-helpers.sh
│   ├── docker-helpers.sh
│   ├── helpers.sh
│   ├── python-helpers.sh
│   └── terraform-helpers.sh
└── version

43 directories, 68 files
```

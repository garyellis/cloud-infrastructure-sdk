# cloud infrastructure sdk
Scaffolding framework cli for cloud infrastructure related projects.

* ansible role (not started)
* ansible/terraform project (not started)
* terragrunt/terraform live project (in progress)
* terraform module (not started)
* packer build (not started)
* docker image (not started)


## Usage
Create a teraform live project (apigateway) with an app root module (mulefoo) with five environments (dev, qa, uat, cat, and prod.)

```bash
$cloud-infrastructure-sdk terraform-live init --project-name apigateway --app-name mulefoo -e dev -e qa -e uat -e cat -e prod

$ tree apigateway
apigateway
├── README.md
├── scripts
│   └── helpers.sh
├── terraform
│   ├── live
│   │   ├── cat
│   │   │   ├── mulefoo
│   │   │   │   └── terragrunt.hcl
│   │   │   ├── terragrunt.hcl
│   │   │   └── vars.yaml
│   │   ├── dev
│   │   │   ├── mulefoo
│   │   │   │   └── terragrunt.hcl
│   │   │   ├── terragrunt.hcl
│   │   │   └── vars.yaml
│   │   ├── metadata.yaml.sh
│   │   ├── prod
│   │   │   ├── mulefoo
│   │   │   │   └── terragrunt.hcl
│   │   │   ├── terragrunt.hcl
│   │   │   └── vars.yaml
│   │   ├── qa
│   │   │   ├── mulefoo
│   │   │   │   └── terragrunt.hcl
│   │   │   ├── terragrunt.hcl
│   │   │   └── vars.yaml
│   │   └── uat
│   │       ├── mulefoo
│   │       │   └── terragrunt.hcl
│   │       ├── terragrunt.hcl
│   │       └── vars.yaml
│   └── modules
│       └── mulefoo
│           └── aws
│               ├── locals.tf
│               ├── main.tf
│               ├── outputs.tf
│               └── variables.tf
└── version

```

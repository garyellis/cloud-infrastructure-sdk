# ansible-terraform
The ansible-terraform command is used to create or update a project that leverages terraform for IAAS/cloud configuration management and ansible for OS and app configuration management. It maintains instances of or "environments" for an application. 


# Terraform



# Ansibe



# Usage

### Customize your project config.yaml
Customize your ansible roles to suite your OS and application deployment. Ansible roles are grouped by OS and application. OS roles in most cases are the same for all projects. application roles are in most cases unique per project.
```
$ vi config.yaml

---

ansible_terraform:
  s3_bucket_region: us-west-2
  s3_bucket_name_prefix: ews-works
  ansible_roles:
    os_roles:
    - name: geerlingguy.ntp
      src: geerlingguy.ntp
      version: 2.1.0
    - name: geerlingguy.repo-epel
      src: geerlingguy.repo-epel
      version: master
    app_roles:
    - name: hello-world-pretend-role
      src: hello-world-pretend-role
      version: v0.1.0
```

### Create a project
```
$ cloud-infrastructure-sdk ansible-teraform init \
    --config ./config.yaml \
    --project-name my-project \
    --app-name foo \
    --dc-name dc1 \
    -e dev1 \
    -e dev2 \
    -e dev3 \
    -e dev4 \
    --infra-provider aws \
    --aws-region us-west-2 \
    --s3-bucket-region us-west-2 \
    --s3-bucket-name my-bucket \
    --vault-addr https://my-hashicorp-vault-server


$ cd my-project
```

### Install your project tools dependencies
```
make dependencies
```

### Configure your terragrunt environment variables

```
$ vim iaas/terraform/live/dc1/dev1/vars.yaml

---
name: foo-dc1-dev1


region: "us-west-2"
allowed_account_ids:
 - "<my-aws-account-id>"
tags:
  dcname: dc1
  environment: dev1

foo:
  nodes_count: 1
  nodes_instance_type: t3.medium
  ami_id: my-ami-id
  key_name: "my-key-name"
  disable_api_termination: false
  instance_auto_recovery_enabled: false
  vpc_id: "my-vpc-id"
  nodes_subnet_ids:
   - "my-subnet-id1-aza"
   - "my-subnet-id1-azb"
  sg_attachments: []
  sg_egress_cidr_rules: []
  sg_ingress_cidr_rules:
    - { desc: "mgmt ssh",  from_port: "22", to_port: "22", protocol: "tcp", cidr_blocks:  "my-cidr1,my-cidr2" }
  lb_subnet_ids:
   - "my-subnet-id1-aza"
   - "my-subnet-id1-azb"

dns_domain: <my-domain-name>
dns_zone_id: <my-dns-zone-id>

vault_addr: "https://my-hashicorp-vault-server"
vault_ssh_ca_path: "ssh-client-signer"
```

### Set your local environment variables
The following example selects the dev1 environment.
```
# set the environment to work on
$ export ENV=./env/dc1/dev1.sh

# source the ENV file to set the vault address environmental variable
$ source ./env/dc1/dev1.sh

# login to vault
$ vault login -method=<my-login-method> username=<my-vault-username>

# setup aws credentials for the environment as needed.
. ./scripts/aws-helpers.sh get_temporary_credentials
```

### Run terragrunt

```
$ make terragrunt-init
$ make terragrunt-plan
$ make terragrunt-apply
```

### Generate the ansible inventory file

```
$ make generate-ansible-inventory
```

### Run the ansible playbooks on the inventory
```
$ make ansible-playbook
```
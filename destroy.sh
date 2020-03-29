#!/bin/bash

export TF_VAR_do_token=$(< do_token)

export GOOGLE_APPLICATION_CREDENTIALS=/home/terraform/service-account-key.json

export TF_VAR_ssh_fingerprint="hello"

envsubst < main.tf | tee main.tf

terraform init
terraform destroy -auto-approve
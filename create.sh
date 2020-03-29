#!/bin/bash

printf "%s" "$pubkey" > "config/pubkey.b64"

eval `ssh-agent` # create the process

# Get ssh fingerprint
ssh-keygen -f /home/terraform/.ssh/id_rsa -q -N ""
ssh-add /home/terraform/.ssh/id_rsa
ssh-add -L

export TF_VAR_ssh_fingerprint=$(ssh-add -l -E md5| awk '{print $2}'|cut -d':' -f 2-)
export TF_VAR_do_token=$(< config/do_token)

# Add ssh pub key to do
export auth="Authorization: Bearer "$TF_VAR_do_token
export payload="{\"name\":\"$TF_VAR_cluster_id\",\"public_key\":\"$(cat /home/terraform/.ssh/id_rsa.pub)\"}"
curl -X POST -H "Content-Type: application/json" -H "$auth" -d "$payload" "https://api.digitalocean.com/v2/account/keys" 

export GOOGLE_APPLICATION_CREDENTIALS=/home/terraform/config/service-account-key.json

stk create

envsubst < main.tf | tee main.tf

terraform init
terraform plan -out create.plan 
terraform apply -auto-approve create.plan

stk encrypt

if [ ! -z "$HOW_LONG" ]
then
    sleep "$HOW_LONG"m
    terraform destroy -auto-approve
fi

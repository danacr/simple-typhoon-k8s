# Simple Typhoon K8s

### Tiny bash setup script for Typhoon K8s on Digital Ocean

The purpose of this repository is to provision a cluster in a really simple and fast way. Result of [this](https://twitter.com/errordeveloper/status/1240262848351211520) twitter thread.

This script automates the setup for [Typhoon K8s](https://github.com/poseidon/typhoon) which is an awesome upstream Kubernetes distribution. It uses upstream [hyperkube](https://typhoon.psdn.io/architecture/operating-systems/#kubernetes-properties) for the worker and control plane images.

> Requires [terraform](https://github.com/hashicorp/terraform) and [jq](https://github.com/stedolan/jq)

Create a cluster

```sh
# Set the token for digital ocean
export TF_VAR_do_token=DO_TOKEN_WITH_WRITE_PERMISSIONS

# Run setup script
bash setup.sh

# Apply configuration (ETA is 3 minutes)
terraform apply -auto-approve

# Setup kubeconfig
export KUBECONFIG=./cluster-config

# You're done
kubectl get nodes
```

Destroy a cluster:

```sh
terraform destroy -auto-approve
```

> Note: ETA for your cluster can be between 3 and 30 minutes. It depends on the region and time of day.

Test:

```
docker run -e HOW_LONG=1 -e CLUSTER_VERSION=1.17.4 -e TF_VAR_cluster_region=nyc3 -e TF_VAR_cluster_id=70a45d58-a13d-4f39-bef1-707b19ebfe55 -it test
```

Dev:

```
docker build . -t test
docker run -e CLUSTER_VERSION=1.17.4 -e TF_VAR_cluster_region=nyc3 -e TF_VAR_cluster_id=70a45d58-a13d-4f39-bef1-707b19ebfe55 -it test bash -c "source create.sh"
docker run -e CLUSTER_VERSION=1.17.4 -e TF_VAR_cluster_region=nyc3 -e TF_VAR_cluster_id=70a45d58-a13d-4f39-bef1-707b19ebfe55 -it test bash -c "source destroy.sh"
https://storage.googleapis.com/70a45d58-a13d-4f39-bef1-707b19ebfe55/cluster-config.gpg
```

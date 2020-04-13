module "cluster" {
  source = "git::https://github.com/poseidon/typhoon//digital-ocean/container-linux/kubernetes?ref=$CLUSTER_VERSION"

  # Digital Ocean
  cluster_name = var.cluster_id
  region       = var.cluster_region
  dns_zone     = "k8stfw.com"
  # controller_type = "s-4vcpu-8gb"
  # worker_type     = "s-2vcpu-2gb"

  # configuration
  ssh_fingerprints = [var.ssh_fingerprint]

  # optional
  worker_count = 2
}

terraform {
  backend "gcs" {
    bucket = "$TF_VAR_cluster_id"
    prefix = "terraform/state"
  }
}

# Obtain cluster kubeconfig
resource "local_file" "kubeconfig-cluster" {
  content  = module.cluster.kubeconfig-admin
  filename = "./cluster-config"
}


provider "digitalocean" {
  version = "1.15.1"
  token   = var.do_token
}

provider "google" {
  project = "k8sftw"
  region  = "us-east1"
}

provider "ct" {
  version = "0.5.0"
}

variable "cluster_id" {}
variable "cluster_region" {}
variable "do_token" {}
variable "ssh_fingerprint" {}

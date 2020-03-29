
provider "digitalocean" {
  version = "1.14.0"
  token   = var.do_token
}

provider "google" {
  project = "k8sftw"
  region  = "us-east1"
}

provider "ct" {
  version = "0.4.0"
}

variable "cluster_id" {}
variable "cluster_region" {}
variable "do_token" {}
variable "ssh_fingerprint" {}

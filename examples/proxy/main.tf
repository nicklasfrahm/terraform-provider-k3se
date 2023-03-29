terraform {
  required_providers {
    k3se = {
      source = "registry.terraform.io/nicklasfrahm/k3se"
    }
  }
}

provider "k3se" {}

resource "k3se_cluster" "proxy" {
  # Version may either be a specific k3s version or a release channel
  # as listed here: https://update.k3s.io/v1-release/channels
  version = "stable"

  # Cluster provides cluster-wide settings that should be applied
  # to all nodes in the cluster. All options are equivalent to the
  # commmand line options of the `k3s` command.
  cluster = {
    server = {
      # It is highly recommended to always specify this option as it
      # is used to determine the server URL of the cluster.
      tls_san = [
        "192.168.56.11",
      ]
      write_kubeconfig_mode = "644"
      node_label = [
        "example=proxy",
      ]
    }
  }

  # A list of all nodes in the cluster and their connection information.
  nodes = [
    {
      role = "server"
      ssh = {
        host     = "192.168.56.11"
        user     = "vagrant"
        key_file = "~/.ssh/id_ed25519"
      }
      server = {
        node_label = [
          "mylabel=a",
        ]
      }
    },
  ]

  # An SSH proxy, also known as jumpbox or a bastion host
  # can be used to access nodes in a private network.
  ssh_proxy = {
    host     = "192.168.56.11"
    user     = "vagrant"
    key_file = "~/.ssh/id_ed25519"
  }
}

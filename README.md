# Terraform Provider for k3se

This repository implements a Terraform provider for [k3se][k3se-github].

This provider uses the recommended way to write a Terraform provider using the [Terraform Plugin Framework][tf-plugin-framework] and is written in [Go][golang].

## Usage

```hcl
# The does not take any arguments.
provider "k3se" {}

# Create a new cluster.
resource "k3se_cluster" "single_node" {
  name = "example"
}
```

## License

This project is and will always be licensed under the MIT license. See [LICENSE.md](LICENSE.md) for more information.

[k3se-github]: https://gitub.com/nicklasfrahm/k3se
[tf-plugin-framework]: https://developer.hashicorp.com/terraform/plugin/framework
[golang]: https://go.dev/

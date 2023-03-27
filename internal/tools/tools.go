//go:build tools

// This is a hack to ensure that the documentation generation tool is not removed by `go mod tidy`.
// Reference: https://stackoverflow.com/questions/60233481/go-mod-tidy-removes-linters-from-go-mod
package tools

import (
	_ "github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs"
)

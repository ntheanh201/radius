// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package register

import (
	"github.com/project-radius/radius/pkg/cli/cmd/credential/common"
	credential_register_azure "github.com/project-radius/radius/pkg/cli/cmd/credential/register/azure"
	"github.com/project-radius/radius/pkg/cli/framework"
	"github.com/spf13/cobra"
)

// NewCommand creates an instance of the command for the `rad provider create` command.
func NewCommand(factory framework.Factory) *cobra.Command {
	// This command is not runnable, and thus has no runner.
	cmd := &cobra.Command{
		Use:   "register",
		Short: "Register(Add or update) cloud provider credential for a Radius installation.",
		Long:  "Register (Add or update) cloud provider configuration for a Radius installation." + common.LongDescriptionBlurb,
		Example: `
# Register (Add or update) cloud provider credential for Azure with service principal authentication
rad credential register azure --client-id <client id> --client-secret <client secret> --tenant-id <tenant id> --subscription <subscription id> --resource-group <resource group name>		
`,
	}

	azure, _ := credential_register_azure.NewCommand(factory)
	cmd.AddCommand(azure)

	return cmd
}

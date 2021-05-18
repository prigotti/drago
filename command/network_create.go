package command

import (
	"context"
	"flag"
	"fmt"
	"strings"

	structs "github.com/seashell/drago/drago/structs"
	cli "github.com/seashell/drago/pkg/cli"
)

// NetworkCreateCommand :
type NetworkCreateCommand struct {
	UI cli.UI

	// Parsed flags
	name         string
	addressRange string

	Command
}

func (c *NetworkCreateCommand) FlagSet() *flag.FlagSet {

	flags := c.Command.FlagSet(c.Name())

	flags.Usage = func() { c.UI.Output("\n" + c.Help() + "\n") }

	// General options
	flags.StringVar(&c.name, "name", "", "")
	flags.StringVar(&c.addressRange, "range", "", "")

	return flags
}

// Name :
func (c *NetworkCreateCommand) Name() string {
	return "network create"
}

// Synopsis :
func (c *NetworkCreateCommand) Synopsis() string {
	return "Create a new network"
}

// Run :
func (c *NetworkCreateCommand) Run(ctx context.Context, args []string) int {

	flags := c.FlagSet()

	if err := flags.Parse(args); err != nil {
		return 1
	}

	args = flags.Args()
	if len(args) > 0 {
		c.UI.Error("This command takes no arguments")
		return 1
	}

	// Get the HTTP client
	api, err := c.Command.APIClient()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error setting up API client: %s", err))
		return 1
	}

	err = api.Networks().Create(&structs.Network{
		Name:         c.name,
		AddressRange: c.addressRange,
	})
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error creating network: %s", err))
		return 1
	}

	return 0
}

// Help :
func (c *NetworkCreateCommand) Help() string {
	h := `
Usage: drago network create [options]

  Create a new Drago network.

  If ACLs are enabled, this option requires a token with the 'network:write' capability.

General Options:
` + GlobalOptions() + `

Network List Options:

  -name="string"
    Sets the human readable name for the network.

  -range="cidr"
    Sets the address range of the network, in CIDR notation.

 `
	return strings.TrimSpace(h)
}

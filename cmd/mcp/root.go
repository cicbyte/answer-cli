package mcp

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func runMcp(cmd *cobra.Command, args []string) error {
	fmt.Fprintln(os.Stderr, "MCP Server not yet implemented")
	return nil
}

func GetMcpCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "mcp",
		Short: "Start MCP Server",
		RunE:  runMcp,
	}
}

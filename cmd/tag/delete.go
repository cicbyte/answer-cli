package tag

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var tagDeleteYes bool

var deleteCmd = &cobra.Command{
	Use:   "delete <slug-name>",
	Short: "Delete a tag",
	Long: `Delete a tag. This action cannot be undone.

Examples:
  answer-cli tag delete go
  answer-cli tag delete go --yes`,
	Args: cobra.ExactArgs(1),
	Run:  runDelete,
}

func init() {
	deleteCmd.Flags().BoolVarP(&tagDeleteYes, "yes", "y", false, "Skip confirmation")
}

func runDelete(cmd *cobra.Command, args []string) {
	slugName := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if !tagDeleteYes {
		fmt.Printf("Delete tag '%s'? [y/N]: ", slugName)
		var input string
		fmt.Scanln(&input)
		if input != "y" && input != "yes" && input != "Y" {
			fmt.Println("Cancelled.")
			return
		}
	}

	if err := cli.Tag.Delete(context.Background(), slugName); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Tag '%s' deleted.\n", slugName)
}

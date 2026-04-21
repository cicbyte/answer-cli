package comment

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var commentDeleteYes bool

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a comment",
	Long: `Delete a comment. This action cannot be undone.

Examples:
  answer-cli comment delete 123
  answer-cli comment delete 123 --yes`,
	Args: cobra.ExactArgs(1),
	Run:  runDelete,
}

func init() {
	deleteCmd.Flags().BoolVarP(&commentDeleteYes, "yes", "y", false, "Skip confirmation")
}

func runDelete(cmd *cobra.Command, args []string) {
	id := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if !commentDeleteYes {
		fmt.Printf("Delete comment #%s? [y/N]: ", id)
		var input string
		fmt.Scanln(&input)
		if input != "y" && input != "yes" && input != "Y" {
			fmt.Println("Cancelled.")
			return
		}
	}

	if err := cli.Comment.Delete(context.Background(), id); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Comment #%s deleted.\n", id)
}

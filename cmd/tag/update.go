package tag

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/spf13/cobra"
)

var (
	tagUpdateName        string
	tagUpdateDescription string
)

var updateCmd = &cobra.Command{
	Use:   "update <slug-name>",
	Short: "Update a tag",
	Long: `Update an existing tag.

Examples:
  answer-cli tag update go --name="Go Language"
  answer-cli tag update go --description="Updated description"
  answer-cli tag update go --name="Go" --description="New description"`,
	Args: cobra.ExactArgs(1),
	Run:  runUpdate,
}

func init() {
	updateCmd.Flags().StringVar(&tagUpdateName, "name", "", "New display name")
	updateCmd.Flags().StringVarP(&tagUpdateDescription, "description", "d", "", "New description")
}

func runUpdate(cmd *cobra.Command, args []string) {
	slugName := args[0]

	if tagUpdateName == "" && tagUpdateDescription == "" {
		fmt.Fprintln(os.Stderr, "Error: specify at least --name or --description")
		os.Exit(1)
	}

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	req := &models.TagUpdateReq{
		SlugName:     slugName,
		DisplayName:  tagUpdateName,
		OriginalText: tagUpdateDescription,
	}

	if err := cli.Tag.Update(context.Background(), req); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Tag '%s' updated successfully.\n", slugName)
}

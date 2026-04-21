package tag

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/spf13/cobra"
)

var (
	tagCreateSlug        string
	tagCreateName        string
	tagCreateDescription string
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a tag",
	Long: `Create a new tag.

Examples:
  answer-cli tag create --slug=go --name="Go Language"
  answer-cli tag create --slug=go --name="Go" --description="Go programming language"`,
	Run: runCreate,
}

func init() {
	createCmd.Flags().StringVar(&tagCreateSlug, "slug", "", "Tag slug (required)")
	createCmd.Flags().StringVar(&tagCreateName, "name", "", "Tag display name (required)")
	createCmd.Flags().StringVarP(&tagCreateDescription, "description", "d", "", "Tag description")
}

func runCreate(cmd *cobra.Command, args []string) {
	if tagCreateSlug == "" {
		fmt.Fprintln(os.Stderr, "Error: --slug is required")
		os.Exit(1)
	}
	if tagCreateName == "" {
		fmt.Fprintln(os.Stderr, "Error: --name is required")
		os.Exit(1)
	}

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	req := &models.TagAddReq{
		SlugName:     tagCreateSlug,
		DisplayName:  tagCreateName,
		OriginalText: tagCreateDescription,
	}

	result, err := cli.Tag.Add(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Tag created successfully!\n")
	fmt.Printf("  Slug:  %s\n", result.SlugName)
	fmt.Printf("  Name:  %s\n", result.DisplayName)
}

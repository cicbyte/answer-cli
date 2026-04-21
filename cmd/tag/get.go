package tag

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var tagGetJSON bool

var getCmd = &cobra.Command{
	Use:   "get <slug-name>",
	Short: "Get tag detail",
	Long: `Get detailed information about a tag.

Examples:
  answer-cli tag get go
  answer-cli tag get go --json`,
	Args: cobra.ExactArgs(1),
	Run:  runGet,
}

func init() {
	getCmd.Flags().BoolVar(&tagGetJSON, "json", false, "Output in JSON format")
}

func runGet(cmd *cobra.Command, args []string) {
	slugName := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	tag, err := cli.Tag.Get(context.Background(), slugName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if tagGetJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(tag)
		return
	}

	fmt.Printf("Slug:       %s\n", tag.SlugName)
	fmt.Printf("Name:       %s\n", tag.DisplayName)
	fmt.Printf("Questions:  %d\n", tag.QuestionCount)
	fmt.Printf("Followers:  %d\n", tag.FollowCount)
	fmt.Printf("Status:     %d\n", tag.Status)
	fmt.Printf("Recommend:  %v\n", tag.Recommend)
	fmt.Printf("Reserved:   %v\n", tag.Reserved)
	if tag.CreatedAt > 0 {
		fmt.Printf("Created:    %s\n", models.FormatTimestamp(tag.CreatedAt).Format("2006-01-02 15:04:05"))
	}
	if tag.OriginalText != "" {
		fmt.Println()
		fmt.Println("Description:")
		fmt.Println(tag.OriginalText)
	}
}

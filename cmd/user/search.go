package user

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var userSearchJSON bool

var searchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search users",
	Long: `Search for users by username or display name.

Examples:
  answer-cli user search "john"
  answer-cli user search "john" --json`,
	Args: cobra.ExactArgs(1),
	Run:  runSearch,
}

func init() {
	searchCmd.Flags().BoolVar(&userSearchJSON, "json", false, "Output in JSON format")
}

func runSearch(cmd *cobra.Command, args []string) {
	query := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	users, err := cli.User.SearchUsers(context.Background(), query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if userSearchJSON || output.GetOutputFormat("") == "json" {
		output.PrintJSON(users)
		return
	}

	if len(users) == 0 {
		fmt.Println("No users found.")
		return
	}

	headers := []string{"Username", "Display Name"}
	var rows [][]string
	for _, u := range users {
		displayName := u.DisplayName
		if displayName == "" {
			displayName = u.Username
		}
		rows = append(rows, []string{
			u.Username,
			displayName,
		})
	}

	output.PrintTable(headers, rows)
	fmt.Printf("\nTotal: %d\n", len(users))
}

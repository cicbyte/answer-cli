package user

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	userInfoUsername string
	userInfoJSON     bool
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Get user info",
	Long: `Get user information. Defaults to current user if no username specified.

Examples:
  answer-cli user info
  answer-cli user info --username=john
  answer-cli user info --json`,
	Run: runInfo,
}

func init() {
	infoCmd.Flags().StringVarP(&userInfoUsername, "username", "u", "", "Username (default: current user)")
	infoCmd.Flags().BoolVar(&userInfoJSON, "json", false, "Output in JSON format")
}

func runInfo(cmd *cobra.Command, args []string) {
	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if userInfoUsername != "" {
		users, err := cli.User.SearchUsers(context.Background(), userInfoUsername)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if len(users) == 0 {
			fmt.Fprintf(os.Stderr, "User '%s' not found.\n", userInfoUsername)
			os.Exit(1)
		}
		u := users[0]
		if userInfoJSON || output.GetOutputFormat("") == "json" {
			output.PrintJSON(u)
			return
		}
		printUserBasicInfo(u)
	} else {
		u, err := cli.User.GetByUsername(context.Background())
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if userInfoJSON || output.GetOutputFormat("") == "json" {
			output.PrintJSON(u)
			return
		}
		printUserBasicInfo(u)
	}
}

func printUserBasicInfo(u *models.UserBasicInfo) {
	fmt.Printf("Username: %s\n", u.Username)
	fmt.Printf("Display:  %s\n", u.DisplayName)
	if u.Rank > 0 {
		fmt.Printf("Rank:     %d\n", u.Rank)
	}
}

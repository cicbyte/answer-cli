package notification

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	notificationReadAll  bool
	notificationReadType string
)

var readCmd = &cobra.Command{
	Use:   "read [id]",
	Short: "Mark notification(s) as read",
	Long: `Mark one or all notifications as read.

Examples:
  answer-cli notification read --all
  answer-cli notification read --all --type=inbox
  answer-cli notification read 12345`,
	Run: runRead,
}

func init() {
	readCmd.Flags().BoolVar(&notificationReadAll, "all", false, "Mark all as read")
	readCmd.Flags().StringVar(&notificationReadType, "type", "", "Notification type for --all (inbox|achievement)")
}

func runRead(cmd *cobra.Command, args []string) {
	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if notificationReadAll {
		if err := cli.Notification.ClearAll(context.Background(), notificationReadType); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("All notifications marked as read.")
		return
	}

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Error: provide notification ID or use --all")
		os.Exit(1)
	}

	id := args[0]
	if err := cli.Notification.ClearID(context.Background(), id); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Notification #%s marked as read.\n", id)
}

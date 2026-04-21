package question

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	questionCloseType int
	questionCloseMsg  string
)

func getCloseCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "close <id>",
		Short: "关闭问题",
		Long: `关闭问题。

示例:
  answer-cli question close 123
  answer-cli question close 123 --type=2 --msg="Duplicate of..."`,
		Args: cobra.ExactArgs(1),
		Run:  runClose,
	}

	cmd.Flags().IntVarP(&questionCloseType, "type", "t", 1, "关闭类型 (1=问题已解决, 2=重复问题, 3=非主题问题, 4=其他)")
	cmd.Flags().StringVarP(&questionCloseMsg, "msg", "m", "", "关闭原因说明")

	return cmd
}

func runClose(cmd *cobra.Command, args []string) {
	id := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	if err := cli.Question.Close(context.Background(), id, questionCloseType, questionCloseMsg); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Question #%s closed successfully!\n", id)
}

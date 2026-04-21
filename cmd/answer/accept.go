package answer

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var answerAcceptQuestionID string

func getAcceptCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "accept <answer-id>",
		Short: "采纳回答",
		Long: `将回答标记为已采纳。

示例:
  answer-cli answer accept 456 --question-id=123
  answer-cli answer accept 456`,
		Args: cobra.ExactArgs(1),
		Run:  runAccept,
	}

	cmd.Flags().StringVarP(&answerAcceptQuestionID, "question-id", "q", "", "问题 ID (如未指定则自动检测)")

	return cmd
}

func runAccept(cmd *cobra.Command, args []string) {
	answerID := args[0]

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// If question-id not provided, try to get it from the answer info
	questionID := answerAcceptQuestionID
	if questionID == "" {
		a, err := cli.Answer.Get(context.Background(), answerID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: cannot auto-detect question id: %v\n", err)
			fmt.Fprintln(os.Stderr, "Please provide --question-id")
			os.Exit(1)
		}
		questionID = a.QuestionID
	}

	if err := cli.Answer.Accept(context.Background(), questionID, answerID); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Answer #%s accepted for question #%s!\n", answerID, questionID)
}

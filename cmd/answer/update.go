package answer

import (
	"context"
	"fmt"
	"os"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	answerUpdateContent string
	answerUpdateFile    string
)

func getUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <id>",
		Short: "更新回答",
		Long: `更新已有的回答。

示例:
  answer-cli answer update 456 --content="Updated answer..."
  answer-cli answer update 456 --file=answer.md
  echo "new content" | answer-cli answer update 456`,
		Args: cobra.ExactArgs(1),
		Run:  runUpdate,
	}

	cmd.Flags().StringVarP(&answerUpdateContent, "content", "c", "", "新内容")
	cmd.Flags().StringVarP(&answerUpdateFile, "file", "f", "", "从文件读取内容")

	return cmd
}

func runUpdate(cmd *cobra.Command, args []string) {
	id := args[0]

	if answerUpdateContent == "" && answerUpdateFile == "" {
		pipeContent, err := output.ReadPipeOrFile("")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading pipe input: %v\n", err)
			os.Exit(1)
		}
		if pipeContent != "" {
			answerUpdateContent = pipeContent
		}
	}

	if answerUpdateContent == "" && answerUpdateFile == "" {
		fmt.Fprintln(os.Stderr, "Error: specify --content, --file, or pipe input")
		os.Exit(1)
	}

	content := answerUpdateContent

	if answerUpdateFile != "" {
		data, err := os.ReadFile(answerUpdateFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
		content = string(data)
	}

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	req := &models.AnswerUpdateReq{
		ID:      id,
		Content: content,
	}

	if err := cli.Answer.Update(context.Background(), req); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Answer #%s updated successfully!\n", id)
}

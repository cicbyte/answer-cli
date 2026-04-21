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
	answerCreateContent string
	answerCreateFile    string
)

func getCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create <question-id>",
		Short:   "创建回答",
		Aliases: []string{"new", "add"},
		Long: `为指定问题创建回答。

示例:
  answer-cli answer create 123 --content="This is my answer..."
  answer-cli answer create 123 --file=answer.md
  echo "answer content" | answer-cli answer create 123`,
		Args: cobra.ExactArgs(1),
		Run:  runCreate,
	}

	cmd.Flags().StringVarP(&answerCreateContent, "content", "c", "", "回答内容")
	cmd.Flags().StringVarP(&answerCreateFile, "file", "f", "", "从文件读取内容")

	return cmd
}

func runCreate(cmd *cobra.Command, args []string) {
	questionID := args[0]

	content := answerCreateContent

	if answerCreateFile != "" {
		data, err := os.ReadFile(answerCreateFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
			os.Exit(1)
		}
		content = string(data)
	}

	if content == "" {
		pipeContent, err := output.ReadPipeOrFile("")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading pipe input: %v\n", err)
			os.Exit(1)
		}
		if pipeContent != "" {
			content = pipeContent
		}
	}

	if content == "" {
		fmt.Fprintln(os.Stderr, "Error: answer content is required (use --content, --file, or pipe input)")
		os.Exit(1)
	}

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	req := &models.AnswerAddReq{
		QuestionID: questionID,
		Content:    content,
	}

	result, err := cli.Answer.Add(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Answer created successfully!\n")
	fmt.Printf("  ID:          %s\n", result.ID)
	fmt.Printf("  Question ID: %s\n", result.QuestionID)
}

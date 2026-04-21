package question

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/cicbyte/answer-cli/internal/models"
	"github.com/cicbyte/answer-cli/internal/output"
	"github.com/spf13/cobra"
)

var (
	questionUpdateTitle   string
	questionUpdateContent string
	questionUpdateTags    string
	questionUpdateFile    string
)

func getUpdateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update <id>",
		Short: "更新问题",
		Long: `更新已有的问题。

示例:
  answer-cli question update 123 --title="New title"
  answer-cli question update 123 --content="Updated content"
  answer-cli question update 123 --tags=go,web
  answer-cli question update 123 --file=content.md`,
		Args: cobra.ExactArgs(1),
		Run:  runUpdate,
	}

	cmd.Flags().StringVarP(&questionUpdateTitle, "title", "t", "", "新标题")
	cmd.Flags().StringVarP(&questionUpdateContent, "content", "c", "", "新内容")
	cmd.Flags().StringVar(&questionUpdateTags, "tags", "", "标签 (逗号分隔)")
	cmd.Flags().StringVarP(&questionUpdateFile, "file", "f", "", "从文件读取内容")

	return cmd
}

func runUpdate(cmd *cobra.Command, args []string) {
	id := args[0]

	if questionUpdateTitle == "" && questionUpdateContent == "" && questionUpdateTags == "" && questionUpdateFile == "" {
		fmt.Fprintln(os.Stderr, "Error: specify at least one of --title, --content, --tags, or --file")
		os.Exit(1)
	}

	content := questionUpdateContent

	if questionUpdateFile != "" {
		data, err := os.ReadFile(questionUpdateFile)
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

	cli, err := getClient()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	req := &models.QuestionUpdateReq{
		ID:      id,
		Title:   questionUpdateTitle,
		Content: content,
	}

	if questionUpdateTags != "" {
		tags := strings.Split(questionUpdateTags, ",")
		for i, tag := range tags {
			tags[i] = strings.TrimSpace(tag)
		}
		req.Tags = tags
	}

	if err := cli.Question.Update(context.Background(), req); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Question #%s updated successfully!\n", id)
}

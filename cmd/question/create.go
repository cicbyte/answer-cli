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
	questionCreateTitle   string
	questionCreateContent string
	questionCreateTags    string
	questionCreateFile    string
)

func getCreateCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "create",
		Short:   "创建问题",
		Aliases: []string{"new", "add"},
		Long: `创建新问题。

示例:
  answer-cli question create --title="How to..." --content="..."
  answer-cli question create --title="Bug report" --tags=bug,go
  answer-cli question create --title="Issue" --file=content.md
  echo "content" | answer-cli question create --title="Question"`,
		Run: runCreate,
	}

	cmd.Flags().StringVarP(&questionCreateTitle, "title", "t", "", "问题标题 (必填)")
	cmd.Flags().StringVarP(&questionCreateContent, "content", "c", "", "问题内容")
	cmd.Flags().StringVar(&questionCreateTags, "tags", "", "标签 (逗号分隔)")
	cmd.Flags().StringVarP(&questionCreateFile, "file", "f", "", "从文件读取内容")

	return cmd
}

func runCreate(cmd *cobra.Command, args []string) {
	if questionCreateTitle == "" {
		fmt.Fprintln(os.Stderr, "Error: --title is required")
		os.Exit(1)
	}

	content := questionCreateContent

	if questionCreateFile != "" {
		data, err := os.ReadFile(questionCreateFile)
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

	req := &models.QuestionAddReq{
		Title:   questionCreateTitle,
		Content: content,
	}

	if questionCreateTags != "" {
		tags := strings.Split(questionCreateTags, ",")
		for i, tag := range tags {
			tags[i] = strings.TrimSpace(tag)
		}
		req.Tags = tags
	}

	result, err := cli.Question.Add(context.Background(), req)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Question created successfully!\n")
	fmt.Printf("  ID:    %s\n", result.ID)
	fmt.Printf("  Title: %s\n", result.Title)
}

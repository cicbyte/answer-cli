package search

import (
	"github.com/cicbyte/answer-cli/cmd/question"
	"github.com/spf13/cobra"
)

// GetSearchCommand 返回 search 命令，实际复用 question list 的逻辑
func GetSearchCommand() *cobra.Command {
	cmd := question.NewListAsSearch()
	cmd.Use = "search [query]"
	cmd.Aliases = nil
	cmd.Short = "搜索问题和回答（等同于 question list <query>）"
	return cmd
}

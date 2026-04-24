package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/glamour"
	"golang.org/x/term"
)

var globalFormat string

func SetFormat(f string) {
	globalFormat = f
}

func GetFormat(override string) string {
	if override != "" {
		return override
	}
	return globalFormat
}

func GetOutputFormat(override string) string {
	return GetFormat(override)
}

func IsJSON(override string) bool {
	f := GetFormat(override)
	return f == "json" || f == "jsonl"
}

func IsJSONL(override string) bool {
	return GetFormat(override) == "jsonl"
}

func PrintJSON(data any) {
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(data)
}

func PrintJSONL(items []map[string]any) {
	enc := json.NewEncoder(os.Stdout)
	for _, item := range items {
		enc.Encode(item)
	}
}

type Item struct {
	Title    string
	Subtitle string
	Tags     []string
}

func PrintItems(items []Item, footer string) {
	if len(items) == 0 {
		fmt.Println("暂无数据")
		return
	}

	w, _, _ := GetTermSize()
	contentW := w - 4

	sep := "─"

	for i, item := range items {
		title := item.Title
		runes := []rune(title)
		if len(runes) > contentW {
			title = string(runes[:contentW-3]) + "..."
		}

		fmt.Println(title)

		parts := make([]string, 0, 3)
		if item.Subtitle != "" {
			parts = append(parts, item.Subtitle)
		}
		if len(item.Tags) > 0 {
			tags := make([]string, len(item.Tags))
			for j, t := range item.Tags {
				tags[j] = "[" + t + "]"
			}
			parts = append(parts, strings.Join(tags, " "))
		}
		if len(parts) > 0 {
			fmt.Printf("  %s\n", Dim(strings.Join(parts, "  ")))
		}

		if i < len(items)-1 {
			fmt.Println(Dim(strings.Repeat(sep, min(contentW, 60))))
		}
	}

	if footer != "" {
		fmt.Println()
		fmt.Println(Dim(footer))
	}
}

func PrintTable(headers []string, rows [][]string) {
	items := make([]Item, len(rows))
	for i, row := range rows {
		items[i] = Item{
			Title:    row[0],
			Subtitle: strings.Join(row[1:], "  "),
		}
	}
	PrintItems(items, "")
}

func Truncate(s string, maxLen int) string {
	s = strings.ReplaceAll(s, "\n", " ")
	runes := []rune(s)
	if len(runes) <= maxLen {
		return string(runes)
	}
	return string(runes[:maxLen-1]) + "…"
}

func ReadPipeOrFile(filePath string) (string, error) {
	if filePath != "" {
		data, err := os.ReadFile(filePath)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}

	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			return "", err
		}
		return string(data), nil
	}

	return "", nil
}

func RenderMarkdown(content string) string {
	w, _, _ := GetTermSize()
	return RenderMarkdownWidth(content, w)
}

func RenderMarkdownWidth(content string, width int) string {
	if content == "" {
		return ""
	}
	if width <= 0 {
		width = 80
	}
	r, err := glamour.NewTermRenderer(
		glamour.WithAutoStyle(),
		glamour.WithWordWrap(width),
	)
	if err != nil {
		return content
	}
	out, err := r.Render(content)
	if err != nil {
		return content
	}
	return out
}

func Dim(s string) string {
	return "\x1b[2m" + s + "\x1b[0m"
}

func GetTermSize() (int, int, error) {
	w, h, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		return 80, 24, nil
	}
	return w, h, nil
}

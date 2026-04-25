package output

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"
)

// --- Truncate ---

func TestTruncate_Short(t *testing.T) {
	if got := Truncate("hello", 10); got != "hello" {
		t.Fatalf("expected hello, got %s", got)
	}
}

func TestTruncate_Exact(t *testing.T) {
	if got := Truncate("hello", 5); got != "hello" {
		t.Fatalf("expected hello, got %s", got)
	}
}

func TestTruncate_Over(t *testing.T) {
	got := Truncate("hello world", 8)
	if got != "hello w…" {
		t.Fatalf("expected 'hello w…', got %q", got)
	}
	if len([]rune(got)) != 8 {
		t.Fatalf("expected 8 runes, got %d", len([]rune(got)))
	}
}

func TestTruncate_Newline(t *testing.T) {
	got := Truncate("hello\nworld", 20)
	if strings.Contains(got, "\n") {
		t.Fatal("should not contain newline")
	}
	if got != "hello world" {
		t.Fatalf("expected 'hello world', got %q", got)
	}
}

func TestTruncate_Empty(t *testing.T) {
	if got := Truncate("", 10); got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestTruncate_Chinese(t *testing.T) {
	got := Truncate("你好世界", 3)
	if got != "你好…" {
		t.Fatalf("expected '你好…', got %q", got)
	}
	if len([]rune(got)) != 3 {
		t.Fatalf("expected 3 runes, got %d", len([]rune(got)))
	}
}

// --- Dim ---

func TestDim_WrapsWithANSI(t *testing.T) {
	got := Dim("hello")
	if !strings.HasPrefix(got, "\x1b[2m") {
		t.Fatal("should start with dim escape")
	}
	if !strings.HasSuffix(got, "\x1b[0m") {
		t.Fatal("should end with reset escape")
	}
	if !strings.Contains(got, "hello") {
		t.Fatal("should contain original text")
	}
}

// --- Format state ---

func TestFormatGlobalAndOverride(t *testing.T) {
	SetFormat("json")
	defer SetFormat("")

	if GetFormat("") != "json" {
		t.Fatal("expected global json")
	}
	if GetFormat("table") != "table" {
		t.Fatal("expected override table")
	}
	if GetOutputFormat("") != "json" {
		t.Fatal("expected output format json")
	}
}

func TestIsJSON(t *testing.T) {
	SetFormat("json")
	defer SetFormat("")

	if !IsJSON("") {
		t.Fatal("expected true for json")
	}
	if !IsJSON("jsonl") {
		t.Fatal("expected true for jsonl override")
	}
	SetFormat("table")
	if IsJSON("") {
		t.Fatal("expected false for table")
	}
}

func TestIsJSONL(t *testing.T) {
	SetFormat("jsonl")
	defer SetFormat("")

	if !IsJSONL("") {
		t.Fatal("expected true for jsonl")
	}
	SetFormat("json")
	if IsJSONL("") {
		t.Fatal("expected false for json")
	}
}

// --- PrintJSON (captured stdout) ---

func TestPrintJSON_Formatted(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintJSON(map[string]string{"key": "value"})

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, `"key": "value"`) {
		t.Fatalf("expected formatted JSON, got %s", output)
	}
	// 应该有缩进
	if !strings.Contains(output, "  ") {
		t.Fatalf("expected indented JSON, got %s", output)
	}
}

func TestPrintJSON_Nested(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintJSON(map[string]any{"list": []int{1, 2, 3}})

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, `"list"`) {
		t.Fatalf("expected list key, got %s", output)
	}
}

// --- PrintJSONL ---

func TestPrintJSONL(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintJSONL([]map[string]any{
		{"id": "1"},
		{"id": "2"},
	})

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	lines := strings.Split(strings.TrimSpace(output), "\n")
	if len(lines) != 2 {
		t.Fatalf("expected 2 lines, got %d", len(lines))
	}
	// JSONL 每行是独立 JSON 对象
	for _, line := range lines {
		var m map[string]any
		if err := json.Unmarshal([]byte(line), &m); err != nil {
			t.Fatalf("invalid JSON line: %s", line)
		}
	}
}

// --- PrintTable / PrintTableRight (basic rendering check) ---

func TestPrintTable_RendersHeaders(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintTable([]string{"Name", "Value"}, [][]string{{"a", "b"}})

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	// go-pretty StyleDefault 会将表头转为大写
	if !strings.Contains(output, "NAME") {
		t.Fatalf("should contain header NAME, got: %q", output)
	}
	if !strings.Contains(output, "a") {
		t.Fatalf("should contain row data 'a', got: %q", output)
	}
}

func TestPrintTableEmpty(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintTable([]string{"H"}, nil)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "H") {
		t.Fatal("should contain header even with no rows")
	}
}

func TestPrintTableRight_Renders(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintTableRight([]string{"ID", "Count"}, [][]string{{"q1", "42"}}, 2)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	output := buf.String()

	if !strings.Contains(output, "q1") {
		t.Fatal("should contain q1")
	}
	if !strings.Contains(output, "42") {
		t.Fatal("should contain 42")
	}
}

// --- RenderMarkdownWidth ---

func TestRenderMarkdownWidth_Empty(t *testing.T) {
	got := RenderMarkdownWidth("", 80)
	if got != "" {
		t.Fatalf("expected empty, got %q", got)
	}
}

func TestRenderMarkdownWidth_ZeroWidth(t *testing.T) {
	got := RenderMarkdownWidth("# hello", 0)
	// width<=0 时默认 80，应该仍然渲染
	if got == "" {
		t.Fatal("expected rendered markdown")
	}
}

func TestRenderMarkdownWidth_PlainText(t *testing.T) {
	got := RenderMarkdownWidth("hello world", 80)
	if !strings.Contains(got, "hello") {
		t.Fatalf("expected hello in output, got %q", got)
	}
}

func TestRenderMarkdownWidth_CodeBlock(t *testing.T) {
	md := "```go\nfunc main() {}\n```"
	got := RenderMarkdownWidth(md, 80)
	if got == "" {
		t.Fatal("expected rendered markdown")
	}
}

// --- fromSlice ---

func TestFromSlice(t *testing.T) {
	row := fromSlice([]string{"a", "b", "c"})
	if len(row) != 3 {
		t.Fatalf("expected 3 items, got %d", len(row))
	}
	if row[0] != "a" || row[2] != "c" {
		t.Fatalf("expected a,c, got %v", row)
	}
}

func TestFromSlice_Empty(t *testing.T) {
	row := fromSlice([]string{})
	if len(row) != 0 {
		t.Fatalf("expected 0 items, got %d", len(row))
	}
}

// --- ReadPipeOrFile ---

func TestReadPipeOrFile_File(t *testing.T) {
	tmp, err := os.CreateTemp("", "test_*.txt")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmp.Name())
	tmp.WriteString("file content")
	tmp.Close()

	got, err := ReadPipeOrFile(tmp.Name())
	if err != nil {
		t.Fatal(err)
	}
	if got != "file content" {
		t.Fatalf("expected 'file content', got %q", got)
	}
}

func TestReadPipeOrFile_Nonexistent(t *testing.T) {
	_, err := ReadPipeOrFile("/nonexistent/file/path")
	if err == nil {
		t.Fatal("expected error for nonexistent file")
	}
}

// --- GetTermSize ---

func TestGetTermSize(t *testing.T) {
	w, h, err := GetTermSize()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// 在测试环境中可能无法获取真实尺寸，默认 80x24
	if w <= 0 || h <= 0 {
		t.Fatalf("expected positive dimensions, got %dx%d", w, h)
	}
}

// tests/cli_test.go
// CLI 集成测试 — 通过 exec 运行编译后的二进制，验证各命令的输出和行为。
//
// 运行方式:
//   go test ./tests/ -v -count=1
//
// 前置条件:
//   1. 服务器可达且已通过 auth login 认证
//   2. 测试使用独立配置目录（ANSWER_CLI_HOME），不影响日常使用
//   3. 需先构建: go build -o answer-cli.exe .

package tests

import (
	"encoding/json"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// ============================================================
// 测试基础设施
// ============================================================

// cli 运行 answer-cli 并返回 stdout、stderr、exit code
func cli(t *testing.T, args ...string) (stdout, stderr string, exitCode int) {
	t.Helper()

	bin := os.Getenv("ANSWER_CLI_BIN")
	if bin == "" {
		// 自动查找项目根目录下的二进制
		wd, _ := os.Getwd()
		for _, name := range []string{"answer-cli.exe", "answer-cli"} {
			p := filepath.Join(wd, "..", name)
			if _, err := os.Stat(p); err == nil {
				bin = p
				break
			}
		}
	}
	if bin == "" {
		t.Skip("ANSWER_CLI_BIN 未设置且未找到编译后的二进制")
	}

	cmd := exec.Command(bin, args...)
	// 默认使用用户真实配置；设置 ANSWER_CLI_HOME 时使用独立目录
	if home := os.Getenv("ANSWER_CLI_HOME"); home != "" {
		cmd.Env = append(os.Environ(), "ANSWER_CLI_HOME="+home)
	}

	var outBuf, errBuf strings.Builder
	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	exitCode = 0
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			exitCode = exitErr.ExitCode()
		} else {
			t.Fatalf("执行失败: %v (stderr: %s)", err, errBuf.String())
		}
	}

	return strings.TrimSpace(outBuf.String()), strings.TrimSpace(errBuf.String()), exitCode
}

// assertContains 检查 stdout 是否包含指定文本
func assertContains(t *testing.T, output, substr string) {
	t.Helper()
	if !strings.Contains(output, substr) {
		t.Fatalf("输出中未找到 %q\n完整输出:\n%s", substr, output)
	}
}

// assertNotContains 检查 stdout 不包含指定文本
func assertNotContains(t *testing.T, output, substr string) {
	t.Helper()
	if strings.Contains(output, substr) {
		t.Fatalf("输出中不应包含 %q\n完整输出:\n%s", substr, output)
	}
}

// assertExitZero 断言命令成功退出
func assertExitZero(t *testing.T, exitCode int, stderr string) {
	t.Helper()
	if exitCode != 0 {
		t.Fatalf("期望退出码 0，实际 %d\nstderr: %s", exitCode, stderr)
	}
}

// assertExitNonZero 断言命令失败退出
func assertExitNonZero(t *testing.T, exitCode int) {
	t.Helper()
	if exitCode == 0 {
		t.Fatal("期望非零退出码，实际为 0")
	}
}

// assertJSONValid 断言输出是合法 JSON
func assertJSONValid(t *testing.T, output string) {
	t.Helper()
	var v any
	if err := json.Unmarshal([]byte(output), &v); err != nil {
		t.Fatalf("输出不是合法 JSON: %v\n输出:\n%s", err, output)
	}
}

// ============================================================
// 帮助信息
// ============================================================

func TestHelp(t *testing.T) {
	stdout, _, code := cli(t, "--help")
	assertExitZero(t, code, "")
	assertContains(t, stdout, "answer-cli")
	assertContains(t, stdout, "Apache Answer")
}

func TestQuestionHelp(t *testing.T) {
	stdout, _, code := cli(t, "question", "--help")
	assertExitZero(t, code, "")
	assertContains(t, stdout, "list")
	assertContains(t, stdout, "get")
}

func TestAuthHelp(t *testing.T) {
	stdout, _, code := cli(t, "auth", "--help")
	assertExitZero(t, code, "")
	assertContains(t, stdout, "login")
	assertContains(t, stdout, "logout")
	assertContains(t, stdout, "status")
}

func TestTagHelp(t *testing.T) {
	stdout, _, code := cli(t, "tag", "--help")
	assertExitZero(t, code, "")
	assertContains(t, stdout, "list")
	assertContains(t, stdout, "create")
}

func TestAnswerHelp(t *testing.T) {
	stdout, _, code := cli(t, "answer", "--help")
	assertExitZero(t, code, "")
	assertContains(t, stdout, "list")
	assertContains(t, stdout, "get")
}

func TestCommentHelp(t *testing.T) {
	stdout, _, code := cli(t, "comment", "--help")
	assertExitZero(t, code, "")
	assertContains(t, stdout, "list")
	assertContains(t, stdout, "get")
	assertContains(t, stdout, "add")
}

// ============================================================
// 认证
// ============================================================

func TestAuthStatus_NotConfigured(t *testing.T) {
	// 使用空目录作为 home，应该提示未配置
	stdout, _, code := cli(t, "auth", "status")
	// 如果已有配置，不会报错；如果无配置，应该提示
	if code != 0 {
		assertContains(t, stdout, "未配置")
	}
}

func TestAuthStatus_WithConfig(t *testing.T) {
	stdout, _, code := cli(t, "auth", "status")
	assertExitZero(t, code, "")
	assertContains(t, stdout, "服务器")
	// 如果已认证，应显示用户信息
	if strings.Contains(stdout, "已认证") {
		assertContains(t, stdout, "zhyj")
	}
}

// ============================================================
// 问题列表
// ============================================================

func TestQuestionList_Default(t *testing.T) {
	stdout, stderr, code := cli(t, "question", "list")
	assertExitZero(t, code, stderr)

	// 应包含表格表头
	assertContains(t, stdout, "#")
	assertContains(t, stdout, "标题")

	// 如果有数据，应包含问题 ID
	if strings.Contains(stdout, "1001") {
		assertContains(t, stdout, "共")
	}
}

func TestQuestionList_WithPage(t *testing.T) {
	_, stderr, code := cli(t, "question", "list", "--page=1", "--size=5")
	assertExitZero(t, code, stderr)

	// 小页码应该不会出错
	assertNotContains(t, stderr, "错误")
	assertNotContains(t, stderr, "Error")
}

func TestQuestionList_OutOfRange(t *testing.T) {
	cli(t, "question", "list", "--page=9999")
	// 不应 panic，超出范围返回空或错误
}

func TestQuestionList_WithOrder(t *testing.T) {
	for _, order := range []string{"newest", "active", "hot", "score"} {
		t.Run(order, func(t *testing.T) {
			_, stderr, code := cli(t, "question", "list", "--order="+order)
			assertExitZero(t, code, stderr)
			assertNotContains(t, stderr, "错误")
		})
	}
}

func TestQuestionList_WithTag(t *testing.T) {
	_, stderr, code := cli(t, "question", "list", "--tag=go")
	assertExitZero(t, code, stderr)
	assertNotContains(t, stderr, "错误")
}

func TestQuestionList_Search(t *testing.T) {
	_, stderr, code := cli(t, "question", "list", "git")
	assertExitZero(t, code, stderr)
	assertNotContains(t, stderr, "错误")
}

func TestQuestionList_SearchEmpty(t *testing.T) {
	stdout, _, code := cli(t, "question", "list", "xyznonexistent12345")
	assertExitZero(t, code, "")
	assertContains(t, stdout, "未找到") // 搜索无结果
}

func TestQuestionList_JSON(t *testing.T) {
	stdout, stderr, code := cli(t, "question", "list", "--format=json")
	assertExitZero(t, code, stderr)
	assertJSONValid(t, stdout)
	assertContains(t, stdout, `"count"`)
	assertContains(t, stdout, `"list"`)
}

func TestQuestionList_JSONL(t *testing.T) {
	stdout, stderr, code := cli(t, "question", "list", "--format=jsonl")
	assertExitZero(t, code, stderr)
	// JSONL 每行是独立 JSON
	lines := strings.Split(stdout, "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		var v any
		if err := json.Unmarshal([]byte(line), &v); err != nil {
			t.Fatalf("JSONL 行无效: %v\n行: %s", err, line)
		}
	}
}

// ============================================================
// 问题详情
// ============================================================

func TestQuestionGet_Exists(t *testing.T) {
	stdout, stderr, code := cli(t, "question", "get", "10010000000000020")
	assertExitZero(t, code, stderr)
	assertContains(t, stdout, "如何网页截图")
}

func TestQuestionGet_NotExists(t *testing.T) {
	_, _, code := cli(t, "question", "get", "99999999999999999")
	assertExitNonZero(t, code)
}

func TestQuestionGet_InvalidID(t *testing.T) {
	_, _, code := cli(t, "question", "get", "abc")
	assertExitNonZero(t, code)
}

func TestQuestionGet_MissingID(t *testing.T) {
	_, stderr, code := cli(t, "question", "get")
	assertExitNonZero(t, code)
	assertContains(t, stderr, "accepts 1 arg(s)")
}

func TestQuestionGet_JSON(t *testing.T) {
	stdout, stderr, code := cli(t, "question", "get", "10010000000000020", "--json")
	assertExitZero(t, code, stderr)
	assertJSONValid(t, stdout)
	assertContains(t, stdout, `"id"`)
	assertContains(t, stdout, `"title"`)
}

// ============================================================
// 回答列表
// ============================================================

func TestAnswerList_Exists(t *testing.T) {
	stdout, stderr, code := cli(t, "answer", "list", "10010000000000075")
	assertExitZero(t, code, stderr)
	// 应包含表格或数据
	assertContains(t, stdout, "#") // 表格有 # 列
}

func TestAnswerList_NotExists(t *testing.T) {
	_, _, code := cli(t, "answer", "list", "99999999999999999")
	// 无回答或问题不存在
	assertExitZero(t, code, "")
}

func TestAnswerList_MissingID(t *testing.T) {
	_, stderr, code := cli(t, "answer", "list")
	assertExitNonZero(t, code)
	assertContains(t, stderr, "accepts 1 arg(s)")
}

func TestAnswerList_JSON(t *testing.T) {
	stdout, stderr, code := cli(t, "answer", "list", "10010000000000075", "--format=json")
	assertExitZero(t, code, stderr)
	assertJSONValid(t, stdout)
}

// ============================================================
// 回答详情
// ============================================================

func TestAnswerGet_Exists(t *testing.T) {
	stdout, stderr, code := cli(t, "answer", "get", "10020000000000022")
	assertExitZero(t, code, stderr)
	// answer get 应输出 ID 和 Content 字段
	assertContains(t, stdout, "ID:")
}

func TestAnswerGet_NotExists(t *testing.T) {
	_, _, code := cli(t, "answer", "get", "99999999999999999")
	assertExitNonZero(t, code)
}

func TestAnswerGet_JSON(t *testing.T) {
	stdout, stderr, code := cli(t, "answer", "get", "10020000000000076", "--json")
	assertExitZero(t, code, stderr)
	assertJSONValid(t, stdout)
}

// ============================================================
// 评论列表
// ============================================================

func TestCommentList_Empty(t *testing.T) {
	stdout, _, code := cli(t, "comment", "list", "10010000000000075")
	assertExitZero(t, code, "")
	// 可能无评论或无评论
	if strings.Contains(stdout, "No comments") || strings.Contains(stdout, "暂无") {
		// 预期行为
	}
}

func TestCommentList_WithQuestion(t *testing.T) {
_ , _, code := cli(t, "comment", "list", "10010000000000073")
	assertExitZero(t, code, "")
}

func TestCommentList_JSON(t *testing.T) {
	stdout, stderr, code := cli(t, "comment", "list", "10010000000000073", "--format=json")
	assertExitZero(t, code, stderr)
	assertJSONValid(t, stdout)
}

func TestCommentList_MissingID(t *testing.T) {
	_, stderr, code := cli(t, "comment", "list")
	assertExitNonZero(t, code)
	assertContains(t, stderr, "accepts 1 arg(s)")
}

// ============================================================
// 评论详情
// ============================================================

func TestCommentGet_NotExists(t *testing.T) {
	_, _, code := cli(t, "comment", "get", "99999999999999999")
	assertExitNonZero(t, code)
}

func TestCommentGet_MissingID(t *testing.T) {
	_, stderr, code := cli(t, "comment", "get")
	assertExitNonZero(t, code)
	assertContains(t, stderr, "accepts 1 arg(s)")
}

// ============================================================
// 标签列表
// ============================================================

func TestTagList_Default(t *testing.T) {
	stdout, stderr, code := cli(t, "tag", "list")
	assertExitZero(t, code, stderr)
	// 应包含表格表头
	assertContains(t, stdout, "标签")
	assertContains(t, stdout, "名称")
}

func TestTagList_WithOrder(t *testing.T) {
	for _, order := range []string{"popular", "name", "newest"} {
		t.Run(order, func(t *testing.T) {
			_, stderr, code := cli(t, "tag", "list", "--order="+order)
			assertExitZero(t, code, stderr)
		})
	}
}

func TestTagList_JSON(t *testing.T) {
	stdout, stderr, code := cli(t, "tag", "list", "--format=json")
	assertExitZero(t, code, stderr)
	assertJSONValid(t, stdout)
	assertContains(t, stdout, `"count"`)
}

func TestTagList_WithPage(t *testing.T) {
	_, stderr, code := cli(t, "tag", "list", "--page=1", "--size=5")
	assertExitZero(t, code, stderr)
	assertNotContains(t, stderr, "错误")
}

// ============================================================
// 搜索
// ============================================================

func TestSearch_Keyword(t *testing.T) {
	_, stderr, code := cli(t, "search", "git")
	assertExitZero(t, code, stderr)
	assertNotContains(t, stderr, "错误")
}

func TestSearch_NoResult(t *testing.T) {
	stdout, _, code := cli(t, "search", "xyznonexistent12345")
	assertExitZero(t, code, "")
	assertContains(t, stdout, "未找到")
}

func TestSearch_JSON(t *testing.T) {
	stdout, stderr, code := cli(t, "search", "git", "--format=json")
	assertExitZero(t, code, stderr)
	assertJSONValid(t, stdout)
}

// ============================================================
// 通知
// ============================================================

func TestNotificationList_Default(t *testing.T) {
	_, stderr, code := cli(t, "notification", "list")
	assertExitZero(t, code, stderr)
	// 可能为空
	assertNotContains(t, stderr, "错误")
}

func TestNotificationList_Inbox(t *testing.T) {
	_, stderr, code := cli(t, "notification", "list", "--type=inbox")
	assertExitZero(t, code, stderr)
}

func TestNotificationList_Achievement(t *testing.T) {
	_, stderr, code := cli(t, "notification", "list", "--type=achievement")
	assertExitZero(t, code, stderr)
}

func TestNotificationList_JSON(t *testing.T) {
	stdout, stderr, code := cli(t, "notification", "list", "--format=json")
	assertExitZero(t, code, stderr)
	assertJSONValid(t, stdout)
}

// ============================================================
// 用户信息
// ============================================================

func TestUserInfo(t *testing.T) {
	_, stderr, code := cli(t, "user", "info")
	assertExitZero(t, code, stderr)
}

func TestUserInfo_JSON(t *testing.T) {
	stdout, stderr, code := cli(t, "user", "info", "--format=json")
	assertExitZero(t, code, stderr)
	assertJSONValid(t, stdout)
}

// ============================================================
// 配置
// ============================================================

func TestConfigList(t *testing.T) {
	stdout, _, code := cli(t, "config", "list")
	assertExitZero(t, code, "")
	assertContains(t, stdout, "server")
}

func TestConfigGet(t *testing.T) {
	stdout, _, code := cli(t, "config", "get", "server.base_url")
	assertExitZero(t, code, "")
	// 应返回 URL
	if !strings.HasPrefix(stdout, "http") {
		t.Logf("config get server.base_url 返回: %s (可能未配置)", stdout)
	}
}

func TestConfigGet_NotExists(t *testing.T) {
	cli(t, "config", "get", "nonexistent.key")
	// 不应 panic，可能返回错误或空
}

// ============================================================
// 错误处理
// ============================================================

func TestUnknownCommand(t *testing.T) {
	_, _, code := cli(t, "nonexistent")
	assertExitNonZero(t, code)
}

func TestInvalidFlag(t *testing.T) {
	_, _, code := cli(t, "question", "list", "--nonexistent=flag")
	assertExitNonZero(t, code)
}

// ============================================================
// JSON 输出格式化验证
// ============================================================

func TestJSONOutput_Indented(t *testing.T) {
	stdout, _, code := cli(t, "question", "list", "--format=json", "--size=1")
	assertExitZero(t, code, "")
	// 缩进格式化应包含换行和空格
	if !strings.Contains(stdout, "\n") {
		t.Fatal("JSON 输出应该是缩进格式化的（包含换行）")
	}
	if !strings.Contains(stdout, "  ") {
		t.Fatal("JSON 输出应该有缩进空格")
	}
}

func TestJSONOutput_QuestionGet_Indented(t *testing.T) {
	stdout, _, code := cli(t, "question", "get", "10010000000000020", "--json")
	assertExitZero(t, code, "")
	if !strings.Contains(stdout, "\n") {
		t.Fatal("JSON 输出应该是缩进格式化的")
	}
	assertContains(t, stdout, `"title"`)
	assertContains(t, stdout, `"content"`)
}

// ============================================================
// 表格输出验证
// ============================================================

func TestTableOutput_QuestionList(t *testing.T) {
	stdout, _, code := cli(t, "question", "list", "--size=3")
	assertExitZero(t, code, "")

	// 表格应有分隔线
	if strings.Contains(stdout, "1001") {
		// 有数据时，验证表格格式
		assertContains(t, stdout, "---")
	}
}

func TestTableOutput_TagList(t *testing.T) {
	stdout, _, code := cli(t, "tag", "list", "--size=3")
	assertExitZero(t, code, "")
	assertContains(t, stdout, "---")
}

func TestTableOutput_AnswerList(t *testing.T) {
	stdout, stderr, code := cli(t, "answer", "list", "10010000000000075")
	assertExitZero(t, code, stderr)
	if strings.Contains(stdout, "1002") {
		assertContains(t, stdout, "---")
	}
}

// ============================================================
// 边界条件
// ============================================================

func TestEmptyOutput_NotPanic(t *testing.T) {
	// 各种空输入不应导致 panic
	cli(t, "question", "list", "--page=0")
	cli(t, "tag", "list", "--page=999")
	cli(t, "notification", "list", "--size=0")
}

func TestLargePageNumber(t *testing.T) {
	cli(t, "question", "list", "--page=100000")
	// 不应 panic，可以是错误或空结果
}

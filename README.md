# answer-cli

> Apache Answer Q&A 社区的命令行工具 — 浏览问答、搜索内容、AI 对话、MCP 集成，全部在终端完成。



## 功能特性

- **TUI 浏览模式** — 无参数启动进入 Bubbletea 终端 UI，搜索、浏览问题详情与回答，支持 Markdown 渲染
- **完整 CLI 命令** — 问答、评论、标签、通知、用户、投票的 CRUD 操作，适合脚本和自动化
- **AI 对话** — 基于 API 搜索的 tool calling，LLM 自动检索社区内容并生成回答
- **MCP Server** — stdio 模式，让 Claude Desktop、Cherry Studio 等 AI 客户端直接操作社区数据
- **多格式输出** — 支持 pretty / JSON / JSONL 格式，`--format` 全局切换
- **多 AI 提供商** — OpenAI、Ollama、智谱等 OpenAI 兼容 API

## 安装

### 从源码构建

```bash
git clone https://github.com/cicbyte/answer-cli.git
cd answer-cli
go build -o answer-cli .
```

交叉编译（需 Python 3，可选 UPX）：

```bash
python scripts/build.py --local    # 当前平台
python scripts/build.py             # 全平台（Windows/Linux/macOS）
```

### 环境要求

- Go >= 1.24
- Apache Answer 实例（自建或 [answer.apache.org](https://answer.apache.org)）

## 快速开始

```bash
# 配置服务器地址
answer-cli config set server.base_url https://your-answer-site.com

# 登录
answer-cli auth login

# 列出问题（TUI 浏览模式）
answer-cli

# 搜索
answer-cli search "golang 并发"

# AI 对话
answer-cli chat "如何处理 goroutine 泄露？"
```

## 使用方法

### TUI 浏览模式

无参数启动进入 TUI 界面：

```bash
answer-cli
```

| 按键 | 功能 |
|------|------|
| `↑↓` / `j/k` | 移动光标 / 滚动 |
| `Enter` | 进入详情 |
| `/` | 搜索 |
| `Tab` | 切换排序（最新/活跃/热门/评分） |
| `n` / `p` | 翻页 |
| `v` | 投票 |
| `r` | 写回答 |
| `c` | 写评论 |
| `g` / `G` | 跳顶部 / 底部 |
| `Esc` | 返回 |
| `q` | 退出 |

### 问题

```bash
answer-cli question list              # 列出问题
answer-cli question list --order hot  # 按热门排序
answer-cli question get <id>          # 查看详情（含回答、Markdown 渲染）
answer-cli question create --title="..." --content="..."
answer-cli question update <id> --title="..."
answer-cli question delete <id>
```

### 搜索

```bash
answer-cli search "关键词"
```

### 评论 / 标签 / 通知 / 用户 / 投票

```bash
answer-cli comment list --object-id <id> --object-type question
answer-cli comment add --object-id <id> --object-type question --content="..."
answer-cli tag list
answer-cli tag search "golang"
answer-cli notification list
answer-cli user get <username>
answer-cli vote up --object-id <id> --object-type question
```

### AI 对话

```bash
answer-cli chat "问题"                   # 单轮对话
answer-cli chat -i                       # 多轮交互对话
```

AI Agent 通过 5 个 function tools 检索社区内容（搜索问题、获取详情、列出回答、搜索标签、搜索用户），自动选择检索策略后生成回答。

### 全部命令

```bash
answer-cli --help
```

| 命令 | 说明 |
|------|------|
| `auth login` / `logout` / `status` | 认证管理 |
| `config list` / `get` / `set` | 配置管理 |
| `question list` / `get` / `create` / `update` / `delete` | 问题 CRUD |
| `answer list` / `get` / `create` / `update` / `delete` | 回答 CRUD |
| `comment list` / `add` / `update` / `delete` | 评论 CRUD |
| `tag list` / `search` / `get` / `info` | 标签查询 |
| `search <query>` | 全文搜索 |
| `notification list` | 通知列表 |
| `user get` / `search` | 用户查询 |
| `vote up` / `down` | 投票 |
| `chat [question]` | AI 对话 |
| `mcp` | 启动 MCP Server |

## 配置

配置文件路径：`~/.cicbyte/answer-cli/config/config.yaml`（首次运行自动创建）。

```yaml
server:
  base_url: https://your-answer-site.com
token: your-access-token

ai:
  provider: openai          # openai / ollama / zhipu
  base_url: https://api.openai.com/v1
  api_key: sk-xxx
  model: gpt-4o
```

也可通过命令设置：

```bash
answer-cli config set server.base_url https://your-answer-site.com
answer-cli config list
```

## MCP Server

`answer-cli mcp` 以 stdio 模式运行 MCP Server，注册 6 个工具：`question_search`、`question_get`、`answer_list`、`tag_search`、`tag_get`、`user_search`。

**Claude Desktop 配置：**

```json
{
  "mcpServers": {
    "answer": {
      "command": "answer-cli",
      "args": ["mcp"]
    }
  }
}
```

**Cherry Studio 配置：**

设置 → 模型服务 → MCP 服务器 → 添加：
- 名称：`answer`
- 命令：`answer-cli`
- 参数：`mcp`

## 技术栈

- Go 1.24
- [Bubbletea](https://github.com/charmbracelet/bubbletea) + [Bubbles](https://github.com/charmbracelet/bubbles) + [Lipgloss](https://github.com/charmbracelet/lipgloss) + [Glamour](https://github.com/charmbracelet/glamour) — TUI 框架
- [Cobra](https://github.com/spf13/cobra) — CLI 框架
- [mcp-go](https://github.com/mark3labs/mcp-go) — MCP Server
- [go-openai](https://github.com/sashabaranov/go-openai) — OpenAI 兼容 API（function calling）
- [Resty](https://github.com/go-resty/resty) — HTTP 客户端
- [go-pretty](https://github.com/jedib0t/go-pretty) — 终端表格

## 许可证

[MIT](LICENSE) © 2025 cicbyte

# answer-cli

> A CLI tool for [Apache Answer](https://github.com/apache/answer) Q&A communities — TUI browsing, CLI operations, AI chat, and MCP integration, all in your terminal.

简体中文 | [中文](README.md)

![Release](https://img.shields.io/github/v/release/cicbyte/answer-cli?style=flat)
![Go Report Card](https://goreportcard.com/badge/github.com/cicbyte/answer-cli)
![License](https://img.shields.io/github/license/cicbyte/answer-cli)
![Last Commit](https://img.shields.io/github/last-commit/cicbyte/answer-cli)

## Features

- **TUI Browser** — Launch a Bubbletea terminal UI with no arguments to search, browse questions and answers, with Markdown rendering
- **CLI Commands** — Query and write operations for questions, answers, comments, tags, and notifications
- **AI Chat** — LLM-powered tool calling that automatically searches community content and generates answers
- **MCP Server** — stdio mode, letting Claude Desktop, Cherry Studio and other AI clients interact with your community data
- **Multiple Output Formats** — pretty / JSON / JSONL, switch globally with `--format`
- **Multi-Provider AI** — OpenAI, Ollama, Zhipu and any OpenAI-compatible API

## Installation

```bash
git clone https://github.com/cicbyte/answer-cli.git
cd answer-cli
go build -o answer-cli .
```

Cross-compile (requires Python 3, optional UPX):

```bash
python scripts/build.py --local    # Current platform
python scripts/build.py             # All platforms (Windows/Linux/macOS)
```

**Requirements:** Go >= 1.24, Apache Answer instance

## Quick Start

```bash
answer-cli config set server.base_url https://your-answer-site.com
answer-cli auth login
answer-cli                              # TUI browser mode
answer-cli search "goroutine leak"      # Search
answer-cli chat "How to handle goroutine leaks?"  # AI chat
```

## Commands

| Command | Description |
|---------|-------------|
| `auth login` / `logout` / `status` | Authentication |
| `config list` / `get` / `set` | Configuration |
| `question list` / `get` / `create` / `update` / `delete` / `close` / `reopen` | Questions |
| `answer list` / `get` / `create` / `update` / `delete` | Answers |
| `comment list` / `get` / `add` / `update` / `delete` | Comments |
| `tag list` / `get` / `create` / `update` / `delete` | Tags |
| `notification list` | Notifications |
| `search <query>` | Full-text search |
| `chat [question]` | AI chat |
| `mcp` | Start MCP Server |

### Questions

```bash
answer-cli question list                          # List questions
answer-cli question list --order hot              # Sort by hot
answer-cli question list --tag go                  # Filter by tag
answer-cli question get <id>                      # View detail (with answers, Markdown)
answer-cli question create -t "Title" -c "Content" --tags=go
answer-cli question delete <id> --yes
```

### Answers / Comments / Tags

```bash
answer-cli answer list <question-id>               # List answers
answer-cli answer get <id>                        # View answer detail
answer-cli comment list <object-id>               # List comments
answer-cli comment add --object-id <id> -c "Comment"
answer-cli tag list                               # List tags
```

### AI Chat

```bash
answer-cli chat "question"                        # Single-turn
answer-cli chat -i                                # Interactive multi-turn
```

The AI Agent uses 5 function tools to search community content (search questions, get details, list answers, search tags, search users), automatically selecting the best retrieval strategy before generating an answer.

### Global Options

```bash
answer-cli question list --format json           # Indented JSON output
answer-cli question list --format jsonl          # JSONL line-delimited output
```

## TUI Browser Mode

Launch the Bubbletea terminal UI with no arguments:

```bash
answer-cli
```

| Key | Action |
|-----|--------|
| `↑↓` / `j/k` | Move cursor / Scroll |
| `Enter` | Open detail |
| `/` | Search |
| `Tab` | Cycle sort order (newest/active/hot/score) |
| `n` / `p` | Next / Previous page |
| `Home` / `End` | Jump to top / bottom |
| `Esc` | Go back |
| `q` | Quit |

## Configuration

Config file: `~/.cicbyte/answer-cli/config/config.yaml` (auto-created on first run)

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

```bash
answer-cli config set server.base_url https://your-answer-site.com
answer-cli config list
```

## MCP Server

`answer-cli mcp` runs an MCP Server in stdio mode, registering 6 tools: `question_search`, `question_get`, `answer_list`, `tag_search`, `tag_get`, `user_search`.

**Claude Desktop:**

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

**Cherry Studio:** Settings → Model Services → MCP Servers → Add, command `answer-cli`, args `mcp`

## Tech Stack

- Go 1.24
- [Bubbletea](https://github.com/charmbracelet/bubbletea) + [Bubbles](https://github.com/charmbracelet/bubbles) + [Lipgloss](https://github.com/charmbracelet/lipgloss) + [Glamour](https://github.com/charmbracelet/glamour) — TUI
- [Cobra](https://github.com/spf13/cobra) — CLI framework
- [mcp-go](https://github.com/mark3labs/mcp-go) — MCP Server
- [go-openai](https://github.com/sashabaranov/go-openai) — OpenAI-compatible API
- [Resty](https://github.com/go-resty/resty) — HTTP client
- [go-pretty](https://github.com/jedib0t/go-pretty) — Terminal tables

## License

[MIT](LICENSE) © 2025 cicbyte

---
name: answer-cli
description: answer-cli 使用指南。Apache Answer Q&A 社区 CLI 工具，支持问题/回答/评论/标签管理、AI 对话、MCP Server、TUI 浏览。在以下场景触发：(1) 用户想通过 CLI 操作 Answer 社区（搜索问题、查看详情、创建内容），(2) 用户想用 AI 对话检索社区内容，(3) 用户想配置 MCP Server 集成到 AI 客户端，(4) 用户想管理应用配置
---

# answer-cli 使用指南

answer-cli 是 Apache Answer Q&A 社区的命令行工具。

## AI 使用约束

- **禁止启动交互式模式**：`answer-cli`（无参数）会启动 Bubbletea TUI，这是给人类使用的，AI 绝不能执行此命令
- **禁止使用交互式输入命令**：不带参数的 `answer-cli auth login`、`answer-cli chat`（无参数）等会进入交互式等待输入，AI 不能执行。必须通过 flags 或位置参数提供所有参数
- **禁止修改配置**：`answer-cli config set`、`answer-cli auth login` 等配置操作必须由人类完成，AI 不能代为执行。如果前置条件未满足，应提示用户自行配置
- **AI 只能执行非交互式只读/写入命令**：如 `question list`、`question get <id>`、`answer list <qid>`、`search "关键词"`、`chat "问题"`

## 前置条件（需人类提前配置）

1. 配置服务器并登录：`answer-cli config set server.base_url <url>` + `answer-cli auth login`
2. AI 功能需要额外配置 LLM 服务（见 [配置参考](references/config.md)）

## 命令速查

### AI 可用

| 命令 | 说明 |
|------|------|
| `answer-cli question list` | 列出问题 |
| `answer-cli question get <id>` | 查看问题详情 |
| `answer-cli answer list <qid>` | 列出问题的回答 |
| `answer-cli answer get <id>` | 查看回答详情 |
| `answer-cli comment list <oid>` | 列出评论 |
| `answer-cli tag list` | 列出标签 |
| `answer-cli search "关键词"` | 搜索 |
| `answer-cli notification list` | 通知列表 |
| `answer-cli chat "问题"` | AI 对话（单轮） |
| `answer-cli question create -t "标题" -c "内容" --tags=go` | 创建问题 |
| `answer-cli answer create <qid> -c "回答内容"` | 创建回答 |
| `answer-cli comment add <oid> -c "评论"` | 添加评论 |
| `answer-cli auth status` | 查看认证状态 |
| `answer-cli config list` | 查看配置 |
| `answer-cli config get <key>` | 查看配置项 |

### AI 禁止使用（需人类操作）

| 命令 | 原因 |
|------|------|
| `answer-cli` | 启动交互式 TUI |
| `answer-cli chat`（无参数） | 多轮交互式对话 |
| `answer-cli auth login`（无 flags） | 交互式登录 |
| `answer-cli config set` | 修改配置 |

## 典型工作流

### 搜索与浏览

```bash
answer-cli question list                          # 列出问题
answer-cli question list --order hot              # 按热门排序
answer-cli question list --tag go                  # 按标签过滤
answer-cli question get <id>                       # 查看详情（含回答）
answer-cli search "关键词"                         # 全文搜索
```

### AI 对话

```bash
answer-cli chat "如何处理 goroutine 泄露？"        # 单轮对话
```

### 写操作

```bash
answer-cli question create -t "标题" -c "内容" --tags=go
answer-cli answer create <qid> -c "回答内容"
answer-cli comment add <oid> -c "评论"
```

## 模块详细文档

- [AI 对话](references/chat.md) — AI Agent、工具调用、多轮对话
- [配置参考](references/config.md) — 服务器、AI、日志配置
- [MCP Server](references/mcp.md) — 工具列表、客户端配置

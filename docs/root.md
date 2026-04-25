# answer-cli

Apache Answer Q&A 社区的命令行工具。

## 用法

```bash
answer-cli [flags]
answer-cli [command]
```

## 描述

直接运行 `answer-cli` 不带任何参数时，启动基于 Bubbletea 的交互式 TUI 浏览界面。也可以使用子命令进行直接操作。

## 子命令

| 命令 | 说明 |
|------|------|
| [auth](./auth.md) | 认证管理（登录/登出/状态） |
| [config](./config.md) | 管理应用配置（服务器/AI/日志） |
| [question](./question.md) | 问题管理（列表/详情/创建/更新/删除/关闭/重开） |
| [answer](./answer.md) | 回答管理（列表/详情/创建/更新/删除/采纳） |
| [comment](./comment.md) | 评论管理（列表/详情/添加/更新/删除） |
| [tag](./tag.md) | 标签管理（列表/详情/创建/更新/删除） |
| [notification](./notification.md) | 通知管理（列表/已读/状态） |
| [search](./search.md) | 全文搜索（等同于 question list 的搜索模式） |
| [chat](./chat.md) | AI 对话（基于社区数据回答问题） |
| [mcp](./mcp.md) | 启动 MCP Server |

## 全局选项

| 标志 | 说明 |
|------|------|
| `--format` | 输出格式：`pretty`（默认）/ `json` / `jsonl` |

## 示例

```bash
# 启动 TUI 浏览界面
answer-cli

# 列出问题
answer-cli question list

# 查看问题详情
answer-cli question get 10010000000000020

# 搜索
answer-cli search "golang 并发"

# AI 对话
answer-cli chat "如何处理 goroutine 泄露？"

# 多轮对话
answer-cli chat

# 启动 MCP Server
answer-cli mcp
```

## 配置目录

所有数据存储在 `~/.cicbyte/answer-cli/` 下：

```
~/.cicbyte/answer-cli/
├── config/
│   └── config.yaml    # 应用配置
└── logs/
    └── app.log        # 日志文件
```

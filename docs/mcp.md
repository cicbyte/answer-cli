# mcp

MCP (Model Context Protocol) Server，以 stdio 模式运行。

## 用法

```bash
answer-cli mcp
```

## 描述

启动 stdio 模式的 MCP Server，让 AI 客户端（Claude Desktop、Cherry Studio 等）能够搜索和操作 Answer 社区数据。

## 注册的 Tools

| Tool | 描述 | 参数 |
|------|------|------|
| `question_search` | 搜索问题 | `keyword`、`tag`、`order`、`limit` |
| `question_get` | 获取问题详情 | `question_id` |
| `question_create` | 创建问题 | `title`、`content`、`tags` |
| `question_update` | 更新问题 | `question_id`、`title`、`content`、`tags` |
| `answer_list` | 列出问题的回答 | `question_id`、`limit` |
| `answer_create` | 创建回答 | `question_id`、`content` |
| `answer_update` | 更新回答 | `answer_id`、`content` |
| `comment_add` | 添加评论 | `object_id`、`content` |
| `tag_search` | 搜索标签 | `query` |

## 客户端配置

### Claude Desktop

在 `claude_desktop_config.json` 中添加：

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

### Cherry Studio

设置 → 模型服务 → MCP 服务器 → 添加：
- 命令：`answer-cli`
- 参数：`mcp`

## 前置条件

- 需要已配置 Answer 服务器并完成认证

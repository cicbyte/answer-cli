# MCP Server

以 stdio 模式运行，让 AI 客户端（Claude Desktop、Cherry Studio 等）搜索和操作 Answer 社区数据。

## 启动

```bash
answer-cli mcp
```

## 注册的 Tools

| Tool | 描述 | 参数 |
|------|------|------|
| `question_search` | 搜索问题 | `keyword`、`tag`、`order`、`limit` |
| `question_get` | 获取问题详情 | `question_id` |
| `answer_list` | 列出问题的回答 | `question_id`、`limit` |
| `tag_search` | 搜索标签 | `query` |
| `tag_get` | 获取标签详情 | `slug_name` |
| `user_search` | 搜索用户 | `username` |

## 客户端配置

### Claude Desktop

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

- 已配置 Answer 服务器并完成认证

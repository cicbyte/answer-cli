# AI 对话

基于 AI Agent 模式与 Answer 社区对话，LLM 自动搜索相关问题并生成回答。

## 用法

```bash
answer-cli chat [问题] [flags]
```

## 选项

| 标志 | 说明 |
|------|------|
| `--non-stream` | 使用非流式输出 |

## AI Agent 工具

Agent 通过 function calling 自动检索社区内容：

1. **question_search** — 按关键词搜索问题（自动获取 top 结果详情和回答）
2. **question_get** — 获取问题完整详情
3. **answer_list** — 获取问题的所有回答
4. **tag_search** — 搜索标签
5. **user_search** — 搜索用户

搜索后会自动预取最相关问题的详情和回答，一次调用即可获得完整上下文。

## 多轮对话

无参数进入交互模式：

- 输入 `/quit`、`/exit` 或 `/q` 退出
- 输入 `/clear` 清除对话上下文

## 前置条件

- 需要 AI 服务配置（`answer-cli config set ai.provider/model/base_url/api_key`）

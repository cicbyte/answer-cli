# chat

基于 AI Agent 模式与 Answer 社区对话。LLM 自动搜索相关问题、获取详情和回答内容后生成回答。

## 用法

```bash
answer-cli chat [question] [flags]
```

## 描述

chat 命令内置 RAG Agent，通过 function calling 让 LLM 自主检索社区内容：

1. **question_search** — 按关键词搜索问题（自动获取 top 结果的详情和回答）
2. **question_get** — 获取问题完整详情
3. **answer_list** — 获取问题的所有回答
4. **tag_search** — 搜索标签
5. **user_search** — 搜索用户

Agent 会在搜索后自动获取最相关问题的详情和回答，然后基于完整内容生成回答。

### 交互模式

无参数进入多轮对话模式：

```
  AI 对话模式
  输入问题开始对话，/quit 退出，/clear 清除上下文

  user > 如何开启远程 SSH？
  ▸ question_search
  ✓ 搜索并获取详情

  （Markdown 渲染的回答内容）

  Tokens: 1234 + 567 · 3.2s

  user > /quit
  再见!
```

## 选项

| 标志 | 说明 |
|------|------|
| `--non-stream` | 使用非流式输出（等待完整回答后一次性显示） |

## 多轮对话

无参数进入交互模式，支持上下文连续提问：

- `/quit`、`/exit`、`/q` 退出
- `/clear` 清除对话上下文

## 示例

```bash
# 单轮对话
answer-cli chat "如何处理 goroutine 泄露？"

# 多轮对话
answer-cli chat
```

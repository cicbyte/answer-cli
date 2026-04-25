# question

管理 Apache Answer 社区中的问题。

## 用法

```bash
answer-cli question [command]
```

## 子命令

| 命令 | 别名 | 说明 |
|------|------|------|
| [list](#question-list) | `ls` | 列出问题，支持排序、标签过滤和关键词搜索 |
| [get](#question-get) | — | 查看问题详情（含 Markdown 渲染的回答） |
| [create](#question-create) | — | 创建新问题 |
| [update](#question-update) | — | 更新问题 |
| [delete](#question-delete) | — | 删除问题 |
| [close](#question-close) | — | 关闭问题 |
| [reopen](#question-reopen) | — | 重新打开问题 |

---

## question list

列出问题，支持排序、标签过滤和关键词搜索。

```bash
answer-cli question list [query] [flags]
```

### 选项

| 标志 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--order` | `-o` | `newest` | 排序：`newest` / `active` / `hot` / `score` / `unanswered` |
| `--tag` | `-t` | — | 按标签 slug 过滤 |
| `--page` | `-p` | 1 | 页码 |
| `--size` | `-s` | 20 | 每页数量 |

### 示例

```bash
answer-cli question list
answer-cli question list --order hot
answer-cli question list --tag go
answer-cli question list "git 使用"
```

---

## question get

查看问题详情，包含 Markdown 渲染的问题内容和回答列表。

```bash
answer-cli question get <id> [flags]
```

### 选项

| 标志 | 说明 |
|------|------|
| `--json` | 以 JSON 格式输出 |

### 示例

```bash
answer-cli question get 10010000000000020
answer-cli question get 10010000000000020 --json
```

---

## question create

创建新问题。

```bash
answer-cli question create [flags]
```

### 选项

| 标志 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--title` | `-t` | — | 问题标题（必填） |
| `--content` | `-c` | — | 问题内容（必填） |
| `--tags` | — | — | 标签 slug，逗号分隔 |

### 示例

```bash
answer-cli question create -t "如何配置 Go 模块代理？" -c "在国内使用 Go 时..." --tags=go
```

---

## question update

更新问题标题和内容。

```bash
answer-cli question update <id> [flags]
```

### 选项

| 标志 | 简写 | 说明 |
|------|------|------|
| `--title` | `-t` | 新标题 |
| `--content` | `-c` | 新内容 |
| `--tags` | — | 标签 slug，逗号分隔 |

---

## question delete

删除问题。

```bash
answer-cli question delete <id> --yes
```

需要 `--yes` 确认。

---

## question close

关闭问题。

```bash
answer-cli question close <id>
```

---

## question reopen

重新打开已关闭的问题。

```bash
answer-cli question reopen <id>
```

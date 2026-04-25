# answer

管理 Apache Answer 社区中的回答。

## 用法

```bash
answer-cli answer [command]
```

## 子命令

| 命令 | 说明 |
|------|------|
| [list](#answer-list) | 列出问题的回答 |
| [get](#answer-get) | 查看回答详情 |
| [create](#answer-create) | 创建回答 |
| [update](#answer-update) | 更新回答 |
| [delete](#answer-delete) | 删除回答 |
| [accept](#answer-accept) | 采纳回答 |

---

## answer list

列出指定问题的所有回答。

```bash
answer-cli answer list <question-id> [flags]
```

### 选项

| 标志 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--page` | `-p` | 1 | 页码 |
| `--size` | `-s` | 20 | 每页数量 |

### 示例

```bash
answer-cli answer list 10010000000000020
```

---

## answer get

查看回答详情。

```bash
answer-cli answer get <id> [flags]
```

### 选项

| 标志 | 说明 |
|------|------|
| `--json` | 以 JSON 格式输出 |

---

## answer create

为指定问题创建回答。

```bash
answer-cli answer create <question-id> --content="..."
```

### 选项

| 标志 | 简写 | 说明 |
|------|------|------|
| `--content` | `-c` | 回答内容（必填） |

---

## answer update

更新回答内容。

```bash
answer-cli answer update <id> --content="..."
```

---

## answer delete

删除回答。

```bash
answer-cli answer delete <id> --yes
```

---

## answer accept

采纳回答。

```bash
answer-cli answer accept <answer-id> --question-id=<question-id>
```

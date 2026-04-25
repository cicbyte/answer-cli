# comment

管理问题和回答的评论。

## 用法

```bash
answer-cli comment [command]
```

## 子命令

| 命令 | 说明 |
|------|------|
| [list](#comment-list) | 列出对象的评论 |
| [get](#comment-get) | 查看评论详情 |
| [add](#comment-add) | 添加评论 |
| [update](#comment-update) | 更新评论 |
| [delete](#comment-delete) | 删除评论 |

---

## comment list

列出指定对象（问题或回答）的评论。

```bash
answer-cli comment list <object-id> [flags]
```

### 选项

| 标志 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--page` | `-p` | 1 | 页码 |
| `--size` | `-s` | 20 | 每页数量 |

---

## comment get

查看评论详情。

```bash
answer-cli comment get <id>
```

---

## comment add

为指定对象添加评论。

```bash
answer-cli comment add <object-id> --text="..."
```

### 选项

| 标志 | 说明 |
|------|------|
| `--text` | 评论内容（必填） |

---

## comment update

更新评论内容。

```bash
answer-cli comment update <id> --text="..."
```

---

## comment delete

删除评论。

```bash
answer-cli comment delete <id> --yes
```

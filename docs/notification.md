# notification

管理通知。

## 用法

```bash
answer-cli notification [command]
```

## 子命令

| 命令 | 说明 |
|------|------|
| [list](#notification-list) | 列出通知 |
| [status](#notification-status) | 查看未读通知状态 |
| [read](#notification-read) | 标记通知为已读 |

---

## notification list

列出通知。

```bash
answer-cli notification list [flags]
```

### 选项

| 标志 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--type` | `-t` | — | 通知类型：`inbox` / `achievement` |
| `--page` | `-p` | 1 | 页码 |
| `--size` | `-s` | 20 | 每页数量 |

---

## notification status

显示未读通知数量。

```bash
answer-cli notification status
```

---

## notification read

标记通知为已读。

```bash
answer-cli notification read <id>
answer-cli notification read --all
```

### 选项

| 标志 | 说明 |
|------|------|
| `--all` | 标记所有通知为已读 |

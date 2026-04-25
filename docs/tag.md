# tag

管理社区标签。

## 用法

```bash
answer-cli tag [command]
```

## 子命令

| 命令 | 说明 |
|------|------|
| [list](#tag-list) | 列出标签 |
| [get](#tag-get) | 查看标签详情 |
| [create](#tag-create) | 创建标签 |
| [update](#tag-update) | 更新标签 |
| [delete](#tag-delete) | 删除标签 |

---

## tag list

列出社区标签。

```bash
answer-cli tag list [flags]
```

### 选项

| 标志 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--order` | `-o` | `popular` | 排序：`popular` / `name` / `newest` |
| `--page` | `-p` | 1 | 页码 |
| `--size` | `-s` | 20 | 每页数量 |

### 示例

```bash
answer-cli tag list
answer-cli tag list --order name
```

---

## tag get

查看标签详情。

```bash
answer-cli tag get <slug-name>
```

---

## tag create

创建标签。

```bash
answer-cli tag create --slug=<slug> --name=<display-name>
```

---

## tag update

更新标签。

```bash
answer-cli tag update <slug-name> --name=<new-name>
```

---

## tag delete

删除标签。

```bash
answer-cli tag delete <slug-name> --yes
```

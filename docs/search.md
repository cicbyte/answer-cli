# search

全文搜索问题和回答。等同于 `question list` 的搜索模式。

## 用法

```bash
answer-cli search [query] [flags]
```

## 描述

按关键词搜索社区中的问题和回答，使用 Answer 服务器的全文搜索引擎。默认按相关性排序。

## 参数

| 参数 | 说明 |
|------|------|
| `query` | 搜索关键词 |

## 选项

| 标志 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--keyword` | `-k` | — | 搜索关键词（等同于位置参数） |
| `--order` | — | `newest` | 排序：`newest` / `active` / `hot` / `score` / `relevance` |
| `--tag` | `-t` | — | 按标签过滤 |
| `--page` | `-p` | 1 | 页码 |
| `--size` | `-s` | 20 | 每页数量 |

## 示例

```bash
answer-cli search "golang 并发"
answer-cli search --keyword="docker" --order=relevance
answer-cli search --tag=go
answer-cli search "ssh" --format json
```

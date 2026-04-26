---
name: answer-cli
description: answer-cli 使用指南。Apache Answer Q&A 社区 CLI 工具，支持问题/回答/评论/标签管理、搜索、AI 对话、通知管理。在以下场景触发：(1) 用户想通过 CLI 操作 Answer 社区（搜索、浏览、创建内容），(2) 用户想用 AI 对话检索社区内容
---

# answer-cli 使用指南

answer-cli 是 Apache Answer Q&A 社区的命令行工具。

## AI 使用约束

- **禁止启动交互式模式**：`answer-cli`（无参数）会启动 TUI，AI 绝不能执行
- **禁止使用交互式输入命令**：不带参数的 `answer-cli auth login`、`answer-cli chat` 等会进入交互式等待输入，AI 不能执行
- **禁止修改配置**：`answer-cli config set`、`answer-cli auth login` 等配置操作必须由人类完成
- **删除命令必须加 --yes**：`delete` 命令默认要求交互式确认，AI 必须传 `--yes` 跳过

## 前置条件（需人类提前配置）

1. 配置服务器并登录：`answer-cli config set server.base_url <url>` + `answer-cli auth login`
2. AI 功能需要额外配置 LLM 服务

## 全局选项

- `--format` — 输出格式：`table`（默认）、`json`、`jsonl`

## 搜索

```bash
answer-cli search "关键词"                        # 全文搜索
answer-cli search "关键词" --order relevance       # 按相关性排序
answer-cli search "关键词" --page=2 --size=10      # 翻页
answer-cli search "关键词" --format json            # JSON 输出
```

`search` 等同于 `question list "关键词"`。

## 问题 (question)

### 列表

```bash
answer-cli question list                           # 列出问题（按最新）
answer-cli question list --order hot               # 按热门排序
answer-cli question list --order active             # 按活跃排序
answer-cli question list --order score              # 按分数排序
answer-cli question list --order unanswered         # 未回答
answer-cli question list --tag go                   # 按标签过滤
answer-cli question list --page=2 --size=10        # 翻页
answer-cli question list "chrome"                   # 关键词搜索
```

参数：`--order`（newest|active|hot|score|unanswered|relevance）、`--tag`、`-k/--keyword`、`-p/--page`、`-s/--size`

### 详情

```bash
answer-cli question get <id>                       # 查看详情（含回答列表）
answer-cli question get <id> --json                # JSON 输出
```

### 创建

```bash
answer-cli question create -t "标题" -c "内容" --tags=go,web
answer-cli question create -t "标题" -f content.md --tags=go  # 从文件读取内容
echo "内容" | answer-cli question create -t "标题" --tags=go  # 管道输入
```

参数：`-t/--title`（必填）、`-c/--content`、`--tags`（必填，逗号分隔）、`-f/--file`

内容来源优先级：`--content` > `--file` > 管道输入

### 更新

```bash
answer-cli question update <id> -t "新标题"
answer-cli question update <id> -c "新内容"
answer-cli question update <id> --tags=go,web
answer-cli question update <id> -f content.md
```

### 删除

```bash
answer-cli question delete <id> --yes              # 必须加 --yes 跳过确认
```

### 关闭/重开

```bash
answer-cli question close <id>                     # 关闭问题
answer-cli question close <id> --type=2 -m "重复"  # 指定关闭类型和原因
answer-cli question reopen <id>                    # 重新打开
```

关闭类型：`1`=已解决（默认）、`2`=重复、`3`=非主题、`4`=其他

## 回答 (answer)

### 列表

```bash
answer-cli answer list <question-id>               # 列出回答
answer-cli answer list <question-id> --page=2      # 翻页
```

### 详情

```bash
answer-cli answer get <id>                         # 查看回答详情
answer-cli answer get <id> --json                  # JSON 输出
```

### 创建

```bash
answer-cli answer create <question-id> -c "回答内容"
answer-cli answer create <question-id> -f answer.md # 从文件读取
echo "内容" | answer-cli answer create <qid>       # 管道输入
```

### 更新

```bash
answer-cli answer update <id> -c "新内容"
answer-cli answer update <id> -f answer.md
echo "新内容" | answer-cli answer update <id>
```

### 采纳

```bash
answer-cli answer accept <answer-id>               # 自动检测关联问题
answer-cli answer accept <answer-id> -q <question-id>  # 指定问题 ID
```

### 删除

```bash
answer-cli answer delete <id> --yes
```

## 评论 (comment)

评论的对象可以是问题或回答，通过 `<object-id>` 指定。

### 列表

```bash
answer-cli comment list <object-id>                # 列出评论
answer-cli comment list <object-id> --page=2       # 翻页
```

### 详情

```bash
answer-cli comment get <id>                        # 查看评论详情
```

### 添加

```bash
answer-cli comment add <object-id> -t "评论内容"
answer-cli comment add <object-id> -t "回复" --reply-to=<comment-id>  # 回复评论
echo "评论" | answer-cli comment add <object-id>   # 管道输入
```

参数：`-t/--text`、`--reply-to`

### 更新

```bash
answer-cli comment update <id> -t "新内容"
echo "新内容" | answer-cli comment update <id>     # 管道输入
```

### 删除

```bash
answer-cli comment delete <id> --yes
```

## 标签 (tag)

### 列表

```bash
answer-cli tag list                                # 列出标签（按热门）
answer-cli tag list --order name                   # 按名称排序
answer-cli tag list --order newest                 # 按最新
```

参数：`--order`（popular|name|newest）

### 详情

```bash
answer-cli tag get <slug-name>                     # 查看标签详情
answer-cli tag get go --json                       # JSON 输出
```

### 创建

```bash
answer-cli tag create --slug=go --name="Go Language"
answer-cli tag create --slug=go --name="Go" -d "Go 编程语言"
```

参数：`--slug`（必填）、`--name`（必填）、`-d/--description`

### 更新

```bash
answer-cli tag update <slug-name> --name="新名称"
answer-cli tag update <slug-name> -d "新描述"
```

### 删除

```bash
answer-cli tag delete <slug-name> --yes
```

## 通知 (notification)

### 列表

```bash
answer-cli notification list                       # 收件箱通知
answer-cli notification list --type=achievement    # 成就通知
answer-cli notification list --page=2              # 翻页
```

参数：`--type`（inbox|achievement）

### 标记已读

```bash
answer-cli notification read --all                 # 全部标记已读
answer-cli notification read --all --type=inbox    # 指定类型全部已读
answer-cli notification read <id>                  # 单条标记已读
```

## AI 对话

```bash
answer-cli chat "如何处理 goroutine 泄露？"         # 单轮对话
```

AI Agent 会自动搜索社区内容，基于搜索结果生成回答。

# config

管理 answer-cli 应用配置（服务器、AI、日志等参数）。

## 用法

```bash
answer-cli config [command]
```

## 子命令

| 命令 | 说明 |
|------|------|
| [list](#config-list) | 列出所有配置项及当前值 |
| [get](#config-get) | 查看单个配置项的值 |
| [set](#config-set) | 设置配置项 |

---

## config list

以表格形式列出所有配置项和当前值。

```bash
answer-cli config list
```

---

## config get

查看指定配置项的值。

```bash
answer-cli config get <key>
```

### 示例

```bash
answer-cli config get server.base_url
answer-cli config get ai.model
```

---

## config set

设置指定配置项的值。

```bash
answer-cli config set <key> <value>
```

### 常用配置项

| 配置项 | 说明 | 示例 |
|--------|------|------|
| `server.base_url` | Answer 服务器地址 | `https://answer.example.com` |
| `server.token` | 认证 Token | （通过 auth login 设置） |
| `ai.provider` | LLM 提供商 | `openai` / `ollama` / `zhipu` |
| `ai.base_url` | LLM API 地址 | `https://api.openai.com/v1` |
| `ai.model` | LLM 模型名称 | `gpt-4o` |
| `ai.api_key` | LLM API 密钥 | `sk-xxx` |

### 示例

```bash
answer-cli config set server.base_url https://answer.example.com
answer-cli config set ai.provider openai
answer-cli config set ai.model gpt-4o
answer-cli config set ai.api_key sk-xxx
```

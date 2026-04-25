# 配置参考

配置文件路径：`~/.cicbyte/answer-cli/config/config.yaml`

## config 命令

```bash
answer-cli config list              # 列出所有配置
answer-cli config get <key>         # 查看配置值
answer-cli config set <key> <value> # 设置配置
```

## 配置项

### Server

| 配置项 | 说明 |
|--------|------|
| `server.base_url` | Answer 服务器地址 |
| `server.token` | 认证 Token（通过 auth login 设置） |

### AI（LLM 对话）

| 配置项 | 说明 |
|--------|------|
| `ai.provider` | 提供商：`openai`/`ollama`/`zhipu` |
| `ai.base_url` | API 地址 |
| `ai.model` | 模型名称 |
| `ai.api_key` | API 密钥 |

## 常用配置示例

```bash
# OpenAI
answer-cli config set ai.provider openai
answer-cli config set ai.base_url https://api.openai.com/v1
answer-cli config set ai.model gpt-4o
answer-cli config set ai.api_key sk-xxx

# Ollama 本地模型
answer-cli config set ai.provider ollama
answer-cli config set ai.base_url http://localhost:11434/v1
answer-cli config set ai.model llama3

# 智谱
answer-cli config set ai.provider zhipu
answer-cli config set ai.base_url https://open.bigmodel.cn/api/paas/v4
answer-cli config set ai.model glm-4
answer-cli config set ai.api_key your-key
```

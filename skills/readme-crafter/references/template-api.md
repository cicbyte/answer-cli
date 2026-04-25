# API 服务模板

适用于：REST API、GraphQL 服务、微服务、后端服务等。

```markdown
# {ServiceName}

> {一句话说明：什么 API/服务、提供什么能力、给谁用}

{徽章：构建状态、运行时间、版本号、许可证}

## ✨ 功能特性

- **{功能 1}** — {带来的好处}
- **{功能 2}** — {带来的好处}
- **{功能 3}** — {带来的好处}

## 🚀 快速开始

### 自托管

\`\`\`bash
git clone {仓库地址}
cd {服务名}
{安装命令}
{启动命令}
\`\`\`

### Docker

\`\`\`bash
docker run -p {端口}:{端口} {镜像}
\`\`\`

### 托管版 / 云服务

{如有托管版本，链接到注册页}

## 🔐 认证方式

{描述认证方式}

\`\`\`bash
# API Key
curl -H "Authorization: Bearer {token}" {基础URL}/{接口}

# OAuth2
\`\`\`

| 方式 | 说明 |
|---|---|
| API Key | {如何获取} |
| OAuth 2.0 | {支持的授权流程} |
| JWT | {详情} |

## 📖 API 参考

### `{METHOD} /{接口}`

{一句话描述}

**请求：**

\`\`\`bash
curl -X {METHOD} {基础URL}/{接口} \\
  -H "Authorization: Bearer {token}" \\
  -H "Content-Type: application/json" \\
  -d '{ "key": "value" }'
\`\`\`

**响应** `200 OK`：

\`\`\`json
{
  "key": "value"
}
\`\`\`

| 参数 | 类型 | 位置 | 说明 |
|---|---|---|---|
| `param` | `{type}` | {body/query/path} | {说明} |

### `{METHOD} /{接口2}`

{描述}

\`\`\`bash
curl -X {METHOD} {基础URL}/{接口2}
\`\`\`

**响应** `200 OK`：

\`\`\`json
{...}
\`\`\`

### 错误码

| 状态码 | 说明 |
|---|---|
| `400` | {说明} |
| `401` | {说明} |
| `404` | {说明} |
| `429` | {说明} |
| `500` | {说明} |

## 🔌 SDK / 客户端库

{如有客户端 SDK}

| 语言 | 安装方式 |
|---|---|
| JavaScript | `npm install {package}` |
| Python | `pip install {package}` |
| Go | `go get {module}` |

### JavaScript 示例

\`\`\`javascript
import { Client } from '{package}';

const client = new Client({ apiKey: '{token}' });
const result = await client.{方法}({ 参数 });
\`\`\`

## ⚙️ 配置

| 变量 | 说明 | 默认值 |
|---|---|---|
| `PORT` | {说明} | `3000` |
| `DATABASE_URL` | {说明} | — |

## 📊 速率限制

{如适用}

| 套餐 | 请求次数/分钟 |
|---|---|
| 免费版 | {n} |
| 专业版 | {n} |
| 企业版 | {n} |

## 🤝 参与贡献

参见 [贡献指南](CONTRIBUTING.md)。

## 📄 开源许可证

[MIT](LICENSE) © {年份} {作者}
```

### API 服务 README 写作要点

- 每个接口都提供 curl 示例 — 这是最通用的展示方式
- 同时展示请求体和响应体
- 认证章节要醒目 — 用户需要立即知道如何认证
- 速率限制很重要 — 放在用户能找到的位置
- SDK 示例降低接入门槛

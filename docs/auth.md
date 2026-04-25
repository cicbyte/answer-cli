# auth

认证管理，处理与 Answer 服务器的登录、登出和状态查询。

## 用法

```bash
answer-cli auth [command]
```

## 子命令

| 命令 | 说明 |
|------|------|
| [login](#auth-login) | 登录到 Answer 服务器 |
| [logout](#auth-logout) | 从服务器登出 |
| [status](#auth-status) | 查看当前认证状态 |

---

## auth login

登录到 Answer 服务器。支持 Token 认证。

```bash
answer-cli auth login [flags]
```

### 选项

| 标志 | 简写 | 默认值 | 说明 |
|------|------|--------|------|
| `--url` | `-u` | — | 服务器 URL，默认使用已配置的服务器 |
| `--token` | `-t` | — | 直接使用 Token 认证 |

### 示例

```bash
# 交互式登录
answer-cli auth login

# 使用 Token 登录
answer-cli auth login --token=your-access-token

# 指定服务器登录
answer-cli auth login --url=https://answer.example.com --token=your-token
```

---

## auth logout

从服务器登出，清除本地存储的认证 Token。

```bash
answer-cli auth logout
```

---

## auth status

显示当前认证状态，包括服务器信息和用户信息。

```bash
answer-cli auth status
```

### 输出示例

```
服务器  https://answer.example.com
状态    已认证
用户    zhyj
```

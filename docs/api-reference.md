# Apache Answer API 参考

> 基于 `swag init` 从源码自动生成，详见 `参考/answer/docs/swagger.json`

## 认证方式

- **ApiKeyAuth**：通过 HTTP Header `Authorization` 传递 Token
- 绝大部分写操作和用户相关操作需要认证

## API 基础路径

- 用户端：`/answer/api/v1/`
- 管理端：`/answer/admin/api/`
- 安装：`/installation/`

## 端点总览

共 **199 个端点**，按功能分为 27 个分组。

---

### 1. Question — 问题管理（17 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/question/page` | 分页获取问题列表 |
| GET | `/answer/api/v1/question/info` | 获取问题详情 |
| POST | `/answer/api/v1/question` | 添加问题 |
| PUT | `/answer/api/v1/question` | 更新问题 |
| DELETE | `/answer/api/v1/question` | 删除问题 |
| POST | `/answer/api/v1/question/answer` | 同时添加问题和回答 |
| POST | `/answer/api/v1/question/recover` | 恢复已删除问题 |
| PUT | `/answer/api/v1/question/reopen` | 重新打开问题 |
| PUT | `/answer/api/v1/question/status` | 关闭问题 |
| PUT | `/answer/api/v1/question/operation` | 问题操作 |
| GET | `/answer/api/v1/question/recommend/page` | 推荐问题列表 |
| GET | `/answer/api/v1/question/similar` | 标题模糊搜索相似问题 |
| GET | `/answer/api/v1/question/similar/tag` | 按标签搜索相似问题 |
| GET | `/answer/api/v1/question/link` | 获取问题链接 |
| GET | `/answer/api/v1/question/invite` | 获取邀请信息 |
| PUT | `/answer/api/v1/question/invite` | 邀请用户回答 |

### 2. Answer — 回答管理（7 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/answer/page` | 回答列表 |
| GET | `/answer/api/v1/answer/info` | 回答详情 |
| POST | `/answer/api/v1/answer` | 添加回答 |
| PUT | `/answer/api/v1/answer` | 更新回答 |
| DELETE | `/answer/api/v1/answer` | 删除回答 |
| POST | `/answer/api/v1/answer/acceptance` | 采纳回答 |
| POST | `/answer/api/v1/answer/recover` | 恢复已删除回答 |

### 3. Tag — 标签管理（12 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/tags` | 获取标签列表 |
| GET | `/answer/api/v1/tags/page` | 分页获取标签 |
| GET | `/answer/api/v1/tags/following` | 获取已关注标签 |
| GET | `/answer/api/v1/tag` | 获取单个标签 |
| POST | `/answer/api/v1/tag` | 添加标签 |
| PUT | `/answer/api/v1/tag` | 更新标签 |
| DELETE | `/answer/api/v1/tag` | 删除标签 |
| POST | `/answer/api/v1/tag/merge` | 合并标签 |
| POST | `/answer/api/v1/tag/recover` | 恢复已删除标签 |
| PUT | `/answer/api/v1/tag/synonym` | 更新标签同义词 |
| GET | `/answer/api/v1/tag/synonyms` | 获取标签同义词 |
| GET | `/answer/api/v1/question/tags` | 获取问题标签列表 |

### 4. User — 用户管理（21 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/answer/api/v1/user/register/email` | 邮箱注册 |
| POST | `/answer/api/v1/user/login/email` | 邮箱登录 |
| GET | `/answer/api/v1/user/logout` | 登出 |
| GET | `/answer/api/v1/user/info` | 获取当前用户信息 |
| PUT | `/answer/api/v1/user/info` | 更新用户信息 |
| PUT | `/answer/api/v1/user/interface` | 更新用户界面配置 |
| PUT | `/answer/api/v1/user/password` | 修改密码 |
| POST | `/answer/api/v1/user/password/reset` | 找回密码 |
| POST | `/answer/api/v1/user/password/replacement` | 重置密码 |
| PUT | `/answer/api/v1/user/email` | 邮箱变更验证 |
| POST | `/answer/api/v1/user/email/change/code` | 发送邮箱变更验证码 |
| POST | `/answer/api/v1/user/email/verification` | 邮箱验证 |
| POST | `/answer/api/v1/user/email/verification/send` | 发送邮箱验证 |
| GET | `/answer/api/v1/user/ranking` | 用户排名 |
| GET | `/answer/api/v1/user/staff` | 获取版主/工作人员 |
| GET | `/answer/api/v1/user/info/search` | 按名称搜索用户 |
| GET | `/answer/api/v1/user/action/record` | 用户操作记录 |
| GET | `/answer/api/v1/user/notification/config` | 获取通知配置 |
| POST | `/answer/api/v1/user/notification/config` | 更新通知配置 |
| PUT | `/answer/api/v1/user/notification/unsubscribe` | 取消订阅通知 |

### 5. Comment — 评论系统（8 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/comment/page` | 评论列表 |
| GET | `/answer/api/v1/comment` | 获取评论详情 |
| POST | `/answer/api/v1/comment` | 添加评论 |
| PUT | `/answer/api/v1/comment` | 更新评论 |
| DELETE | `/answer/api/v1/comment` | 删除评论 |
| GET | `/answer/api/v1/activity/timeline` | 对象时间线 |
| GET | `/answer/api/v1/activity/timeline/detail` | 时间线详情 |
| GET | `/answer/api/v1/personal/comment/page` | 个人评论列表 |

### 6. Activity — 投票与关注（5 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/answer/api/v1/vote/up` | 点赞 |
| POST | `/answer/api/v1/vote/down` | 踩 |
| GET | `/answer/api/v1/personal/vote/page` | 个人投票记录 |
| POST | `/answer/api/v1/follow` | 关注/取消关注 |
| PUT | `/answer/api/v1/follow/tags` | 更新关注标签 |

### 7. Notification — 通知系统（5 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/notification/page` | 通知列表 |
| PUT | `/answer/api/v1/notification/read/state` | 标记已读 |
| PUT | `/answer/api/v1/notification/read/state/all` | 全部标记已读 |
| GET | `/answer/api/v1/notification/status` | 获取红点状态 |
| PUT | `/answer/api/v1/notification/status` | 清除红点 |

### 8. AI — AI 对话（6 个端点）

**用户端：**

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/ai/conversation` | 获取对话详情 |
| GET | `/answer/api/v1/ai/conversation/page` | 对话列表 |
| POST | `/answer/api/v1/ai/conversation/vote` | 对话投票 |

**管理端：**

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/admin/api/ai/conversation/page` | 管理员查看对话列表 |
| GET | `/answer/admin/api/ai/conversation` | 管理员查看对话详情 |
| DELETE | `/answer/admin/api/ai/conversation` | 管理员删除对话 |

### 9. Search — 搜索（2 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/search` | 搜索对象 |
| GET | `/answer/api/v1/search/desc` | 获取搜索描述 |

### 10. Report — 举报（3 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/answer/api/v1/report` | 添加举报 |
| PUT | `/answer/api/v1/report/review` | 审核举报 |
| GET | `/answer/api/v1/report/unreviewed/post` | 未审核举报列表 |

### 11. Review — 内容审核（2 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/review/pending/post/page` | 待审核内容列表 |
| PUT | `/answer/api/v1/review/pending/post` | 更新审核状态 |

### 12. Revision — 版本修订（5 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/revisions` | 修订列表 |
| GET | `/answer/api/v1/revisions/unreviewed` | 未审核修订列表 |
| PUT | `/answer/api/v1/revisions/audit` | 修订审核 |
| GET | `/answer/api/v1/revisions/edit/check` | 检查是否可编辑修订 |
| GET | `/answer/api/v1/reviewing/type` | 获取审核类型 |

### 13. Collection — 收藏（2 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/answer/api/v1/collection/switch` | 添加/取消收藏 |
| GET | `/answer/api/v1/personal/collection/page` | 个人收藏列表 |

### 14. Badge — 徽章系统（7 个端点）

**用户端：**

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/badge` | 获取徽章信息 |
| GET | `/answer/api/v1/badge/awards/page` | 徽章奖励列表 |
| GET | `/answer/api/v1/badge/user/awards` | 用户徽章奖励 |
| GET | `/answer/api/v1/badge/user/awards/recent` | 用户最近徽章 |
| GET | `/answer/api/v1/badges` | 所有徽章（按组分类） |

**管理端：**

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/admin/api/badges` | 徽章列表 |
| PUT | `/answer/admin/api/badge/status` | 更新徽章状态 |

### 15. Upload — 文件上传（2 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/answer/api/v1/file` | 上传文件 |
| POST | `/answer/api/v1/post/render` | 渲染内容（预览） |

### 16. Plugin — 插件系统（15 个端点）

**前端插件：**

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/plugin/status` | 获取所有插件状态 |
| GET | `/answer/api/v1/embed/config` | 获取嵌入插件配置 |
| GET | `/answer/api/v1/render/config` | 获取渲染配置 |
| GET | `/answer/api/v1/user/plugin/config` | 获取用户插件配置 |
| PUT | `/answer/api/v1/user/plugin/config` | 更新用户插件配置 |
| GET | `/answer/api/v1/user/plugin/configs` | 用户可用插件列表 |

**外部登录连接器：**

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/connector/info` | 获取已启用的连接器 |
| GET | `/answer/api/v1/connector/user/info` | 获取用户连接器信息 |
| DELETE | `/answer/api/v1/connector/user/unbinding` | 解绑外部登录 |
| POST | `/answer/api/v1/connector/binding/email` | 绑定外部登录发送邮箱 |

**管理端：**

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/admin/api/plugins` | 插件列表 |
| GET | `/answer/admin/api/plugin/config` | 获取插件配置 |
| PUT | `/answer/admin/api/plugin/config` | 更新插件配置 |
| PUT | `/answer/admin/api/plugin/status` | 更新插件状态 |

### 17. Personal — 个人中心（4 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/personal/answer/page` | 个人回答列表 |
| GET | `/personal/question/page` | 个人问题列表 |
| GET | `/answer/api/v1/personal/rank/page` | 个人排名 |
| GET | `/answer/api/v1/personal/qa/top` | 用户问答排行 |

### 18. Meta — 表情反应（2 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/meta/reaction` | 获取表情反应 |
| PUT | `/answer/api/v1/meta/reaction` | 添加/更新表情反应 |

### 19. Permission — 权限（1 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/permission` | 检查用户权限 |

### 20. Lang — 多语言（5 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/language/options` | 获取语言选项 |
| GET | `/answer/api/v1/language/config` | 获取语言配置映射 |
| GET | `/answer/admin/api/language/options` | 管理员获取语言选项 |
| GET | `/installation/language/options` | 安装页语言选项 |
| GET | `/installation/language/config` | 安装页语言配置 |

### 21. Site — 站点公开信息（4 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/api/v1/siteinfo` | 站点信息 |
| GET | `/answer/api/v1/siteinfo/legal` | 法律信息 |
| GET | `/custom.css` | 自定义 CSS |
| GET | `/robots.txt` | robots.txt |

### 22. Installation — 系统安装（5 个端点）

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/` | 检查配置文件，不存在则重定向到安装页 |
| POST | `/installation/base-info` | 初始化基础信息 |
| POST | `/installation/config-file/check` | 检查配置文件是否存在 |
| POST | `/installation/db/check` | 检查数据库是否存在 |
| POST | `/installation/init` | 初始化环境 |

---

## 管理后台 — admin（59 个端点）

### AI 配置

| 方法 | 路径 | 说明 |
|------|------|------|
| GET/PUT | `/answer/admin/api/ai-config` | 获取/更新 AI 配置 |
| POST | `/answer/admin/api/ai-models` | 获取 AI 模型列表 |
| GET | `/answer/admin/api/ai-provider` | 获取 AI 提供商配置 |
| GET/PUT | `/answer/admin/api/mcp-config` | 获取/更新 MCP 配置 |

### API Key 管理

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/admin/api/api-key/all` | 获取所有 API Key |
| POST | `/answer/admin/api/api-key` | 添加 API Key |
| PUT | `/answer/admin/api/api-key` | 更新 API Key |
| DELETE | `/answer/admin/api/api-key` | 删除 API Key |

### 仪表盘与内容管理

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/admin/api/dashboard` | 仪表盘信息 |
| GET | `/answer/admin/api/question/page` | 管理问题列表 |
| PUT | `/answer/admin/api/question/status` | 更新问题状态 |
| GET | `/answer/admin/api/answer/page` | 管理回答列表 |
| PUT | `/answer/admin/api/answer/status` | 更新回答状态 |
| DELETE | `/answer/admin/api/delete/permanently` | 永久删除 |

### 站点设置

| 方法 | 路径 | 说明 |
|------|------|------|
| GET/PUT | `/answer/admin/api/siteinfo/general` | 通用信息 |
| GET/PUT | `/answer/admin/api/siteinfo/branding` | 品牌/Logo |
| GET/PUT | `/answer/admin/api/siteinfo/interface` | 界面设置 |
| GET/PUT | `/answer/admin/api/siteinfo/theme` | 主题配置 |
| GET/PUT | `/answer/admin/api/siteinfo/login` | 登录配置 |
| GET/PUT | `/answer/admin/api/siteinfo/security` | 安全设置 |
| GET/PUT | `/answer/admin/api/siteinfo/seo` | SEO 配置 |
| GET/PUT | `/answer/admin/api/siteinfo/polices` | 政策配置 |
| GET/PUT | `/answer/admin/api/siteinfo/tag` | 标签设置 |
| GET/PUT | `/answer/admin/api/siteinfo/question` | 问题设置 |
| GET/PUT | `/answer/admin/api/siteinfo/users` | 用户设置 |
| GET/PUT | `/answer/admin/api/siteinfo/users-settings` | 用户详细设置 |
| GET/PUT | `/answer/admin/api/siteinfo/custom-css-html` | 自定义 CSS/HTML |
| GET/PUT | `/answer/admin/api/siteinfo/advanced` | 高级设置 |

### 用户管理

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/admin/api/users/page` | 用户列表 |
| POST | `/answer/admin/api/user` | 添加用户 |
| POST | `/answer/admin/api/users` | 批量添加用户 |
| PUT | `/answer/admin/api/user/status` | 更新用户状态 |
| PUT | `/answer/admin/api/user/role` | 更新用户角色 |
| PUT | `/answer/admin/api/user/profile` | 编辑用户资料 |
| PUT | `/answer/admin/api/user/password` | 重置用户密码 |
| GET | `/answer/admin/api/user/activation` | 获取用户激活信息 |
| POST | `/answer/admin/api/users/activation` | 发送用户激活邮件 |

### 其他管理

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/answer/admin/api/roles` | 角色列表 |
| GET/PUT | `/answer/admin/api/setting/privileges` | 权限配置 |
| GET/PUT | `/answer/admin/api/setting/smtp` | SMTP 配置 |
| GET/PUT | `/answer/admin/api/theme/options` | 主题选项 |
| GET | `/answer/admin/api/reasons` | 举报原因管理 |

---

## 核心数据模型

| 实体 | 关键 Schema | 说明 |
|------|-------------|------|
| Question | `QuestionInfoResp` (32 字段), `QuestionAdd`, `QuestionUpdate` | 问题（核心实体） |
| Answer | `AnswerInfo` (15 字段), `AnswerAddReq`, `AnswerUpdateReq` | 回答 |
| User | `GetCurrentLoginUserInfoResp` (27 字段), `UserBasicInfo` (10 字段) | 用户 |
| Tag | `GetTagResp` (17 字段), `TagItem`, `TagSynonym` | 标签及同义词 |
| Comment | `GetCommentResp` (18 字段), `AddCommentReq` | 评论 |
| Vote | `VoteReq`, `VoteResp` | 投票 |
| Collection | `CollectionSwitchReq` | 收藏 |
| Notification | `NotificationChannelConfig` | 通知 |
| Badge | `GetBadgeInfoResp` (8 字段), `BadgeListInfo` | 徽章 |
| Report | `GetReportListPageResp` (21 字段) | 举报 |
| Revision | `GetRevisionResp` (10 字段) | 版本修订 |
| Search | `SearchObject` (12 字段), `SearchResp` | 搜索 |
| AI | `AIConversationRecord`, `SiteAIReq`, `SiteMCPReq` | AI 对话与配置 |
| Site | `SiteInfoResp` (18 字段) 及各种 `Site*Req/Resp` | 站点配置 |
| Plugin | `GetPluginListResp` (7 字段), `ConfigField` | 插件系统 |

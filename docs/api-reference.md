# Apache Answer API 参考

> 从 `docs/swagger.json` 自动生成，共 **190 个端点**，204 个数据模型

## 认证方式

- **Auth**：通过 HTTP Header `Authorization` 传递 Token，写操作和用户相关操作需要认证
- **ApiKeyAuth**：管理员 API Key 认证
- **Public**：无需认证的公开端点

## API 基础路径

- 用户端：`/answer/api/v1/`
- 管理端：`/answer/admin/api/`

---

### AI 对话 — `ai-conversation`（3 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/ai/conversation` |  | query: conversation_id | `AIConversationDetailResp` | get conversation detail |
| GET | `/ai/conversation/page` |  | query: page, page_size | `PageModel` | get conversation list |
| POST | `/ai/conversation/vote` |  | AIConversationVoteReq | `` | vote record |

### AI 对话（管理） — `ai-conversation-admin`（3 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/ai/conversation` |  | query: conversation_id | `AIConversationAdminDetailResp` | get conversation detail for admin |
| DELETE | `/ai/conversation` |  | AIConversationAdminDeleteReq | `` | delete conversation for admin |
| GET | `/ai/conversation/page` |  | query: page, page_size | `PageModel` | get conversation list for admin |

### site — `site`（2 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/siteinfo` |  |  | `SiteInfoResp` | get site info |
| GET | `/siteinfo/legal` |  | query: info_type | `GetSiteLegalInfoResp` | get site legal info |

### 个人中心 — `Personal`（1 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/personal/answer/page` | 🔑 | query: username, order, page, page_size | `` | list personal answers |

### 举报 — `Report`（3 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| POST | `/report` | 🔑 | AddReportReq | `` | add report |
| PUT | `/report/review` | 🔑 | ReviewReportReq | `` | review report |
| GET | `/report/unreviewed/post` | 🔑 | query: page | `PageModel` | get unreviewed report post page |

### 举报原因 — `reason`（2 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/reasons` | 🔑 | query: object_type, action | `` | get reasons by object type and action |
| GET | `/reasons` | 🔑 | query: object_type, action | `` | get reasons by object type and action |

### 内容审核 — `Review`（2 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| PUT | `/review/pending/post` | 🔑 | UpdateReviewReq | `` | update review |
| GET | `/review/pending/post/page` | 🔑 | query: page, object_id | `PageModel` | get unreviewed post page |

### 回答管理 — `Answer`（7 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| PUT | `/answer` | 🔑 | AnswerUpdateReq | `` | Update Answer |
| POST | `/answer` | 🔑 | AnswerAddReq | `` | Add Answer |
| DELETE | `/answer` | 🔑 | RemoveAnswerReq | `` | delete answer |
| POST | `/answer/acceptance` | 🔑 | AcceptAnswerReq | `` | Accept Answer |
| GET | `/answer/info` |  | query: id | `GetAnswerInfoResp` | Get Answer Detail |
| GET | `/answer/page` |  | query: question_id, order, page, page_size | `` | AnswerList |
| POST | `/answer/recover` | 🔑 | RecoverAnswerReq | `` | recover answer |

### 外部登录连接器 — `PluginConnector`（4 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| POST | `/connector/binding/email` |  | ExternalLoginBindingUserSendEmailReq | `ExternalLoginBindingUserSendEmailResp` | external login binding user send email |
| GET | `/connector/info` | 🔑 |  | `` | get all enabled connectors |
| GET | `/connector/user/info` | 🔑 |  | `` | get all connectors info about user |
| DELETE | `/connector/user/unbinding` | 🔑 | ExternalLoginUnbindingReq | `` | unbind external user login |

### 多语言 — `Lang`（3 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/language/options` | 🔑 |  | `` | Get language options |
| GET | `/language/config` |  | header: Accept-Language | `` | get language config mapping |
| GET | `/language/options` |  |  | `` | Get language options |

### 徽章管理 — `AdminBadge`（2 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| PUT | `/badge/status` | 🔑 | UpdateBadgeStatusReq | `` | update badge status |
| GET | `/badges` | 🔑 | query: page, page_size, status, q | `` | list all badges by page |

### 徽章系统 — `api-badge`（5 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/badge` |  | query: id | `GetBadgeInfoResp` | get badge info |
| GET | `/badge/awards/page` |  | query: page, page_size, badge_id, username | `GetBadgeInfoResp` | get badge award list |
| GET | `/badge/user/awards` |  | query: username | `` | get user badge award list |
| GET | `/badge/user/awards/recent` |  | query: username | `` | get user badge award list |
| GET | `/badges` |  |  | `` | list all badges group by group |

### 投票与关注 — `Activity`（5 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| POST | `/follow` | 🔑 | FollowReq | `FollowResp` | follow object or cancel follow operation |
| PUT | `/follow/tags` | 🔑 | UpdateFollowTagsReq | `` | update user follow tags |
| GET | `/personal/vote/page` | 🔑 | query: page, page_size | `PageModel` | get user personal votes |
| POST | `/vote/down` | 🔑 | VoteReq | `VoteResp` | vote down |
| POST | `/vote/up` | 🔑 | VoteReq | `VoteResp` | vote up |

### 排名 — `Rank`（1 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/personal/rank/page` |  | query: page, page_size, username | `PageModel` | user personal rank list |

### 插件渲染 — `PluginRender`（1 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/render/config` |  |  | `RenderConfig` | GetRenderConfig |

### 插件管理 — `AdminPlugin`（4 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/plugin/config` | 🔑 | query: plugin_slug_name | `GetPluginConfigResp` | get plugin config |
| PUT | `/plugin/config` | 🔑 | UpdatePluginConfigReq | `` | update plugin config |
| PUT | `/plugin/status` | 🔑 | UpdatePluginStatusReq | `` | update plugin status |
| GET | `/plugins` | 🔑 | query: status, have_config | `` | get plugin list |

### 插件系统 — `Plugin`（2 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/embed/config` |  |  | `` | get embed plugin config |
| GET | `/plugin/status` |  |  | `` | get all plugins status |

### 搜索 — `Search`（2 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/search` | 🔑 | query: q, order | `SearchResp` | search object |
| GET | `/search/desc` |  |  | `SearchResp` | get search description |

### 收藏 — `Collection`（2 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| POST | `/collection/switch` | 🔑 | CollectionSwitchReq | `CollectionSwitchResp` | add collection |
| GET | `/personal/collection/page` | 🔑 | query: page, page_size | `` | list personal collections |

### 文件上传 — `Upload`（2 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| POST | `/file` | 🔑 | formData: source, file | `` | upload file |
| POST | `/post/render` | 🔑 | PostRenderReq | `` | render post content |

### 权限 — `Permission`（1 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/permission` | 🔑 | header: Authorization, action | `` | check user permission |

### 标签管理 — `Tag`（12 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/question/tags` | 🔑 | query: tag | `` | get tag list |
| GET | `/tag` |  | query: tag_id, tag_name | `GetTagResp` | get tag one |
| PUT | `/tag` | 🔑 | UpdateTagReq | `` | update tag |
| POST | `/tag` | 🔑 | AddTagReq | `` | add tag |
| DELETE | `/tag` | 🔑 | RemoveTagReq | `` | delete tag |
| POST | `/tag/merge` | 🔑 | AddTagReq | `` | merge tag |
| POST | `/tag/recover` | 🔑 | RecoverTagReq | `` | recover delete tag |
| PUT | `/tag/synonym` | 🔑 | UpdateTagSynonymReq | `` | update tag |
| GET | `/tag/synonyms` |  | query: tag_id | `GetTagSynonymsResp` | get tag synonyms |
| GET | `/tags` |  | query: tags | `` | get tags list |
| GET | `/tags/following` | 🔑 |  | `` | get following tag list |
| GET | `/tags/page` |  | query: page, page_size, slug_name, query_cond | `PageModel` | get tag page |

### 版本修订 — `Revision`（5 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/reviewing/type` | 🔑 |  | `` | get reviewing type |
| GET | `/revisions` |  | query: object_id | `` | get revision list |
| PUT | `/revisions/audit` | 🔑 | RevisionAuditReq | `` | revision audit |
| GET | `/revisions/edit/check` | 🔑 | query: id | `` | check can update revision |
| GET | `/revisions/unreviewed` | 🔑 | query: page | `PageModel` | get unreviewed revision list |

### 用户插件 — `UserPlugin`（3 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/user/plugin/config` | 🔑 | query: plugin_slug_name | `GetPluginConfigResp` | get user plugin config |
| PUT | `/user/plugin/config` | 🔑 | UpdateUserPluginConfigReq | `` | update user plugin config |
| GET | `/user/plugin/configs` | 🔑 |  | `` | get plugin list that used for user. |

### 用户管理 — `User`（21 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/personal/user/info` | 🔑 | query: username | `GetOtherUserInfoResp` | GetOtherUserInfoByUsername |
| GET | `/user/action/record` | 🔑 | query: action | `ActionRecordResp` | ActionRecord |
| PUT | `/user/email` | 🔑 | UserChangeEmailVerifyReq | `` | user change email verification |
| POST | `/user/email/change/code` | 🔑 | UserChangeEmailSendCodeReq | `` | send email to the user email then change their email |
| POST | `/user/email/verification` |  | query: code | `UserLoginResp` | UserVerifyEmail |
| POST | `/user/email/verification/send` | 🔑 | query: captcha_id, captcha_code | `` | UserVerifyEmailSend |
| GET | `/user/info` | 🔑 |  | `GetCurrentLoginUserInfoResp` | GetUserInfoByUserID |
| PUT | `/user/info` | 🔑 | header: Authorization, data | `` | UserUpdateInfo update user info |
| GET | `/user/info/search` | 🔑 | query: username | `GetOtherUserInfoResp` | SearchUserListByName |
| PUT | `/user/interface` | 🔑 | header: Authorization, data | `` | UserUpdateInterface update user interface config |
| POST | `/user/login/email` |  | UserEmailLoginReq | `UserLoginResp` | UserEmailLogin |
| GET | `/user/logout` | 🔑 |  | `` | user logout |
| PUT | `/user/notification/config` | 🔑 | UpdateUserNotificationConfigReq | `` | update user's notification config |
| POST | `/user/notification/config` | 🔑 |  | `GetUserNotificationConfigResp` | get user's notification config |
| PUT | `/user/notification/unsubscribe` |  | UserUnsubscribeNotificationReq | `` | unsubscribe notification |
| PUT | `/user/password` | 🔑 | UserModifyPasswordReq | `` | UserModifyPassWord |
| POST | `/user/password/replacement` |  | UserRePassWordRequest | `` | UseRePassWord |
| POST | `/user/password/reset` |  | UserRetrievePassWordRequest | `` | RetrievePassWord |
| GET | `/user/ranking` |  |  | `UserRankingResp` | get user ranking |
| POST | `/user/register/email` |  | UserRegisterReq | `UserLoginResp` | UserRegisterByEmail |
| GET | `/user/staff` |  | query: username, page_size | `GetUserStaffResp` | get user staff |

### 管理后台 — `admin`（59 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/ai-config` | 🔑 |  | `SiteAIResp` | get AI configuration |
| PUT | `/ai-config` | 🔑 | SiteAIReq | `` | update AI configuration |
| POST | `/ai-models` | 🔑 |  | `` | get AI models |
| GET | `/ai-provider` | 🔑 |  | `` | get AI provider configuration |
| GET | `/answer/page` | 🔑 | query: page, page_size, status, query, question_id | `` | AdminAnswerPage admin answer page |
| PUT | `/answer/status` | 🔑 | AdminUpdateAnswerStatusReq | `` | update answer status |
| PUT | `/api-key` | 🔑 | UpdateAPIKeyReq | `` | update apikey |
| POST | `/api-key` | 🔑 | AddAPIKeyReq | `AddAPIKeyResp` | add apikey |
| DELETE | `/api-key` | 🔑 | DeleteAPIKeyReq | `` | delete apikey |
| GET | `/api-key/all` | 🔑 |  | `` | get all api keys |
| GET | `/dashboard` | 🔑 |  | `` | DashboardInfo |
| DELETE | `/delete/permanently` | 🔑 | DeletePermanentlyReq | `` | delete permanently |
| GET | `/mcp-config` | 🔑 |  | `SiteMCPResp` | get MCP configuration |
| PUT | `/mcp-config` | 🔑 | SiteMCPReq | `` | update MCP configuration |
| GET | `/question/page` | 🔑 | query: page, page_size, status, query | `` | AdminQuestionPage admin question page |
| PUT | `/question/status` | 🔑 | AdminUpdateQuestionStatusReq | `` | update question status |
| GET | `/roles` | 🔑 |  | `` | get role list |
| GET | `/setting/privileges` | 🔑 |  | `GetPrivilegesConfigResp` | GetPrivilegesConfig get privileges config |
| PUT | `/setting/privileges` | 🔑 | UpdatePrivilegesConfigReq | `` | update privileges config |
| GET | `/setting/smtp` | 🔑 |  | `GetSMTPConfigResp` | GetSMTPConfig get smtp config |
| PUT | `/setting/smtp` | 🔑 | UpdateSMTPConfigReq | `` | update smtp config |
| GET | `/siteinfo/advanced` | 🔑 |  | `SiteAdvancedResp` | get site advanced setting |
| PUT | `/siteinfo/advanced` | 🔑 | SiteAdvancedReq | `` | update site advanced info |
| GET | `/siteinfo/branding` | 🔑 |  | `SiteBrandingResp` | get site interface |
| PUT | `/siteinfo/branding` | 🔑 | SiteBrandingReq | `` | update site info branding |
| GET | `/siteinfo/custom-css-html` | 🔑 |  | `SiteCustomCssHTMLResp` | get site info custom html css config |
| PUT | `/siteinfo/custom-css-html` | 🔑 | SiteCustomCssHTMLReq | `` | update site custom css html config |
| GET | `/siteinfo/general` | 🔑 |  | `SiteGeneralResp` | get site general information |
| PUT | `/siteinfo/general` | 🔑 | SiteGeneralReq | `` | update site general information |
| GET | `/siteinfo/interface` | 🔑 |  | `SiteInterfaceSettingsResp` | get site interface |
| PUT | `/siteinfo/interface` | 🔑 | SiteInterfaceReq | `` | update site info interface |
| GET | `/siteinfo/login` | 🔑 |  | `SiteLoginResp` | get site info login config |
| PUT | `/siteinfo/login` | 🔑 | SiteLoginReq | `` | update site login |
| GET | `/siteinfo/polices` | 🔑 |  | `SitePoliciesResp` | Get the policies information for the site |
| PUT | `/siteinfo/polices` | 🔑 | SitePoliciesReq | `` | update site policies configuration |
| GET | `/siteinfo/question` | 🔑 |  | `SiteQuestionsResp` | get site questions setting |
| PUT | `/siteinfo/question` | 🔑 | SiteQuestionsReq | `` | update site question settings |
| GET | `/siteinfo/security` | 🔑 |  | `SiteSecurityResp` | Get the security information for the site |
| PUT | `/siteinfo/security` | 🔑 | SiteSecurityReq | `` | update site security configuration |
| GET | `/siteinfo/seo` | 🔑 |  | `SiteSeoResp` | get site seo information |
| PUT | `/siteinfo/seo` | 🔑 | SiteSeoReq | `` | update site seo information |
| GET | `/siteinfo/tag` | 🔑 |  | `SiteTagsResp` | get site tags setting |
| PUT | `/siteinfo/tag` | 🔑 | SiteTagsReq | `` | update site tag settings |
| GET | `/siteinfo/theme` | 🔑 |  | `SiteThemeResp` | get site info theme config |
| PUT | `/siteinfo/theme` | 🔑 | SiteThemeReq | `` | update site custom css html config |
| GET | `/siteinfo/users` | 🔑 |  | `SiteUsersResp` | get site user config |
| PUT | `/siteinfo/users` | 🔑 | SiteUsersReq | `` | update site info config about users |
| GET | `/siteinfo/users-settings` | 🔑 |  | `SiteUsersSettingsResp` | get site interface |
| PUT | `/siteinfo/users-settings` | 🔑 | SiteUsersSettingsReq | `` | update site info users settings |
| GET | `/theme/options` | 🔑 |  | `` | Get theme options |
| POST | `/user` | 🔑 | AddUserReq | `` | add user |
| GET | `/user/activation` | 🔑 | query: user_id | `GetUserActivationResp` | get user activation |
| PUT | `/user/password` | 🔑 | UpdateUserPasswordReq | `` | update user password |
| PUT | `/user/profile` | 🔑 | EditUserProfileReq | `` | edit user profile |
| PUT | `/user/role` | 🔑 | UpdateUserRoleReq | `` | update user role |
| PUT | `/user/status` | 🔑 | UpdateUserStatusReq | `` | update user |
| POST | `/users` | 🔑 | AddUsersReq | `` | add users |
| POST | `/users/activation` | 🔑 | SendUserActivationReq | `` | send user activation |
| GET | `/users/page` | 🔑 | query: page, page_size, query, staff, status | `PageModel` | get user page |

### 系统安装 — `installation`（1 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/` |  |  | `` | if config file not exist try to redirect to install page |

### 表情反应 — `Meta`（2 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/meta/reaction` |  | query: object_id | `ReactionRespItem` | get reaction |
| PUT | `/meta/reaction` | 🔑 | UpdateReactionReq | `` | add or update reaction |

### 评论系统 — `Comment`（8 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/activity/timeline` |  | query: object_id, tag_slug_name, object_type, show_vote | `GetObjectTimelineResp` | get object timeline |
| GET | `/activity/timeline/detail` |  | query: revision_id | `GetObjectTimelineResp` | get object timeline detail |
| GET | `/comment` |  | query: id | `PageModel` | get comment by id |
| PUT | `/comment` | 🔑 | UpdateCommentReq | `` | update comment |
| POST | `/comment` | 🔑 | AddCommentReq | `GetCommentResp` | add comment |
| DELETE | `/comment` | 🔑 | RemoveCommentReq | `` | remove comment |
| GET | `/comment/page` |  | query: page, page_size, object_id, query_cond | `PageModel` | get comment page |
| GET | `/personal/comment/page` |  | query: page, page_size, username | `PageModel` | user personal comment list |

### 通知系统 — `Notification`（5 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/notification/page` | 🔑 | query: page, page_size, type, inbox_type | `` | get notification list |
| PUT | `/notification/read/state` | 🔑 | NotificationClearIDRequest | `` | ClearUnRead |
| PUT | `/notification/read/state/all` | 🔑 | NotificationClearRequest | `` | ClearUnRead |
| GET | `/notification/status` | 🔑 |  | `` | GetRedDot |
| PUT | `/notification/status` | 🔑 | NotificationClearRequest | `` | DelRedDot |

### 问题管理 — `Question`（17 个端点）

| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |
|------|------|------|------|------|------|
| GET | `/personal/qa/top` |  | query: username | `` | UserTop |
| PUT | `/question` | 🔑 | QuestionUpdate | `` | update question |
| POST | `/question` | 🔑 | QuestionAdd | `` | add question |
| DELETE | `/question` | 🔑 | RemoveQuestionReq | `` | delete question |
| POST | `/question/answer` | 🔑 | QuestionAddByAnswer | `` | add question and answer |
| GET | `/question/info` |  | query: id | `` | get question details |
| GET | `/question/invite` |  | query: id | `` | get question invite user info |
| PUT | `/question/invite` | 🔑 | QuestionUpdateInviteUser | `` | update question invite user |
| GET | `/question/link` |  | query: in_days, order, page, page_size, question_id | `PageModel` | get question link |
| PUT | `/question/operation` | 🔑 | OperationQuestionReq | `` | Operation question |
| GET | `/question/page` |  | QuestionPageReq | `PageModel` | get questions by page |
| GET | `/question/recommend/page` |  | QuestionPageReq | `PageModel` | get recommend questions by page |
| POST | `/question/recover` | 🔑 | QuestionRecoverReq | `` | recover deleted question |
| PUT | `/question/reopen` | 🔑 | ReopenQuestionReq | `` | reopen question |
| GET | `/question/similar` | 🔑 | query: title | `` | fuzzy query similar questions based on title |
| GET | `/question/similar/tag` |  | query: question_id | `` | Search Similar Question |
| PUT | `/question/status` | 🔑 | CloseQuestionReq | `` | Close question |

---

## 核心数据模型

详细模型定义见 `docs/api/*.models.md`，以下按包分组列出被 API 引用的模型。

### 响应包装（1 个）

`RespBody`

### 分页（1 个）

`PageModel`

### 插件（2 个）

`EmbedConfig`, `RenderConfig`

### 核心业务模型（155 个）

`AIConversationAdminDeleteReq`, `AIConversationAdminDetailResp`, `AIConversationAdminListItem`, `AIConversationDetailResp`, `AIConversationListItem`, `AIConversationVoteReq`
`AcceptAnswerReq`, `ActionRecordResp`, `AddAPIKeyReq`, `AddAPIKeyResp`, `AddCommentReq`, `AddReportReq`
`AddTagReq`, `AddUserReq`, `AddUsersReq`, `AdminUpdateAnswerStatusReq`, `AdminUpdateQuestionStatusReq`, `AnswerAddReq`
`AnswerUpdateReq`, `CloseQuestionReq`, `CollectionSwitchReq`, `CollectionSwitchResp`, `ConnectorInfoResp`, `ConnectorUserInfoResp`
`DeleteAPIKeyReq`, `DeletePermanentlyReq`, `EditUserProfileReq`, `ExternalLoginBindingUserSendEmailReq`, `ExternalLoginBindingUserSendEmailResp`, `ExternalLoginUnbindingReq`
`FollowReq`, `FollowResp`, `GetAIModelResp`, `GetAIProviderResp`, `GetAPIKeyResp`, `GetAnswerInfoResp`
`GetBadgeInfoResp`, `GetBadgeListPagedResp`, `GetBadgeListResp`, `GetCommentPersonalWithPageResp`, `GetCommentResp`, `GetCurrentLoginUserInfoResp`
`GetFollowingTagsResp`, `GetObjectTimelineResp`, `GetOtherUserInfoResp`, `GetPluginConfigResp`, `GetPluginListResp`, `GetPrivilegesConfigResp`
`GetRankPersonalPageResp`, `GetReportListPageResp`, `GetReviewingTypeResp`, `GetRevisionResp`, `GetRoleResp`, `GetSMTPConfigResp`
`GetSiteLegalInfoResp`, `GetTagBasicResp`, `GetTagPageResp`, `GetTagResp`, `GetTagSynonymsResp`, `GetUnreviewedPostPageResp`
`GetUnreviewedRevisionResp`, `GetUserActivationResp`, `GetUserBadgeAwardListResp`, `GetUserNotificationConfigResp`, `GetUserPageResp`, `GetUserPluginListResp`
`GetUserStaffResp`, `GetVoteWithPageResp`, `NotificationClearIDRequest`, `NotificationClearRequest`, `OperationQuestionReq`, `PostRenderReq`
`QuestionAdd`, `QuestionAddByAnswer`, `QuestionPageReq`, `QuestionPageResp`, `QuestionRecoverReq`, `QuestionUpdate`
`QuestionUpdateInviteUser`, `ReactionRespItem`, `RecoverAnswerReq`, `RecoverTagReq`, `RemoveAnswerReq`, `RemoveCommentReq`
`RemoveQuestionReq`, `RemoveTagReq`, `ReopenQuestionReq`, `ReviewReportReq`, `RevisionAuditReq`, `SearchResp`
`SendUserActivationReq`, `SiteAIReq`, `SiteAIResp`, `SiteAdvancedReq`, `SiteAdvancedResp`, `SiteBrandingReq`
`SiteBrandingResp`, `SiteCustomCssHTMLReq`, `SiteCustomCssHTMLResp`, `SiteGeneralReq`, `SiteGeneralResp`, `SiteInfoResp`
`SiteInterfaceReq`, `SiteInterfaceSettingsResp`, `SiteLoginReq`, `SiteLoginResp`, `SiteMCPReq`, `SiteMCPResp`
`SitePoliciesReq`, `SitePoliciesResp`, `SiteQuestionsReq`, `SiteQuestionsResp`, `SiteSecurityReq`, `SiteSecurityResp`
`SiteSeoReq`, `SiteSeoResp`, `SiteTagsReq`, `SiteTagsResp`, `SiteThemeReq`, `SiteThemeResp`
`SiteUsersReq`, `SiteUsersResp`, `SiteUsersSettingsReq`, `SiteUsersSettingsResp`, `UpdateAPIKeyReq`, `UpdateBadgeStatusReq`
`UpdateCommentReq`, `UpdateFollowTagsReq`, `UpdateInfoRequest`, `UpdatePluginConfigReq`, `UpdatePluginStatusReq`, `UpdatePrivilegesConfigReq`
`UpdateReactionReq`, `UpdateReviewReq`, `UpdateSMTPConfigReq`, `UpdateTagReq`, `UpdateTagSynonymReq`, `UpdateUserInterfaceRequest`
`UpdateUserNotificationConfigReq`, `UpdateUserPasswordReq`, `UpdateUserPluginConfigReq`, `UpdateUserRoleReq`, `UpdateUserStatusReq`, `UserChangeEmailSendCodeReq`
`UserChangeEmailVerifyReq`, `UserEmailLoginReq`, `UserLoginResp`, `UserModifyPasswordReq`, `UserRankingResp`, `UserRePassWordRequest`
`UserRegisterReq`, `UserRetrievePassWordRequest`, `UserUnsubscribeNotificationReq`, `VoteReq`, `VoteResp`


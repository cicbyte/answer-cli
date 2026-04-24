# admin

## `GET` `/ai-config`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteAIResp`

---

## `PUT` `/ai-config`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteAIReq`
> AI config
  - `ai_providers` []``SiteAIProvider``
  - `chosen_provider` `string`
  - `enabled` `boolean`
  - `prompt_config` `AIPromptConfig`

**响应:**

- `200`:
`RespBody`

---

## `POST` `/ai-models`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` []``GetAIModelResp``

---

## `GET` `/ai-provider`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` []``GetAIProviderResp``

---

## `GET` `/answer/page`

**认证**: 管理员 | **参数方式**: query

> Status:[available,deleted,pending]

**请求参数:**

  - `page` (`integer`) — page size
  - `page_size` (`integer`) — page size
  - `status` (`string`) — user status
  - `query` (`string`) — answer id or question title
  - `question_id` (`string`) — question id

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/answer/status`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `AdminUpdateAnswerStatusReq`
> AdminUpdateAnswerStatusReq
  - `answer_id` `string` **(必填)**
  - `status` `string` enum: `available`, `deleted` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/api-key`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateAPIKeyReq`
> apikey
  - `description` `string` **(必填)**
  - `id` `integer` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `POST` `/api-key`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `AddAPIKeyReq`
> apikey
  - `description` `string` **(必填)**
  - `scope` `string` enum: `read-only`, `global` **(必填)**

**响应:**

- `200`:
`RespBody`
  - `data` `AddAPIKeyResp`

---

## `DELETE` `/api-key`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `DeleteAPIKeyReq`
> apikey
  - `id` `integer`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/api-key/all`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` []``GetAPIKeyResp``

---

## `GET` `/dashboard`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`

---

## `DELETE` `/delete/permanently`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `DeletePermanentlyReq`
> DeletePermanentlyReq
  - `type` `string` enum: `users`, `questions`, `answers` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/mcp-config`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteMCPResp`

---

## `PUT` `/mcp-config`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteMCPReq`
> MCP config
  - `enabled` `boolean`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/question/page`

**认证**: 管理员 | **参数方式**: query

> Status:[available,closed,deleted,pending]

**请求参数:**

  - `page` (`integer`) — page size
  - `page_size` (`integer`) — page size
  - `status` (`string`) — user status
  - `query` (`string`) — question id or title

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/question/status`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `AdminUpdateQuestionStatusReq`
> AdminUpdateQuestionStatusReq
  - `question_id` `string` **(必填)**
  - `status` `string` enum: `available`, `closed`, `deleted` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/roles`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` []``GetRoleResp``

---

## `GET` `/setting/privileges`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `GetPrivilegesConfigResp`

---

## `PUT` `/setting/privileges`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdatePrivilegesConfigReq`
> config
  - `custom_privileges` []``Privilege``
  - `level` `PrivilegeLevel` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/setting/smtp`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `GetSMTPConfigResp`

---

## `PUT` `/setting/smtp`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateSMTPConfigReq`
> smtp config
  - `encryption` `string` enum: `SSL`, `TLS` — "" SSL TLS
  - `from_email` `string`
  - `from_name` `string`
  - `smtp_authentication` `boolean`
  - `smtp_host` `string`
  - `smtp_password` `string`
  - `smtp_port` `integer`
  - `smtp_username` `string`
  - `test_email_recipient` `string`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/advanced`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteAdvancedResp`

---

## `PUT` `/siteinfo/advanced`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteAdvancedReq`
> advanced settings
  - `authorized_attachment_extensions` []``string``
  - `authorized_image_extensions` []``string``
  - `max_attachment_size` `integer`
  - `max_image_megapixel` `integer`
  - `max_image_size` `integer`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/branding`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteBrandingResp`

---

## `PUT` `/siteinfo/branding`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteBrandingReq`
> branding info
  - `favicon` `string`
  - `logo` `string`
  - `mobile_logo` `string`
  - `square_icon` `string`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/custom-css-html`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteCustomCssHTMLResp`

---

## `PUT` `/siteinfo/custom-css-html`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteCustomCssHTMLReq`
> login info
  - `custom_css` `string`
  - `custom_footer` `string`
  - `custom_head` `string`
  - `custom_header` `string`
  - `custom_sidebar` `string`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/general`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteGeneralResp`

---

## `PUT` `/siteinfo/general`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteGeneralReq`
> general
  - `contact_email` `string` **(必填)**
  - `description` `string`
  - `name` `string` **(必填)**
  - `short_description` `string`
  - `site_url` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/interface`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteInterfaceSettingsResp`

---

## `PUT` `/siteinfo/interface`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteInterfaceReq`
> general
  - `language` `string` **(必填)**
  - `time_zone` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/login`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteLoginResp`

---

## `PUT` `/siteinfo/login`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteLoginReq`
> login info
  - `allow_email_domains` []``string``
  - `allow_email_registrations` `boolean`
  - `allow_new_registrations` `boolean`
  - `allow_password_login` `boolean`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/polices`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SitePoliciesResp`

---

## `PUT` `/siteinfo/polices`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SitePoliciesReq`
> write info
  - `privacy_policy_original_text` `string`
  - `privacy_policy_parsed_text` `string`
  - `terms_of_service_original_text` `string`
  - `terms_of_service_parsed_text` `string`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/question`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteQuestionsResp`

---

## `PUT` `/siteinfo/question`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteQuestionsReq`
> questions settings
  - `min_content` `integer`
  - `min_tags` `integer`
  - `restrict_answer` `boolean`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/security`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteSecurityResp`

---

## `PUT` `/siteinfo/security`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteSecurityReq`
> write info
  - `check_update` `boolean`
  - `external_content_display` `string` enum: `always_display`, `ask_before_display` **(必填)**
  - `login_required` `boolean`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/seo`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteSeoResp`

---

## `PUT` `/siteinfo/seo`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteSeoReq`
> seo
  - `permalink` `integer` **(必填)**
  - `robots` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/tag`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteTagsResp`

---

## `PUT` `/siteinfo/tag`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteTagsReq`
> tags settings
  - `recommend_tags` []``SiteWriteTag``
  - `required_tag` `boolean`
  - `reserved_tags` []``SiteWriteTag``

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/theme`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteThemeResp`

---

## `PUT` `/siteinfo/theme`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteThemeReq`
> login info
  - `color_scheme` `string`
  - `layout` `string` enum: `Full-width`, `Fixed-width`
  - `theme` `string` **(必填)**
  - `theme_config` `object`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/users`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteUsersResp`

---

## `PUT` `/siteinfo/users`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteUsersReq`
> users info
  - `allow_update_avatar` `boolean`
  - `allow_update_bio` `boolean`
  - `allow_update_display_name` `boolean`
  - `allow_update_location` `boolean`
  - `allow_update_username` `boolean`
  - `allow_update_website` `boolean`
  - `default_avatar` `string` enum: `system`, `gravatar` **(必填)**
  - `gravatar_base_url` `string`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/siteinfo/users-settings`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `SiteUsersSettingsResp`

---

## `PUT` `/siteinfo/users-settings`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SiteUsersSettingsReq`
> general
  - `default_avatar` `string` enum: `system`, `gravatar` **(必填)**
  - `gravatar_base_url` `string`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/theme/options`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`

---

## `POST` `/user`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `AddUserReq`
> user
  - `display_name` `string` **(必填)**
  - `email` `string` **(必填)**
  - `password` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/user/activation`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `user_id` (`string` **(必填)**) — user id

**响应:**

- `200`:
`RespBody`
  - `data` `GetUserActivationResp`

---

## `PUT` `/user/password`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateUserPasswordReq`
> user
  - `password` `string` **(必填)**
  - `user_id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/user/profile`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `EditUserProfileReq`
> user
  - `display_name` `string` **(必填)**
  - `email` `string` **(必填)**
  - `user_id` `string` **(必填)**
  - `username` `string`

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/user/role`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateUserRoleReq`
> user
  - `role_id` `integer` **(必填)** — role id
  - `user_id` `string` **(必填)** — user id

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/user/status`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateUserStatusReq`
> user
  - `remove_all_content` `boolean`
  - `status` `string` enum: `normal`, `suspended`, `deleted`, `inactive` **(必填)**
  - `suspend_duration` `string` enum: `24h`, `48h`, `72h`, `7d`, `14d`, `1m`, `2m`, `3m`, `6m`, `1y`, `forever`
  - `user_id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `POST` `/users`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `AddUsersReq`
> user
  - `users` `string` — users info line by line

**响应:**

- `200`:
`RespBody`

---

## `POST` `/users/activation`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `SendUserActivationReq`
> SendUserActivationReq
  - `user_id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/users/page`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `page` (`integer`) — page size
  - `page_size` (`integer`) — page size
  - `query` (`string`) — search query: email, username or id:[id]
  - `staff` (`boolean`) — staff user
  - `status` (`string`) — user status

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `records` 

---

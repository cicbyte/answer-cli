# admin — 数据模型

### `RespBody`

- `code` `integer` — http code
- `data`  — response data
- `msg` `string` — response message
- `reason` `string` — reason key

---

### `PageModel`

- `count` `integer`
- `list` 

---

### `AddAPIKeyReq`

- `description` `string` **(必填)**
- `scope` `string` enum: `read-only`, `global` **(必填)**

---

### `AddAPIKeyResp`

- `access_key` `string`

---

### `AddUserReq`

- `display_name` `string` **(必填)**
- `email` `string` **(必填)**
- `password` `string` **(必填)**

---

### `AddUsersReq`

- `users` `string` — users info line by line

---

### `AdminUpdateAnswerStatusReq`

- `answer_id` `string` **(必填)**
- `status` `string` enum: `available`, `deleted` **(必填)**

---

### `AdminUpdateQuestionStatusReq`

- `question_id` `string` **(必填)**
- `status` `string` enum: `available`, `closed`, `deleted` **(必填)**

---

### `DeleteAPIKeyReq`

- `id` `integer`

---

### `DeletePermanentlyReq`

- `type` `string` enum: `users`, `questions`, `answers` **(必填)**

---

### `EditUserProfileReq`

- `display_name` `string` **(必填)**
- `email` `string` **(必填)**
- `user_id` `string` **(必填)**
- `username` `string`

---

### `GetAIModelResp`

- `created` `integer`
- `id` `string`
- `object` `string`
- `owned_by` `string`

---

### `GetAIProviderResp`

- `default_api_host` `string`
- `display_name` `string`
- `name` `string`

---

### `GetAPIKeyResp`

- `access_key` `string`
- `created_at` `integer`
- `description` `string`
- `id` `integer`
- `last_used_at` `integer`
- `scope` `string`

---

### `GetPrivilegesConfigResp`

- `options` []``PrivilegeOption``
- `selected_level` `PrivilegeLevel`

---

### `GetRoleResp`

- `description` `string`
- `id` `integer`
- `name` `string`

---

### `GetSMTPConfigResp`

- `encryption` `string` — "" SSL TLS
- `from_email` `string`
- `from_name` `string`
- `smtp_authentication` `boolean`
- `smtp_host` `string`
- `smtp_password` `string`
- `smtp_port` `integer`
- `smtp_username` `string`

---

### `GetUserActivationResp`

- `activation_url` `string`

---

### `SendUserActivationReq`

- `user_id` `string` **(必填)**

---

### `SiteAIReq`

- `ai_providers` []``SiteAIProvider``
- `chosen_provider` `string`
- `enabled` `boolean`
- `prompt_config` `AIPromptConfig`

---

### `SiteAIResp`

- `ai_providers` []``SiteAIProvider``
- `chosen_provider` `string`
- `enabled` `boolean`
- `prompt_config` `AIPromptConfig`

---

### `SiteAdvancedReq`

- `authorized_attachment_extensions` []``string``
- `authorized_image_extensions` []``string``
- `max_attachment_size` `integer`
- `max_image_megapixel` `integer`
- `max_image_size` `integer`

---

### `SiteAdvancedResp`

- `authorized_attachment_extensions` []``string``
- `authorized_image_extensions` []``string``
- `max_attachment_size` `integer`
- `max_image_megapixel` `integer`
- `max_image_size` `integer`

---

### `SiteBrandingReq`

- `favicon` `string`
- `logo` `string`
- `mobile_logo` `string`
- `square_icon` `string`

---

### `SiteBrandingResp`

- `favicon` `string`
- `logo` `string`
- `mobile_logo` `string`
- `square_icon` `string`

---

### `SiteCustomCssHTMLReq`

- `custom_css` `string`
- `custom_footer` `string`
- `custom_head` `string`
- `custom_header` `string`
- `custom_sidebar` `string`

---

### `SiteCustomCssHTMLResp`

- `custom_css` `string`
- `custom_footer` `string`
- `custom_head` `string`
- `custom_header` `string`
- `custom_sidebar` `string`

---

### `SiteGeneralReq`

- `contact_email` `string` **(必填)**
- `description` `string`
- `name` `string` **(必填)**
- `short_description` `string`
- `site_url` `string` **(必填)**

---

### `SiteGeneralResp`

- `contact_email` `string` **(必填)**
- `description` `string`
- `name` `string` **(必填)**
- `short_description` `string`
- `site_url` `string` **(必填)**

---

### `SiteInterfaceReq`

- `language` `string` **(必填)**
- `time_zone` `string` **(必填)**

---

### `SiteInterfaceSettingsResp`

- `language` `string` **(必填)**
- `time_zone` `string` **(必填)**

---

### `SiteLoginReq`

- `allow_email_domains` []``string``
- `allow_email_registrations` `boolean`
- `allow_new_registrations` `boolean`
- `allow_password_login` `boolean`

---

### `SiteLoginResp`

- `allow_email_domains` []``string``
- `allow_email_registrations` `boolean`
- `allow_new_registrations` `boolean`
- `allow_password_login` `boolean`

---

### `SiteMCPReq`

- `enabled` `boolean`

---

### `SiteMCPResp`

- `enabled` `boolean`
- `http_header` `string`
- `type` `string`
- `url` `string`

---

### `SitePoliciesReq`

- `privacy_policy_original_text` `string`
- `privacy_policy_parsed_text` `string`
- `terms_of_service_original_text` `string`
- `terms_of_service_parsed_text` `string`

---

### `SitePoliciesResp`

- `privacy_policy_original_text` `string`
- `privacy_policy_parsed_text` `string`
- `terms_of_service_original_text` `string`
- `terms_of_service_parsed_text` `string`

---

### `SiteQuestionsReq`

- `min_content` `integer`
- `min_tags` `integer`
- `restrict_answer` `boolean`

---

### `SiteQuestionsResp`

- `min_content` `integer`
- `min_tags` `integer`
- `restrict_answer` `boolean`

---

### `SiteSecurityReq`

- `check_update` `boolean`
- `external_content_display` `string` enum: `always_display`, `ask_before_display` **(必填)**
- `login_required` `boolean`

---

### `SiteSecurityResp`

- `check_update` `boolean`
- `external_content_display` `string` enum: `always_display`, `ask_before_display` **(必填)**
- `login_required` `boolean`

---

### `SiteSeoReq`

- `permalink` `integer` **(必填)**
- `robots` `string` **(必填)**

---

### `SiteSeoResp`

- `permalink` `integer` **(必填)**
- `robots` `string` **(必填)**

---

### `SiteTagsReq`

- `recommend_tags` []``SiteWriteTag``
- `required_tag` `boolean`
- `reserved_tags` []``SiteWriteTag``

---

### `SiteTagsResp`

- `recommend_tags` []``SiteWriteTag``
- `required_tag` `boolean`
- `reserved_tags` []``SiteWriteTag``

---

### `SiteThemeReq`

- `color_scheme` `string`
- `layout` `string` enum: `Full-width`, `Fixed-width`
- `theme` `string` **(必填)**
- `theme_config` `object`

---

### `SiteThemeResp`

- `color_scheme` `string`
- `layout` `string`
- `theme` `string`
- `theme_config` `object`
- `theme_options` []``ThemeOption``

---

### `SiteUsersReq`

- `allow_update_avatar` `boolean`
- `allow_update_bio` `boolean`
- `allow_update_display_name` `boolean`
- `allow_update_location` `boolean`
- `allow_update_username` `boolean`
- `allow_update_website` `boolean`
- `default_avatar` `string` enum: `system`, `gravatar` **(必填)**
- `gravatar_base_url` `string`

---

### `SiteUsersResp`

- `allow_update_avatar` `boolean`
- `allow_update_bio` `boolean`
- `allow_update_display_name` `boolean`
- `allow_update_location` `boolean`
- `allow_update_username` `boolean`
- `allow_update_website` `boolean`
- `default_avatar` `string` enum: `system`, `gravatar` **(必填)**
- `gravatar_base_url` `string`

---

### `SiteUsersSettingsReq`

- `default_avatar` `string` enum: `system`, `gravatar` **(必填)**
- `gravatar_base_url` `string`

---

### `SiteUsersSettingsResp`

- `default_avatar` `string` enum: `system`, `gravatar` **(必填)**
- `gravatar_base_url` `string`

---

### `UpdateAPIKeyReq`

- `description` `string` **(必填)**
- `id` `integer` **(必填)**

---

### `UpdatePrivilegesConfigReq`

- `custom_privileges` []``Privilege``
- `level` `PrivilegeLevel` **(必填)**

---

### `UpdateSMTPConfigReq`

- `encryption` `string` enum: `SSL`, `TLS` — "" SSL TLS
- `from_email` `string`
- `from_name` `string`
- `smtp_authentication` `boolean`
- `smtp_host` `string`
- `smtp_password` `string`
- `smtp_port` `integer`
- `smtp_username` `string`
- `test_email_recipient` `string`

---

### `UpdateUserPasswordReq`

- `password` `string` **(必填)**
- `user_id` `string` **(必填)**

---

### `UpdateUserRoleReq`

- `role_id` `integer` **(必填)** — role id
- `user_id` `string` **(必填)** — user id

---

### `UpdateUserStatusReq`

- `remove_all_content` `boolean`
- `status` `string` enum: `normal`, `suspended`, `deleted`, `inactive` **(必填)**
- `suspend_duration` `string` enum: `24h`, `48h`, `72h`, `7d`, `14d`, `1m`, `2m`, `3m`, `6m`, `1y`, `forever`
- `user_id` `string` **(必填)**

---

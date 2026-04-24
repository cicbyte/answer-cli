# site — 数据模型

### `RespBody`

- `code` `integer` — http code
- `data`  — response data
- `msg` `string` — response message
- `reason` `string` — reason key

---

### `GetSiteLegalInfoResp`

- `privacy_policy_original_text` `string`
- `privacy_policy_parsed_text` `string`
- `terms_of_service_original_text` `string`
- `terms_of_service_parsed_text` `string`

---

### `SiteInfoResp`

- `ai_enabled` `boolean`
- `branding` `SiteBrandingResp`
- `custom_css_html` `SiteCustomCssHTMLResp`
- `general` `SiteGeneralResp`
- `interface` `SiteInterfaceSettingsResp`
- `login` `SiteLoginResp`
- `mcp_enabled` `boolean`
- `revision` `string`
- `site_advanced` `SiteAdvancedResp`
- `site_legal` `SiteLegalSimpleResp`
- `site_questions` `SiteQuestionsResp`
- `site_security` `SiteSecurityResp`
- `site_seo` `SiteSeoResp`
- `site_tags` `SiteTagsResp`
- `site_users` `SiteUsersResp`
- `theme` `SiteThemeResp`
- `users_settings` `SiteUsersSettingsResp`
- `version` `string`

---

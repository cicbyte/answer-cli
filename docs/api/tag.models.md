# Tag — 数据模型

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

### `AddTagReq`

- `display_name` `string` **(必填)** — display_name
- `original_text` `string` **(必填)** — original text
- `slug_name` `string` **(必填)** — slug_name

---

### `GetFollowingTagsResp`

- `display_name` `string` — display name
- `main_tag_slug_name` `string` — if main tag slug name is not empty, this tag is synonymous with the main tag
- `recommend` `boolean`
- `reserved` `boolean`
- `slug_name` `string` — slug name
- `tag_id` `string` — tag id

---

### `GetTagBasicResp`

- `display_name` `string`
- `recommend` `boolean`
- `reserved` `boolean`
- `slug_name` `string`
- `tag_id` `string`

---

### `GetTagResp`

- `created_at` `integer`
- `description` `string`
- `display_name` `string`
- `excerpt` `string`
- `follow_count` `integer`
- `is_follower` `boolean`
- `main_tag_slug_name` `string` — if main tag slug name is not empty, this tag is synonymous with the main tag
- `member_actions` []``PermissionMemberAction``
- `original_text` `string`
- `parsed_text` `string`
- `question_count` `integer`
- `recommend` `boolean`
- `reserved` `boolean`
- `slug_name` `string`
- `status` `string`
- `tag_id` `string`
- `updated_at` `integer`

---

### `GetTagSynonymsResp`

- `member_actions` []``PermissionMemberAction`` — MemberActions
- `synonyms` []``TagSynonym`` — synonyms

---

### `RecoverTagReq`

- `tag_id` `string` **(必填)**

---

### `RemoveTagReq`

- `tag_id` `string` **(必填)** — tag_id

---

### `UpdateTagReq`

- `display_name` `string` — display_name
- `edit_summary` `string` — edit summary
- `original_text` `string` — original text
- `slug_name` `string` — slug_name
- `tag_id` `string` **(必填)** — tag_id

---

### `UpdateTagSynonymReq`

- `synonym_tag_list` []``TagItem`` **(必填)** — synonym tag list
- `tag_id` `string` **(必填)** — tag_id

---

# Revision — 数据模型

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

### `GetReviewingTypeResp`

- `label` `string`
- `name` `string`
- `todo_amount` `integer`

---

### `GetRevisionResp`

- `content` 
- `create_at` `integer`
- `id` `string`
- `object_id` `string`
- `reason` `string`
- `status` `integer`
- `title` `string`
- `url_title` `string`
- `use_id` `string`
- `user_info` `UserBasicInfo`

---

### `RevisionAuditReq`

- `id` `string` **(必填)** — object id
- `operation` `string` **(必填)** — approve or reject

---

# Report — 数据模型

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

### `AddReportReq`

- `captcha_code` `string`
- `captcha_id` `string` — captcha_id
- `content` `string` — report content
- `object_id` `string` **(必填)** — object id
- `report_type` `integer` **(必填)** — report type

---

### `ReviewReportReq`

- `close_msg` `string`
- `close_type` `integer`
- `content` `string`
- `flag_id` `string` **(必填)**
- `operation_type` `string` enum: `edit_post`, `close_post`, `delete_post`, `unlist_post`, `ignore_report` **(必填)**
- `tags` []``TagItem``
- `title` `string`

---

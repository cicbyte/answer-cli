# Review — 数据模型

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

### `UpdateReviewReq`

- `review_id` `integer` **(必填)**
- `status` `string` enum: `approve`, `reject` **(必填)**

---

# ai-conversation-admin — 数据模型

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

### `AIConversationAdminDeleteReq`

- `conversation_id` `string` **(必填)**

---

### `AIConversationAdminDetailResp`

- `conversation_id` `string`
- `created_at` `integer`
- `records` []``AIConversationRecord``
- `topic` `string`
- `user_info` `AIConversationUserInfo`

---

# ai-conversation — 数据模型

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

### `AIConversationDetailResp`

- `conversation_id` `string`
- `created_at` `integer`
- `records` []``AIConversationRecord``
- `topic` `string`
- `updated_at` `integer`

---

### `AIConversationVoteReq`

- `cancel` `boolean`
- `chat_completion_id` `string` **(必填)**
- `vote_type` `string` enum: `helpful`, `unhelpful` **(必填)**

---

# Meta — 数据模型

### `RespBody`

- `code` `integer` — http code
- `data`  — response data
- `msg` `string` — response message
- `reason` `string` — reason key

---

### `ReactionRespItem`

- `count` `integer` — Count is the number of users who reacted
- `emoji` `string` — Emoji is the reaction emoji
- `is_active` `boolean` — IsActive is if current user has reacted
- `tooltip` `string` — Tooltip is the user's name who reacted

---

### `UpdateReactionReq`

- `emoji` `string` enum: `heart`, `smile`, `frown` **(必填)**
- `object_id` `string` **(必填)**
- `reaction` `string` enum: `activate`, `deactivate` **(必填)**

---

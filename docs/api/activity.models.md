# Activity — 数据模型

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

### `FollowReq`

- `is_cancel` `boolean` — is cancel
- `object_id` `string` **(必填)** — object id

---

### `FollowResp`

- `follows` `integer` — the followers of object
- `is_followed` `boolean` — if user is followed object will be true,otherwise false

---

### `UpdateFollowTagsReq`

- `slug_name_list` []``string`` — tag slug name list

---

### `VoteReq`

- `captcha_code` `string`
- `captcha_id` `string`
- `is_cancel` `boolean`
- `object_id` `string` **(必填)**

---

### `VoteResp`

- `down_votes` `integer`
- `up_votes` `integer`
- `vote_status` `string`
- `votes` `integer`

---

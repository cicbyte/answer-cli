# Comment — 数据模型

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

### `AddCommentReq`

- `captcha_code` `string`
- `captcha_id` `string`
- `mention_username_list` []``string`` — @ user id list
- `object_id` `string` **(必填)** — object id
- `original_text` `string` **(必填)** — original comment content
- `reply_comment_id` `string` — reply comment id

---

### `GetCommentResp`

- `comment_id` `string` — comment id
- `created_at` `integer` — create time
- `is_vote` `boolean` — current user if already vote this comment
- `member_actions` []``PermissionMemberAction`` — MemberActions
- `object_id` `string` — object id
- `original_text` `string` — original comment content
- `parsed_text` `string` — parsed comment content
- `reply_comment_id` `string` — reply comment id
- `reply_user_display_name` `string` — reply user display name
- `reply_user_id` `string` — reply user id
- `reply_user_status` `string` — reply user status
- `reply_username` `string` — reply user username
- `user_avatar` `string` — user avatar
- `user_display_name` `string` — user display name
- `user_id` `string` — user id
- `user_status` `string` — user status
- `username` `string` — username
- `vote_count` `integer` — user vote amount

---

### `GetObjectTimelineResp`

- `object_info` `ActObjectInfo`
- `timeline` []``ActObjectTimeline``

---

### `RemoveCommentReq`

- `captcha_code` `string`
- `captcha_id` `string`
- `comment_id` `string` **(必填)** — comment id

---

### `UpdateCommentReq`

- `captcha_code` `string`
- `captcha_id` `string` — whether user can delete it
- `comment_id` `string` **(必填)** — comment id
- `original_text` `string` **(必填)** — original comment content

---

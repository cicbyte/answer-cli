# Question — 数据模型

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

### `CloseQuestionReq`

- `close_msg` `string` — close_type
- `close_type` `integer` — close_type
- `id` `string` **(必填)**

---

### `OperationQuestionReq`

- `id` `string` **(必填)**
- `operation` `string` — operation [pin unpin hide show]

---

### `QuestionAdd`

- `captcha_code` `string`
- `captcha_id` `string` — captcha_id
- `content` `string` — content
- `tags` []``TagItem`` — tags
- `title` `string` **(必填)** — question title

---

### `QuestionAddByAnswer`

- `answer_content` `string` **(必填)**
- `captcha_code` `string`
- `captcha_id` `string` — captcha_id
- `content` `string` — content
- `mention_username_list` []``string``
- `tags` []``TagItem`` — tags
- `title` `string` **(必填)** — question title

---

### `QuestionPageReq`

- `in_days` `integer`
- `order` `string` enum: `newest`, `active`, `hot`, `score`, `unanswered`, `recommend`, `frequent`
- `page` `integer`
- `page_size` `integer`
- `tag` `string`
- `username` `string`

---

### `QuestionRecoverReq`

- `question_id` `string` **(必填)**

---

### `QuestionUpdate`

- `captcha_code` `string`
- `captcha_id` `string` — captcha_id
- `content` `string` — content
- `edit_summary` `string` — edit summary
- `id` `string` **(必填)** — question id
- `invite_user` []``string``
- `tags` []``TagItem`` — tags
- `title` `string` **(必填)** — question title

---

### `QuestionUpdateInviteUser`

- `captcha_code` `string`
- `captcha_id` `string` — captcha_id
- `id` `string` **(必填)**
- `invite_user` []``string``

---

### `RemoveQuestionReq`

- `captcha_code` `string`
- `captcha_id` `string` — captcha_id
- `id` `string` **(必填)** — question id

---

### `ReopenQuestionReq`

- `question_id` `string`

---

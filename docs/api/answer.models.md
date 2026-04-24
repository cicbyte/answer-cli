# Answer — 数据模型

### `RespBody`

- `code` `integer` — http code
- `data`  — response data
- `msg` `string` — response message
- `reason` `string` — reason key

---

### `AcceptAnswerReq`

- `answer_id` `string`
- `question_id` `string` **(必填)**

---

### `AnswerAddReq`

- `captcha_code` `string`
- `captcha_id` `string`
- `content` `string` **(必填)**
- `question_id` `string`

---

### `AnswerUpdateReq`

- `captcha_code` `string`
- `captcha_id` `string`
- `content` `string` **(必填)**
- `edit_summary` `string`
- `id` `string`
- `title` `string`

---

### `GetAnswerInfoResp`

- `info` `AnswerInfo`
- `question` `QuestionInfoResp`

---

### `RecoverAnswerReq`

- `answer_id` `string` **(必填)**

---

### `RemoveAnswerReq`

- `captcha_code` `string`
- `captcha_id` `string`
- `id` `string` **(必填)**

---

# Question

## `GET` `/personal/qa/top`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `username` (`string` **(必填)**) — username

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/question`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `QuestionUpdate`
> question
  - `captcha_code` `string`
  - `captcha_id` `string` — captcha_id
  - `content` `string` — content
  - `edit_summary` `string` — edit summary
  - `id` `string` **(必填)** — question id
  - `invite_user` []``string``
  - `tags` []``TagItem`` — tags
  - `title` `string` **(必填)** — question title

**响应:**

- `200`:
`RespBody`

---

## `POST` `/question`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `QuestionAdd`
> question
  - `captcha_code` `string`
  - `captcha_id` `string` — captcha_id
  - `content` `string` — content
  - `tags` []``TagItem`` — tags
  - `title` `string` **(必填)** — question title

**响应:**

- `200`:
`RespBody`

---

## `DELETE` `/question`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `RemoveQuestionReq`
> question
  - `captcha_code` `string`
  - `captcha_id` `string` — captcha_id
  - `id` `string` **(必填)** — question id

**响应:**

- `200`:
`RespBody`

---

## `POST` `/question/answer`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `QuestionAddByAnswer`
> question
  - `answer_content` `string` **(必填)**
  - `captcha_code` `string`
  - `captcha_id` `string` — captcha_id
  - `content` `string` — content
  - `mention_username_list` []``string``
  - `tags` []``TagItem`` — tags
  - `title` `string` **(必填)** — question title

**响应:**

- `200`:
`RespBody`

---

## `GET` `/question/info`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `id` (`string` **(必填)**) — Question TagID

**响应:**

- `200`:
`string`

---

## `GET` `/question/invite`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `id` (`string` **(必填)**) — Question ID

**响应:**

- `200`:
`string`

---

## `PUT` `/question/invite`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `QuestionUpdateInviteUser`
> question
  - `captcha_code` `string`
  - `captcha_id` `string` — captcha_id
  - `id` `string` **(必填)**
  - `invite_user` []``string``

**响应:**

- `200`:
`RespBody`

---

## `GET` `/question/link`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `in_days` (`integer`)
  - `order` (`string`)
  - `page` (`integer`)
  - `page_size` (`integer`)
  - `question_id` (`string` **(必填)**)

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `list` 

---

## `PUT` `/question/operation`

**认证**: 管理员 | **参数方式**: body

> Operation question \n operation [pin unpin hide show]

**请求参数:**

**Body**: `OperationQuestionReq`
> question
  - `id` `string` **(必填)**
  - `operation` `string` — operation [pin unpin hide show]

**响应:**

- `200`:
`RespBody`

---

## `GET` `/question/page`

**认证**: 公开 | **参数方式**: body

**请求参数:**

**Body**: `QuestionPageReq`
> QuestionPageReq
  - `in_days` `integer`
  - `order` `string` enum: `newest`, `active`, `hot`, `score`, `unanswered`, `recommend`, `frequent`
  - `page` `integer`
  - `page_size` `integer`
  - `tag` `string`
  - `username` `string`

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `list` 

---

## `GET` `/question/recommend/page`

**认证**: 公开 | **参数方式**: body

**请求参数:**

**Body**: `QuestionPageReq`
> QuestionPageReq
  - `in_days` `integer`
  - `order` `string` enum: `newest`, `active`, `hot`, `score`, `unanswered`, `recommend`, `frequent`
  - `page` `integer`
  - `page_size` `integer`
  - `tag` `string`
  - `username` `string`

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `list` 

---

## `POST` `/question/recover`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `QuestionRecoverReq`
> question
  - `question_id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/question/reopen`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `ReopenQuestionReq`
> question
  - `question_id` `string`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/question/similar`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `title` (`string` **(必填)**) — title

**响应:**

- `200`:
`RespBody`

---

## `GET` `/question/similar/tag`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `question_id` (`string` **(必填)**) — question_id

**响应:**

- `200`:
`string`

---

## `PUT` `/question/status`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `CloseQuestionReq`
> question
  - `close_msg` `string` — close_type
  - `close_type` `integer` — close_type
  - `id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

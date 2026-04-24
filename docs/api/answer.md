# Answer

## `PUT` `/answer`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `AnswerUpdateReq`
> AnswerUpdateReq
  - `captcha_code` `string`
  - `captcha_id` `string`
  - `content` `string` **(必填)**
  - `edit_summary` `string`
  - `id` `string`
  - `title` `string`

**响应:**

- `200`:
`RespBody`

---

## `POST` `/answer`

**认证**: 管理员 | **参数方式**: body

> add answer

**请求参数:**

**Body**: `AnswerAddReq`
> add answer request
  - `captcha_code` `string`
  - `captcha_id` `string`
  - `content` `string` **(必填)**
  - `question_id` `string`

**响应:**

- `200`:
`RespBody`

---

## `DELETE` `/answer`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `RemoveAnswerReq`
> answer
  - `captcha_code` `string`
  - `captcha_id` `string`
  - `id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `POST` `/answer/acceptance`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `AcceptAnswerReq`
> AcceptAnswerReq
  - `answer_id` `string`
  - `question_id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/answer/info`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `id` (`string` **(必填)**) — id

**响应:**

- `200`:
`RespBody`
  - `data` `GetAnswerInfoResp`

---

## `GET` `/answer/page`

**认证**: 公开 | **参数方式**: query

> AnswerList <br> <b>order</b> (default or updated)

**请求参数:**

  - `question_id` (`string` **(必填)**) — question_id
  - `order` (`string` **(必填)**) — order
  - `page` (`string` **(必填)**) — page
  - `page_size` (`string` **(必填)**) — page_size

**响应:**

- `200`:
`string`

---

## `POST` `/answer/recover`

**认证**: 管理员 | **参数方式**: body

> recover the deleted answer

**请求参数:**

**Body**: `RecoverAnswerReq`
> answer
  - `answer_id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

# Report

## `POST` `/report`

**认证**: 管理员 | **参数方式**: body

> add report <br> source (question, answer, comment, user)

**请求参数:**

**Body**: `AddReportReq`
> report
  - `captcha_code` `string`
  - `captcha_id` `string` — captcha_id
  - `content` `string` — report content
  - `object_id` `string` **(必填)** — object id
  - `report_type` `integer` **(必填)** — report type

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/report/review`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `ReviewReportReq`
> flag
  - `close_msg` `string`
  - `close_type` `integer`
  - `content` `string`
  - `flag_id` `string` **(必填)**
  - `operation_type` `string` enum: `edit_post`, `close_post`, `delete_post`, `unlist_post`, `ignore_report` **(必填)**
  - `tags` []``TagItem``
  - `title` `string`

**响应:**

- `200`:
`RespBody`

---

## `GET` `/report/unreviewed/post`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `page` (`integer`) — page

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `list` 

---

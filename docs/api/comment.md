# Comment

## `GET` `/activity/timeline`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `object_id` (`string`) — object id
  - `tag_slug_name` (`string`) — tag slug name
  - `object_type` (`string`) — object type
  - `show_vote` (`boolean`) — is show vote

**响应:**

- `200`:
`RespBody`
  - `data` `GetObjectTimelineResp`

---

## `GET` `/activity/timeline/detail`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `revision_id` (`string` **(必填)**) — revision id

**响应:**

- `200`:
`RespBody`
  - `data` `GetObjectTimelineResp`

---

## `GET` `/comment`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `id` (`string` **(必填)**) — id

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `list` 

---

## `PUT` `/comment`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateCommentReq`
> comment
  - `captcha_code` `string`
  - `captcha_id` `string` — whether user can delete it
  - `comment_id` `string` **(必填)** — comment id
  - `original_text` `string` **(必填)** — original comment content

**响应:**

- `200`:
`RespBody`

---

## `POST` `/comment`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `AddCommentReq`
> comment
  - `captcha_code` `string`
  - `captcha_id` `string`
  - `mention_username_list` []``string`` — @ user id list
  - `object_id` `string` **(必填)** — object id
  - `original_text` `string` **(必填)** — original comment content
  - `reply_comment_id` `string` — reply comment id

**响应:**

- `200`:
`RespBody`
  - `data` `GetCommentResp`

---

## `DELETE` `/comment`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `RemoveCommentReq`
> comment
  - `captcha_code` `string`
  - `captcha_id` `string`
  - `comment_id` `string` **(必填)** — comment id

**响应:**

- `200`:
`RespBody`

---

## `GET` `/comment/page`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `page` (`integer`) — page
  - `page_size` (`integer`) — page size
  - `object_id` (`string` **(必填)**) — object id
  - `query_cond` (`string`) — query condition

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `list` 

---

## `GET` `/personal/comment/page`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `page` (`integer`) — page
  - `page_size` (`integer`) — page size
  - `username` (`string`) — username

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `list` 

---

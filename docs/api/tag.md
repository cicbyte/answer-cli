# Tag

## `GET` `/question/tags`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `tag` (`string`) — tag

**响应:**

- `200`:
`RespBody`
  - `data` []``GetTagBasicResp``

---

## `GET` `/tag`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `tag_id` (`string` **(必填)**) — tag id
  - `tag_name` (`string` **(必填)**) — tag name

**响应:**

- `200`:
`RespBody`
  - `data` `GetTagResp`

---

## `PUT` `/tag`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateTagReq`
> tag
  - `display_name` `string` — display_name
  - `edit_summary` `string` — edit summary
  - `original_text` `string` — original text
  - `slug_name` `string` — slug_name
  - `tag_id` `string` **(必填)** — tag_id

**响应:**

- `200`:
`RespBody`

---

## `POST` `/tag`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `AddTagReq`
> tag
  - `display_name` `string` **(必填)** — display_name
  - `original_text` `string` **(必填)** — original text
  - `slug_name` `string` **(必填)** — slug_name

**响应:**

- `200`:
`RespBody`

---

## `DELETE` `/tag`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `RemoveTagReq`
> tag
  - `tag_id` `string` **(必填)** — tag_id

**响应:**

- `200`:
`RespBody`

---

## `POST` `/tag/merge`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `AddTagReq`
> tag
  - `display_name` `string` **(必填)** — display_name
  - `original_text` `string` **(必填)** — original text
  - `slug_name` `string` **(必填)** — slug_name

**响应:**

- `200`:
`RespBody`

---

## `POST` `/tag/recover`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `RecoverTagReq`
> tag
  - `tag_id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/tag/synonym`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateTagSynonymReq`
> tag
  - `synonym_tag_list` []``TagItem`` **(必填)** — synonym tag list
  - `tag_id` `string` **(必填)** — tag_id

**响应:**

- `200`:
`RespBody`

---

## `GET` `/tag/synonyms`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `tag_id` (`integer` **(必填)**) — tag id

**响应:**

- `200`:
`RespBody`
  - `data` `GetTagSynonymsResp`

---

## `GET` `/tags`

**认证**: 公开 | **参数方式**: query

> get tags list by slug name

**请求参数:**

  - `tags` (`array`) — string collection

**响应:**

- `200`:
`RespBody`
  - `data` []``GetTagBasicResp``

---

## `GET` `/tags/following`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` []``GetFollowingTagsResp``

---

## `GET` `/tags/page`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `page` (`integer`) — page size
  - `page_size` (`integer`) — page size
  - `slug_name` (`string`) — slug_name
  - `query_cond` (`string`) — query condition

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `list` 

---

# Activity

## `POST` `/follow`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `FollowReq`
> follow
  - `is_cancel` `boolean` — is cancel
  - `object_id` `string` **(必填)** — object id

**响应:**

- `200`:
`RespBody`
  - `data` `FollowResp`

---

## `PUT` `/follow/tags`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateFollowTagsReq`
> follow
  - `slug_name_list` []``string`` — tag slug name list

**响应:**

- `200`:
`RespBody`

---

## `GET` `/personal/vote/page`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `page` (`integer`) — page size
  - `page_size` (`integer`) — page size

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `list` 

---

## `POST` `/vote/down`

**认证**: 管理员 | **参数方式**: body

> add vote

**请求参数:**

**Body**: `VoteReq`
> vote
  - `captcha_code` `string`
  - `captcha_id` `string`
  - `is_cancel` `boolean`
  - `object_id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`
  - `data` `VoteResp`

---

## `POST` `/vote/up`

**认证**: 管理员 | **参数方式**: body

> add vote

**请求参数:**

**Body**: `VoteReq`
> vote
  - `captcha_code` `string`
  - `captcha_id` `string`
  - `is_cancel` `boolean`
  - `object_id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`
  - `data` `VoteResp`

---

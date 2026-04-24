# Review

## `PUT` `/review/pending/post`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateReviewReq`
> review
  - `review_id` `integer` **(必填)**
  - `status` `string` enum: `approve`, `reject` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/review/pending/post/page`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `page` (`integer`) — page
  - `object_id` (`string`) — object_id

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `list` 

---

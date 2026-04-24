# api-badge

## `GET` `/badge`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `id` (`string` **(必填)**) — id

**响应:**

- `200`:
`RespBody`
  - `data` `GetBadgeInfoResp`

---

## `GET` `/badge/awards/page`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `page` (`integer`) — page
  - `page_size` (`integer`) — page size
  - `badge_id` (`string` **(必填)**) — badge id
  - `username` (`string`) — only list the award by username

**响应:**

- `200`:
`RespBody`
  - `data` `GetBadgeInfoResp`

---

## `GET` `/badge/user/awards`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `username` (`string` **(必填)**) — user name

**响应:**

- `200`:
`RespBody`
  - `data` []``GetUserBadgeAwardListResp``

---

## `GET` `/badge/user/awards/recent`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `username` (`string` **(必填)**) — user name

**响应:**

- `200`:
`RespBody`
  - `data` []``GetUserBadgeAwardListResp``

---

## `GET` `/badges`

**认证**: 公开

**响应:**

- `200`:
`RespBody`
  - `data` []``GetBadgeListResp``

---

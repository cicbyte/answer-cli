# Search

## `GET` `/search`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `q` (`string` **(必填)**) — query string
  - `order` (`string` **(必填)**) — order

**响应:**

- `200`:
`RespBody`
  - `data` `SearchResp`

---

## `GET` `/search/desc`

**认证**: 公开

**响应:**

- `200`:
`RespBody`
  - `data` `SearchResp`

---

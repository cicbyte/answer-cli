# AdminBadge

## `PUT` `/badge/status`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateBadgeStatusReq`
> UpdateBadgeStatusReq
  - `id` `string` **(必填)** — badge id
  - `status` `BadgeStatus` **(必填)** — badge status

**响应:**

- `200`:
`RespBody`

---

## `GET` `/badges`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `page` (`integer`) — page
  - `page_size` (`integer`) — page size
  - `status` (`string`) — badge status
  - `q` (`string`) — search param

**响应:**

- `200`:
`RespBody`
  - `data` []``GetBadgeListPagedResp``

---

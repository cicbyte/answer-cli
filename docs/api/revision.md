# Revision

## `GET` `/reviewing/type`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` []``GetReviewingTypeResp``

---

## `GET` `/revisions`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `object_id` (`string` **(必填)**) — object id

**响应:**

- `200`:
`RespBody`
  - `data` []``GetRevisionResp``

---

## `PUT` `/revisions/audit`

**认证**: 管理员 | **参数方式**: body

> revision audit operation:approve or reject

**请求参数:**

**Body**: `RevisionAuditReq`
> audit
  - `id` `string` **(必填)** — object id
  - `operation` `string` **(必填)** — approve or reject

**响应:**

- `200`:
`RespBody`

---

## `GET` `/revisions/edit/check`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `id` (`string` **(必填)**) — id

**响应:**

- `200`:
`RespBody`

---

## `GET` `/revisions/unreviewed`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `page` (`string` **(必填)**) — page id

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `list` 

---

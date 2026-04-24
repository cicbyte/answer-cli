# site

## `GET` `/siteinfo`

**认证**: 公开

**响应:**

- `200`:
`RespBody`
  - `data` `SiteInfoResp`

---

## `GET` `/siteinfo/legal`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `info_type` (`string` **(必填)**) — legal information type

**响应:**

- `200`:
`RespBody`
  - `data` `GetSiteLegalInfoResp`

---

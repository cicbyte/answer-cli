# Upload

## `POST` `/file`

**认证**: 管理员 | **参数方式**: mixed

**请求参数:**

  - `source` (`string` **(必填)**) — identify the source of the file upload
  - `file` (`file` **(必填)**) — file

**响应:**

- `200`:
`RespBody`
  - `data` `string`

---

## `POST` `/post/render`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `PostRenderReq`
> PostRenderReq
  - `content` `string`

**响应:**

- `200`:
`RespBody`

---

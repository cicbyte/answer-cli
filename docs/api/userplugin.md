# UserPlugin

## `GET` `/user/plugin/config`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `plugin_slug_name` (`string` **(必填)**) — plugin_slug_name

**响应:**

- `200`:
`RespBody`
  - `data` `GetPluginConfigResp`

---

## `PUT` `/user/plugin/config`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateUserPluginConfigReq`
> UpdatePluginConfigReq
  - `config_fields` `object`
  - `plugin_slug_name` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/user/plugin/configs`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` []``GetUserPluginListResp``

---

# AdminPlugin

## `GET` `/plugin/config`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `plugin_slug_name` (`string` **(必填)**) — plugin_slug_name

**响应:**

- `200`:
`RespBody`
  - `data` `GetPluginConfigResp`

---

## `PUT` `/plugin/config`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdatePluginConfigReq`
> UpdatePluginConfigReq
  - `config_fields` `object`
  - `plugin_slug_name` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/plugin/status`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdatePluginStatusReq`
> UpdatePluginStatusReq
  - `enabled` `boolean`
  - `plugin_slug_name` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/plugins`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `status` (`string`) — status: active/inactive
  - `have_config` (`boolean`) — have config

**响应:**

- `200`:
`RespBody`
  - `data` []``GetPluginListResp``

---

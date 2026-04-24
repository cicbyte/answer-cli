# PluginConnector

## `POST` `/connector/binding/email`

**认证**: 公开 | **参数方式**: body

**请求参数:**

**Body**: `ExternalLoginBindingUserSendEmailReq`
> external login binding user send email
  - `binding_key` `string` **(必填)**
  - `email` `string` **(必填)**
  - `must` `boolean` — If must is true, whatever email if exists, try to bind user.
If must is false, when email exist, will only be prompted with a warning.

**响应:**

- `200`:
`RespBody`
  - `data` `ExternalLoginBindingUserSendEmailResp`

---

## `GET` `/connector/info`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` []``ConnectorInfoResp``

---

## `GET` `/connector/user/info`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` []``ConnectorUserInfoResp``

---

## `DELETE` `/connector/user/unbinding`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `ExternalLoginUnbindingReq`
> ExternalLoginUnbindingReq
  - `external_id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

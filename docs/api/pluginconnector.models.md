# PluginConnector — 数据模型

### `RespBody`

- `code` `integer` — http code
- `data`  — response data
- `msg` `string` — response message
- `reason` `string` — reason key

---

### `ConnectorInfoResp`

- `icon` `string`
- `link` `string`
- `name` `string`

---

### `ConnectorUserInfoResp`

- `binding` `boolean`
- `external_id` `string`
- `icon` `string`
- `link` `string`
- `name` `string`

---

### `ExternalLoginBindingUserSendEmailReq`

- `binding_key` `string` **(必填)**
- `email` `string` **(必填)**
- `must` `boolean` — If must is true, whatever email if exists, try to bind user.
If must is false, when email exist, will only be prompted with a warning.

---

### `ExternalLoginBindingUserSendEmailResp`

- `access_token` `string`
- `email_exist_and_must_be_confirmed` `boolean`

---

### `ExternalLoginUnbindingReq`

- `external_id` `string` **(必填)**

---

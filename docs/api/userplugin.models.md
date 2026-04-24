# UserPlugin — 数据模型

### `RespBody`

- `code` `integer` — http code
- `data`  — response data
- `msg` `string` — response message
- `reason` `string` — reason key

---

### `GetPluginConfigResp`

- `config_fields` []``ConfigField``
- `description` `string`
- `name` `string`
- `slug_name` `string`
- `version` `string`

---

### `GetUserPluginListResp`

- `name` `string`
- `slug_name` `string`

---

### `UpdateUserPluginConfigReq`

- `config_fields` `object`
- `plugin_slug_name` `string` **(必填)**

---

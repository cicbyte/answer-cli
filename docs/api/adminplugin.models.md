# AdminPlugin — 数据模型

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

### `GetPluginListResp`

- `description` `string`
- `enabled` `boolean`
- `have_config` `boolean`
- `link` `string`
- `name` `string`
- `slug_name` `string`
- `version` `string`

---

### `UpdatePluginConfigReq`

- `config_fields` `object`
- `plugin_slug_name` `string` **(必填)**

---

### `UpdatePluginStatusReq`

- `enabled` `boolean`
- `plugin_slug_name` `string` **(必填)**

---

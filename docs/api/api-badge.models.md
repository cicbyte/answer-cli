# api-badge — 数据模型

### `RespBody`

- `code` `integer` — http code
- `data`  — response data
- `msg` `string` — response message
- `reason` `string` — reason key

---

### `GetBadgeInfoResp`

- `award_count` `integer` — badge award count
- `description` `string` — badge description
- `earned_count` `integer` — badge earned count
- `icon` `string` — badge icon
- `id` `string` — badge id
- `is_single` `boolean` — badge is single or multiple
- `level` `BadgeLevel` — badge level
- `name` `string` — badge name

---

### `GetBadgeListResp`

- `badges` []``BadgeListInfo`` — badge list info
- `group_name` `string` — badge group name

---

### `GetUserBadgeAwardListResp`

- `earned_count` `integer` — badge award count
- `icon` `string` — badge icon
- `id` `string` — badge id
- `level` `BadgeLevel` — badge level
- `name` `string` — badge name

---

# AdminBadge — 数据模型

### `RespBody`

- `code` `integer` — http code
- `data`  — response data
- `msg` `string` — response message
- `reason` `string` — reason key

---

### `GetBadgeListPagedResp`

- `award_count` `integer` — badge award count
- `description` `string` — badge description
- `earned` `boolean` — badge earned count
- `group_name` `string` — badge group name
- `icon` `string` — badge icon
- `id` `string` — badge id
- `level` `BadgeLevel` — badge level
- `name` `string` — badge name
- `status` `BadgeStatus` — badge status

---

### `UpdateBadgeStatusReq`

- `id` `string` **(必填)** — badge id
- `status` `BadgeStatus` **(必填)** — badge status

---

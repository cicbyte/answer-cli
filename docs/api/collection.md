# Collection

## `POST` `/collection/switch`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `CollectionSwitchReq`
> collection
  - `bookmark` `boolean`
  - `group_id` `string` **(必填)**
  - `object_id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`
  - `data` `CollectionSwitchResp`

---

## `GET` `/personal/collection/page`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `page` (`string` **(必填)**) — page
  - `page_size` (`string` **(必填)**) — page_size

**响应:**

- `200`:
`RespBody`

---

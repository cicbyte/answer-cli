# Meta

## `GET` `/meta/reaction`

**认证**: 公开 | **参数方式**: query

> get reaction for an object

**请求参数:**

  - `object_id` (`string` **(必填)**) — object_id

**响应:**

- `200`:
`RespBody`
  - `data` `ReactionRespItem`

---

## `PUT` `/meta/reaction`

**认证**: 管理员 | **参数方式**: body

> update reaction. if not exist, add one

**请求参数:**

**Body**: `UpdateReactionReq`
> reaction
  - `emoji` `string` enum: `heart`, `smile`, `frown` **(必填)**
  - `object_id` `string` **(必填)**
  - `reaction` `string` enum: `activate`, `deactivate` **(必填)**

**响应:**

- `200`:
`RespBody`

---

# Notification

## `GET` `/notification/page`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `page` (`integer`) — page size
  - `page_size` (`integer`) — page size
  - `type` (`string` **(必填)**) — type
  - `inbox_type` (`string` **(必填)**) — inbox_type

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/notification/read/state`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `NotificationClearIDRequest`
> NotificationClearIDRequest
  - `id` `string`

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/notification/read/state/all`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `NotificationClearRequest`
> NotificationClearRequest
  - `type` `string` enum: `inbox`, `achievement` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/notification/status`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/notification/status`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `NotificationClearRequest`
> NotificationClearRequest
  - `type` `string` enum: `inbox`, `achievement` **(必填)**

**响应:**

- `200`:
`RespBody`

---

# ai-conversation-admin

## `GET` `/ai/conversation`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `conversation_id` (`string` **(必填)**) — conversation id

**响应:**

- `200`:
`RespBody`
  - `data` `AIConversationAdminDetailResp`

---

## `DELETE` `/ai/conversation`

**认证**: 公开 | **参数方式**: body

> delete conversation and its related records for admin

**请求参数:**

**Body**: `AIConversationAdminDeleteReq`
> apikey
  - `conversation_id` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `GET` `/ai/conversation/page`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `page` (`integer`) — page
  - `page_size` (`integer`) — page size

**响应:**

- `200`:
`RespBody`
  - `data` `PageModel`
  - `list` 

---

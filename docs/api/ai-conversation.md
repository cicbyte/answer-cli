# ai-conversation

## `GET` `/ai/conversation`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `conversation_id` (`string` **(必填)**) — conversation id

**响应:**

- `200`:
`RespBody`
  - `data` `AIConversationDetailResp`

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

## `POST` `/ai/conversation/vote`

**认证**: 公开 | **参数方式**: body

**请求参数:**

**Body**: `AIConversationVoteReq`
> vote request
  - `cancel` `boolean`
  - `chat_completion_id` `string` **(必填)**
  - `vote_type` `string` enum: `helpful`, `unhelpful` **(必填)**

**响应:**

- `200`:
`RespBody`

---

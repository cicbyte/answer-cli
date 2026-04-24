# User

## `GET` `/personal/user/info`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `username` (`string` **(必填)**) — username

**响应:**

- `200`:
`RespBody`
  - `data` `GetOtherUserInfoResp`

---

## `GET` `/user/action/record`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `action` (`string` **(必填)**) — action

**响应:**

- `200`:
`RespBody`
  - `data` `ActionRecordResp`

---

## `PUT` `/user/email`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UserChangeEmailVerifyReq`
> UserChangeEmailVerifyReq
  - `code` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `POST` `/user/email/change/code`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UserChangeEmailSendCodeReq`
> UserChangeEmailSendCodeReq
  - `captcha_code` `string`
  - `captcha_id` `string`
  - `e_mail` `string` **(必填)**
  - `pass` `string`

**响应:**

- `200`:
`RespBody`

---

## `POST` `/user/email/verification`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `code` (`string` **(必填)**) — code

**响应:**

- `200`:
`RespBody`
  - `data` `UserLoginResp`

---

## `POST` `/user/email/verification/send`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `captcha_id` (`string`) — captcha_id
  - `captcha_code` (`string`) — captcha_code

**响应:**

- `200`:
`string`

---

## `GET` `/user/info`

**认证**: 管理员

> get user info, if user no login response http code is 200, but user info is null

**响应:**

- `200`:
`RespBody`
  - `data` `GetCurrentLoginUserInfoResp`

---

## `PUT` `/user/info`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

  - `Authorization` (`string` **(必填)**) — access-token
  - `data` (`string` **(必填)**) — UpdateInfoRequest

**响应:**

- `200`:
`RespBody`

---

## `GET` `/user/info/search`

**认证**: 管理员 | **参数方式**: query

**请求参数:**

  - `username` (`string` **(必填)**) — username

**响应:**

- `200`:
`RespBody`
  - `data` `GetOtherUserInfoResp`

---

## `PUT` `/user/interface`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

  - `Authorization` (`string` **(必填)**) — access-token
  - `data` (`string` **(必填)**) — UpdateInfoRequest

**响应:**

- `200`:
`RespBody`

---

## `POST` `/user/login/email`

**认证**: 公开 | **参数方式**: body

**请求参数:**

**Body**: `UserEmailLoginReq`
> UserEmailLogin
  - `captcha_code` `string`
  - `captcha_id` `string`
  - `e_mail` `string` **(必填)**
  - `pass` `string` **(必填)**

**响应:**

- `200`:
`RespBody`
  - `data` `UserLoginResp`

---

## `GET` `/user/logout`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/user/notification/config`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UpdateUserNotificationConfigReq`
> UpdateUserNotificationConfigReq
  - `all_new_question` `NotificationChannelConfig`
  - `all_new_question_for_following_tags` `NotificationChannelConfig`
  - `inbox` `NotificationChannelConfig`

**响应:**

- `200`:
`RespBody`

---

## `POST` `/user/notification/config`

**认证**: 管理员

**响应:**

- `200`:
`RespBody`
  - `data` `GetUserNotificationConfigResp`

---

## `PUT` `/user/notification/unsubscribe`

**认证**: 公开 | **参数方式**: body

**请求参数:**

**Body**: `UserUnsubscribeNotificationReq`
> UserUnsubscribeNotificationReq
  - `code` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `PUT` `/user/password`

**认证**: 管理员 | **参数方式**: body

**请求参数:**

**Body**: `UserModifyPasswordReq`
> UserModifyPasswordReq
  - `captcha_code` `string`
  - `captcha_id` `string`
  - `old_pass` `string`
  - `pass` `string` **(必填)**

**响应:**

- `200`:
`RespBody`

---

## `POST` `/user/password/replacement`

**认证**: 公开 | **参数方式**: body

**请求参数:**

**Body**: `UserRePassWordRequest`
> UserRePassWordRequest
  - `code` `string` **(必填)**
  - `pass` `string` **(必填)**

**响应:**

- `200`:
`string`

---

## `POST` `/user/password/reset`

**认证**: 公开 | **参数方式**: body

**请求参数:**

**Body**: `UserRetrievePassWordRequest`
> UserRetrievePassWordRequest
  - `captcha_code` `string`
  - `captcha_id` `string`
  - `e_mail` `string` **(必填)**

**响应:**

- `200`:
`string`

---

## `GET` `/user/ranking`

**认证**: 公开

**响应:**

- `200`:
`RespBody`
  - `data` `UserRankingResp`

---

## `POST` `/user/register/email`

**认证**: 公开 | **参数方式**: body

**请求参数:**

**Body**: `UserRegisterReq`
> UserRegisterReq
  - `captcha_code` `string`
  - `captcha_id` `string`
  - `e_mail` `string` **(必填)**
  - `name` `string` **(必填)**
  - `pass` `string` **(必填)**

**响应:**

- `200`:
`RespBody`
  - `data` `UserLoginResp`

---

## `GET` `/user/staff`

**认证**: 公开 | **参数方式**: query

**请求参数:**

  - `username` (`string` **(必填)**) — username
  - `page_size` (`string` **(必填)**) — page_size

**响应:**

- `200`:
`RespBody`
  - `data` `GetUserStaffResp`

---

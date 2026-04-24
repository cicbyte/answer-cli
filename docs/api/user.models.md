# User — 数据模型

### `RespBody`

- `code` `integer` — http code
- `data`  — response data
- `msg` `string` — response message
- `reason` `string` — reason key

---

### `ActionRecordResp`

- `captcha_id` `string`
- `captcha_img` `string`
- `verify` `boolean`

---

### `GetCurrentLoginUserInfoResp`

- `access_token` `string` — access token
- `answer_count` `integer` — answer count
- `authority_group` `integer` — authority group
- `avatar` `AvatarInfo`
- `bio` `string` — bio markdown
- `bio_html` `string` — bio html
- `color_scheme` `string` — Color scheme
- `created_at` `integer` — create time
- `display_name` `string` — display name
- `e_mail` `string` — email
- `follow_count` `integer` — follow count
- `have_password` `boolean` — user have password
- `id` `string` — user id
- `language` `string` — language
- `last_login_date` `integer` — last login date
- `location` `string` — location
- `mail_status` `integer` — mail status(1 pass 2 to be verified)
- `mobile` `string` — mobile
- `notice_status` `integer` — notice status(1 on 2off)
- `question_count` `integer` — question count
- `rank` `integer` — rank
- `role_id` `integer` — role id
- `status` `string` — user status
- `suspended_until` `integer` — suspended until timestamp
- `username` `string` — username
- `visit_token` `string` — visit token
- `website` `string` — website

---

### `GetOtherUserInfoResp`

- `info` `GetOtherUserInfoByUsernameResp`

---

### `GetUserNotificationConfigResp`

- `all_new_question` `NotificationChannelConfig`
- `all_new_question_for_following_tags` `NotificationChannelConfig`
- `inbox` `NotificationChannelConfig`

---

### `GetUserStaffResp`

- `avatar` `string` — avatar
- `display_name` `string` — display name
- `username` `string` — username

---

### `UpdateInfoRequest`

- `avatar` `AvatarInfo`
- `bio` `string`
- `display_name` `string`
- `location` `string`
- `username` `string`
- `website` `string`

---

### `UpdateUserInterfaceRequest`

- `color_scheme` `string` **(必填)** — Color scheme
- `language` `string` **(必填)** — language

---

### `UpdateUserNotificationConfigReq`

- `all_new_question` `NotificationChannelConfig`
- `all_new_question_for_following_tags` `NotificationChannelConfig`
- `inbox` `NotificationChannelConfig`

---

### `UserChangeEmailSendCodeReq`

- `captcha_code` `string`
- `captcha_id` `string`
- `e_mail` `string` **(必填)**
- `pass` `string`

---

### `UserChangeEmailVerifyReq`

- `code` `string` **(必填)**

---

### `UserEmailLoginReq`

- `captcha_code` `string`
- `captcha_id` `string`
- `e_mail` `string` **(必填)**
- `pass` `string` **(必填)**

---

### `UserLoginResp`

- `access_token` `string` — access token
- `answer_count` `integer` — answer count
- `authority_group` `integer` — authority group
- `avatar` `string` — avatar
- `bio` `string` — bio markdown
- `bio_html` `string` — bio html
- `color_scheme` `string` — Color scheme
- `created_at` `integer` — create time
- `display_name` `string` — display name
- `e_mail` `string` — email
- `follow_count` `integer` — follow count
- `have_password` `boolean` — user have password
- `id` `string` — user id
- `language` `string` — language
- `last_login_date` `integer` — last login date
- `location` `string` — location
- `mail_status` `integer` — mail status(1 pass 2 to be verified)
- `mobile` `string` — mobile
- `notice_status` `integer` — notice status(1 on 2off)
- `question_count` `integer` — question count
- `rank` `integer` — rank
- `role_id` `integer` — role id
- `status` `string` — user status
- `suspended_until` `integer` — suspended until timestamp
- `username` `string` — username
- `visit_token` `string` — visit token
- `website` `string` — website

---

### `UserModifyPasswordReq`

- `captcha_code` `string`
- `captcha_id` `string`
- `old_pass` `string`
- `pass` `string` **(必填)**

---

### `UserRankingResp`

- `staffs` []``UserRankingSimpleInfo``
- `users_with_the_most_reputation` []``UserRankingSimpleInfo``
- `users_with_the_most_vote` []``UserRankingSimpleInfo``

---

### `UserRePassWordRequest`

- `code` `string` **(必填)**
- `pass` `string` **(必填)**

---

### `UserRegisterReq`

- `captcha_code` `string`
- `captcha_id` `string`
- `e_mail` `string` **(必填)**
- `name` `string` **(必填)**
- `pass` `string` **(必填)**

---

### `UserRetrievePassWordRequest`

- `captcha_code` `string`
- `captcha_id` `string`
- `e_mail` `string` **(必填)**

---

### `UserUnsubscribeNotificationReq`

- `code` `string` **(必填)**

---

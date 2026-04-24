# Apache Answer API Reference

共 190 个端点，204 个数据模型

## Activity (5)

- 🔑`POST` `/follow` — follow object or cancel follow operation
- 🔑`PUT` `/follow/tags` — update user follow tags
- 🔑`GET` `/personal/vote/page` — get user personal votes
- 🔑`POST` `/vote/down` — vote down
- 🔑`POST` `/vote/up` — vote up

## AdminBadge (2)

- 🔑`PUT` `/badge/status` — update badge status
- 🔑`GET` `/badges` — list all badges by page

## AdminPlugin (4)

- 🔑`GET` `/plugin/config` — get plugin config
- 🔑`PUT` `/plugin/config` — update plugin config
- 🔑`PUT` `/plugin/status` — update plugin status
- 🔑`GET` `/plugins` — get plugin list

## Answer (7)

- 🔑`PUT` `/answer` — Update Answer
- 🔑`POST` `/answer` — Add Answer
- 🔑`DELETE` `/answer` — delete answer
- 🔑`POST` `/answer/acceptance` — Accept Answer
- `GET` `/answer/info` — Get Answer Detail
- `GET` `/answer/page` — AnswerList
- 🔑`POST` `/answer/recover` — recover answer

## Collection (2)

- 🔑`POST` `/collection/switch` — add collection
- 🔑`GET` `/personal/collection/page` — list personal collections

## Comment (8)

- `GET` `/activity/timeline` — get object timeline
- `GET` `/activity/timeline/detail` — get object timeline detail
- `GET` `/comment` — get comment by id
- 🔑`PUT` `/comment` — update comment
- 🔑`POST` `/comment` — add comment
- 🔑`DELETE` `/comment` — remove comment
- `GET` `/comment/page` — get comment page
- `GET` `/personal/comment/page` — user personal comment list

## Lang (3)

- 🔑`GET` `/language/options` — Get language options
- `GET` `/language/config` — get language config mapping
- `GET` `/language/options` — Get language options

## Meta (2)

- `GET` `/meta/reaction` — get reaction
- 🔑`PUT` `/meta/reaction` — add or update reaction

## Notification (5)

- 🔑`GET` `/notification/page` — get notification list
- 🔑`PUT` `/notification/read/state` — ClearUnRead
- 🔑`PUT` `/notification/read/state/all` — ClearUnRead
- 🔑`GET` `/notification/status` — GetRedDot
- 🔑`PUT` `/notification/status` — DelRedDot

## Permission (1)

- 🔑`GET` `/permission` — check user permission

## Personal (1)

- 🔑`GET` `/personal/answer/page` — list personal answers

## Plugin (2)

- `GET` `/embed/config` — get embed plugin config
- `GET` `/plugin/status` — get all plugins status

## PluginConnector (4)

- `POST` `/connector/binding/email` — external login binding user send email
- 🔑`GET` `/connector/info` — get all enabled connectors
- 🔑`GET` `/connector/user/info` — get all connectors info about user
- 🔑`DELETE` `/connector/user/unbinding` — unbind external user login

## PluginRender (1)

- `GET` `/render/config` — GetRenderConfig

## Question (17)

- `GET` `/personal/qa/top` — UserTop
- 🔑`PUT` `/question` — update question
- 🔑`POST` `/question` — add question
- 🔑`DELETE` `/question` — delete question
- 🔑`POST` `/question/answer` — add question and answer
- `GET` `/question/info` — get question details
- `GET` `/question/invite` — get question invite user info
- 🔑`PUT` `/question/invite` — update question invite user
- `GET` `/question/link` — get question link
- 🔑`PUT` `/question/operation` — Operation question
- `GET` `/question/page` — get questions by page
- `GET` `/question/recommend/page` — get recommend questions by page
- 🔑`POST` `/question/recover` — recover deleted question
- 🔑`PUT` `/question/reopen` — reopen question
- 🔑`GET` `/question/similar` — fuzzy query similar questions based on title
- `GET` `/question/similar/tag` — Search Similar Question
- 🔑`PUT` `/question/status` — Close question

## Rank (1)

- `GET` `/personal/rank/page` — user personal rank list

## Report (3)

- 🔑`POST` `/report` — add report
- 🔑`PUT` `/report/review` — review report
- 🔑`GET` `/report/unreviewed/post` — get unreviewed report post page

## Review (2)

- 🔑`PUT` `/review/pending/post` — update review
- 🔑`GET` `/review/pending/post/page` — get unreviewed post page

## Revision (5)

- 🔑`GET` `/reviewing/type` — get reviewing type
- `GET` `/revisions` — get revision list
- 🔑`PUT` `/revisions/audit` — revision audit
- 🔑`GET` `/revisions/edit/check` — check can update revision
- 🔑`GET` `/revisions/unreviewed` — get unreviewed revision list

## Search (2)

- 🔑`GET` `/search` — search object
- `GET` `/search/desc` — get search description

## Tag (12)

- 🔑`GET` `/question/tags` — get tag list
- `GET` `/tag` — get tag one
- 🔑`PUT` `/tag` — update tag
- 🔑`POST` `/tag` — add tag
- 🔑`DELETE` `/tag` — delete tag
- 🔑`POST` `/tag/merge` — merge tag
- 🔑`POST` `/tag/recover` — recover delete tag
- 🔑`PUT` `/tag/synonym` — update tag
- `GET` `/tag/synonyms` — get tag synonyms
- `GET` `/tags` — get tags list
- 🔑`GET` `/tags/following` — get following tag list
- `GET` `/tags/page` — get tag page

## Upload (2)

- 🔑`POST` `/file` — upload file
- 🔑`POST` `/post/render` — render post content

## User (21)

- 🔑`GET` `/personal/user/info` — GetOtherUserInfoByUsername
- 🔑`GET` `/user/action/record` — ActionRecord
- 🔑`PUT` `/user/email` — user change email verification
- 🔑`POST` `/user/email/change/code` — send email to the user email then change their email
- `POST` `/user/email/verification` — UserVerifyEmail
- 🔑`POST` `/user/email/verification/send` — UserVerifyEmailSend
- 🔑`GET` `/user/info` — GetUserInfoByUserID
- 🔑`PUT` `/user/info` — UserUpdateInfo update user info
- 🔑`GET` `/user/info/search` — SearchUserListByName
- 🔑`PUT` `/user/interface` — UserUpdateInterface update user interface config
- `POST` `/user/login/email` — UserEmailLogin
- 🔑`GET` `/user/logout` — user logout
- 🔑`PUT` `/user/notification/config` — update user's notification config
- 🔑`POST` `/user/notification/config` — get user's notification config
- `PUT` `/user/notification/unsubscribe` — unsubscribe notification
- 🔑`PUT` `/user/password` — UserModifyPassWord
- `POST` `/user/password/replacement` — UseRePassWord
- `POST` `/user/password/reset` — RetrievePassWord
- `GET` `/user/ranking` — get user ranking
- `POST` `/user/register/email` — UserRegisterByEmail
- `GET` `/user/staff` — get user staff

## UserPlugin (3)

- 🔑`GET` `/user/plugin/config` — get user plugin config
- 🔑`PUT` `/user/plugin/config` — update user plugin config
- 🔑`GET` `/user/plugin/configs` — get plugin list that used for user.

## admin (59)

- 🔑`GET` `/ai-config` — get AI configuration
- 🔑`PUT` `/ai-config` — update AI configuration
- 🔑`POST` `/ai-models` — get AI models
- 🔑`GET` `/ai-provider` — get AI provider configuration
- 🔑`GET` `/answer/page` — AdminAnswerPage admin answer page
- 🔑`PUT` `/answer/status` — update answer status
- 🔑`PUT` `/api-key` — update apikey
- 🔑`POST` `/api-key` — add apikey
- 🔑`DELETE` `/api-key` — delete apikey
- 🔑`GET` `/api-key/all` — get all api keys
- 🔑`GET` `/dashboard` — DashboardInfo
- 🔑`DELETE` `/delete/permanently` — delete permanently
- 🔑`GET` `/mcp-config` — get MCP configuration
- 🔑`PUT` `/mcp-config` — update MCP configuration
- 🔑`GET` `/question/page` — AdminQuestionPage admin question page
- 🔑`PUT` `/question/status` — update question status
- 🔑`GET` `/roles` — get role list
- 🔑`GET` `/setting/privileges` — GetPrivilegesConfig get privileges config
- 🔑`PUT` `/setting/privileges` — update privileges config
- 🔑`GET` `/setting/smtp` — GetSMTPConfig get smtp config
- 🔑`PUT` `/setting/smtp` — update smtp config
- 🔑`GET` `/siteinfo/advanced` — get site advanced setting
- 🔑`PUT` `/siteinfo/advanced` — update site advanced info
- 🔑`GET` `/siteinfo/branding` — get site interface
- 🔑`PUT` `/siteinfo/branding` — update site info branding
- 🔑`GET` `/siteinfo/custom-css-html` — get site info custom html css config
- 🔑`PUT` `/siteinfo/custom-css-html` — update site custom css html config
- 🔑`GET` `/siteinfo/general` — get site general information
- 🔑`PUT` `/siteinfo/general` — update site general information
- 🔑`GET` `/siteinfo/interface` — get site interface
- 🔑`PUT` `/siteinfo/interface` — update site info interface
- 🔑`GET` `/siteinfo/login` — get site info login config
- 🔑`PUT` `/siteinfo/login` — update site login
- 🔑`GET` `/siteinfo/polices` — Get the policies information for the site
- 🔑`PUT` `/siteinfo/polices` — update site policies configuration
- 🔑`GET` `/siteinfo/question` — get site questions setting
- 🔑`PUT` `/siteinfo/question` — update site question settings
- 🔑`GET` `/siteinfo/security` — Get the security information for the site
- 🔑`PUT` `/siteinfo/security` — update site security configuration
- 🔑`GET` `/siteinfo/seo` — get site seo information
- 🔑`PUT` `/siteinfo/seo` — update site seo information
- 🔑`GET` `/siteinfo/tag` — get site tags setting
- 🔑`PUT` `/siteinfo/tag` — update site tag settings
- 🔑`GET` `/siteinfo/theme` — get site info theme config
- 🔑`PUT` `/siteinfo/theme` — update site custom css html config
- 🔑`GET` `/siteinfo/users` — get site user config
- 🔑`PUT` `/siteinfo/users` — update site info config about users
- 🔑`GET` `/siteinfo/users-settings` — get site interface
- 🔑`PUT` `/siteinfo/users-settings` — update site info users settings
- 🔑`GET` `/theme/options` — Get theme options
- 🔑`POST` `/user` — add user
- 🔑`GET` `/user/activation` — get user activation
- 🔑`PUT` `/user/password` — update user password
- 🔑`PUT` `/user/profile` — edit user profile
- 🔑`PUT` `/user/role` — update user role
- 🔑`PUT` `/user/status` — update user
- 🔑`POST` `/users` — add users
- 🔑`POST` `/users/activation` — send user activation
- 🔑`GET` `/users/page` — get user page

## ai-conversation (3)

- `GET` `/ai/conversation` — get conversation detail
- `GET` `/ai/conversation/page` — get conversation list
- `POST` `/ai/conversation/vote` — vote record

## ai-conversation-admin (3)

- `GET` `/ai/conversation` — get conversation detail for admin
- `DELETE` `/ai/conversation` — delete conversation for admin
- `GET` `/ai/conversation/page` — get conversation list for admin

## api-badge (5)

- `GET` `/badge` — get badge info
- `GET` `/badge/awards/page` — get badge award list
- `GET` `/badge/user/awards` — get user badge award list
- `GET` `/badge/user/awards/recent` — get user badge award list
- `GET` `/badges` — list all badges group by group

## installation (1)

- `GET` `/` — if config file not exist try to redirect to install page

## reason (2)

- 🔑`GET` `/reasons` — get reasons by object type and action
- 🔑`GET` `/reasons` — get reasons by object type and action

## site (2)

- `GET` `/siteinfo` — get site info
- `GET` `/siteinfo/legal` — get site legal info

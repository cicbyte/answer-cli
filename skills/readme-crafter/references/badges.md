# 徽章目录

常用徽章按类别整理。只使用能提供真实信息的徽章。

## 构建与 CI

| 徽章 | Markdown |
|---|---|
| GitHub Actions | `![Build](https://img.shields.io/github/actions/workflow/status/{owner}/{repo}/{workflow-file}?branch=main)` |
| Travis CI | `![Build](https://img.shields.io/travis/{owner}/{repo})` |
| CircleCI | `![CircleCI](https://img.shields.io/circleci/build/github/{owner}/{repo})` |
| Jenkins | `![Jenkins](https://img.shields.io/jenkins/build?jobUrl={url})` |

## 代码质量

| 徽章 | Markdown |
|---|---|
| 覆盖率 | `![Coverage](https://img.shields.io/codecov/c/github/{owner}/{repo})` |
| 覆盖率 (Codecov) | `![Coverage](https://codecov.io/gh/{owner}/{repo}/branch/main/graph/badge.svg)` |
| 代码检查 | `![Lint](https://img.shields.io/github/actions/workflow/status/{owner}/{repo}/lint.yml)` |

## 版本与发布

| 徽章 | Markdown |
|---|---|
| npm 版本 | `![npm](https://img.shields.io/npm/v/{package})` |
| PyPI 版本 | `![PyPI](https://img.shields.io/pypi/v/{package})` |
| Docker 拉取量 | `![Docker Pulls](https://img.shields.io/docker/pulls/{image})` |
| GitHub Release | `![Release](https://img.shields.io/github/v/release/{owner}/{repo})` |
| Maven Central | `![Maven Central](https://img.shields.io/maven-central/v/{groupId}/{artifactId})` |
| Go Report Card | `![Go Report Card](https://goreportcard.com/badge/github/{owner}/{repo})` |
| Crates.io | `![Crates.io](https://img.shields.io/crates/v/{crate})` |
| RubyGem | `![Gem](https://img.shields.io/gem/v/{gem})` |

## 开源许可证

| 徽章 | Markdown |
|---|---|
| MIT | `![MIT](https://img.shields.io/github/license/{owner}/{repo})` |
| Apache 2.0 | `![Apache-2.0](https://img.shields.io/github/license/{owner}/{repo})` |
| GPL v3 | `![GPL-3.0](https://img.shields.io/github/license/{owner}/{repo})` |

## 平台与统计

| 徽章 | Markdown |
|---|---|
| Star 数 | `![Stars](https://img.shields.io/github/stars/{owner}/{repo}?style=social)` |
| 下载量 (npm) | `![npm downloads](https://img.shields.io/npm/dm/{package})` |
| 下载量 (PyPI) | `![PyPI - Downloads](https://img.shields.io/pypi/dm/{package})` |
| 主语言 | `![Language](https://img.shields.io/github/languages/top/{owner}/{repo})` |
| 最近提交 | `![Last Commit](https://img.shields.io/github/last-commit/{owner}/{repo})` |

## 使用准则

- **单行最多 5 个徽章** — 只选最有意义的
- **优先级顺序**：构建状态 > 版本号 > 许可证 > 下载量 > Star 数
- **不要放**：尚不存在的指标徽章（空 CI、零下载）
- **风格**：使用默认（flat）样式；除非项目非常热门否则避免 `for-the-badge`
- **位置**：徽章放在项目标题正下方，排列在同一行

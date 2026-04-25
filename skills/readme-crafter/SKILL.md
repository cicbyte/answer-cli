---
name: readme-crafter
description: 为任何项目创建或优化高质量的 README.md 文件。当用户要求创建、编写、改进、润色或重新生成 README 时触发；当用户说"帮我写个 README"、"优化我的 README"、"写个项目文档"或类似表述时触发。也适用于用户希望让项目在 GitHub/GitLab 上更专业、更有吸引力时使用。
---

# Readme Crafter

生成专业、精致的 README.md 文件，突出项目核心价值，吸引目标用户。

## 调用方式

```
/readme-crafter                    # 默认：自动检测语言，无额外特性
/readme-crafter --lang zh          # 指定中文
/readme-crafter --lang zh,en       # 中英双语
/readme-crafter --lang en          # 英文
/readme-crafter --lang zh --feat badges,compare,community,toc,repo-meta,diff
```

### 参数说明

| 参数 | 格式 | 说明 | 默认值 |
|---|---|---|---|
| `--lang` | 逗号分隔的语言代码（`zh`/`en`） | README 输出语言。`zh,en` 生成双语版 | 自动检测 |
| `--feat` | 逗号分隔的特性名 | 启用的可选特性（见下方列表） | 无 |

### 可选特性（`--feat`）

| 特性名 | 说明 |
|---|---|
| `auto-badges` | 从项目文件自动推断并推荐徽章（CI 平台、包管理器、许可证等） |
| `compare` | 生成与同类项目的对比表格（需用户提供竞品名称） |
| `community` | 添加社区生态章节（Discord/论坛链接、贡献者、Star History） |
| `toc` | 自动生成目录导航（README 超过 100 行时自动启用） |
| `repo-meta` | 输出 GitHub repo 元数据建议（description、topics、social image） |
| `diff` | 优化模式下展示「修改前 vs 修改后」对比，便于审查改动 |

## 工作流程

```
0. 收集参数  →  解析 --lang 和 --feat，确定语言和可选特性
1. 检测 README →  已存在？【优化模式】 : 不存在？【创建模式】
2. 分析项目  →  扫描结构、依赖、入口、测试、配置
3. 判定类型  →  库/CLI/Web 应用/框架/工具/API 服务
4. 提炼定位  →  名称、标语、核心卖点、差异化优势
5. 生成文档  →  按模板 + 可选特性生成完整 README
6. 写入结果  →  创建或更新 README.md
```

## 第零步：收集参数

解析用户提供的参数，未指定的参数使用默认值。

### 语言参数（`--lang`）

| 用户输入 | 行为 |
|---|---|
| 未指定 | 根据已有文档和项目元数据自动检测主语言 |
| `zh` | 全中文 README |
| `en` | 全英文 README |
| `zh,en` | 双语 README，使用「主语言在上 / Secondary language below」分隔 |
| `en,zh` | 双语 README，英文在上 |

### 特性参数（`--feat`）

按逗号分隔解析。每个特性对应一个可选增强模块，详见 [可选特性实现指南](#可选特性实现指南)。

## 第一步：检测已有 README

在项目根目录查找 README 文件（`README.md`、`README.rst`、`README.txt`、`README`）。

### 不存在 → 创建模式

直接进入「第二步：分析项目」，按完整流程生成全新 README。

### 已存在 → 优化模式

先读取并诊断已有 README，再针对性改进。

#### 诊断清单

逐项检查已有 README 是否存在以下问题：

| 检查项 | 问题表现 | 改进动作 |
|---|---|---|
| **首屏吸引力** | 开头没有标语；前 10 行看不出项目价值 | 补充一句话定位 + 核心卖点 |
| **功能描述** | 功能罗列笼统；缺少差异化说明 | 按主题分组，附价值说明 |
| **快速开始** | 缺失；步骤过多；无法复制粘贴即用 | 补充 3-5 步最小可运行示例 |
| **安装方式** | 只有一种；缺少包管理器/平台选项 | 补充所有适用安装途径 |
| **代码示例** | 缺失；代码块无语言标签；示例不可运行 | 补充带语言标签的可运行示例 |
| **徽章** | 缺失；过多；包含无意义徽章 | 参照 badges.md 选取最多 5 个有意义的 |
| **截图/演示** | 缺失 | 添加截图占位或实际截图 |
| **章节完整性** | 缺少配置、API、贡献指南、许可证等 | 按项目类型补充缺失章节 |
| **格式规范** | 段落过长；缺少表格；层级混乱 | 重构为可扫描的结构化格式 |
| **空泛描述** | 含"革命性"、"前沿"、"下一代"等词 | 替换为具体数据和事实 |
| **过时内容** | 版本号、命令、链接已失效 | 更新为当前实际状态 |
| **语言混杂** | 同一段落内混用多种语言 | 按指定语言统一或拆分为双语 |

#### 优化策略

1. **保留有价值内容** — 已有 README 中准确、详实的内容不应丢弃，整合到新结构中
2. **补缺失** — 按诊断清单补全缺失的章节和要素
3. **改结构** — 按项目类型对应的模板重新组织章节顺序
4. **升质量** — 将描述性文字改为代码示例；将长段落改为表格/列表
5. **去冗余** — 删除空泛营销词、过时信息、重复内容
6. **强首屏** — 确保前 10 行能瞬间传达项目价值

## 第二步：分析项目

并行读取以下文件以理解项目全貌：

| 分析维度 | 检查文件 |
|---|---|
| 项目身份 | `package.json`、`go.mod`、`Cargo.toml`、`pyproject.toml`、`setup.py`、`pom.xml`、`*.gemspec`、`composer.json` |
| 目录结构 | 目录树（前 2 层）、主入口文件 |
| 核心功能 | 源码文件、测试文件、CLI 参数解析、路由定义 |
| 技术栈 | 依赖文件、Dockerfile、docker-compose、CI 配置 |
| 配置项 | `.env.example`、配置文件、Makefile、justfile |
| 现有文档 | 已有 README、CHANGELOG、CONTRIBUTING、docs/ 目录 |
| 徽章/CI | `.github/workflows/`、`.gitlab-ci.yml`、Makefile targets |

需要回答的关键问题：
- 这个项目解决什么问题？（一句话概括）
- 目标用户是谁？
- 和同类项目相比有什么不同？
- 最简上手示例是什么？

## 第三步：判定项目类型

根据项目类型选择对应的 README 模板：

| 类型 | 模板 | 重点章节 |
|---|---|---|
| **库/SDK** | `references/template-library.md` | 安装、快速上手、用法示例、API 参考 |
| **CLI 工具** | `references/template-cli.md` | 安装、命令用法、示例、配置 |
| **Web 应用** | `references/template-webapp.md` | 功能特性、截图、部署、技术栈 |
| **框架** | `references/template-framework.md` | 设计理念、快速开始、核心概念、架构 |
| **API 服务** | `references/template-api.md` | 接口文档、认证方式、示例、SDK |
| **通用/其他** | `references/template-general.md` | 概述、功能、快速开始、使用方法 |

## 第四步：生成 README

按下方章节结构生成。不适用的章节直接省略，根据项目类型调整章节优先级。启用 `--feat` 的特性在对应位置插入。

### 章节结构（按优先级排序）

1. **标题区** — 项目名 + 一句话标语
2. **徽章** — 构建状态、版本号、许可证、覆盖率（参考 [references/badges.md](references/badges.md)）← `auto-badges` 在此生效
3. **目录** — 超过 100 行时自动生成 ← `toc` 在此生效
4. **展示区** — 截图、GIF 或在线演示链接
5. **功能特性** — 按主题分组的核心能力列表
6. **快速开始** — 最小可运行示例（最多 3-5 条命令）
7. **安装方式** — 所有安装途径（npm、pip、brew、docker 等）
8. **使用方法** — 按场景分组的代码示例
9. **与同类方案对比** — 对比表格 ← `compare` 在此生效
10. **配置说明** — 环境变量、配置文件、选项表格
11. **API/CLI 参考** — 精简的参数/选项文档（仅适用时保留）
12. **架构设计** — 简要的架构图或描述（仅复杂项目保留）
13. **社区与生态** — Discord、贡献者、Star History ← `community` 在此生效
14. **参与贡献** — 简要说明 + CONTRIBUTING.md 链接
15. **开源许可证** — SPDX 标识符 + 链接
16. **致谢** — 致谢、灵感来源（可选）

### 写作原则

- **价值先行**：前 10 行必须回答"我为什么要关注这个项目？"
- **用代码说话**：优先展示代码示例，而非长篇描述
- **段落极简**：每个章节引言最多一段话 — 然后直接进入代码/列表
- **标注语言**：代码块始终标注语言标签 ```python、```bash、```typescript
- **拒绝废话**：去掉"革命性"、"前沿"、"下一代"等空泛营销词
- **可扫描性**：对比用表格，特性用列表，而非大段文字
- **徽章克制**：单行最多 5 个徽章，只放有实际意义的
- **截图占位**：没有截图时添加 `<!-- screenshot: 在此处添加演示截图 -->`

### 多语言输出

根据 `--lang` 参数决定输出语言：

**单语言（`zh` 或 `en`）：**
- 全文使用指定语言
- 章节标题、描述、注释统一为该语言

**双语（`zh,en` 或 `en,zh`）：**
- 按参数顺序排列，主语言在前
- 使用分割线分隔两个语言版本
- 同一章节内不混用语言

```
<!-- zh,en 示例 -->
# ProjectName
> 中文标语...

（中文内容完整版）

---

> English tagline...

（英文内容完整版）
```

### 格式规范示例

```markdown
<!-- 正确：清晰的标题+标语 -->
# ProjectName
> 一句话说明它做什么以及为什么重要。

<!-- 正确：分组展示功能 -->
## 功能特性

- **零配置** — 开箱即用，无需额外设置
- **极致性能** — 比 [替代方案](link) 快 10 倍，得益于 [原因]
- **类型安全** — 完整的 TypeScript 支持，自动生成类型

<!-- 正确：最简快速开始 -->
## 快速开始

\`\`\`bash
npx create-project my-app
cd my-app
npm run dev
\`\`\`

<!-- 错误：堆砌空泛描述 -->
## 关于
本项目是一个革命性的、前沿的、下一代解决方案...
```

## 第五步：写入 README

### 创建模式

1. 从 `references/` 读取对应模板
2. 根据项目实际内容适配模板
3. 按启用的 `--feat` 特性插入额外章节
4. 按 `--lang` 指定的语言生成
5. 将完整 README.md 写入项目根目录

### 优化模式

1. 读取已有 README.md 全文
2. 对照诊断清单，列出所有需要改进的点
3. 向用户展示诊断结果
4. 如果启用了 `diff` 特性，展示「修改前 vs 修改后」对比
5. 从 `references/` 读取对应模板作为结构参考
6. 在已有内容基础上增量改写：补缺失、改结构、升质量、去冗余
7. 按启用的 `--feat` 特性插入额外章节
8. 按 `--lang` 指定的语言调整
9. 将优化后的 README.md 写回项目根目录

## 可选特性实现指南

### `auto-badges` — 自动推荐徽章

从项目文件推断并推荐徽章，而非让 Claude 从静态目录中挑选：

| 检测到的文件 | 推荐徽章 |
|---|---|
| `.github/workflows/*.yml` | GitHub Actions Build |
| `package.json` | npm version + npm downloads |
| `pyproject.toml` / `setup.py` | PyPI version |
| `go.mod` | Go Report Card |
| `Cargo.toml` | Crates.io version |
| `LICENSE`（MIT/Apache/GPL） | 许可证徽章 |
| `codecov.yml` / `.github` 中含 coverage | 覆盖率徽章 |

输出推荐的 Markdown 徽章代码，用户确认后插入。

### `compare` — 竞品对比表格

生成结构化的对比表格，帮助用户理解差异。如果用户未指定竞品，从项目描述和功能推断可能的替代品并询问确认。

```markdown
## 与同类方案对比

| 特性 | {本项目} | {竞品 A} | {竞品 B} |
|---|---|---|---|
| {维度 1} | ✅ | ✅ | ❌ |
| {维度 2} | ✅ | ❌ | ✅ |
| {维度 3} | {值} | {值} | {值} |
| {维度 4} | {值} | {值} | {值} |
```

对比维度应客观、可验证，避免主观评价。突出本项目在关键维度的优势。

### `community` — 社区生态章节

在「参与贡献」之前插入社区相关内容：

```markdown
## 社区

### 交流

- [Discord](链接) — 提问、讨论、分享
- [GitHub Discussions](链接) — 功能建议、Bug 反馈
- [论坛/邮件列表](链接) — 深度讨论

### 贡献者

感谢所有参与贡献的开发者！

<!-- all-contributors 规范 -->
<a href="https://github.com/{owner}/{repo}/graphs/contributors">
  <img src="https://contrib.rocks/image?repo={owner}/{repo}" />
</a>

### Star History

[![Star History Chart](https://api.star-history.com/svg?repos={owner}/{repo}&type=Date)](https://star-history.com/#{owner}/{repo}&Date)
```

如果项目没有 Discord/论坛链接，只生成贡献者和 Star History 部分。

### `toc` — 自动目录

README 超过 100 行时，在标题和徽章之后、正文之前插入目录：

```markdown
## 目录

- [功能特性](#功能特性)
- [快速开始](#快速开始)
- [安装方式](#安装方式)
- [使用方法](#使用方法)
- [配置说明](#配置说明)
- [参与贡献](#参与贡献)
- [开源许可证](#开源许可证)
```

使用 Markdown 锚点链接到各章节标题。章节列表根据实际生成的章节动态调整。

### `repo-meta` — GitHub 仓库元数据建议

在 README 写入完成后，额外输出一份 GitHub 仓库元数据建议，方便用户复制到 GitHub 仓库设置中：

```markdown
## GitHub 仓库元数据建议

### Description（仓库描述）
{一句话，不超过 350 字符，用于 GitHub 搜索和社交预览}

### Topics（标签）
{逗号分隔的标签，建议 5-10 个}

### Social Image（社交预览图）
{建议尺寸 1280x640，建议使用 og:image 标签}
{如果没有，提示用户可以使用 https://opengraph.xyz 生成}
```

这些信息不写入 README，而是作为附加输出展示给用户。

### `diff` — 修改前后对比

优化模式下，在执行改写之前生成对比展示：

1. 将已有 README 标记为「修改前」
2. 将优化后的内容标记为「修改后」
3. 用表格或并排格式展示关键改动点：

```
### 诊断结果与改进计划

| 改进项 | 修改前 | 修改后 |
|---|---|---|
| 标题区 | 缺少标语 | 添加一句话定位 |
| 徽章 | 无 | 添加 Build + License + Version |
| 快速开始 | 缺失 | 补充 3 步最小示例 |
| ... | ... | ... |
```

用户确认后再执行实际写入。

## 资源文件

### references/
- [template-library.md](references/template-library.md) — 库/SDK 模板
- [template-cli.md](references/template-cli.md) — CLI 工具模板
- [template-webapp.md](references/template-webapp.md) — Web 应用模板
- [template-framework.md](references/template-framework.md) — 框架模板
- [template-api.md](references/template-api.md) — API 服务模板
- [template-general.md](references/template-general.md) — 通用模板
- [badges.md](references/badges.md) — 徽章目录与使用指南

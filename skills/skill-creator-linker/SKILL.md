---
name: skill-creator-linker
description: 创建项目级 skill 并链接到 Claude Code。基于 skill-creator，skill 文件存放在项目 skills/ 目录（git 版本控制），通过 junction 链接到 .claude/skills/ 供 Claude Code 使用。在以下场景触发：(1) 用户想创建新的项目级 skill，(2) 用户想将已有 .claude/skills 中的 skill 迁移到项目目录，(3) 用户想更新或迭代已有的项目级 skill
license: Complete terms in LICENSE.txt
---

# Skill Creator Linker

创建项目级 skill，文件存放在 `skills/<name>/` 随 git 提交，通过 junction 链接到 `.claude/skills/<name>/` 供 Claude Code 使用。

## 目录结构

```
project/
├── skills/                    ← git 版本控制，可分发
│   └── my-skill/
│       ├── SKILL.md
│       └── references/
└── .claude/
    └── skills/
        └── my-skill           ← junction → skills/my-skill/
```

编辑任一处，两边同步更新。

## 创建流程

### Step 1: 理解 skill 需求

明确 skill 的用途和触发场景。需要了解：
- skill 解决什么问题
- 具体使用示例
- 什么情况下应触发此 skill

### Step 2: 规划内容

分析示例，确定需要哪些可复用资源：
- **scripts/** — 需要确定性执行或频繁重写的代码
- **references/** — 需要按需加载的参考文档、API 文档、领域知识
- **assets/** — 输出中使用的模板、图片等

### Step 3: 创建并链接

```bash
# 1. 在项目 skills/ 目录创建 skill 结构
mkdir -p skills/<name>/references

# 2. 编写 SKILL.md 和资源文件（见下方规范）

# 3. 创建 junction 链接（Windows，无需管理员权限）
powershell -Command "New-Item -ItemType Junction -Path '.claude\skills\<name>' -Target 'skills\<name>'"
```

### Step 4: 编写 SKILL.md

#### Frontmatter（YAML）

```yaml
---
name: skill-name
description: 一句话描述 + 触发场景列表。这是 Claude 判断是否加载 skill 的唯一依据。
---
```

- `name` 和 `description` 是必填项
- `description` 必须包含完整的触发条件，不要在 body 中写"何时使用"
- 不要添加其他 frontmatter 字段

#### Body（Markdown）

- 始终使用祈使句/不定式
- 只写 Claude 不具备的知识，默认 Claude 已经很聪明
- 保持简洁，优先用示例替代冗长解释
- 控制在 500 行以内，超出则拆分到 references/

### Step 5: 迭代

使用 skill 处理真实任务 → 发现不足 → 更新 SKILL.md 或资源文件 → 再次测试。

## Skill 设计原则

### 渐进式加载

三级加载机制管理上下文开销：
1. **Metadata**（name + description）— 始终在上下文中
2. **SKILL.md body** — skill 触发时加载
3. **references/** — Claude 按需加载

### 适度自由度

- **高自由度**（文本指令）：多种方法都可行时
- **中自由度**（伪代码/带参数的脚本）：有偏好模式但允许变化时
- **低自由度**（具体脚本，少参数）：操作脆弱、必须一致时

### 拆分模式

**按领域拆分**：skill 涉及多个独立领域时
```
skill/
├── SKILL.md          # 概览 + 导航
└── references/
    ├── domain-a.md   # 领域 A
    └── domain-b.md   # 领域 B
```

**按变体拆分**：skill 支持多种框架/方案时
```
skill/
├── SKILL.md          # 核心流程 + 选择指引
└── references/
    ├── variant-a.md  # 方案 A
    └── variant-b.md  # 方案 B
```

**条件详情**：基础内容放 SKILL.md，高级内容放 references
```markdown
## 基础用法
[简要说明]

## 高级功能
- **功能 A**: 见 [references/advanced.md](references/advanced.md)
```

### 不要包含的文件

- README.md、CHANGELOG.md、INSTALLATION_GUIDE.md 等辅助文档
- skill 只包含 AI 执行任务所需的信息

## 迁移已有 skill

将 `.claude/skills/` 中的 skill 迁移到项目目录：

```bash
# 1. 复制到项目 skills/ 目录
cp -r .claude/skills/<name> skills/<name>

# 2. 删除原目录
rm -rf .claude/skills/<name>

# 3. 创建 junction 链接
powershell -Command "New-Item -ItemType Junction -Path '.claude\skills\<name>' -Target 'skills\<name>'"
```

## 注意事项

- 使用 **Junction**（非 Symlink），Windows 下无需管理员权限
- `.claude/` 通常在 `.gitignore` 中，junction 本身不会被 git 追踪
- `skills/` 目录应纳入 git 版本控制
- 删除 junction 用 `rmdir` 或 PowerShell `Remove-Item`，不要用 `rm -rf` 递归删除
- references 文件保持一级深度，避免嵌套引用
- 超过 100 行的 reference 文件在顶部添加目录

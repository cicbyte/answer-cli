---
name: skill-linker
description: 将项目级 skill 链接到 .claude/skills 目录，实现 git 版本控制与 Claude Code 同时使用。在以下场景触发：(1) 创建了新的项目级 skill 需要链接到 .claude/skills，(2) 需要让 skill 随项目分发同时能在 Claude Code 中生效
---

# Skill Linker

项目级 skill 存放于 `skills/<name>/`（随 git 提交），通过 junction 链接到 `.claude/skills/<name>/` 供 Claude Code 使用。

## 工作原理

```
skills/answer-cli/           ← 实际文件，git 版本控制，可分发
├── SKILL.md
└── references/

.claude/skills/answer-cli    ← junction 链接，指向 skills/answer-cli/
```

编辑任一处，两边同步更新。

## 操作步骤

### 1. 创建新 skill 并链接

```bash
# 在项目 skills/ 目录创建 skill
mkdir -p skills/<name>
# 编写 SKILL.md ...

# 创建 junction 链接（Windows，无需管理员权限）
powershell -Command "New-Item -ItemType Junction -Path '.claude\skills\<name>' -Target 'skills\<name>'"
```

### 2. 已有 .claude/skills 中的 skill 迁移到项目目录

```bash
# 移动到项目 skills/ 目录
cp -r .claude/skills/<name> skills/<name>

# 删除原目录
rm -rf .claude/skills/<name>

# 创建 junction 链接
powershell -Command "New-Item -ItemType Junction -Path '.claude\skills\<name>' -Target 'skills\<name>'"
```

### 3. 验证

```bash
ls .claude/skills/<name>/SKILL.md   # 应能正常访问
```

## 注意事项

- 使用 **Junction**（非 Symlink），Windows 下无需管理员权限
- `.claude/` 目录通常在 `.gitignore` 中，junction 本身不会被 git 追踪
- `skills/` 目录应纳入 git 版本控制
- 删除 junction 使用 `rmdir`（不加斜杠）或 PowerShell `Remove-Item`，不要用 `rm -rf` 递归删除

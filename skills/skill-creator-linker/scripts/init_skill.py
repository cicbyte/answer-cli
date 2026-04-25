#!/usr/bin/env python3
import sys
if sys.stdout.encoding != 'utf-8':
    sys.stdout.reconfigure(encoding='utf-8')
if sys.stderr.encoding != 'utf-8':
    sys.stderr.reconfigure(encoding='utf-8')

"""
Skill 初始化脚本 - 从模板创建新 skill

用法:
    init_skill.py <skill-name> --path <path>

示例:
    init_skill.py my-new-skill --path skills
    init_skill.py my-api-helper --path ./skills
"""

import sys
from pathlib import Path


SKILL_TEMPLATE = """---
name: {skill_name}
description: "TODO: 完整描述 skill 的功能和触发场景。必须包含何时使用信息，具体场景、文件类型或任务。"
---

# {skill_title}

## 概述

[TODO: 1-2 句话说明此 skill 的用途]

## 结构选择

[TODO: 根据需求选择最合适的结构模式：

**1. 工作流型**（适合顺序流程）
- 有清晰的步骤化流程时使用
- 结构: ## 概述 → ## 工作流 → ## 步骤 1 → ## 步骤 2...

**2. 任务型**（适合工具集合）
- 提供多种独立操作时使用
- 结构: ## 概述 → ## 快速开始 → ## 任务 1 → ## 任务 2...

**3. 参考规范型**（适合标准或规范）
- 品牌指南、编码规范、需求文档时使用
- 结构: ## 概述 → ## 规范 → ## 规格 → ## 用法...

**4. 能力型**（适合集成系统）
- 提供多个关联功能时使用
- 结构: ## 概述 → ## 核心能力 → ### 功能 1 → ### 功能 2...

模式可以混合使用。完成后删除本段。]

## [TODO: 替换为第一个主要章节]

[TODO: 添加内容。参考示例：
- 技术类 skill 提供代码示例
- 复杂工作流提供决策树
- 真实用户请求的具体示例
- 按需引用 scripts/templates/references]

## 资源目录

### scripts/
可执行代码（Python/Bash 等），用于需要确定性执行或频繁重写的操作。
不是每个 skill 都需要，不需要时删除。

### references/
按需加载的参考文档，保持 SKILL.md 精简。
不是每个 skill 都需要，不需要时删除。

### assets/
输出中使用的模板、图片等文件，不加载到上下文。
不是每个 skill 都需要，不需要时删除。
"""

EXAMPLE_SCRIPT = '''#!/usr/bin/env python3
"""
{skill_name} 示例脚本

这是占位脚本，可替换为实际实现或删除。
"""

def main():
    print("这是 {skill_name} 的示例脚本")
    # TODO: 添加实际逻辑

if __name__ == "__main__":
    main()
'''

EXAMPLE_REFERENCE = """# {skill_title} 参考文档

这是占位参考文档，可替换为实际内容或删除。

参考文档适用于：
- 详细的 API 文档
- 复杂的多步流程指南
- 不适合放在 SKILL.md 中的长篇幅内容
"""

EXAMPLE_ASSET = """# 示例资源文件

这是占位资源，可替换为实际文件（模板、图片、字体等）或删除。

资源文件不会被加载到上下文，而是在输出中使用。

常见资源类型：
- 模板: .pptx, .docx, 项目目录
- 图片: .png, .jpg, .svg
- 字体: .ttf, .otf, .woff2
- 数据: .csv, .json, .yaml
"""


def title_case_skill_name(skill_name):
    """将 kebab-case 转为 Title Case 用于显示。"""
    return ' '.join(word.capitalize() for word in skill_name.split('-'))


def init_skill(skill_name, path):
    """初始化新 skill 目录。"""
    skill_dir = Path(path).resolve() / skill_name

    if skill_dir.exists():
        print(f"错误: skill 目录已存在: {skill_dir}")
        return None

    try:
        skill_dir.mkdir(parents=True, exist_ok=False)
        print(f"已创建 skill 目录: {skill_dir}")
    except Exception as e:
        print(f"错误: 创建目录失败: {e}")
        return None

    skill_title = title_case_skill_name(skill_name)
    skill_content = SKILL_TEMPLATE.format(
        skill_name=skill_name,
        skill_title=skill_title
    )

    skill_md_path = skill_dir / 'SKILL.md'
    try:
        skill_md_path.write_text(skill_content, encoding='utf-8')
        print("已创建 SKILL.md")
    except Exception as e:
        print(f"错误: 创建 SKILL.md 失败: {e}")
        return None

    try:
        scripts_dir = skill_dir / 'scripts'
        scripts_dir.mkdir(exist_ok=True)
        example_script = scripts_dir / 'example.py'
        example_script.write_text(EXAMPLE_SCRIPT.format(skill_name=skill_name), encoding='utf-8')
        example_script.chmod(0o755)
        print("已创建 scripts/example.py")

        references_dir = skill_dir / 'references'
        references_dir.mkdir(exist_ok=True)
        example_reference = references_dir / 'api_reference.md'
        example_reference.write_text(EXAMPLE_REFERENCE.format(skill_title=skill_title), encoding='utf-8')
        print("已创建 references/api_reference.md")

        assets_dir = skill_dir / 'assets'
        assets_dir.mkdir(exist_ok=True)
        example_asset = assets_dir / 'example_asset.txt'
        example_asset.write_text(EXAMPLE_ASSET, encoding='utf-8')
        print("已创建 assets/example_asset.txt")
    except Exception as e:
        print(f"错误: 创建资源目录失败: {e}")
        return None

    print(f"\nSkill '{skill_name}' 初始化成功: {skill_dir}")
    print("\n后续步骤:")
    print("1. 编辑 SKILL.md，完成 TODO 项并更新 description")
    print("2. 自定义或删除 scripts/、references/、assets/ 中的示例文件")
    print("3. 创建 junction 链接到 .claude/skills/（如需在 Claude Code 中使用）")

    return skill_dir


def main():
    if len(sys.argv) < 4 or sys.argv[2] != '--path':
        print("用法: init_skill.py <skill-name> --path <path>")
        print("\nSkill 名称要求:")
        print("  - kebab-case 格式（如 'my-data-analyzer'）")
        print("  - 仅小写字母、数字和连字符")
        print("  - 最多 64 个字符")
        print("\n示例:")
        print("  init_skill.py my-new-skill --path skills")
        print("  init_skill.py my-api-helper --path ./skills")
        sys.exit(1)

    skill_name = sys.argv[1]
    path = sys.argv[3]

    print(f"正在初始化 skill: {skill_name}")
    print(f"  位置: {path}")
    print()

    result = init_skill(skill_name, path)

    sys.exit(0 if result else 1)


if __name__ == "__main__":
    main()

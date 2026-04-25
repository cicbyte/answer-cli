#!/usr/bin/env python3
import sys
if sys.stdout.encoding != 'utf-8':
    sys.stdout.reconfigure(encoding='utf-8')
if sys.stderr.encoding != 'utf-8':
    sys.stderr.reconfigure(encoding='utf-8')

"""
Skill 快速验证脚本 - 基础版本
"""

import sys
import os
import re
import yaml
from pathlib import Path

def validate_skill(skill_path):
    """基础 skill 验证"""
    skill_path = Path(skill_path)

    skill_md = skill_path / 'SKILL.md'
    if not skill_md.exists():
        return False, "未找到 SKILL.md"

    content = skill_md.read_text(encoding='utf-8')
    if not content.startswith('---'):
        return False, "未找到 YAML frontmatter"

    match = re.match(r'^---\n(.*?)\n---', content, re.DOTALL)
    if not match:
        return False, "frontmatter 格式无效"

    frontmatter_text = match.group(1)

    try:
        frontmatter = yaml.safe_load(frontmatter_text)
        if not isinstance(frontmatter, dict):
            return False, "frontmatter 必须是 YAML 字典"
    except yaml.YAMLError as e:
        return False, f"YAML 解析失败: {e}"

    ALLOWED_PROPERTIES = {'name', 'description', 'license', 'allowed-tools', 'metadata', 'compatibility'}

    unexpected_keys = set(frontmatter.keys()) - ALLOWED_PROPERTIES
    if unexpected_keys:
        return False, (
            f"frontmatter 包含不允许的字段: {', '.join(sorted(unexpected_keys))}。"
            f"允许的字段: {', '.join(sorted(ALLOWED_PROPERTIES))}"
        )

    if 'name' not in frontmatter:
        return False, "缺少 'name' 字段"
    if 'description' not in frontmatter:
        return False, "缺少 'description' 字段"

    name = frontmatter.get('name', '')
    if not isinstance(name, str):
        return False, f"name 必须是字符串，当前为 {type(name).__name__}"
    name = name.strip()
    if name:
        if not re.match(r'^[a-z0-9-]+$', name):
            return False, f"name '{name}' 应使用 kebab-case（仅小写字母、数字和连字符）"
        if name.startswith('-') or name.endswith('-') or '--' in name:
            return False, f"name '{name}' 不能以连字符开头/结尾或包含连续连字符"
        if len(name) > 64:
            return False, f"name 过长（{len(name)} 字符），最大 64 字符"

    description = frontmatter.get('description', '')
    if not isinstance(description, str):
        return False, f"description 必须是字符串，当前为 {type(description).__name__}"
    description = description.strip()
    if description:
        if '<' in description or '>' in description:
            return False, "description 不能包含尖括号（< 或 >）"
        if len(description) > 1024:
            return False, f"description 过长（{len(description)} 字符），最大 1024 字符"

    compatibility = frontmatter.get('compatibility', '')
    if compatibility:
        if not isinstance(compatibility, str):
            return False, f"compatibility 必须是字符串，当前为 {type(compatibility).__name__}"
        if len(compatibility) > 500:
            return False, f"compatibility 过长（{len(compatibility)} 字符），最大 500 字符"

    return True, "验证通过"

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("用法: python quick_validate.py <skill-directory>")
        sys.exit(1)

    valid, message = validate_skill(sys.argv[1])
    print(message)
    sys.exit(0 if valid else 1)

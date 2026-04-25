#!/usr/bin/env python3
import io, sys
if sys.stdout.encoding != 'utf-8':
    sys.stdout.reconfigure(encoding='utf-8')
if sys.stderr.encoding != 'utf-8':
    sys.stderr.reconfigure(encoding='utf-8')

"""
Skill 打包脚本 - 将 skill 文件夹打包为可分发的 .skill 文件

用法:
    python package_skill.py <skill-path> [output-dir]

示例:
    python package_skill.py skills/my-skill
    python package_skill.py skills/my-skill ./dist
"""

import sys
import zipfile
from pathlib import Path
from quick_validate import validate_skill


def package_skill(skill_path, output_dir=None):
    """将 skill 文件夹打包为 .skill 文件。"""
    skill_path = Path(skill_path).resolve()

    if not skill_path.exists():
        print(f"错误: skill 目录不存在: {skill_path}")
        return None

    if not skill_path.is_dir():
        print(f"错误: 路径不是目录: {skill_path}")
        return None

    skill_md = skill_path / "SKILL.md"
    if not skill_md.exists():
        print(f"错误: 未找到 SKILL.md: {skill_path}")
        return None

    print("正在验证 skill...")
    valid, message = validate_skill(skill_path)
    if not valid:
        print(f"验证失败: {message}")
        print("  请修复验证错误后再打包。")
        return None
    print(f"验证通过: {message}\n")

    skill_name = skill_path.name
    if output_dir:
        output_path = Path(output_dir).resolve()
        output_path.mkdir(parents=True, exist_ok=True)
    else:
        output_path = Path.cwd()

    skill_filename = output_path / f"{skill_name}.skill"

    try:
        with zipfile.ZipFile(skill_filename, 'w', zipfile.ZIP_DEFLATED) as zipf:
            for file_path in skill_path.rglob('*'):
                if file_path.is_file():
                    arcname = file_path.relative_to(skill_path.parent)
                    zipf.write(file_path, arcname)
                    print(f"  已添加: {arcname}")

        print(f"\n打包成功: {skill_filename}")
        return skill_filename

    except Exception as e:
        print(f"打包失败: {e}")
        return None


def main():
    if len(sys.argv) < 2:
        print("用法: python package_skill.py <skill-path> [output-dir]")
        print("\n示例:")
        print("  python package_skill.py skills/my-skill")
        print("  python package_skill.py skills/my-skill ./dist")
        sys.exit(1)

    skill_path = sys.argv[1]
    output_dir = sys.argv[2] if len(sys.argv) > 2 else None

    print(f"正在打包 skill: {skill_path}")
    if output_dir:
        print(f"  输出目录: {output_dir}")
    print()

    result = package_skill(skill_path, output_dir)
    sys.exit(0 if result else 1)


if __name__ == "__main__":
    main()

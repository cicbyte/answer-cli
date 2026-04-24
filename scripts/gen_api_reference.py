#!/usr/bin/env python3
"""从 swagger.json 生成 docs/api-reference.md，保持手工整理的结构化格式。"""

import json
import os
import re
import sys
from collections import defaultdict


def load_swagger(path: str) -> dict:
    with open(path, "r", encoding="utf-8") as f:
        return json.load(f)


def ref_name(ref: str) -> str:
    m = re.match(r"^#/definitions/(?:[\w.]+\.)?(\w+)$", ref)
    return m.group(1) if m else ref


def resolve_ref(definitions: dict, ref: str) -> dict | None:
    m = re.match(r"^#/definitions/(.+)$", ref)
    if not m:
        return None
    return definitions.get(m.group(1))


def detect_auth(spec: dict) -> str:
    sec = spec.get("security", [])
    if not sec:
        return "public"
    if any("ApiKeyAuth" in s for s in sec):
        return "admin"
    return "auth"


def short_path(path: str) -> str:
    for prefix in ["/answer/api/v1", "/answer/admin/api", "/installation"]:
        if path.startswith(prefix):
            return path[len(prefix):]
    return path


def get_param_info(spec: dict, definitions: dict) -> str:
    """返回参数摘要，如 'query: q, order' 或 'body: QuestionPageReq'"""
    params = spec.get("parameters", [])
    if not params:
        return ""
    first_in = params[0].get("in", "query")
    if first_in == "body":
        schema = params[0].get("schema", {})
        if "$ref" in schema:
            return ref_name(schema["$ref"])
        return "body"
    names = [p.get("name", "") for p in params]
    return f"{first_in}: {', '.join(names)}"


def get_response_data_type(spec: dict, definitions: dict) -> str:
    """提取响应 data 字段的类型名"""
    resp200 = spec.get("responses", {}).get("200", {})
    schema = resp200.get("schema", {})
    if not schema:
        return ""
    if "allOf" in schema:
        for sub in schema["allOf"]:
            props = sub.get("properties", {}).get("data", {})
            if "$ref" in props:
                return ref_name(props["$ref"])
            # nested allOf
            nested = props.get("allOf", [])
            for n in nested:
                if "$ref" in n:
                    return ref_name(n["$ref"])
    return ""


def collect_all_refs(obj, depth=0):
    """收集所有 $ref 引用的短名称"""
    if depth > 12 or not isinstance(obj, (dict, list)):
        return set()
    refs = set()
    items = obj.values() if isinstance(obj, dict) else obj
    for v in items:
        if isinstance(v, str) and v.startswith("#/definitions/"):
            refs.add(ref_name(v))
        elif isinstance(v, (dict, list)):
            refs |= collect_all_refs(v, depth + 1)
    return refs


# tag 到中文名称和分组顺序的映射
TAG_META = {
    "Question": ("问题管理", 1),
    "Answer": ("回答管理", 2),
    "Tag": ("标签管理", 3),
    "User": ("用户管理", 4),
    "Comment": ("评论系统", 5),
    "Activity": ("投票与关注", 6),
    "Notification": ("通知系统", 7),
    "ai-conversation": ("AI 对话", 8),
    "ai-conversation-admin": ("AI 对话（管理）", 8),
    "Search": ("搜索", 9),
    "Report": ("举报", 10),
    "Review": ("内容审核", 11),
    "Revision": ("版本修订", 12),
    "Collection": ("收藏", 13),
    "api-badge": ("徽章系统", 14),
    "AdminBadge": ("徽章管理", 14),
    "Upload": ("文件上传", 15),
    "Plugin": ("插件系统", 16),
    "PluginConnector": ("外部登录连接器", 16),
    "PluginRender": ("插件渲染", 16),
    "UserPlugin": ("用户插件", 16),
    "AdminPlugin": ("插件管理", 16),
    "Personal": ("个人中心", 17),
    "Meta": ("表情反应", 18),
    "Permission": ("权限", 19),
    "Lang": ("多语言", 20),
    "Site": ("站点公开信息", 21),
    "installation": ("系统安装", 22),
    "admin": ("管理后台", 99),
    "Rank": ("排名", 23),
    "reason": ("举报原因", 24),
}


def main():
    swagger_path = os.path.join(os.path.dirname(__file__), "..", "docs", "swagger.json")
    if len(sys.argv) > 1:
        swagger_path = sys.argv[1]

    swagger = load_swagger(swagger_path)
    definitions = swagger.get("definitions", {})
    paths = swagger.get("paths", {})

    # 按分组收集端点
    groups: dict[str, list[tuple[str, str, dict]]] = defaultdict(list)
    all_refs: set[str] = set()

    for path, methods in paths.items():
        if "/answer/" not in path and path != "/":
            continue
        for method, spec in methods.items():
            tags = spec.get("tags", ["other"])
            tag = tags[0] if tags else "other"
            groups[tag].append((path, method.upper(), spec))
            all_refs |= collect_all_refs(spec)

    # 统计
    total = sum(len(v) for v in groups.values())
    total_defs = len(definitions)

    lines = []
    lines.append("# Apache Answer API 参考")
    lines.append("")
    lines.append(f"> 从 `docs/swagger.json` 自动生成，共 **{total} 个端点**，{total_defs} 个数据模型")
    lines.append("")
    lines.append("## 认证方式")
    lines.append("")
    lines.append("- **Auth**：通过 HTTP Header `Authorization` 传递 Token，写操作和用户相关操作需要认证")
    lines.append("- **ApiKeyAuth**：管理员 API Key 认证")
    lines.append("- **Public**：无需认证的公开端点")
    lines.append("")
    lines.append("## API 基础路径")
    lines.append("")
    lines.append("- 用户端：`/answer/api/v1/`")
    lines.append("- 管理端：`/answer/admin/api/`")
    lines.append("")
    lines.append("---")
    lines.append("")

    # 按 TAG_META 排序输出
    sorted_tags = sorted(groups.keys(), key=lambda t: TAG_META.get(t, (t, 50)))

    for tag in sorted_tags:
        endpoints = groups[tag]
        meta = TAG_META.get(tag, (tag, 50))
        cn_name = meta[0]

        # 检查是否是管理端
        has_admin = any("/admin/" in p for p, _, _ in endpoints)
        has_public = any("/admin/" not in p for p, _, _ in endpoints)

        lines.append(f"### {cn_name} — `{tag}`（{len(endpoints)} 个端点）")
        lines.append("")
        lines.append("| 方法 | 路径 | 认证 | 参数 | 响应 | 说明 |")
        lines.append("|------|------|------|------|------|------|")

        for path, method, spec in endpoints:
            auth = detect_auth(spec)
            summary = spec.get("summary", "")
            sp = short_path(path)
            param_info = get_param_info(spec, definitions)
            resp_type = get_response_data_type(spec, definitions)

            auth_label = {"public": "", "auth": "🔒", "admin": "🔑"}.get(auth, "")
            lines.append(f"| {method} | `{sp}` | {auth_label} | {param_info} | `{resp_type}` | {summary} |")

        lines.append("")

    # 核心数据模型
    lines.append("---")
    lines.append("")
    lines.append("## 核心数据模型")
    lines.append("")
    lines.append("详细模型定义见 `docs/api/*.models.md`，以下按包分组列出被 API 引用的模型。")
    lines.append("")

    # 按包前缀分组
    pkg_models: dict[str, list[str]] = defaultdict(list)
    for full_key in definitions:
        parts = full_key.split(".", 1)
        if len(parts) == 2:
            pkg, name = parts
        else:
            pkg, name = "other", parts[0]
        if name in all_refs:
            pkg_models[pkg].append(name)

    pkg_cn = {
        "schema": "核心业务模型",
        "handler": "响应包装",
        "pager": "分页",
        "plugin": "插件",
        "constant": "常量",
        "install": "安装",
        "entity": "实体",
        "other": "其他",
    }

    for pkg in sorted(pkg_models.keys()):
        models = sorted(pkg_models[pkg])
        cn = pkg_cn.get(pkg, pkg)
        lines.append(f"### {cn}（{len(models)} 个）")
        lines.append("")
        # 每行最多 6 个，避免行太长
        for i in range(0, len(models), 6):
            chunk = models[i:i+6]
            lines.append(", ".join(f"`{n}`" for n in chunk))
        lines.append("")

    lines.append("")

    out_path = os.path.join(os.path.dirname(__file__), "..", "docs", "api-reference.md")
    with open(out_path, "w", encoding="utf-8") as f:
        f.write("\n".join(lines))
    print(f"写入 {out_path} ({total} 个端点, {len(all_refs)} 个引用模型)")


if __name__ == "__main__":
    main()

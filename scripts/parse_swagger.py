#!/usr/bin/env python3
"""解析 Apache Answer swagger.json，按 tag 拆分为 Markdown 文档。"""

import json
import os
import re
import sys
from collections import defaultdict


def load_swagger(path: str) -> dict:
    with open(path, "r", encoding="utf-8") as f:
        return json.load(f)


def resolve_ref(definitions: dict, ref: str) -> dict | None:
    """解析 $ref 引用，返回定义内容。"""
    m = re.match(r"^#/definitions/(.+)$", ref)
    if not m:
        return None
    return definitions.get(m.group(1))


def ref_name(ref: str) -> str:
    """从 $ref 中提取短名称（去掉包前缀）。"""
    m = re.match(r"^#/definitions/(?:[\w.]+\.)?(\w+)$", ref)
    return m.group(1) if m else ref


def resolve_schema(definitions: dict, schema: dict, depth: int = 0) -> str:
    """递归解析 schema，返回 Markdown 描述。最多递归 3 层。"""
    if depth > 3 or not schema:
        return ""

    if "$ref" in schema:
        name = ref_name(schema["$ref"])
        return f"`{name}`"

    parts = []

    t = schema.get("type", "")
    if t == "array":
        items_str = resolve_schema(definitions, schema.get("items", {}), depth + 1)
        return f"[]`{items_str}`" if items_str else "`array`"

    if t == "object" and "properties" in schema:
        props = schema["properties"]
        for pname, pspec in props.items():
            desc = pspec.get("description", "")
            type_str = resolve_schema(definitions, pspec, depth + 1)
            req = ""
            # 检查 required 列表
            if pname in schema.get("required", []):
                req = " **(必填)**"
            line = f"  - `{pname}` {type_str}{req}"
            if desc:
                line += f" — {desc}"
            parts.append(line)
        return "\n".join(parts)

    if "allOf" in schema:
        for sub in schema["allOf"]:
            s = resolve_schema(definitions, sub, depth + 1)
            if s:
                parts.append(s)
        return "\n".join(parts)

    enum = schema.get("enum", [])
    if enum:
        return f"`{t}` enum: {', '.join(f'`{v}`' for v in enum)}"

    return f"`{t}`" if t else ""


def detect_auth(spec: dict) -> str:
    sec = spec.get("security", [])
    if not sec:
        return "public"
    if any("ApiKeyAuth" in s for s in sec):
        return "admin"
    return "auth"


def detect_param_style(spec: dict) -> str:
    params = spec.get("parameters", [])
    if not params:
        return "none"
    ins = {p.get("in", "") for p in params}
    if "body" in ins:
        return "body"
    if "query" in ins:
        return "query"
    if "path" in ins:
        return "path"
    return "mixed"


def format_params(spec: dict, definitions: dict) -> str:
    params = spec.get("parameters", [])
    if not params:
        return ""

    lines = []
    in_style = params[0].get("in", "query")

    if in_style == "body":
        schema = params[0].get("schema", {})
        desc = params[0].get("description", "")
        if "$ref" in schema:
            name = ref_name(schema["$ref"])
            resolved = resolve_ref(definitions, schema["$ref"])
            lines.append(f"**Body**: `{name}`")
            if desc:
                lines.append(f"> {desc}")
            if resolved:
                props = resolved.get("properties", {})
                required = resolved.get("required", [])
                for pname, pspec in props.items():
                    ptype = resolve_schema(definitions, pspec)
                    req = " **(必填)**" if pname in required else ""
                    pdesc = pspec.get("description", "")
                    line = f"  - `{pname}` {ptype}{req}"
                    if pdesc:
                        line += f" — {pdesc}"
                    lines.append(line)
        elif "type" in schema:
            lines.append(f"**Body**: {schema['type']}")
        return "\n".join(lines)

    for p in params:
        name = p.get("name", "")
        ptype = p.get("type", "string")
        desc = p.get("description", "")
        loc = p.get("in", "query")
        req = " **(必填)**" if p.get("required") else ""
        line = f"  - `{name}` (`{ptype}`{req})"
        if desc:
            line += f" — {desc}"
        lines.append(line)
    return "\n".join(lines)


def format_response(spec: dict, definitions: dict) -> str:
    responses = spec.get("responses", {})
    if not responses:
        return ""

    lines = []
    for code, rspec in responses.items():
        schema = rspec.get("schema", {})
        if not schema:
            lines.append(f"- `{code}`: {rspec.get('description', '')}")
            continue
        lines.append(f"- `{code}`:")
        s = resolve_schema(definitions, schema)
        if s:
            lines.append(s)
    return "\n".join(lines)


def format_definition(name: str, defn: dict, definitions: dict) -> str:
    lines = [f"### `{name}`", ""]
    desc = defn.get("description", "")
    if desc:
        lines.append(desc)
        lines.append("")

    props = defn.get("properties", {})
    required = defn.get("required", [])
    if props:
        for pname, pspec in props.items():
            ptype = resolve_schema(definitions, pspec)
            req = " **(必填)**" if pname in required else ""
            pdesc = pspec.get("description", "")
            line = f"- `{pname}` {ptype}{req}"
            if pdesc:
                line += f" — {pdesc}"
            lines.append(line)
    return "\n".join(lines)


def short_path(path: str) -> str:
    """去掉 /answer/api/v1 或 /answer/admin/api 前缀，保留有意义的部分。"""
    for prefix in ["/answer/api/v1", "/answer/admin/api"]:
        if path.startswith(prefix):
            return path[len(prefix):]
    return path


def is_api_path(path: str) -> bool:
    return "/answer/" in path or path == "/"


def main():
    swagger_path = os.path.join(os.path.dirname(__file__), "..", "docs", "swagger.json")
    if len(sys.argv) > 1:
        swagger_path = sys.argv[1]

    out_dir = os.path.join(os.path.dirname(__file__), "..", "docs", "api")
    os.makedirs(out_dir, exist_ok=True)

    swagger = load_swagger(swagger_path)
    definitions = swagger.get("definitions", {})
    paths = swagger.get("paths", {})

    # 按 tag 分组
    by_tag: dict[str, list[tuple[str, str, dict]]] = defaultdict(list)
    for path, methods in paths.items():
        if not is_api_path(path):
            continue
        for method, spec in methods.items():
            tags = spec.get("tags", ["other"])
            tag = tags[0] if tags else "other"
            by_tag[tag].append((path, method.upper(), spec))

    # 写 index
    index_lines = ["# Apache Answer API Reference", ""]
    index_lines.append(f"共 {sum(len(v) for v in by_tag.values())} 个端点，{len(definitions)} 个数据模型\n")

    for tag in sorted(by_tag.keys()):
        endpoints = by_tag[tag]
        index_lines.append(f"## {tag} ({len(endpoints)})")
        index_lines.append("")
        for path, method, spec in endpoints:
            auth = detect_auth(spec)
            summary = spec.get("summary", "")
            sp = short_path(path)
            auth_badge = {"public": "", "auth": "🔒", "admin": "🔑"}.get(auth, "")
            index_lines.append(f"- {auth_badge}`{method}` `{sp}` — {summary}")
        index_lines.append("")

    index_path = os.path.join(out_dir, "index.md")
    with open(index_path, "w", encoding="utf-8") as f:
        f.write("\n".join(index_lines))
    print(f"写入 {index_path}")

    # 写各 tag 文件
    # 收集所有被引用的 definition
    used_defs: dict[str, set[str]] = defaultdict(set)

    for tag, endpoints in by_tag.items():
        lines = [f"# {tag}", ""]
        for path, method, spec in endpoints:
            auth = detect_auth(spec)
            summary = spec.get("summary", "")
            desc = spec.get("description", "")
            sp = short_path(path)

            lines.append(f"## `{method}` `{sp}`")
            lines.append("")

            auth_label = {"public": "公开", "auth": "需认证", "admin": "管理员"}.get(auth, auth)
            param_style = detect_param_style(spec)
            meta = f"**认证**: {auth_label}"
            if param_style != "none":
                meta += f" | **参数方式**: {param_style}"
            lines.append(meta)
            lines.append("")

            if desc and desc != summary:
                lines.append(f"> {desc}")
                lines.append("")

            # 请求参数
            params_str = format_params(spec, definitions)
            if params_str:
                lines.append("**请求参数:**")
                lines.append("")
                lines.append(params_str)
                lines.append("")

            # 收集引用的 definition
            def collect_refs(obj, depth=0):
                if depth > 10 or not isinstance(obj, (dict, list)):
                    return
                items = obj.values() if isinstance(obj, dict) else obj
                for v in items:
                    if isinstance(v, str) and v.startswith("#/definitions/"):
                        full_key = v[len("#/definitions/"):]
                        used_defs[tag].add(full_key)
                    elif isinstance(v, (dict, list)):
                        collect_refs(v, depth + 1)
            collect_refs(spec)

            # 响应
            resp_str = format_response(spec, definitions)
            if resp_str:
                lines.append("**响应:**")
                lines.append("")
                lines.append(resp_str)
                lines.append("")

            lines.append("---")
            lines.append("")

        tag_path = os.path.join(out_dir, f"{tag.lower()}.md")
        with open(tag_path, "w", encoding="utf-8") as f:
            f.write("\n".join(lines))
        print(f"写入 {tag_path} ({len(endpoints)} 个端点)")

    # 写各 tag 关联的 definitions
    for tag, def_keys in used_defs.items():
        if not def_keys:
            continue
        def_lines = [f"# {tag} — 数据模型", ""]
        for full_key in sorted(def_keys):
            defn = definitions.get(full_key)
            if not defn:
                continue
            short = ref_name(f"#/definitions/{full_key}")
            def_lines.append(format_definition(short, defn, definitions))
            def_lines.append("")
            def_lines.append("---")
            def_lines.append("")

        def_path = os.path.join(out_dir, f"{tag.lower()}.models.md")
        with open(def_path, "w", encoding="utf-8") as f:
            f.write("\n".join(def_lines))
        print(f"写入 {def_path} ({len(def_keys)} 个模型)")


if __name__ == "__main__":
    main()

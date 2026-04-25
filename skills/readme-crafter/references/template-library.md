# 库/SDK 模板

适用于：npm 包、Python 库、Go 模块、Rust crate、Ruby gem 等。

```markdown
# {PackageName}

> {一句话说明：这个库做什么，为什么开发者需要它}

{徽章：构建状态、版本号、覆盖率、许可证}

{如果库有可视化输出则放截图/GIF，否则省略}

## ✨ 功能特性

- **{功能 1}** — {带来的价值，而非功能描述本身}
- **{功能 2}** — {为什么这对用户很重要}
- **{功能 3}** — {如果有数据支撑则附上具体指标}

## 📦 安装

{选择适用的安装方式}

\`\`\`bash
# npm
npm install {package}

# yarn
yarn add {package}

# pnpm
pnpm add {package}
\`\`\`

\`\`\`bash
# pip
pip install {package}
\`\`\`

\`\`\`bash
# go
go get {module}
\`\`\`

## 🚀 快速开始

{最小可运行示例 — 最多 5-10 行。读者复制粘贴就能看到结果。}

\`\`\`{language}
import { PackageName } from '{package}';

const result = PackageName.doSomething({ input: 'hello' });
console.log(result);
\`\`\`

## 📖 使用方法

### {场景 1：最常用}

{一句话引入}

\`\`\`{language}
{展示最常见的使用模式的代码示例}
\`\`\`

### {场景 2：进阶用法}

{一句话引入}

\`\`\`{language}
{展示进阶或较少见的使用模式的代码示例}
\`\`\`

### {场景 3：特殊场景/集成}

{一句话引入}

\`\`\`{language}
{代码示例}
\`\`\`

## ⚙️ API 参考

{列出主要的导出函数/类/方法。保持精简 — 如有完整文档则链接过去。}

### `functionName(param1, param2)`

{一句话描述}

| 参数 | 类型 | 默认值 | 说明 |
|---|---|---|---|
| `param1` | `{type}` | `{default}` | {说明} |
| `param2` | `{type}` | — | {说明} |

**返回值：** `{returnType}` — {说明}

---

### `ClassName`

{一句话描述}

| 方法 | 参数 | 返回值 | 说明 |
|---|---|---|---|
| `constructor(options)` | `{Options}` | `ClassName` | {说明} |
| `.method()` | — | `{type}` | {说明} |

## 🏗️ 实现原理

{可选：简要说明架构、算法或设计决策。仅当库有用户需要了解的非显而易见的内部机制时保留。}

## 🤝 参与贡献

欢迎贡献！请先阅读 [贡献指南](CONTRIBUTING.md)。

## 📄 开源许可证

[MIT](LICENSE) © {年份} {作者}
```

### 库类 README 写作要点

- 快速开始必须可以复制粘贴后直接运行
- API 参考只覆盖公共接口
- 如果有知名替代品，增加「与 XX 的对比」章节
- 库支持 TypeScript 时展示类型/接口定义
- 支持树摇优化的库应注明
